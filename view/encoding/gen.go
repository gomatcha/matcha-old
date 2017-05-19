package encoding

//go:generate capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std -I$GOPATH/src/github.com/overcyn/mochi -ogo view.capnp
//go:generate capnp compile -I$GOPATH/src/zombiezen.com/go/capnproto2/std -I$GOPATH/src/github.com/overcyn/mochi -oc++ view.capnp
