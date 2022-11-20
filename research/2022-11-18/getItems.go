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
   body := protobuf.Message{
      2: protobuf.Message{
         1: protobuf.Message{
            1: protobuf.Message{
               1: protobuf.String("com.google.android.youtube"),
            },
         },
      },
   }.Marshal()
   req.Body = io.NopCloser(bytes.NewReader(body))
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/getItems"
   req.URL.Scheme = "https"
   req.Header["X-Dfe-Device-Id"] = []string{device}
   field := protobuf.Message{
      3: protobuf.Bytes{0xFF, 0xFF, 0xFF, 0xFF},
      // valid range 0xC0 - 0xFFFF
      4: protobuf.Bytes{0xFF, 0xFF},
   }.Marshal()
   mask := base64.StdEncoding.EncodeToString(field)
   req.Header.Set("X-Dfe-Item-Field-Mask", mask)
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if bytes.Contains(res_body, []byte("Google LLC")) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
   if bytes.Contains(res_body, []byte("Video Players & Editors")) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}
