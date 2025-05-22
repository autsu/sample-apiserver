#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# !!! 在项目根路径下执行该脚本 !!!

# 不管是内部版本还是外部版本都需要生成 deepcopy
TYPES=(
    "transformers"
    "transformers/v1alpha1"
)

# 生成的代码会放到项目的 k8s.io/sample-apiserver/pkg/apis/${TYPE} 目录下
for TYPE in "${TYPES[@]}"; do
    deepcopy-gen -v 2 \
        -h ./hack/boilerplate.go.txt \
        --bounding-dirs . \
        -i k8s.io/sample-apiserver/pkg/apis/${TYPE} \
        -O zz_generated.deepcopy

    # 将生成的代码移动到 pkg/apis/${TYPE} 目录下
    mv k8s.io/sample-apiserver/pkg/apis/${TYPE}/zz_generated.deepcopy.go \
        pkg/apis/${TYPE}
done

# 清空生成的 k8s.io 目录
rm -rf k8s.io
