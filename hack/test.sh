#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail 

if ! command -v http &>/dev/null; then
    echo "httpie not found, installing via Homebrew..."
    brew install httpie
else
    echo "httpie is already installed."
fi

# GET
echo "GET autobot 111"
http --verify=no --cert client.crt --cert-key client.key \
    https://localhost:8443/apis/transformers.example.com/v1alpha1/namespaces/default/autobots/111

# List
echo "LIST autobots"
http --verify=no --cert client.crt --cert-key client.key \
    https://localhost:8443/apis/transformers.example.com/v1alpha1/namespaces/default/autobots

# Create
echo "CREATE autobot"
http --verify=no --cert client.crt --cert-key client.key \
    https://localhost:8443/apis/transformers.example.com/v1alpha1/namespaces/default/autobots \
    Content-Type:application/json \
    apiVersion:transformers.example.com/v1alpha1 \
    kind:Autobot \
    metadata:='{
        "name": "autobot-1"
    }' \
    spec:='{
        "mode": "car",
        "strength": 100
    }'

# Update
echo "UPDATE autobot"
http --verify=no --cert client.crt --cert-key client.key \
    PUT \
    https://localhost:8443/apis/transformers.example.com/v1alpha1/namespaces/default/autobots/autobot-1 \
    Content-Type:application/json \
    apiVersion:transformers.example.com/v1alpha1 \
    kind:Autobot \
    metadata:='{
        "name": "autobot-1",
        "namespace": "default",
        "resourceVersion": "NEEDS_ACTUAL_RESOURCE_VERSION" 
    }' \
    spec:='{
        "mode": "car_updated",
        "strength": 110
    }'

# Delete
echo "DELETE autobot"
http --verify=no --cert client.crt --cert-key client.key \
    DELETE \
    https://localhost:8443/apis/transformers.example.com/v1alpha1/namespaces/default/autobots/autobot-1

# DeleteCollection
echo "DELETE Collection autobots"
http --verify=no --cert client.crt --cert-key client.key \
    DELETE \
    https://localhost:8443/apis/transformers.example.com/v1alpha1/namespaces/default/autobots?labelSelector=env%3Dtest

# Connect
# 用于子资源
echo "CONNECT autobot"
http --verify=no --cert client.crt --cert-key client.key \
    GET \
    https://localhost:8443/apis/transformers.example.com/v1alpha1/namespaces/default/autobots/autobot-1/transform

# Watch
echo "WATCH autobots"
http --stream --verify=no --cert client.crt --cert-key client.key \
    GET \
    https://localhost:8443/apis/transformers.example.com/v1alpha1/namespaces/default/autobots?watch=true
