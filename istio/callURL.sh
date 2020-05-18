#!/bin/bash

function externalIPs {
  echo "external IPs:"
  kubectl get nodes -o jsonpath='{ $.items[*].status.addresses[?(@.type=="ExternalIP")].address }'; echo
}

function internalIPs {
  echo "internal IPs:"
  kubectl get nodes -o jsonpath='{ $.items[*].status.addresses[?(@.type=="InternalIP")].address }'; echo
}

if [ $# -lt 1 ]; then
  echo "You must pass a publicly available ip, try..."
  externalIPs;
  internalIPs;
  exit 2
fi

printf "Using IP: $1\n\n"

echo "Checking ingress-gateway svc..."
kubectl get svc istio-ingressgateway -n istio-system

export INGRESS_PORT=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.spec.ports[?(@.name=="http2")].nodePort}')
export GATEWAY_URL=$1:$INGRESS_PORT

printf "\nUse the following to manually call:\n"
printf "curl -s http://${GATEWAY_URL}/"

printf "\n\nCalling 20x with 2s sleep\n\n"

for run in {1..20}; do curl -s http://${GATEWAY_URL}/; echo; sleep 2; done
