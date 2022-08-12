# Kubernetes workshop

Deployment manifests for Kubernetes workshop

## Getting started

```
minikube start
minikube addons enable ingress

kubectl create namespace myapp-ns 
kubectl config set-context --current --namespace=myapp-ns

kubectl apply -k ./db -n myapp-n

kubectl apply -f ./backend-configmap.yaml -n myapp-ns
kubectl apply -f ./backend-deployment.yaml -n myapp-ns
kubectl apply -f ./backend-service.yaml -n myapp-ns

or

kubectl apply -k ./backend -n myapp-ns

kubectl apply -f ingress.yaml

minikube ip
```

## KubeView

```
git clone https://github.com/benc-uk/kubeview

cd kubeview/charts/

helm install kubeview kubeview

kubectl port-forward svc/kubeview 8888:80
```

## Test rolling update

```
kubectl patch deployments.apps backend-depl -p "{\"spec\":{\"template\":{\"metadata\":{\"labels\":{\"date\":\"`date +'%s'`\"}}}}}" --record

kubectl rollout status deployment backend-depl

kubectl rollout history deployment backend-depl

kubectl rollout undo deployment backend-depl (--to-revision=1)
```

inspire by

[https://github.com/thetkpark/golang-todos](https://github.com/thetkpark/golang-todos)

[https://github.com/thetkpark/argo-cd-example](https://github.com/thetkpark/argo-cd-example)

[มาลองทำ GitOps ด้วย ArgoCD](https://blog.sethanantp.com/try-gitops-with-argocd/)

[How to Deploy to Kubernetes using Argo CD and GitOps](https://www.digitalocean.com/community/tutorials/how-to-deploy-to-kubernetes-using-argo-cd-and-gitops)
