# multiappdemo

Two simple apps, one talks to other. One k8s svc and deployment for each. 

This is a simpler version of Istio's bookinfo app.

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


## Helm chart install

```bash
helm install multiappdemo ./chart/
# or
helm install multiappdemo ./multiappdemo-1.0.0.tgz
# or 
helm repo add mad https://raw.githubusercontent.com/cbron/multiappdemo/master
helm3 install multiappdemo mad/multiappdemo
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

Ambient specific setup instructions:
```
istioctl install -f ./istio/istio-config.yaml --skip-confirmation
kubectl get crd gateways.gateway.networking.k8s.io &> /dev/null || kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.3.0-rc.1/standard-install.yaml
istioctl waypoint apply -n default --enroll-namespace

k label service backend-svc istio.io/use-waypoint=waypoint
k label service frontend-svc istio.io/use-waypoint=waypoint
k label gateway multiappdemo-gateway istio.io/use-waypoint=waypoint
```
If using kind you'll need to use the local load balancer, make sure to setup cloud-provider-kind binary and run it

## Local K3d setup

#### build

in each dir: 

```bash
docker build -t cbron/multiappdemo-backend:latest .
docker push cbron/multiappdemo-backend:latest
```

```bash
docker build -t cbron/multiappdemo-frontend:latest .
docker push cbron/multiappdemo-frontend:latest
```

#### k3d

```bash
k3d cluster create multiappdemo -p "8080:30080@agent[0]" --agents 1
export KUBECONFIG="$(k3d kubeconfig write multiappdemo)"
```

#### deploy

```bash
kubectl apply -f multiappdemo.yaml
```

#### curl from host: 

```bash
curl -s 0.0.0.0:8080
```

## Refreshing files with Kustomize

```bash
kubectl kustomize . > multiappdemo.yaml
kubectl kustomize . > chart/templates/multiappdemo.yaml
```
