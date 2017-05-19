package view

import (
	"os"
	"testing"

	"github.com/overcyn/mochi/layout"
	"github.com/overcyn/mochi/view/encoding"
	"zombiezen.com/go/capnproto2"
)

func TestMarshal(t *testing.T) {
	msg, s, err := capnp.NewMessage(capnp.SingleSegment(nil))
	if err != nil {
		t.Error(err)
	}

	root, err := encoding.NewRootRoot(s)
	if err != nil {
		t.Error(err)
	}

	node, err := (&node{layoutGuide: &layout.Guide{}}).MarshalCapnp(s)
	if err != nil {
		t.Error(err)
	}
	if err = root.SetNode(node); err != nil {
		t.Error(err)
	}

	err = capnp.NewEncoder(os.Stdout).Encode(msg)
	if err != nil {
		t.Error(err)
	}
}
