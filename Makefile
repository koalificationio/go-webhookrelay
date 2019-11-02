
.PHONY: update-swagger get-swagger deps validate generate test

SWAGGER_SPEC=./swagger-webhookrelay/swagger.yaml
OPENAPI_PACCKAGE=./pkg/openapi

update-swagger:
	git submodule update --remote

get-swagger:
	git submodule update --init --recursive

deps:
	GO111MODULE=off go get -u github.com/myitcv/gobin && go mod download

validate: deps get-swagger
	gobin -mod=readonly -run github.com/go-swagger/go-swagger/cmd/swagger validate $(SWAGGER_SPEC)

generate: validate
	gobin -mod=readonly -run github.com/go-swagger/go-swagger/cmd/swagger generate client \
		-f $(SWAGGER_SPEC) \
		-t $(OPENAPI_PACCKAGE) \
		-A openapi \
		--default-scheme=https
	go get -u -f $(OPENAPI_PACCKAGE)/...
	gobin -mod=readonly -run golang.org/x/tools/cmd/goimports -w .

test:
	go test ./...
