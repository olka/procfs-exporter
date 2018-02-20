TITLE=procfs-exporter
GOPATH=$(CURDIR)
GOBIN=$(CURDIR)/bin

build: get verify
	@echo "Building $(TITLE) to ./bin"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o bin/$(TITLE) 
get:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get
run:
	go run $(TITLE).go
start:
	@./bin/$(TITLE)
clean:
	@rm -rf bin/
	@rm -rf src/git*
	@rm -rf pkg
verify:
	go test -a -v -cover
test: verify
	