@0x8d032c18f4c65862;

using Go = import "/go.capnp";
$Go.package("encoding");
$Go.import("github.com/overcyn/mochi/layout/encoding");

struct Point {
    x @0 :Float64;
    y @1 :Float64;
}

struct Rect {
    min @0 :Point;
    max @1 :Point;
}

struct Insets {
    top @0 :Float64;
    left @1 :Float64;
    bottom @2 :Float64;
    right @3 :Float64;
}

struct Guide {
    frame @0 :Rect;
    insets @1 :Insets;
    zIndex @2 :Int64;
}
