GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')

default: clean dependencies test build

workdir:
	mkdir -p workdir

build: workdir/serverless-ui

workdir/serverless-ui: $(GOFILES)
	go build -o workdir/serverless-ui ./serverless-ui

dependencies: 
	@go get github.com/tools/godep
	@cd serverless-ui && dep ensure
bindata:
	./go-bindata -o assets/bindata.go ias/...
test: test-all

test-all:
	@go test -v -cover ./serverless-ui/...

test-min:
	@go test ./...

clean:
	rm -rf serverless-ui/vendor
release:
	# aws s3 cp workdir/amazonian s3://amazonian.package.release/latest/amazonian
	# aws s3 cp workdir/amazonian s3://amazonian.package.release/$(VERSION)/amazonian
