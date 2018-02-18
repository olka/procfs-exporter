TITLE=goproc
GOPATH=$(CURDIR)
GOBIN=$(CURDIR)/bin
build: get verify
	@echo "Building $(TITLE) to ./bin"
	go build -o bin/$(TITLE)
get:
	go get 
run:
	go run $(TITLE).go
start:
	@./bin/$(TITLE)
clean:
	@rm -f bin/$(TITLE);
	@rm -rf src/git*
	@rm -rf pkg
verify:
	go test -a -v -cover
test: verify
	