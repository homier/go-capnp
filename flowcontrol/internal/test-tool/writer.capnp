@0xaca73f831c7ebfdd;

using Go = import "/go.capnp";
$Go.package("main");
$Go.import("github.com/homier/go-capnp/v3/flowcontrol/internal/test-tool");

interface Writer {
  write @0 (data :Data);
}
