@0x8d032c18f4c65862;

using Go = import "/go.capnp";
$Go.package("encoding");
$Go.import("github.com/overcyn/mochi/touch/encoding");

struct TapRecognizer {
  count @0 :Int64;
}
