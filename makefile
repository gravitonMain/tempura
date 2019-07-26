# Go parameters
GO111MODULE=on
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
# change yourself
BINARY_NAME=tempura
# BINARY_UNIX=$(BINARY_NAME)_unix

all: test build
build: init vendor
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
#	rm -f $(BINARY_UNIX)
allclean: clean
	$(GOCLEAN) -modcache
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
init:
ifeq ("$(wildcard ./go.mod)","")
	$(GOMOD) init 
endif
vendor: init
ifeq ("$(wildcard ./vendor/.)","")
	$(GOMOD) vendor
endif
deps:
	$(GOMOD) tidy
# deps:
#         $(GOGET) github.com/markbates/goth
#         $(GOGET) github.com/markbates/pop


# Cross compilation
# build-linux:
#	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
# docker-build:
#	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v    
