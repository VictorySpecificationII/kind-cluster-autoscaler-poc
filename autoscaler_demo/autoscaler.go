package main

import (
    "fmt"
    "os/exec"
    "strconv"
    "strings"
    "time"
)

func main() {
    for {
        // 1. Get node CPU usage
        cmd := exec.Command("kubectl", "top", "nodes")
        output, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Println("Error running kubectl:", err)
            fmt.Println("Output:", string(output))
            return
        }

        lines := strings.Split(string(output), "\n")
        nodeCPU := make(map[string]int)
        minCPU := 1000
        chosenNode := ""

        // 2. Filter for worker nodes and find least-loaded
        for _, line := range lines[1:] {
            if line == "" {
                continue
            }
            parts := strings.Fields(line)
            node := parts[0]

            // Only consider nodes labeled "role=worker"
            labelCmd := exec.Command("kubectl", "get", "node", node, "-o", "jsonpath={.metadata.labels.role}")
            labelOut, err := labelCmd.Output()
            if err != nil || strings.TrimSpace(string(labelOut)) != "worker" {
                continue
            }

            cpuStr := strings.TrimSuffix(parts[1], "m")
            cpu, err := strconv.Atoi(cpuStr)
            if err != nil {
                continue
            }

            nodeCPU[node] = cpu
            if cpu < minCPU {
                minCPU = cpu
                chosenNode = node
            }
        }

        if chosenNode == "" {
            fmt.Println("No worker nodes found, skipping iteration.")
            time.Sleep(5 * time.Second)
            continue
        }

        fmt.Printf("Least-loaded worker node: %s (%dm)\n", chosenNode, minCPU)

        // 3. Observe existing pods with label "app=stress" and reschedule if needed
        podsCmd := exec.Command("kubectl", "get", "pods", "-l", "app=stress", "-o",
            "jsonpath={range .items[*]}{.metadata.name} {.spec.nodeName}{\"\\n\"}{end}")
        podOutput, _ := podsCmd.CombinedOutput()
        podLines := strings.Split(string(podOutput), "\n")

        for _, line := range podLines {
            if line == "" {
                continue
            }
            parts := strings.Fields(line)
            podName := parts[0]
            nodeName := parts[1]

            // Lower threshold to 5m CPU for demo
            if nodeCPU[nodeName] > minCPU+5 {
                fmt.Printf("[MOVE] pod %s from %s -> %s\n", podName, nodeName, chosenNode)
                exec.Command("kubectl", "delete", "pod", podName).Run()

                // Recreate on chosen node
                deployCmd := exec.Command("kubectl", "run", podName, "--image=alpine",
                    "--restart=Never", "--overrides",
                    fmt.Sprintf(`{"apiVersion": "v1","spec":{"nodeSelector":{"kubernetes.io/hostname":"%s"}}}`, chosenNode),
                    "--", "/bin/sh", "-c", "while true; do echo running; sleep 10; done")
                deployCmd.Run()
            } else {
                fmt.Printf("[KEEP] pod %s stays on %s\n", podName, nodeName)
            }
        }

        time.Sleep(5 * time.Second)
    }
}
