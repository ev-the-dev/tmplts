VERSION := $(shell go run scripts/version.go)
# Strip debug info (symbol table and DWARF)
# Also add version info to binary
GO_FLAGS += "-ldflags=-s -w -X 'github.com/ev-the-dev/tmplts/cmd.version=$(VERSION)'"

# Avoid embedding build path in executable
GO_FLAGS += -trimpath

.PHONY version:
version: 
	@echo "Version: $(VERSION)"

# .PHONY clean:
# clean:
# 	rm -r npm/@tmplts/darwin-arm64/bin


  ##############
 # Publishing #
##############
.PHONY: publish-all
publish-all: version
	@npm --version > /dev/null || (echo "The 'npm' command must be in your path to publish" && false)
	@echo "Checking for uncommitted & untracked changes..." && test -z "`git status --porcelain | grep 'M'" || \
		(echo "Cannot publish with uncommited/untracked changes:" && \
		git status --porcelain | grep 'M' && false)
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

platform-linux-x64: version
	@$(MAKE) --no-print-directory GOOS=linux GOARCH=amd64 NPMDIR=npm/@tmplts/linux-x64 platform-unixlike
