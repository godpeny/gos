SHELL := /bin/zsh
REGISTRY ?= idock.daumkakao.io/kargo
VERSION ?= latest

.PHONY: run-server-dev
run-server-dev:
	@echo "Running go generate ./..."
	@go generate ./...
	@echo "Running Server ..."
	@go run cmd/server/main.go

.PHONY: run-client-dev
run-client-dev:
	@go run cmd/client/main.go

.PHONY: run-reverse-proxy-server
run-reverse-proxy-server:
	@echo "Running Reverse Proxy Server ..."
	@go run cmd/reverse-proxy-server/main.go

.PHONY: stub
protoc-ts:
	@protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       delivery/grpc/delivery.proto

.PHONY: reverse-proxy
reverse-proxy:
	@protoc -I . --grpc-gateway_out .. \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
       delivery/grpc/delivery.proto

.PHONY: codegen
codegen:
	@echo "Running codegen"
	@cd ./server/openapi && go run gen_swagger_schemas.go
	@go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen \
		--package=api \
		--generate spec \
		-o ./server/api/api.gen.go ./server/openapi/public/swagger.yaml
	@echo "./server/api/api.gen.go generated from ./server/openapi/public/swagger.yaml"

.PHONY: run-server
run-server: codegen
	source env-local.sh && go run cmd/kargo-server/main.go

.PHONY: build
build:
	DOCKER_BUILDKIT=1 docker build --build-arg GITHUB_USER_TOKEN -t $(REGISTRY)/kargo-server:$(VERSION) -f dockerfiles/Dockerfile .

# generate mock files for unit tests
.PHONY: mockgen
mockgen:
	go generate ./...

# Download OPA cli from here: https://www.openpolicyagent.org/docs/latest/#running-opa
.PHONY: opa-test
opa-test:
	@echo "Start Open Policy Agent unit test"
	opa test deploy -v

.PHONY: test
test: opa-test
	go test --cover -short $$(go list ./... | grep -v mock)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: release
release:
	@hack/release.sh
