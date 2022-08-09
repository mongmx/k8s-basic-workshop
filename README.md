# Kubernetes workshop

Deployment manifests for Kubernetes workshop

## Getting started

```
minikube start
minikube addons enable ingress

kubectl create namespace myapp-ns 
kubectl apply -f backend-configmap.yaml -n myapp-ns
kubectl apply -f backend-deployment.yaml -n myapp-ns
kubectl apply -f backend-service.yaml -n myapp-ns
kubectl apply -f ingress.yaml

minikube ip
```
