GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
APP_NAME=serverlessui
SRC_LOCATION=serverless-ui
BIN_OUTPUT=release
IAAS_PATH=iaas/cloudformation/
IAAS_FILE=iaas.go
IAAS_LOCATION=iaas
MAJOR_VERSION=0
MINOR_VERSION=0

default: clean dependencies  test build

build: serverless-ui

bin-data:
	cd $(SRC_LOCATION) && ./go-bindata -prefix $(IAAS_PATH) -pkg $(IAAS_LOCATION) -o $(IAAS_LOCATION)/$(IAAS_FILE) $(IAAS_PATH)

serverless-ui: $(GOFILES)	
	./build.sh $(APP_NAME) $(SRC_LOCATION) $(BIN_OUTPUT)

dependencies: 
	@go get github.com/tools/godep
	@cd serverless-ui && dep ensure

test: test-all

test-all:
	@go test -v -cover ./$(SRC_LOCATION)/...

test-min:
	@go test ./...

clean:
	rm -rf $(SRC_LOCATION)/vendor
	rm -rf $(BIN_OUTPUT)/$(APP_NAME)*
	# rm -rf $(SRC_LOCATION)/$(IAAS_LOCATION)/$(IAAS_FILE)

publish-release:
	@go get github.com/aktau/github-release
	cd release && ./release.sh "v$(MAJOR_VERSION).$(MINOR_VERSION).$(VERSION)" $(APP_NAME)
