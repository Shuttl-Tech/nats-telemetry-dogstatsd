GIT_COMMIT = $(shell git rev-parse --short HEAD)
PKGS = $(shell go list ./... | grep -v /vendor/)
GOFILES = $(shell find . -name '*.go' -and -not -path "./vendor/*")
UNFMT = $(shell gofmt -l ${GOFILES})

CGO_ENABLED = 0

XC_OS ?= linux darwin
XC_ARCH ?= amd64

lint:
	@golangci-lint run --config golangci.yml

.PHONY: lint

test: fmtcheck lint
	@echo "==> Running tests"
	@go test -v -count=1 -timeout=300s ${PKGS}

.PHONY: test

fmt:
	@echo "==> Fixing code with gofmt"
	@gofmt -s -w ${GOFILES}

.PHONY: fmt

fmtcheck:
	@echo "==> Checking code for gofmt compliance"
	@[ -z "${UNFMT}" ] || ( echo "Following files are not gofmt compliant.\n\n${UNFMT}\n\nRun 'make fmt' for reformat code"; exit 1 )

.PHONY: fmtcheck

build: test
	@go build -a -o _build/nats-telemetry-dogstatsd -ldflags "-s -w -extldflags \"-static\" -X 'main.Version=${GIT_COMMIT}'"

.PHONY: build

define xc-target
  $1/$2:
	@printf "%s%20s %s\n" "-->" "${1}/${2}:" "nats-telemetry-dogstatsd"
	@CGO_ENABLED=0 GOOS="${1}" GOARCH="${2}" \
		go build -a -o "_build/nats-telemetry-dogstatsd-${1}_${2}" \
			-ldflags "-s -w -extldflags \"-static\" -X 'main.Version=${GIT_COMMIT}'"
  .PHONY: $1/$2

  $1:: $1/$2
  .PHONY: $1

  xc:: $1/$2
  .PHONY: xc
endef
$(foreach goarch,$(XC_ARCH),$(foreach goos,$(XC_OS),$(eval $(call xc-target,$(goos),$(goarch)))))