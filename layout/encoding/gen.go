package encoding

///go:generate capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std -ogo layout.capnp
///go:generate capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std -oc++ layout.capnp
//go:generate protoc --gofast_out=$GOPATH/src --proto_path=$GOPATH/src $GOPATH/src/github.com/overcyn/mochi/layout/encoding/layout.proto
