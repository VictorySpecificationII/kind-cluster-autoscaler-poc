//every x seconds
//1. Get nodes that are cordoned off
//  -> create list
//2. Out of list of available nodes, filter for worker nodes
//  -> create list
//3. Observe load on nodes
//4. Reschedule pods on less loaded nodes


package main

import ("fmt"
	"os/exec"
//	"strconv"
	"strings"
//	"time"
       )


func get_cordoned_off_nodes() map[string]string {

	cordoned_off_nodes := make(map[string]string)
	cmd := exec.Command("kubectl", "get", "nodes", "-o", "wide")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running kubectl: ", err)
		fmt.Println("Output: ", string(output))
		return nil
	}

	lines := strings.Split(string(output), "\n")
//	fmt.Println(lines) //debug

	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

	parts := strings.Fields(line)
//	fmt.Println(parts) //debug

	node := parts[0]
	status := parts[1]

//	fmt.Println(status) //debug

	if strings.Contains(status, "SchedulingDisabled"){
		cordoned_off_nodes[node] = status
		}
//	fmt.Println(cordoned_off_nodes) //debug
	}
	return cordoned_off_nodes
}

func get_worker_nodes() map[string]string{

	worker_nodes := make(map[string]string)
	cmd := exec.Command("kubectl", "get", "nodes", "-l", "role=worker")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running kubectl: ", err)
		fmt.Println("Output: ", string(output))
		return nil
	}

	lines := strings.Split(string(output), "\n")
//	fmt.Println(lines) //debug
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

	parts := strings.Fields(line)
//	fmt.Println(parts) //debug

	node := parts[0]
	status := parts[1]

//	fmt.Println(status) //debug
	if !strings.Contains(status, "SchedulingDisabled"){
	worker_nodes[node] = status
	}

//	fmt.Println(cordoned_off_nodes) //debug

	}

	return worker_nodes


}

func observe_cluster_load_on_available_workers() map[string]string {
	worker_node_loads := make(map[string]string)
	cmd := exec.Command("kubectl", "top", "nodes", "-l", "role=worker")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running kubectl: ", err)
		fmt.Println("Output: ", string(output))
		return nil
	}

	lines := strings.Split(string(output), "\n")
//	fmt.Println(lines) //debug

        for _, line := range lines[1:] {
                if line == "" {
                        continue
                }

        parts := strings.Fields(line)
//      fmt.Println(parts) //debug

        node := parts[0]
        load := parts[2]

//      fmt.Println(status) //debug
        if strings.Contains(load, "%"){
        worker_node_loads[node] = load
        }

//	fmt.Println(worker_node_loads) //debug

	}

	return worker_node_loads
}

//func autoscale_pods_on_available_workers() {
//
//}

func main(){
//	cordoned_nodes := get_cordoned_off_nodes()
//	fmt.Println(cordoned_nodes)
//	worker_nodes := get_worker_nodes()
//	fmt.Println(worker_nodes)
	load := observe_cluster_load_on_available_workers()
	fmt.Println(load)
}
