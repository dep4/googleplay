package googleplay

import (
   "github.com/89z/rosso/crypto"
   "github.com/89z/rosso/http"
   "io"
   "net/url"
   "os"
   "strings"
)

var Client = http.Default_Client

func format_query(vals url.Values) string {
   var buf strings.Builder
   for key := range vals {
      val := vals.Get(key)
      buf.WriteString(key)
      buf.WriteByte('=')
      buf.WriteString(val)
      buf.WriteByte('\n')
   }
   return buf.String()
}

// this beats "io.Reader", and also "bytes.Fields"
func parse_query(query string) url.Values {
   vals := make(url.Values)
   for _, field := range strings.Fields(query) {
      key, val, ok := strings.Cut(field, "=")
      if ok {
         vals.Add(key, val)
      }
   }
   return vals
}

type Auth struct {
   url.Values
}

// You can also use host "android.clients.google.com", but it also uses
// TLS fingerprinting.
func New_Auth(email, password string) (*Auth, error) {
   body := url.Values{
      "Email": {email},
      "Passwd": {password},
      "client_sig": {""},
      "droidguard_results": {"."},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/auth", strings.NewReader(body),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   hello, err := crypto.Parse_JA3(crypto.Android_API_26)
   if err != nil {
      return nil, err
   }
   tr := crypto.Transport(hello)
   res, err := Client.Transport(tr).Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   query, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var auth Auth
   auth.Values = parse_query(string(query))
   return &auth, nil
}

func (a Auth) Create(name string) error {
   file, err := os.Open(name)
   if err != nil {
      return err
   }
   defer file.Close()
   query := format_query(a.Values)
   if _, err := file.WriteString(query); err != nil {
      return err
   }
   return nil
}

func (a *Auth) Exchange() error {
   // these values take from Android API 28
   body := url.Values{
      "Token": {a.Get_Token()},
      "app": {"com.android.vending"},
      "client_sig": {"38918a453d07199354f8b19af05ec6562ced5788"},
      "service": {"oauth2:https://www.googleapis.com/auth/googleplay"},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/auth", strings.NewReader(body),
   )
   if err != nil {
      return err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := Client.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   query, err := io.ReadAll(res.Body)
   if err != nil {
      return err
   }
   a.Values = parse_query(string(query))
   return nil
}

func (a Auth) Get_Auth() string {
   return a.Get("Auth")
}

func (a Auth) Get_Token() string {
   return a.Get("Token")
}
