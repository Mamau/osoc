-include .env
export

GIN_MODE=test
BUILD_TIME?=$(shell TZ=${TZ} date '+%Y-%m-%d %H:%M:%S')

LDFLAGS=$(shell echo \
	"-X 'osoc/pkg/application.buildVersionTime=${BUILD_TIME}'" \
)

DEFAULT_GOAL := help
.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z0-9_-]+:.*?##/ { printf "  \033[36m%-27s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: init-utils
init-utils: ## Init utils for work space
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: init
init: ## Initialize basic proto lib
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
	github.com/google/gnostic/cmd/protoc-gen-openapi \
	google.golang.org/protobuf/cmd/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc \
	github.com/envoyproxy/protoc-gen-validate \
	github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/headers

.PHONY: proto-gen
proto-gen: ## Generate codes from proto files
	protoc --proto_path=./proto/ \
	--twirp_out=. \
	--proto_path=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.7 \
	--openapi_out=./ \
	--openapi_opt=default_response=false \
	--go_out=paths=source_relative:./api/v1 \
	--validate_out lang=go:. \
	./proto/*.proto

.PHONY: test
test: ## Test project.
	go test -v ./...

.PHONY: test-short
test-short: ## Test project.
	go test -short -v ./...

.PHONY: test-race
test-race: ## Test project with race detection.
	go test -v -race ./...

.PHONY: test-cover
test-cover: ## Test with coverage and open result in a browser
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

.PHONY: fmt
fmt: ## Format golang files with goimports
	find . -name \*.go -not -path \*/wire_gen.go -exec goimports -w {} \;

.PHONY: wire
wire: ## Actualize dependency-injection
	wire ./...

.PHONY: mod
mod: ## Remove unused modules
	go mod tidy -v

.PHONY: finalcheck
finalcheck: wire fmt mod lint test-short swagger-gen ## Make a final complex check before the commit

.PHONY: run
run: ## Run project for local
	go run -ldflags="${LDFLAGS}" ./cmd/${APP_NAME}/.

.PHONY: start
start: ## Init project
	docker-compose up -d
	make migrate

.PHONY: report
report: ## Generate report by pprof. usage: make report type=heap|profile|block|mutex|trace
	curl -s http://${APP_HOST}:${APP_PORT}/debug/pprof/$(type) > ./$(type).out
ifeq ($(type),trace)
	go tool trace -http=:8080 ./$(type).out
else
	go tool pprof -http=:8080 ./$(type).out
endif

.PHONY: watch
watch: ## Run in live-reload mode
	make stop
	docker-compose up

.PHONY: stop
stop: ## Remove containers but keep volumes
	docker-compose down --remove-orphans

.PHONY: clear
clear: ### Remove containers and volumes
	docker-compose down --remove-orphans --volumes

.PHONY: rebuild
rebuild: ## Rebuild by Docker Compose
	make stop
	docker-compose build --no-cache

# make goose cmd="create database sql"
# make goos cmd="up"
# make goos cmd="down"
# Pay attention, --network parameter must be the same as network in docker-compose.yml file.
.PHONY: goose
goose: ## Work with migration
	docker run -ti -u $(shell id -u) --workdir=/home --network=${APP_NAME}_default -v $(shell pwd):/home jerray/goose goose -dir=migrations mysql "$(MY_USER):$(MY_PASSWORD)@($(MY_HOST):$(MY_PORT))/$(MY_DB_NAME)?parseTime=$(MY_PARSE_TIME)" $(cmd)

.PHONY: migrate
migrate: ## Migration up
	make goose cmd="up"

.PHONY: compile
compile: ## Make binary and docs
	go build -ldflags="${LDFLAGS}" -o bin/${APP_NAME} cmd/${APP_NAME}/main.go cmd/${APP_NAME}/wire_gen.go
