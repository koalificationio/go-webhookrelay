gen:
	docker run --rm -v ${PWD}:/local -v ${GOPATH}/pkg/mod:/go/pkg/mod quay.io/goswagger/swagger:v0.20.1 generate client \
		-f /local/swagger-webhookrelay/swagger.yaml \
		-t /local/pkg/openapi \
		-A openapi \
		--default-scheme=https \
	go get golang.org/x/tools/cmd/goimports
	goimports -w .
