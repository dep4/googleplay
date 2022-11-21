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

var body = protobuf.Message{
   1: protobuf.Slice[protobuf.Message]{
      protobuf.Message{
         10: protobuf.Message{
            1: protobuf.Slice[protobuf.Message]{
               protobuf.Message{
                  1: protobuf.String("android.hardware.camera.autofocus"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.camera.front"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.camera"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.location.network"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.location.gps"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.location"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.microphone"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.screen.landscape"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.screen.portrait"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.sensor.accelerometer"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.telephony"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.touchscreen"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.touchscreen.multitouch"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.usb.host"),
               },
               protobuf.Message{
                  1: protobuf.String("android.hardware.wifi"),
               },
               protobuf.Message{
                  1: protobuf.String("android.software.device_admin"),
               },
            },
            2: protobuf.Slice[protobuf.Raw]{
               protobuf.Raw{
                  Bytes:  []byte{0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x6c, 0x65, 0x67, 0x61, 0x63, 0x79},
                  String: "org.apache.http.legacy",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72},
                  String: "android.test.runner",
               },
            },
            3: protobuf.Slice[protobuf.Raw]{protobuf.Raw{
               Bytes:  []byte{0x61, 0x66},
               String: "af",
            },
               protobuf.Raw{
                  Bytes:  []byte{0x61, 0x66, 0x2d, 0x5a, 0x41},
                  String: "af-ZA",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x61, 0x6d},
                  String: "am",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x61, 0x6d, 0x2d, 0x45, 0x54},
                  String: "am-ET",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x61, 0x72},
                  String: "ar",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x61, 0x72, 0x2d, 0x45, 0x47},
                  String: "ar-EG",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x61, 0x72, 0x2d, 0x58, 0x42},
                  String: "ar-XB",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x61, 0x7a},
                  String: "az",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x61, 0x7a, 0x2d, 0x41, 0x5a},
                  String: "az-AZ",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x62, 0x65},
                  String: "be",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x62, 0x65, 0x2d, 0x42, 0x59},
                  String: "be-BY",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x62, 0x67},
                  String: "bg",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x62, 0x67, 0x2d, 0x42, 0x47},
                  String: "bg-BG",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x62, 0x6e},
                  String: "bn",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x62, 0x6e, 0x2d, 0x42, 0x44},
                  String: "bn-BD",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x62, 0x73, 0x2d, 0x42, 0x41},
                  String: "bs-BA",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x63, 0x61},
                  String: "ca",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x63, 0x61, 0x2d, 0x45, 0x53},
                  String: "ca-ES",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x63, 0x73},
                  String: "cs",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x63, 0x73, 0x2d, 0x43, 0x5a},
                  String: "cs-CZ",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x64, 0x61},
                  String: "da",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x64, 0x61, 0x2d, 0x44, 0x4b},
                  String: "da-DK",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x64, 0x65},
                  String: "de",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x64, 0x65, 0x2d, 0x41, 0x54},
                  String: "de-AT",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x64, 0x65, 0x2d, 0x43, 0x48},
                  String: "de-CH",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x64, 0x65, 0x2d, 0x44, 0x45},
                  String: "de-DE",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x64, 0x65, 0x2d, 0x4c, 0x49},
                  String: "de-LI",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6c},
                  String: "el",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6c, 0x2d, 0x47, 0x52},
                  String: "el-GR",
                  Message: protobuf.Message{
                     12: protobuf.Fixed32(1380396396),
                  },
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6e},
                  String: "en",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6e, 0x2d, 0x41, 0x55},
                  String: "en-AU",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6e, 0x2d, 0x43, 0x41},
                  String: "en-CA",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6e, 0x2d, 0x47, 0x42},
                  String: "en-GB",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6e, 0x2d, 0x49, 0x4e},
                  String: "en-IN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6e, 0x2d, 0x4e, 0x5a},
                  String: "en-NZ",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6e, 0x2d, 0x53, 0x47},
                  String: "en-SG",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6e, 0x2d, 0x55, 0x53},
                  String: "en-US",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x6e, 0x2d, 0x58, 0x41},
                  String: "en-XA",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x73},
                  String: "es",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x73, 0x2d, 0x45, 0x53},
                  String: "es-ES",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x73, 0x2d, 0x55, 0x53},
                  String: "es-US",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x74},
                  String: "et",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x74, 0x2d, 0x45, 0x45},
                  String: "et-EE",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x75},
                  String: "eu",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x65, 0x75, 0x2d, 0x45, 0x53},
                  String: "eu-ES",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x61},
                  String: "fa",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x61, 0x2d, 0x49, 0x52},
                  String: "fa-IR",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x69},
                  String: "fi",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x69, 0x2d, 0x46, 0x49},
                  String: "fi-FI",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x72},
                  String: "fr",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x72, 0x2d, 0x42, 0x45},
                  String: "fr-BE",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x72, 0x2d, 0x43, 0x41},
                  String: "fr-CA",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x72, 0x2d, 0x43, 0x48},
                  String: "fr-CH",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x72, 0x2d, 0x46, 0x52},
                  String: "fr-FR",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x67, 0x6c},
                  String: "gl",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x67, 0x6c, 0x2d, 0x45, 0x53},
                  String: "gl-ES",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x67, 0x75},
                  String: "gu",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x67, 0x75, 0x2d, 0x49, 0x4e},
                  String: "gu-IN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x68, 0x69},
                  String: "hi",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x68, 0x69, 0x2d, 0x49, 0x4e},
                  String: "hi-IN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x68, 0x72},
                  String: "hr",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x68, 0x72, 0x2d, 0x48, 0x52},
                  String: "hr-HR",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x68, 0x75},
                  String: "hu",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x68, 0x75, 0x2d, 0x48, 0x55},
                  String: "hu-HU",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x68, 0x79},
                  String: "hy",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x68, 0x79, 0x2d, 0x41, 0x4d},
                  String: "hy-AM",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x69, 0x64},
                  String: "id",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x69, 0x6e},
                  String: "in",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x69, 0x6e, 0x2d, 0x49, 0x44},
                  String: "in-ID",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x69, 0x73},
                  String: "is",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x69, 0x73, 0x2d, 0x49, 0x53},
                  String: "is-IS",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x69, 0x74},
                  String: "it",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x69, 0x74, 0x2d, 0x43, 0x48},
                  String: "it-CH",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x69, 0x74, 0x2d, 0x49, 0x54},
                  String: "it-IT",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x69, 0x77},
                  String: "iw",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x69, 0x77, 0x2d, 0x49, 0x4c},
                  String: "iw-IL",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6a, 0x61},
                  String: "ja",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6a, 0x61, 0x2d, 0x4a, 0x50},
                  String: "ja-JP",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x61},
                  String: "ka",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x61, 0x2d, 0x47, 0x45},
                  String: "ka-GE",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x6b},
                  String: "kk",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x6b, 0x2d, 0x4b, 0x5a},
                  String: "kk-KZ",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x6d},
                  String: "km",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x6d, 0x2d, 0x4b, 0x48},
                  String: "km-KH",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x6e},
                  String: "kn",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x6e, 0x2d, 0x49, 0x4e},
                  String: "kn-IN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x6f},
                  String: "ko",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x6f, 0x2d, 0x4b, 0x52},
                  String: "ko-KR",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x79},
                  String: "ky",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6b, 0x79, 0x2d, 0x4b, 0x47},
                  String: "ky-KG",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6c, 0x6f},
                  String: "lo",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6c, 0x6f, 0x2d, 0x4c, 0x41},
                  String: "lo-LA",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6c, 0x74},
                  String: "lt",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6c, 0x74, 0x2d, 0x4c, 0x54},
                  String: "lt-LT",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6c, 0x76},
                  String: "lv",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6c, 0x76, 0x2d, 0x4c, 0x56},
                  String: "lv-LV",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x6b},
                  String: "mk",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x6b, 0x2d, 0x4d, 0x4b},
                  String: "mk-MK",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x6c},
                  String: "ml",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x6c, 0x2d, 0x49, 0x4e},
                  String: "ml-IN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x6e},
                  String: "mn",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x6e, 0x2d, 0x4d, 0x4e},
                  String: "mn-MN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x72},
                  String: "mr",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x72, 0x2d, 0x49, 0x4e},
                  String: "mr-IN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x73},
                  String: "ms",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x73, 0x2d, 0x4d, 0x59},
                  String: "ms-MY",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x79},
                  String: "my",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6d, 0x79, 0x2d, 0x4d, 0x4d},
                  String: "my-MM",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6e, 0x62},
                  String: "nb",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6e, 0x62, 0x2d, 0x4e, 0x4f},
                  String: "nb-NO",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6e, 0x65},
                  String: "ne",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6e, 0x65, 0x2d, 0x4e, 0x50},
                  String: "ne-NP",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6e, 0x6c},
                  String: "nl",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6e, 0x6c, 0x2d, 0x42, 0x45},
                  String: "nl-BE",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x6e, 0x6c, 0x2d, 0x4e, 0x4c},
                  String: "nl-NL",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x70, 0x61},
                  String: "pa",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x70, 0x61, 0x2d, 0x49, 0x4e},
                  String: "pa-IN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x70, 0x6c},
                  String: "pl",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x70, 0x6c, 0x2d, 0x50, 0x4c},
                  String: "pl-PL",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x70, 0x74},
                  String: "pt",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x70, 0x74, 0x2d, 0x42, 0x52},
                  String: "pt-BR",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x70, 0x74, 0x2d, 0x50, 0x54},
                  String: "pt-PT",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x72, 0x6f},
                  String: "ro",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x72, 0x6f, 0x2d, 0x52, 0x4f},
                  String: "ro-RO",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x72, 0x75},
                  String: "ru",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x72, 0x75, 0x2d, 0x52, 0x55},
                  String: "ru-RU",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x69},
                  String: "si",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x69, 0x2d, 0x4c, 0x4b},
                  String: "si-LK",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x6b},
                  String: "sk",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x6b, 0x2d, 0x53, 0x4b},
                  String: "sk-SK",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x6c},
                  String: "sl",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x6c, 0x2d, 0x53, 0x49},
                  String: "sl-SI",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x71},
                  String: "sq",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x71, 0x2d, 0x41, 0x4c},
                  String: "sq-AL",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x72},
                  String: "sr",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x72, 0x2d, 0x4c, 0x61, 0x74, 0x6e},
                  String: "sr-Latn",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x72, 0x2d, 0x52, 0x53},
                  String: "sr-RS",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x76},
                  String: "sv",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x76, 0x2d, 0x53, 0x45},
                  String: "sv-SE",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x77},
                  String: "sw",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x73, 0x77, 0x2d, 0x54, 0x5a},
                  String: "sw-TZ",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x74, 0x61},
                  String: "ta",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x74, 0x61, 0x2d, 0x49, 0x4e},
                  String: "ta-IN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x74, 0x65},
                  String: "te",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x74, 0x65, 0x2d, 0x49, 0x4e},
                  String: "te-IN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x74, 0x68},
                  String: "th",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x74, 0x68, 0x2d, 0x54, 0x48},
                  String: "th-TH",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x69, 0x6c},
                  String: "fil",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x66, 0x69, 0x6c, 0x2d, 0x50, 0x48},
                  String: "fil-PH",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x74, 0x72},
                  String: "tr",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x74, 0x72, 0x2d, 0x54, 0x52},
                  String: "tr-TR",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x75, 0x6b},
                  String: "uk",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x75, 0x6b, 0x2d, 0x55, 0x41},
                  String: "uk-UA",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x75, 0x72},
                  String: "ur",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x75, 0x72, 0x2d, 0x50, 0x4b},
                  String: "ur-PK",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x75, 0x7a},
                  String: "uz",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x75, 0x7a, 0x2d, 0x55, 0x5a},
                  String: "uz-UZ",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x76, 0x69},
                  String: "vi",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x76, 0x69, 0x2d, 0x56, 0x4e},
                  String: "vi-VN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x7a, 0x68, 0x2d, 0x43, 0x4e},
                  String: "zh-CN",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x7a, 0x68, 0x2d, 0x48, 0x4b},
                  String: "zh-HK",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x7a, 0x68, 0x2d, 0x54, 0x57},
                  String: "zh-TW",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x7a, 0x75},
                  String: "zu",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x7a, 0x75, 0x2d, 0x5a, 0x41},
                  String: "zu-ZA",
               },
            },
            4: protobuf.Slice[protobuf.Raw]{
               protobuf.Raw{
                  Bytes:  []byte{0x47, 0x4c, 0x5f, 0x4f, 0x45, 0x53, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x45, 0x54, 0x43, 0x31, 0x5f, 0x52, 0x47, 0x42, 0x38, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x75, 0x72, 0x65},
                  String: "GL_OES_compressed_ETC1_RGB8_texture",
               },
               protobuf.Raw{
                  Bytes:  []byte{0x47, 0x4c, 0x5f, 0x4b, 0x48, 0x52, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x73, 0x74, 0x63, 0x5f, 0x6c, 0x64, 0x72},
                  String: "GL_KHR_texture_compression_astc_ldr",
               },
            },
         },
      },
      protobuf.Message{
         19: protobuf.Message{
            2: protobuf.Varint(83032110),
         },
      },
      protobuf.Message{
         20: protobuf.Message{
            2: protobuf.Varint(768),
            3: protobuf.Varint(1280),
            4: protobuf.Varint(2),
            5: protobuf.Varint(320),
            1: protobuf.Varint(3),
         },
      },
   },
}
