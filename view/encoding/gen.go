package encoding

/*
#include <stdlib.h>
*/
import "C"

//go:generate protoc --gofast_out=$GOPATH/src --proto_path=$GOPATH/src $GOPATH/src/github.com/overcyn/mochi/view/encoding/view.proto
