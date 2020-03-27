GO := $(shell which go)
GOPATH ?= $(shell $(GO) env GOPATH)
GOBIN ?= $(GOPATH)/bin
GOSEC ?= $(GOBIN)/gosec
GOLINT ?= $(GOBIN)/golint

all: format lint sec linux macos windows zip

sec:
	$(GOSEC) ./...

lint:
	$(GOLINT) ./...

format:
	$(GO)fmt -w .	

macos:
	GOOS=darwin GOARCH=amd64 $(GO) build -o alertcovid19 main.go

linux:
	GOOS=linux GOARCH=amd64 $(GO) build -o alertcovid19 main.go
	
windows:
	GOOS=windows GOARCH=amd64 $(GO) build -o alertcovid19.exe main.go

zip:
	zip alertcovid19_windows.zip alertcovid19.exe

clean:
	rm -f alertcovid19 alertcovid19.exe alertcovid19_windows.zip