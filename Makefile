SHELL := /bin/bash

GO_ROOT := $(shell go env GOROOT)
GIT_SHA := $(shell git rev-parse HEAD)
GIT_SHA_SHORT := $(shell git rev-parse --short HEAD)
DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
VERSION := $(shell git describe --tags)-$(GIT_SHA_SHORT)
LDFLAGS := -s -w \
	-X 'github.com/stateful/runme/v3/internal/version.BuildDate=$(DATE)' \
	-X 'github.com/stateful/runme/v3/internal/version.BuildVersion=$(subst v,,$(VERSION))' \
	-X 'github.com/stateful/runme/v3/internal/version.Commit=$(GIT_SHA)'

LDTESTFLAGS := -X 'github.com/stateful/runme/v3/internal/version.BuildVersion=$(subst v,,$(VERSION))'

ifeq ($(RUNME_EXT_BASE),)
RUNME_EXT_BASE := "../vscode-runme"
endif

.PHONY: build
build:
	go build -o runme -ldflags="$(LDFLAGS)" main.go

.PHONY: wasm
wasm: WASM_OUTPUT ?= examples/web
wasm:
	cp $(GO_ROOT)/misc/wasm/wasm_exec.js $(WASM_OUTPUT)
	GOOS=js GOARCH=wasm go build -o $(WASM_OUTPUT)/runme.wasm -ldflags="$(LDFLAGS)" ./web

.PHONY: test/execute
test/execute: PKGS ?= "./..."
test/execute: RUN ?= .*
test/execute: RACE ?= false
test/execute: TAGS ?= "" # e.g. TAGS="test_with_docker"
test/execute: build test/prep-git-project
	TZ=UTC go test -ldflags="$(LDTESTFLAGS)" -run="$(RUN)" -tags="$(TAGS)" -timeout=60s -race=$(RACE) -covermode=atomic -coverprofile=cover.out -coverpkg=./... $(PKGS)

.PHONY: test/prep-git-project
test/prep-git-project:
	@cp -r -f internal/project/testdata/git-project/.git.bkp internal/project/testdata/git-project/.git
	@cp -r -f internal/project/testdata/git-project/.gitignore.bkp internal/project/testdata/git-project/.gitignore
	@cp -r -f internal/project/testdata/git-project/nested/.gitignore.bkp internal/project/testdata/git-project/nested/.gitignore

.PHONY: test/clean-git-project
test/clean-git-project:
	@rm -r -f internal/project/testdata/git-project/.git
	@rm -r -f internal/project/testdata/git-project/.gitignore
	@rm -r -f internal/project/testdata/git-project/nested/.gitignore

.PHONY: test
test: test/prep-git-project test/execute test/clean-git-project

.PHONY: test/update-snapshots
test/update-snapshots:
	@TZ=UTC UPDATE_SNAPSHOTS=true go test ./...

.PHONY: test/robustness
test/robustness:
	find "$$GOPATH/pkg/mod/github.com" -name "*.md" | grep -v "\/\." | xargs dirname | uniq | xargs -n1 -I {} ./runme fmt --project {} > /dev/null

.PHONY: coverage/html
test/coverage/html:
	go tool cover -html=cover.out

.PHONY: coverage/func
test/coverage/func:
	go tool cover -func=cover.out

.PHONY: fmt
fmt:
	@gofumpt -w .

.PHONY: lint
lint:
	@revive -config revive.toml -formatter stylish -exclude integration/subject/... ./...

.PHONY: pre-commit
pre-commit: build wasm test lint
	pre-commit run --all-files

.PHONY: install/dev
install/dev:
	go install github.com/mgechev/revive@v1.3.7
	go install github.com/securego/gosec/v2/cmd/gosec@v2.19.0
	go install honnef.co/go/tools/cmd/staticcheck@v0.4.6
	go install mvdan.cc/gofumpt@v0.6.0
	go install github.com/icholy/gomajor@v0.9.5

.PHONY: install/goreleaser
install/goreleaser:
	go install github.com/goreleaser/goreleaser@v1.15.2

.PHONY: proto/generate
proto/generate:
	buf lint
	buf format -w
	buf generate

.PHONY: proto/clean
proto/clean:
	rm -rf internal/gen/proto

.PHONY: proto/dev
proto/dev: build proto/clean proto/generate
	rm -rf $(RUNME_EXT_BASE)/node_modules/@buf/stateful_runme.community_timostamm-protobuf-ts/runme
	cp -vrf internal/gen/proto/ts/runme $(RUNME_EXT_BASE)/node_modules/@buf/stateful_runme.community_timostamm-protobuf-ts

.PHONY: proto/dev/reset
proto/dev/reset:
	rm -rf $(RUNME_EXT_BASE)/node_modules/@buf/stateful_runme.community_timostamm-protobuf-ts
	cd $(RUNME_EXT_BASE) && runme run setup

# Remember to set up buf registry beforehand.
# More: https://docs.buf.build/bsr/authentication
.PHONY: proto/publish
proto/publish:
	@cd ./internal/api && buf push

.PHONY: release
release: install/goreleaser
	@goreleaser check
	@goreleaser release --snapshot --clean

.PHONY: release/publish
release/publish: install/goreleaser
	@goreleaser release

.PHONY: update-gql-schema
update-gql-schema:
	@go run ./cmd/gqltool/main.go > ./client/graphql/schema/introspection_query_result.json
	@cd ./client/graphql/schema && npm run convert

.PHONY: generate
generate:
	go generate ./...

.PHONY: docker
docker:
	CGO_ENABLED=0 make build
	docker build -f Dockerfile.alpine . -t runme:alpine
	docker build -f Dockerfile.ubuntu . -t runme:ubuntu
