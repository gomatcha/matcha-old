package pb

//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. github.com/overcyn/mochi/pb/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. github.com/overcyn/mochi/pb/button/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. github.com/overcyn/mochi/pb/view/tabnav/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. github.com/overcyn/mochi/pb/view/stacknav/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. github.com/overcyn/mochi/pb/view/switchview/*.proto )"

//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=github.com/overcyn/mochi/ios/basic/basic/protobuf github.com/overcyn/mochi/pb/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=github.com/overcyn/mochi/ios/basic/basic/protobuf github.com/overcyn/mochi/pb/button/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=github.com/overcyn/mochi/ios/basic/basic/protobuf github.com/overcyn/mochi/pb/view/tabnav/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=github.com/overcyn/mochi/ios/basic/basic/protobuf github.com/overcyn/mochi/pb/view/stacknav/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=github.com/overcyn/mochi/ios/basic/basic/protobuf github.com/overcyn/mochi/pb/view/switchview/*.proto )"
