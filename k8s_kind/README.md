cluster creation

EDIT: Cluster creation has been automated, run ./spin_up.sh

kind create cluster --config k8s.yaml --name melissa


achristofi@HPCLegion01:~/Desktop$ kind create cluster --config k8s.yaml --name melissa
Creating cluster "melissa" ...
 ✓ Ensuring node image (kindest/node:v1.34.0) 🖼 
 ✓ Preparing nodes 📦 📦 📦 📦 📦 📦  
 ✓ Configuring the external load balancer ⚖️ 
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌 
 ✓ Installing StorageClass 💾 
 ✓ Joining more control-plane nodes 🎮 
 ✓ Joining worker nodes 🚜 
Set kubectl context to "kind-melissa"
You can now use your cluster with:

kubectl cluster-info --context kind-melissa

Have a question, bug, or feature request? Let us know! https://kind.sigs.k8s.io/#community 🙂


metrics server

kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

kubectl edit deployment metrics-server -n kube-system



Under spec.containers[0].args, add:

- --kubelet-insecure-tls

It should look something like:

spec:
  containers:
  - args:
    - --cert-dir=/tmp
    - --secure-port=4443
    - --kubelet-preferred-address-types=InternalIP
    - --kubelet-insecure-tls

Save and exit. The metrics-server pod will restart automatically.

Then label your workers

kubectl label nodes melissa-worker role=worker
kubectl label nodes melissa-worker2 role=worker
kubectl label nodes melissa-worker3 role=worker
