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
            2: protobuf.Slice[protobuf.String]{
               "org.apache.http.legacy",
               "android.test.runner",
            },
            3: protobuf.Slice[protobuf.String]{
               "af",
               "af-ZA",
               "am",
               "am-ET",
               "ar",
               "ar-EG",
               "ar-XB",
               "az",
               "az-AZ",
               "be",
               "be-BY",
               "bg",
               "bg-BG",
               "bn",
               "bn-BD",
               "bs-BA",
               "ca",
               "ca-ES",
               "cs",
               "cs-CZ",
               "da",
               "da-DK",
               "de",
               "de-AT",
               "de-CH",
               "de-DE",
               "de-LI",
               "el",
               "el-GR",
               "en",
               "en-AU",
               "en-CA",
               "en-GB",
               "en-IN",
               "en-NZ",
               "en-SG",
               "en-US",
               "en-XA",
               "es",
               "es-ES",
               "es-US",
               "et",
               "et-EE",
               "eu",
               "eu-ES",
               "fa",
               "fa-IR",
               "fi",
               "fi-FI",
               "fil",
               "fil-PH",
               "fr",
               "fr-BE",
               "fr-CA",
               "fr-CH",
               "fr-FR",
               "gl",
               "gl-ES",
               "gu",
               "gu-IN",
               "hi",
               "hi-IN",
               "hr",
               "hr-HR",
               "hu",
               "hu-HU",
               "hy",
               "hy-AM",
               "id",
               "in",
               "in-ID",
               "is",
               "is-IS",
               "it",
               "it-CH",
               "it-IT",
               "iw",
               "iw-IL",
               "ja",
               "ja-JP",
               "ka",
               "ka-GE",
               "kk",
               "kk-KZ",
               "km",
               "km-KH",
               "kn",
               "kn-IN",
               "ko",
               "ko-KR",
               "ky",
               "ky-KG",
               "lo",
               "lo-LA",
               "lt",
               "lt-LT",
               "lv",
               "lv-LV",
               "mk",
               "mk-MK",
               "ml",
               "ml-IN",
               "mn",
               "mn-MN",
               "mr",
               "mr-IN",
               "ms",
               "ms-MY",
               "my",
               "my-MM",
               "nb",
               "nb-NO",
               "ne",
               "ne-NP",
               "nl",
               "nl-BE",
               "nl-NL",
               "pa",
               "pa-IN",
               "pl",
               "pl-PL",
               "pt",
               "pt-BR",
               "pt-PT",
               "ro",
               "ro-RO",
               "ru",
               "ru-RU",
               "si",
               "si-LK",
               "sk",
               "sk-SK",
               "sl",
               "sl-SI",
               "sq",
               "sq-AL",
               "sr",
               "sr-Latn",
               "sr-RS",
               "sv",
               "sv-SE",
               "sw",
               "sw-TZ",
               "ta",
               "ta-IN",
               "te",
               "te-IN",
               "th",
               "th-TH",
               "tr",
               "tr-TR",
               "uk",
               "uk-UA",
               "ur",
               "ur-PK",
               "uz",
               "uz-UZ",
               "vi",
               "vi-VN",
               "zh-CN",
               "zh-HK",
               "zh-TW",
               "zu",
               "zu-ZA",
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
