# TODO: Look into way for `sub-makes` to maintain
# variable references to the root make. For now I've
# used `TGOOS`, `TGOARCH`, and `TNPMDIR` in the 
# `platform-unixlike` target to solve this issue
VERSION := $(shell go run scripts/version.go)
# Strip debug info (symbol table and DWARF)
# Also add version info to binary
GO_FLAGS += "-ldflags=-s -w -X 'github.com/ev-the-dev/tmplts/cmd.version=$(VERSION)'"

# Avoid embedding build path in executable
GO_FLAGS += -trimpath

.PHONY version:
version: 
	@echo "Version: $(VERSION)"

.PHONY all:
all: | publish-all clean

.PHONY clean:
clean:
	rm -f npm/README.md npm/LICENSE &&\
		rm -rf npm/bin

	cd npm/@tmplts && \
		rm -rf \
		darwin-arm64/bin \
		darwin-x64/bin \
		linux-x64/bin \
		windows-arm64/bin \
		windows-ia32/bin \
		windows-x64/bin

	cd npm/@tmplts && \
		rm -f\
		darwin-arm64/README.md darwin-arm64/LICENSE \
		darwin-x64/README.md darwin-x64/LICENSE \
		linux-x64/README.md linux-x64/LICENSE \
		windows-arm64/README.md windows-arm64/LICENSE \
		windows-ia32/README.md windows-ia32/LICENSE \
		windows-x64/README.md windows-x64/LICENSE

.PHONY copy-files:
copy-files:
	cp README.md LICENSE npm/
	cp README.md LICENSE npm/@tmplts/darwin-arm64/
	cp README.md LICENSE npm/@tmplts/darwin-x64/
	cp README.md LICENSE npm/@tmplts/linux-x64/
	cp README.md LICENSE npm/@tmplts/windows-arm64/
	cp README.md LICENSE npm/@tmplts/windows-ia32/
	cp README.md LICENSE npm/@tmplts/windows-x64/


  ##############
 # Publishing #
##############

.PHONY: publish-all
publish-all: copy-files
	@echo "Attempting to publish all supported binaries..."
	@npm --version > /dev/null || (echo "The 'npm' command must be in your path to publish" && false)
	@echo "Checking for uncommitted & untracked changes..." && test -z "`git status --porcelain | grep -vE '(README\.md|LICENSE)'`" || \
		(echo "Cannot publish with uncommited/untracked changes:" && \
		git status --porcelain | grep -vE '(README\.md|LICENSE)' && false)
	@echo "Checking for main branch..." && test main = "`git rev-parse --abbrev-ref HEAD`" || \
		(echo "Cannot publish from non-main branch `git rev-parse --abbrev-ref HEAD`" && false)
	@echo "Checking for unpushed commits..." && git fetch
	@test "" = "`git cherry`" || (echo "Cannot publish with unpushed commits" && false)

	@$(MAKE) --no-print-directory -j4 \
		publish-default \
		publish-darwin-arm64 \
		publish-darwin-x64 \
		publish-linux-x64 \
		publish-windows-arm64 \
		publish-windows-ia32 \
		publish-windows-x64

publish-default: platform-default
	cd npm && npm publish

publish-darwin-arm64: platform-darwin-arm64
	cd npm/@tmplts/darwin-arm64 && npm publish

publish-darwin-x64: platform-darwin-x64
	cd npm/@tmplts/darwin-x64 && npm publish

publish-linux-x64: platform-linux-x64
	cd npm/@tmplts/linux-x64 && npm publish

publish-windows-arm64: platform-windows-arm64
	cd npm/@tmplts/windows-arm64 && npm publish

publish-windows-ia32: platform-windows-ia32
	cd npm/@tmplts/windows-ia32 && npm publish

publish-windows-x64: platform-windows-x64
	cd npm/@tmplts/windows-x64 && npm publish


  ##################
 # Platform Build #
##################
.PHONY: platform-all:
platform-all:
	@echo "Attempting to generate all supported platform binaries..."

	@$(MAKE) --no-print-directory -j4 \
		platform-default \
		platform-darwin-arm64 \
		platform-darwin-x64 \
		platform-linux-x64 \
		platform-windows-arm64 \
		platform-windows-ia32 \
		platform-windows-x64

platform-unixlike:
	@test -n "$(TGOOS)" || (echo "The environment variable GOOS must be provided" && false)
	@test -n "$(TGOARCH)" || (echo "The environment variable GOARCH must be provided" && false)
	@test -n "$(TNPMDIR)" || (echo "The environment variable NPMDIR must be provided" && false)
	CGO_ENABLED=0 GOOS="$(TGOOS)" GOARCH="$(TGOARCH)" go build $(GO_FLAGS) -o "$(TNPMDIR)/bin/tmplts"

platform-default:
	@$(MAKE) --no-print-directory TGOOS=darwin TGOARCH=arm64 TNPMDIR=npm platform-unixlike

platform-darwin-arm64:
	@$(MAKE) --no-print-directory TGOOS=darwin TGOARCH=arm64 TNPMDIR=npm/@tmplts/darwin-arm64 platform-unixlike

platform-darwin-x64:
	@$(MAKE) --no-print-directory TGOOS=darwin TGOARCH=amd64 TNPMDIR=npm/@tmplts/darwin-x64 platform-unixlike

platform-linux-x64:
	@$(MAKE) --no-print-directory TGOOS=linux TGOARCH=amd64 TNPMDIR=npm/@tmplts/linux-x64 platform-unixlike

platform-windows-arm64:
	CGO_ENABLED=0 GOOS="windows" GOARCH="arm64" go build $(GO_FLAGS) -o "npm/@tmplts/windows-arm64/bin/tmplts.exe"

platform-windows-ia32:
	CGO_ENABLED=0 GOOS="windows" GOARCH="386" go build $(GO_FLAGS) -o "npm/@tmplts/windows-ia32/bin/tmplts.exe"

platform-windows-x64:
	CGO_ENABLED=0 GOOS="windows" GOARCH="amd64" go build $(GO_FLAGS) -o "npm/@tmplts/windows-x64/bin/tmplts.exe"
