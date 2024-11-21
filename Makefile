MODULE_NAME=github.com/dami-i/fsfs
VERSION=v0.1.2

all:
	go mod tidy && go build -o ./bin/fsfs .

update-go-cache:
	GOPROXY=https://proxy.golang.org GO111MODULE=on go install $(MODULE_NAME)@$(VERSION)

list-versions:
	go list -m --versions $(MODULE_NAME)