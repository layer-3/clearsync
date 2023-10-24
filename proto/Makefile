# Env
OS := $(shell uname -s | tr A-Z a-z)
UNAME := $(shell uname)
ifeq ($(UNAME),Darwin)
	PROTO_OS := osx-x86_64
else
ifeq ($(UNAME),Linux)
	PROTO_OS := linux-x86_64
else
	PROTO_OS := win32
endif
endif

GOPATH := $(shell go env GOPATH)

# Tools
PROTO_LINT := $(GOPATH)/bin/protolint
PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go
PROTOC_GEN_GO_GRPC := $(GOPATH)/bin/protoc-gen-go-grpc
PROTOC_GEN_DOC := $(GOPATH)/bin/protoc-gen-doc
PROTOC := ../tmp/bin/protoc
PB_REL_URL=https://github.com/protocolbuffers/protobuf/releases
PB_VER=23.4
TOOLS_PATH=../tmp

# Vars
PROTO_DIR := .
PROTO_SRC := $(sort $(wildcard $(PROTO_DIR)/*.proto))
PROTO_GO_DIR := .
DOCS_DIR := ../docs/grpc

default: proto

# Targets
download: $(PROTOC)

$(PROTO_LINT):
	go install github.com/yoheimuta/protolint/cmd/protolint@v0.35

$(PROTOC_GEN_GO):
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26

$(PROTOC_GEN_GO_GRPC):
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

$(PROTOC_GEN_DOC):
	go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5.1

$(PROTOC):
	mkdir -p $(TOOLS_PATH)
	curl -L $(PB_REL_URL)/download/v$(PB_VER)/protoc-$(PB_VER)-$(PROTO_OS).zip -o $(TOOLS_PATH)/protoc.zip
	unzip $(TOOLS_PATH)/protoc.zip -d $(TOOLS_PATH)
	chmod +x $(PROTOC)
	rm $(TOOLS_PATH)/protoc.zip

lint: $(PROTO_SRC) | $(PROTO_LINT)
	$(PROTO_LINT) lint $(PROTO_SRC)

docs: $(PROTO_SRC) | $(PROTOC) $(PROTOC_GEN_DOC)
	mkdir -p $(DOCS_DIR)
	for i in $(PROTO_SRC) ; do \
	  echo Processing $$i ; \
	  $(PROTOC) -I=$(PROTO_DIR) --doc_out=$(DOCS_DIR) --doc_opt=$(DOCS_DIR)/grpc-md.tmpl,$$i.md $$i; \
	done

$(PROTO_GO_DIR)/%.pb.go: %.proto | $(PROTOC) $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_GRPC)
	mkdir -p $(PROTO_GO_DIR)
	$(PROTOC) -I=$(PROTO_DIR) --go_out=":$(PROTO_GO_DIR)" --go_opt=paths=source_relative --go-grpc_out=$(PROTO_GO_DIR) --go-grpc_opt=paths=source_relative $(PROTO_SRC)

proto: lint docs $(PROTO_SRC:%.proto=$(PROTO_GO_DIR)/%.pb.go)

clean:
	rm -fr $(PROTO_GO_DIR)/*.pb.go $(PROTO_JS_DIR)/*_pb.js $(PROTO_JS_DIR)/*_pb.d.ts
