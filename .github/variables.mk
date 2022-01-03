# ================================================
# GENERIC VARIABLES
# ================================================

BINDIR      		:= $(CURDIR)/bin
BINNAME     		?= loli
CLIENT_VERSION 		:= $(shell git describe --tags --abbrev=0 )
BUILD_DATE 			:= $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
MAIN            	?= $(CURDIR)/cmd/loli/main.go

# ================================================
# GIT VARIABLES
# ================================================

GIT_BRANCH				:= $(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT				:= $(shell git rev-parse HEAD)
GIT_SHORT_COMMIT		:= $(shell git rev-parse --short HEAD)
GIT_TAG					:= $(shell if [ -z "`git status --porcelain`" ]; then git describe --exact-match --tags HEAD 2>/dev/null; fi)
GIT_TREE_STATE			:= $(shell if [ -z "`git status --porcelain`" ]; then echo "clean" ; else echo "dirty"; fi)

# ================================================
# GO VARIABLES
# ================================================

GO_VERSION 	:= $(shell go version)
GOPATH		?= $(shell go env GOPATH)

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif

GO 		:= go
GOOS   	:= $(shell go env GOOS)
GOARCH 	:= $(shell go env GOARCH)

# NOTE: '-race' requires cgo; enable cgo by setting CGO_ENABLED=1
BUILD_FLAG	:= -race
GOBUILD    	:= CGO_ENABLED=1 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build $(BUILD_FLAG)

LDFLAGS := -w -s
LDFLAGS += -X "github.com/lpmatos/loli/internal/version.cliVersion=$(CLIENT_VERSION)"
LDFLAGS += -X "github.com/lpmatos/loli/internal/version.builtDate=$(BUILD_DATE)"
LDFLAGS += -X "github.com/lpmatos/loli/internal/version.builtBy=makefile"
LDFLAGS += -X "github.com/lpmatos/loli/internal/version.commit=$(GIT_COMMIT)"
LDFLAGS += -X "github.com/lpmatos/loli/internal/version.commitShort=$(GIT_SHORT_COMMIT)"
LDFLAGS += -X "github.com/lpmatos/loli/internal/version.commitBranch=$(GIT_BRANCH)"
LDFLAGS += -X "github.com/lpmatos/loli/internal/version.goVersion=$(GO_VERSION)"
