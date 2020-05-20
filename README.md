# multiappdemo

Two simple apps, one talks to other. One k8s svc and deployment for each. 

## Kubectl apply setup

This is great for testing service-mesh functionality:

```bash
kubectl apply -f https://raw.githubusercontent.com/cbron/multiappdemo/master/multiappdemo.yaml
```

or if you'd like to do each part separately:

```bash
kubectl apply -f https://raw.githubusercontent.com/cbron/multiappdemo/master/backend/deployment.yaml
kubectl apply -f https://raw.githubusercontent.com/cbron/multiappdemo/master/backend/service.yaml
kubectl apply -f https://raw.githubusercontent.com/cbron/multiappdemo/master/frontend/deployment.yaml
kubectl apply -f https://raw.githubusercontent.com/cbron/multiappdemo/master/frontend/service.yaml
```


## Service mesh specifics

**Istio**

```bash
kubectl apply -f https://raw.githubusercontent.com/cbron/multiappdemo/master/istio/istio.yaml
```

This creates a 2nd backend service and sets up istio resources. Use the bash script to call it;
```bash
./istio/callURL.sh
```

Produces the following:

![Istio.png](istio/istio.png "Istio.png")


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

### deploy

```bash
kubectl apply -f backend/service.yaml
kubectl apply -f backend/deployment.yaml
kubectl apply -f frontend/service.yaml
kubectl apply -f frontend/deployment.yaml
```

### curl from host: 

```bash
curl -s 0.0.0.0:8080
```


## Kustomize

```bash
kubectl kustomize . > multiappdemo.yaml
```
