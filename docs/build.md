# Build

## Building Matcha


    cd $GOPATH/src/github.com/overcyn/matcha/examples/basic
    gomobile bind -target=ios github.com/overcyn/matcha/app
    goimports -w .; and go generate ./...; and go test ./...; and go vet ./...