# kind-cluster-autoscaler-poc

This is a PoC for a basic autoscaler written in Go, and a Kubernetes cluster running on KinD.

The standard setup is 3 masters and 3 workers to maintain minimum HA.

Since adding nodes manually to an existing cluster is flaky, i've added another 3 nodes which are cordoned off and tainted, to simulate additional capacity that can
be "span up" as load increases and "torn down" as load/node utilization decreases.

Features to add:

[ ] Modify autoscaler to filter out cordoned/tainted nodes and only deploy on schedulable nodes
[ ] Modify autoscaler to dynamically untaint/uncordon nodes to simulate extra capacity coming on
[ ] Modify autoscaler to dynamically taint/cordoned nodes to simulate extra capacity spinning down
[ ] Modify autoscaler to include sanity checks to abide strictly by the minimum requirement for HA, 3 master and 3 worker nodes
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

I have included output to show where the pods are scheduled.

To tear down:

```
cd k8s_kind
./tear_down.sh
```
