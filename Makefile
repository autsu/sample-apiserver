.PHONY: codegen deepcopy-codegen conversion-codegen openapi-codegen

deepcopy-codegen:
	./hack/deepcopy-gen.sh

conversion-codegen:
	./hack/conversion-gen.sh

openapi-codegen:
	./hack/openapi-gen.sh

codegen: deepcopy-codegen conversion-codegen openapi-codegen

run:
	go run main.go --etcd-servers localhost:2379 --kubeconfig ~/.kube/config --authentication-kubeconfig ~/.kube/config --authorization-kubeconfig ~/.kube/config