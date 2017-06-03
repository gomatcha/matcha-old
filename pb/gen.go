package pb

//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. github.com/overcyn/mochi/pb/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. github.com/overcyn/mochi/pb/button/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --go_out=. github.com/overcyn/mochi/pb/view/tabnavigator/*.proto )"

//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=github.com/overcyn/mochi/ios/basic/basic/protobuf github.com/overcyn/mochi/pb/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=github.com/overcyn/mochi/ios/basic/basic/protobuf github.com/overcyn/mochi/pb/button/*.proto )"
//go:generate bash -c "( cd $GOPATH/src && protoc --objc_out=github.com/overcyn/mochi/ios/basic/basic/protobuf github.com/overcyn/mochi/pb/view/tabnavigator/*.proto )"
