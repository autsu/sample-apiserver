#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

openapi-gen -v 2 \
    -h ./hack/boilerplate.go.txt \
    -i k8s.io/sample-apiserver/pkg/apis/transformers/v1alpha1 \
    --input-dirs "k8s.io/apimachinery/pkg/apis/meta/v1" \
    --input-dirs "k8s.io/apimachinery/pkg/runtime" \
    --input-dirs "k8s.io/apimachinery/pkg/version" \
    -O zz_generated_openapi \
    --output-package k8s.io/openapi

mv k8s.io/openapi/zz_generated_openapi.go pkg/generated/openapi/zz_generated.openapi.go
rm -rf k8s.io/openapi
