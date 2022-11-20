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
   var req http.Request
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
   res, err := new(http.Transport).RoundTrip(&req)
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
   req.Header["Accept-Encoding"] = []string{"identity"}
   req.Header["Accept-Language"] = []string{"en-US"}
   req.Header["Connection"] = []string{"Keep-Alive"}
   req.Header["Content-Length"] = []string{"5198"}
   req.Header["Content-Type"] = []string{"application/x-protobuf"}
   req.Header["Host"] = []string{"play-fe.googleapis.com"}
   req.Header["User-Agent"] = []string{"Android-Finsky/30.3.21-21%20%5B0%5D%20%5BPR%5D%20445437866 (api=3,versionCode=83032110,sdk=25,device=vbox86p,hardware=vbox86,product=vbox86p,platformVersionRelease=7.1.1,model=Nexus%204,buildId=NMF26Q,isWideScreen=0,supportedAbis=x86)"}
   req.Header["X-Ad-Id"] = []string{"56a40a11-8d70-42a4-8212-3b35af066f02"}
   req.Header["X-Dfe-Client-Id"] = []string{"am-unknown"}
   req.Header["X-Dfe-Cookie"] = []string{"EAEYACICVVNCFQoFVVMtVFgSDAiL_r-cBhDY8Oi9A0oSCgJVUxIMCIv-v5wGEKDQ670DWAA"}
   req.Header["X-Dfe-Device-Checkin-Consistency-Token"] = []string{"ABFEt1XSZMNapw7-Rr92k25ZprALUTeaZsV7lpxvc1eoq2XCn6G4BGXXWWaXlM1XVYFPIq3gkm4IexR7h2AeLPABIte3SOoolFo3NZBuo6N9o5FEa_ANPC89LbBXZ36Zhk4BspsCC2WA3IcBNg7Ns0Ht-BFeXiTRjN6QmkZC0fAg6yfu3ZGVtF4CSooJR-ELB65HczvfmkQSxR8fW1Uju70hjlOT2Cn42RCsW1WPZXHka141kTM8zxvWw8NVz-RHWvBwgMSV9LU7WNZqHVa-hT8Z1gh_uVS96bsx2OSvgy_COHZAU8oLFImCZmBTYZysNVkmlL5YrNZSDFwdSWgIBMpB-geZvkk4f1mlNeLMp8pdmr7IW4icPrC-E2_xTXQDYH3koabsAITKssc46XZ-LxU4QZLdOFhod4YQoTGozbgtoi1IKyHcFVc"}
   req.Header["X-Dfe-Encoded-Targets"] = []string{"CAE"}
   req.Header["X-Dfe-Mccmnc"] = []string{"310270"}
   req.Header["X-Dfe-Network-Type"] = []string{"4"}
   req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=4000"}
   req.Header["X-Limit-Ad-Tracking-Enabled"] = []string{"false"}
   req.Header["X-Ps-Rh"] = []string{"H4sIAAAAAAAAAOOq46rxMCn2dISCeFPfStMKTyO_It9yv6JgA18Xz3K_KnfDEp_Acl8XR2M3Q5901yxHA98qf2M_R39v95zMREczdy9Xx3iPkqiwoqgiX6-IgPRyiyrvquyg8giPiuSACqPS8mDXgKqC3BTfwMx0C0fPHN3S9KAokH0ABipwe4AAAAA"}
   req.Header["Authorization"] = []string{authorization}
   req.Header["X-Dfe-Device-Id"] = []string{device}
   if _, err := client.Do(req); err != nil {
      return err
   }
   return nil
}
