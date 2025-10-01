#!/bin/bash

docker build -t local-autoscaler:latest .
kind load docker-image local-autoscaler:latest --name melissa

# Apply deployment
kubectl apply -f serviceaccount-autoscaler.yaml
sleep 10s
kubectl apply -f deployment-autoscaler.yaml
sleep 10s
kubectl apply -f deployment-cluster-load.yaml

# Check logs
sleep 10s
kubectl logs -f deploy/autoscaler
