# Strip debug info (symbol table and DWARF)
# Also add version info to binary
GO_FLAGS += "-ldflags=-s -w -X 'github.com/ev-the-dev/tmplts/cmd.version=$(VERSION)'"

# Avoid embedding build path in executable
GO_FLAGS += -trimpath

# CLI Version
VERSION = "9000.0.1"

build: version 
	CGO_ENABLED=0 go build $(GO_FLAGS)

version: npm/package.json
	go run scripts/version.go

.PHONY: publish-all
publish-all:
	@npm --version > /dev/null || (echo "The 'npm' command must be in your path to publish" && false)
	@echo "Checking for uncommitted & untracked changes..." && test -z "`git status --porcelain | grep 'M'" || \
		(echo "Cannot publish with uncommited/untracked changes:" && \
		git status --porcelain | grep 'M' && false)
	@echo "Checking for main branch..." && test main = "`git rev-parse --abbrev-ref HEAD`" || \
		(echo "Cannot publish from non-main branch `git rev-parse --abbrev-ref HEAD`" && false)
	@echo "Checking for unpushed commits..." && git fetch
	@test "" = "`git cherry`" || (echo "Cannot publish with unpushed commits" && false)
	
