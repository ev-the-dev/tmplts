VERSION := $(shell go run scripts/version.go)
# Strip debug info (symbol table and DWARF)
# Also add version info to binary
GO_FLAGS += "-ldflags=-s -w -X 'github.com/ev-the-dev/tmplts/cmd.version=$(VERSION)'"

# Avoid embedding build path in executable
GO_FLAGS += -trimpath

.PHONY version:
version: 
	@echo "Version: $(VERSION)"

.PHONY clean:
clean:
	rm -f npm/README.md npm/LICENSE &&\
		rm -rf npm/bin

	cd npm/@tmplts && \
		rm -rf \
		darwin-arm64/bin \
		darwin-x64/bin \
		linux-x64/bin

	cd npm/@tmplts && \
		rm -f\
		darwin-arm64/README.md darwin-arm64/LICENSE \
		darwin-x64/README.md darwin-x64/LICENSE \
		linux-x64/README.md linux-x64/LICENSE

.PHONY copy-files:
copy-files:
	cp README.md LICENSE npm/
	cp README.md LICENSE npm/@tmplts/darwin-arm64/
	cp README.md LICENSE npm/@tmplts/darwin-x64/
	cp README.md LICENSE npm/@tmplts/linux-x64/

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
		publish-linux-x64

publish-default: platform-default
	cd npm && npm publish

publish-darwin-arm64: platform-darwin-arm64
	cd npm/@tmplts/darwin-arm64 && npm publish

publish-darwin-x64: platform-darwin-x64
	cd npm/@tmplts/darwin-x64 && npm publish

publish-linux-x64: platform-linux-x64
	cd npm/@tmplts/linux-x64 && npm publish


  ##################
 # Platform Build #
##################
platform-unixlike:
	@test -n "$(GOOS)" || (echo "The environment variable GOOS must be provided" && false)
	@test -n "$(GOARCH)" || (echo "The environment variable GOARCH must be provided" && false)
	@test -n "$(NPMDIR)" || (echo "The environment variable NPMDIR must be provided" && false)
	CGO_ENABLED=0 GOOS="$(GOOS)" GOARCH="$(GOARCH)" go build $(GO_FLAGS) -o "$(NPMDIR)/bin/tmplts"

platform-default:
	@$(MAKE) --no-print-directory GOOS=darwin GOARCH=arm64 NPMDIR=npm platform-unixlike

platform-darwin-arm64:
	@$(MAKE) --no-print-directory GOOS=darwin GOARCH=arm64 NPMDIR=npm/@tmplts/darwin-arm64 platform-unixlike

platform-darwin-x64:
	@$(MAKE) --no-print-directory GOOS=darwin GOARCH=amd64 NPMDIR=npm/@tmplts/darwin-x64 platform-unixlike

platform-linux-x64:
	@$(MAKE) --no-print-directory GOOS=linux GOARCH=amd64 NPMDIR=npm/@tmplts/linux-x64 platform-unixlike