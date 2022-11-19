package main

import (
   "fmt"
   "github.com/89z/rosso/protobuf"
   "google.golang.org/protobuf/encoding/protowire"
)

func main() {
   mask := protobuf.Message{
      3: protobuf.Bytes{0, 0, 0, 0xC0},
   }.Marshal()
   // []byte{0x1a, 0x4, 0x0, 0x0, 0x0, 0xc0}
   fmt.Printf("%#v\n", mask)
   mask = protowire.AppendFixed32(nil, 0xC0)
   fmt.Printf("%#v\n", mask)
}
