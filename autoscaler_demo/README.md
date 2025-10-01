created kind cluster
created go app for autoscaler - deploys initial pods for stress then moves pods about as cpu time gets to 50m
created dockerfile
created deployment

deployed

complained about metrics api not being available

installed metrics api

complained about the master nodes being tainted so i modified the deployment to untaint them

deployed again

complained about the standard svcacct not being allowed to do the verbs

created a sa and bound to the deployment

complained it wasn't able to use the nodes and pods verbs

extended the rbac

deployed, works BUT

then i realized i launch the stress pods from the autoscaler itself rather than a separate deployment so now i have to modify it to only launch on the worker nodes

so i have to label the nodes of the kind cluster as workers

did so

now i have to modify the go autoscaler to only deploy on workers

can't see the pods, probably because they die before i get a chance to

tried removing timeout, nothing

new plan - deployment separately for stress pods, modifying autoscaler to reschedule and redeploy

created deployment for stress pods for load


modified autoscaler to observe and reschedule pods, very sensitive so we can observe changes

need to now modify autoscaler to add/remove nodes, min amount is 3M 3W for HA mode

found kindscaler YEA BOI https://github.com/lobuhi/kindscaler
