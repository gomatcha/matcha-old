# Build

## Building Mochi


    cd $GOPATH/src/github.com/overcyn/mochi/examples/basic
    gomobile bind -target=ios github.com/overcyn/mochi/app
    goimports -w .; and go generate ./...; and go test ./...; and go vet ./...