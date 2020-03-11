# multiappdemo

Two simple apps, one talks to other

## Kubectl apply setup

This is great for testing istio functionality like kiali graphs:

```
kubectl apply -f https://raw.githubusercontent.com/cbron/multiappdemo/master/backend/deployment.yaml
kubectl apply -f https://raw.githubusercontent.com/cbron/multiappdemo/master/backend/service.yaml
kubectl apply -f https://raw.githubusercontent.com/cbron/multiappdemo/master/frontend/deployment.yaml
kubectl apply -f https://raw.githubusercontent.com/cbron/multiappdemo/master/frontend/service.yaml
```


## Local K3d setup

### build

in each dir: 

```bash
docker build -t cbron/multiappdemo-backend:latest .
docker push cbron/multiappdemo-backend:latest
```

```bash
docker build -t cbron/multiappdemo-frontend:latest .
docker push cbron/multiappdemo-frontend:latest
```

### k3d

```bash
k3d create -n multiappdemo --publish 8080:30080 --image rancher/k3s:v0.9.1
```

### build

In each dir:

```bash
kubectl apply -f service.yaml
kubectl apply -f deployment.yaml
```

### curl from host: 

```bash
curl -s 0.0.0.0:8080
```
