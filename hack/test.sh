#!/bin/bash

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

