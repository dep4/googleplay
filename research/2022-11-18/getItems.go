package main

import (
   "bytes"
   "encoding/base64"
   "fmt"
   "github.com/89z/rosso/protobuf"
   "io"
   "net/http"
   "net/url"
)

func main() {
   var req http.Request
   req_body := protobuf.Message{
      2:protobuf.Message{
         1:protobuf.Message{
            1: protobuf.Message{
               1: protobuf.String("com.google.android.youtube"),
            },
         },
      },
   }.Marshal()
   req.Body = io.NopCloser(bytes.NewReader(req_body))
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/getItems"
   req.URL.Scheme = "https"
   req.Header["X-Dfe-Device-Id"] = []string{device}
   req.Header["X-Dfe-Item-Field-Mask"] = []string{field_mask()}
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if bytes.Contains(res_body, []byte("VIDEO_PLAYERS")) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

func field_mask() string {
   mask := protobuf.Message{
      //pass
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf, 0xff, 0x7},
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf, 0xff},
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf, 0xfe},
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf, 0xf0},
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf, 0xd0},
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf, 0xc0},
      //3:protobuf.Bytes{0xe6, 0xff, 0xae, 0xc0},
      3:protobuf.Bytes{0xe6, 0xff, 0xa, 0xc0},
      //fail
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf, 0xb0},
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf, 0x90},
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf, 0x10},
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf, 0xF},
      //3:protobuf.Bytes{0xe6, 0xff, 0xaf},
   }.Marshal()
   return base64.StdEncoding.EncodeToString(mask)
}
