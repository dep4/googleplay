package googleplay

import (
   "bytes"
   "encoding/base64"
   "github.com/89z/rosso/protobuf"
   "net/http"
   "os"
)

func (i Items) Category() (string, error) {
   return i.Get(11).Get(2).Get(2).Get(30).Get_String(1)
}

func (h Header) Get_Items(app string) (*Response, error) {
   body := protobuf.Message{
      2: protobuf.Message{
         1: protobuf.Message{
            1: protobuf.Message{
               1: protobuf.String(app),
            },
         },
      },
   }.Marshal()
   req, err := http.NewRequest(
      "POST", "https://play-fe.googleapis.com/fdfe/getItems",
      bytes.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   h.Set_Device(req.Header)
   field := protobuf.Message{
      // valid range 0xC0 - 0xFFFFFFFF
      3: protobuf.Bytes{0xFF, 0xFF, 0xFF, 0xFF},
      // valid range 0xC0 - 0xFFFF
      4: protobuf.Bytes{0xFF, 0xFF},
   }.Marshal()
   mask := base64.StdEncoding.EncodeToString(field)
   req.Header.Set("X-Dfe-Item-Field-Mask", mask)
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   return &Response{res}, nil
}

func Open_Items(name string) (*Items, error) {
   data, err := os.ReadFile(name)
   if err != nil {
      return nil, err
   }
   message, err := protobuf.Unmarshal(data)
   if err != nil {
      return nil, err
   }
   return &Items{message}, nil
}

type Items struct {
   protobuf.Message
}
