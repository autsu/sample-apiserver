.PHONY: codegen deepcopy-codegen conversion-codegen openapi-codegen

deepcopy-codegen:
	./hack/deepcopy-gen.sh

conversion-codegen:
	./hack/conversion-gen.sh

openapi-codegen:
	./hack/openapi-gen.sh

codegen: deepcopy-codegen conversion-codegen openapi-codegen

init:
	openssl req -nodes -new -x509 -keyout ca.key -out ca.crt
	openssl req -out client.csr -new -newkey rsa:4096 -nodes -keyout client.key -subj "/CN=development/O=system:masters"
	openssl x509 -req -days 365 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 01 -sha256 -out client.crt
	openssl pkcs12 -export -in ./client.crt -inkey ./client.key -out client.p12 -passout pass:password

test:
	curl -fv -k --cert-type P12 --cert client.p12:password \
   https://localhost:8443/apis/wardle.example.com/v1alpha1/namespaces/default/flunders

test-mac:
	./hack/test.sh
	
run:
	go run main.go --secure-port 8443 \
	--etcd-servers localhost:2379 \
	--v=7 \
	--client-ca-file ca.crt \
	--kubeconfig ~/.kube/config \
	--authentication-kubeconfig ~/.kube/config \
	--authorization-kubeconfig ~/.kube/config


