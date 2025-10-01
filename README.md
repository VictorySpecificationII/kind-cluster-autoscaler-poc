# kind-cluster-autoscaler-poc

This is a PoC for a basic autoscaler written in Go, and a Kubernetes cluster running on KinD.

The standard setup is 3 masters and 3 workers to maintain minimum HA.

Since adding nodes manually to an existing cluster is flaky, i've added another 3 nodes which are cordoned off and tainted, to simulate additional capacity that can
be "span up" as load increases and "torn down" as load/node utilization decreases.

## Usage
```
cd k8s_kind
./spin_up.sh
cd ..
cd autoscaler_demo
./spin_up.sh
```

To see the autoscaler work, you can run

```
kubectl logs -f deploy/autoscaler
```

I have included output to show where the pods are scheduled
