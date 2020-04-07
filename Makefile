PROJECT = $(notdir $(shell pwd))
ARTIFACTS = bin
THIS_FILE = $(lastword $(MAKEFILE_LIST))
VERSION = $(shell git rev-parse --short HEAD)
COMPANY_IMG_TAG_PREFIX = indiependente
SSH_PRV_KEY = `cat ~/.ssh/id_rsa`
SSH_PUB_KEY = `cat ~/.ssh/id_rsa.pub`
SOURCE_ROOT = cmd/$(PROJECT)/main.go
ifneq ($(OS),Windows_NT)
    OS := $(shell sh -c 'uname -s 2>/dev/null')
endif

ifeq ($(OS),Linux)
    LD_FLAGS = -ldflags="-s -w"
endif

.PHONY: all
all: test fmt lint build

.PHONY: test
test:
	go test -cover -race ./...

.PHONY: build
build:
	$(MAKE) clean
	CGO_ENABLED=0 go build $(LD_FLAGS) -o $(ARTIFACTS)/$(PROJECT) $(SOURCE_ROOT)

.PHONY: clean
clean:
	rm -f $(ARTIFACTS)/*

.PHONY: deps-init
deps-init:
	rm -f go.mod go.sum
	go mod init
	go mod tidy

.PHONY: deps
deps:
	go mod download

.PHONY: docker-build
docker-build:
	@docker build -t $(COMPANY_IMG_TAG_PREFIX)/$(PROJECT) -f Dockerfile .

.PHONY: docker-tag
docker-tag:
	@docker tag $(COMPANY_IMG_TAG_PREFIX)/$(PROJECT) $(COMPANY_IMG_TAG_PREFIX)/$(PROJECT):$(VERSION)

.PHONY: docker-clean
docker-clean:
	$(MAKE) -f $(THIS_FILE) docker-rmi $$(docker images | grep $(PROJECT) | awk '{print $$3}') --force

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	@command -v golangci-lint || (cd /usr/local ; wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest)
	@golangci-lint --version
	golangci-lint run --disable-all \
		--deadline=10m \
		--skip-dirs vendor \
		--skip-files \.*_mock\.*\.go \
		-E errcheck \
		-E govet \
		-E unused \
		-E gocyclo \
		-E golint \
		-E varcheck \
		-E structcheck \
		-E maligned \
		-E ineffassign \
		-E interfacer \
		-E unconvert \
		-E goconst \
		-E gosimple \
		-E staticcheck \
		-E gosec

.PHONY: update-deps
update-deps:
	go mod tidy
