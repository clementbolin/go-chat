######################
## Language: Goalng ##
######################

### COLORS ###
NOC			= \033[0m
BOLD		= \033[1m
UNDERLINE	= \033[4m
BLACK		= \033[1;30m
RED			= \033[1;31m
GREEN		= \033[1;32m
YELLOW		= \033[1;33m
BLUE		= \033[1;34m
VIOLET		= \033[1;35m
CYAN		= \033[1;36m
WHITE		= \033[1;37m

# Go related variables.
GOBASE := $(shell pwd)
GOPATH := $(GOBASE)/vendor:$(GOBASE)
GOBIN := $(GOBASE)/bin
GOFILES := $(wildcard *.go)

# Project Name
PROJECTNAME := $(shell basename "$(PWD)")

# Use linker flags to provide version/build settings
LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

build:
	@echo "$(GREEN)   > Building binary ...\n    $(GOFILES)$(WHITE)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build $(LDFLAGS) -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

clean:
	@echo "$(BLUE)   > Clean Project ...$(WHITE)"
	@rm -rf $(GOBIN)/$(PROJECTNAME)
	@rm -rf bin/
	# @-$(MAKE) go clean

generate:
	@echo "$(RED)   > Generating dependency files ...$(WHITE)"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get $(generate)

setup:
	@echo "$(RED)    > Setup go.mod ...$(WHITE)"
	@go mod init pkg

clean-cache-pkg:
	@echo "$(RED)    > Clean $(GOPATH)/pkg/mod ...$(WHITE)"
	@go clean --modcache

run:
	@echo "$(GREEN)    > Run Project ...$(WHITE)"
	@./bin/$(PROJECTNAME)