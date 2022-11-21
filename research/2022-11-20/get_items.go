package main

import (
   "bytes"
   "encoding/base64"
   "fmt"
   "github.com/89z/googleplay"
   "github.com/89z/rosso/http"
   "github.com/89z/rosso/protobuf"
   "io"
   "net/url"
   "strconv"
   "time"
)

func main() {
   device, err := checkin()
   if err != nil {
      panic(err)
   }
   if err := sync(device); err != nil {
      panic(err)
   }
   fmt.Println("Sleep", sleep)
   time.Sleep(sleep)
   req := new(http.Request)
   req_body := protobuf.Message{
      2: protobuf.Message{
         1: protobuf.Message{
            1: protobuf.Message{
               1: protobuf.String("com.watchfacestudio.md307digital"),
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
   field := protobuf.Message{
      4: protobuf.Bytes{0xFF, 0xFF, 0xFF, 0xFF},
   }.Marshal()
   mask := base64.StdEncoding.EncodeToString(field)
   req.Header.Set("X-Dfe-Item-Field-Mask", mask)
   fmt.Println(device)
   res, err := client.Do(req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   res_body, err := io.ReadAll(res.Body)
   if err != nil {
      panic(err)
   }
   if bytes.Contains(res_body, []byte("config.xhdpi")) {
      fmt.Println("pass")
   } else {
      fmt.Println("fail")
   }
}

func sync(device string) error {
   req := new(http.Request)
   req.Body = io.NopCloser(bytes.NewReader(body.Marshal()))
   req.Header = make(http.Header)
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/sync"
   req.URL.Scheme = "https"
   req.Header["X-Dfe-Device-Id"] = []string{device}
   if _, err := client.Do(req); err != nil {
      return err
   }
   return nil
}

const sleep = 16 * time.Second

var client = http.Default_Client

func checkin() (string, error) {
   res, err := googleplay.Phone.Checkin("x86")
   if err != nil {
      return "", err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return "", err
   }
   message, err := protobuf.Unmarshal(body)
   if err != nil {
      return "", err
   }
   id, err := googleplay.Device{message}.ID()
   if err != nil {
      return "", err
   }
   return strconv.FormatUint(id, 16), nil
}
