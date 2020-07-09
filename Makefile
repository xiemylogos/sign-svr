GOFMT=gofmt
GC=go build


ARCH=$(shell uname -m)
SRC_FILES = $(shell git ls-files | grep -e .go$ | grep -v _test.go)

sign-svr: $(SRC_FILES)
	$(GC)  $(BUILD_NODE_PAR) -o sign-svr main.go
 


sign-svr-cross: sign-svr-windows sign-svr-linux sign-svr-darwin

sign-svr-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o sign-svr-windows-amd64.exe main.go

sign-svr-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o sign-svr-linux-amd64 main.go

sign-svr-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GC) $(BUILD_NODE_PAR) -o sign-svr-darwin-amd64 main.go

tools-cross: tools-windows tools-linux tools-darwin

format:
	$(GOFMT) -w main.go

clean:
	rm -rf *.8 *.o *.out *.6 *exe
	rm -rf sign-svr sign-svr-*
