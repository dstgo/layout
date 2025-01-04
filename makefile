# basic
APP := example-service
MODULE := github.com/dstgo/layout/cmd/layout
APP_VERSION := $(shell git tag --sort=-version:refname | sed -n 1p)
BUILD_TIME := $(shell date +"%Y%m%d%H%M%S")-$(shell git log -1 --format=%h)

# build
MODE := debug
BIN_DIR := $(shell pwd)/bin
HOST_OS := $(shell go env GOHOSTOS)
HOST_ARCH := $(shell go env GOHOSTARCH)
OS := $(HOST_OS)
ARCH := $(HOST_ARCH)
LD_FLAGS := $(nullstring)

# reduce binary size at release mode
ifeq ($(MODE), release)
	LD_FLAGS += -s -w
endif

# inject meta info
ifneq ($(APP), $(nullstring))
	LD_FLAGS += -X main.AppName=$(APP)
endif
ifneq ($(BUILD_TIME), $(nullstring))
	LD_FLAGS += -X main.BuildTime=$(BUILD_TIME)
endif
ifneq ($(APP_VERSION), $(nullstring))
	LD_FLAGS += -X main.Version=$(APP_VERSION)
endif

# binary extension
EXT = $(nullstring)
ifeq ($(OS), windows)
	EXT = .exe
endif

.PHONY: build
build:
	# go lint
	go vet ./...

	# prepare target environment $(os)/$(arch)
	go env -w GOOS=$(OS)
	go env -w GOARCH=$(ARCH)

	# build go module
	go build -trimpath \
		-ldflags="$(LD_FLAGS)" \
		-o $(BIN_DIR)/$(MODE)/$(APP)-$(OS)-$(ARCH)/$(APP)$(EXT) \
		$(MODULE)

	# resume host environment $(host_os)/$(host_arch)
	go env -w GOOS=$(HOST_OS)
	go env -w GOARCH=$(HOST_ARCH)


API_DIR := ./api
API_GEN_DIR := ./api/gen
API_DOC_DIR := $(API_DIR)/doc/v1
API_PB_FILES := $(shell find $(API_DIR) -name "*.proto")

.PHONY: pb, doc
pb:
	# create dir
	mkdir -p $(API_GEN_DIR)
	# generate proto files
	protoc --proto_path=.\
              --proto_path=./third_party \
              --go_out=.\
			  --go-http_out=.\
			  --go-grpc_out=.\
			  --go-errors_out=.\
			  --validate_out=lang=go:.\
              $(API_PB_FILES)

doc:
	# create doc dir
	mkdir -p $(API_DOC_DIR)
	# generate openapi doc
	protoc --proto_path=./\
				  --proto_path=./third_party \
				  --openapi_out=fq_schema_naming=true,default_response=false:$(API_DOC_DIR)\
				  $(API_PB_FILES)





