package main

import (
   "fmt"
   "net/http"
   "net/url"
   "time"
)

type version struct {
   major int
   minor int
   patch int
}

func (v version) String() string {
   b := []byte{'8'}
   b = fmt.Appendf(b, "%02v", v.major)
   b = fmt.Append(b, v.minor)
   b = fmt.Appendf(b, "%02v", v.patch)
   b = append(b, "00"...)
   if len(b) > 8 {
      panic(b)
   }
   return string(b)
}

/*
versionName=6.7.15
versionCode=80671500

versionName=30.4.17
versionCode=83041710
*/
func main() {
   var req http.Request
   req.Header = make(http.Header)
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/delivery"
   val := make(url.Values)
   val["doc"] = []string{"com.android.vending"}
   req.URL.Scheme = "https"
   req.Header["User-Agent"] = []string{"Android-Finsky (sdk=9,versionCode=99999999)"}
   req.Header["X-Dfe-Device-Id"] = []string{device}
   var v version
   for v.major = 6; v.major <= 30; v.major++ {
      for v.minor = 0; v.minor <= 9; v.minor++ {
         for v.patch = 0; v.patch <= 99; v.patch++ {
            val["vc"] = []string{v.String()}
            req.URL.RawQuery = val.Encode()
            res, err := new(http.Transport).RoundTrip(&req)
            if err != nil {
               panic(err)
            }
            if err := res.Body.Close(); err != nil {
               panic(err)
            }
            fmt.Println(res.Status, v.String())
            time.Sleep(99 * time.Millisecond)
         }
      }
   }
}
