# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)


serve: js
	gopherjs serve --http ":8000"

js:
	cd src; gopherjs build toggle.go -o toggle.js

fmt:
	@go fmt src/*.go

install:
	go get -u github.com/gopherjs/gopherjs

