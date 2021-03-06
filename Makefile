
.PHONY: update-swagger get-swagger deps validate generate test

SWAGGER_SPEC=./swagger-webhookrelay/swagger.yaml
OPENAPI_PACCKAGE=./pkg/openapi

update-swagger:
	git submodule sync
	git submodule update --remote

get-swagger:
	git submodule sync
	git submodule update --init --recursive

deps:
	GO111MODULE=off go get -u github.com/myitcv/gobin && go mod download

validate: deps
	gobin -mod=readonly -run github.com/go-swagger/go-swagger/cmd/swagger validate $(SWAGGER_SPEC)

generate: validate
	rm -rf $(OPENAPI_PACCKAGE)
	mkdir -p $(OPENAPI_PACCKAGE)
	gobin -mod=readonly -run github.com/go-swagger/go-swagger/cmd/swagger generate client \
		-f $(SWAGGER_SPEC) \
		-t $(OPENAPI_PACCKAGE) \
		-A openapi \
		--default-scheme=https
	go get -u -f $(OPENAPI_PACCKAGE)/...
	gobin -mod=readonly -run golang.org/x/tools/cmd/goimports -w .

test:
	go test ./...
