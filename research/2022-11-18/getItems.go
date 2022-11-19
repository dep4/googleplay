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

func field_mask() string {
   mask := protobuf.Message{
      // min
      3: protobuf.Bytes{0, 0, 0, 0xc0},
   }.Marshal()
   return base64.StdEncoding.EncodeToString(mask)
}

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
