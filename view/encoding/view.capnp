@0x9abec47167560884;

using Layout = import "/layout/encoding/layout.capnp";

using Go = import "/go.capnp";
$Go.package("encoding");
$Go.import("github.com/overcyn/mochi/view/encoding");

struct Node {
  id @0 :Int64;
  buildId @1 :Int64;
  layoutId @2 :Int64;
  paintId @3 :Int64;
  children @4 :List(Node);
  layoutGuide @5 :Layout.Guide;
  values @6: Map(Text, AnyPointer);
  
  
# Id           mochi.Id
# BuildId      int64
# LayoutId     int64
# PaintId      int64
# Children     map[mochi.Id]*RenderNode
# BridgeName   string
# BridgeState  interface{}
# Values       map[string][]byte

# LayoutGuide  *layout.Guide
# PaintOptions paint.Style
}

struct Root {
    node @0 :Node;
}

struct Map(Key, Value) {
  entries @0 :List(Entry);
  struct Entry {
    key @0 :Key;
    value @1 :Value;
  }
}