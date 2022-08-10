```
kubectl apply -k ./ -n [namespace]

# For testing #
kubectl exec -it pod/[pod name] -n [namespace] -- psql -h localhost -U [username] --password -p 5432 [db name]
```
