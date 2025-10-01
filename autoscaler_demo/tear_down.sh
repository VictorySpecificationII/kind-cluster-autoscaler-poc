#!/bin/bash

kubectl delete -f deployment-cluster-load.yaml
kubectl delete -f deployment-autoscaler.yaml
kubectl delete -f serviceaccount-autoscaler.yaml
