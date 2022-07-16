package googleplay

import (
   "github.com/89z/rosso/crypto"
   "github.com/89z/rosso/http"
   "github.com/89z/rosso/os"
   "io"
   "net/url"
   "strconv"
   "strings"
   "time"
)

const Sleep = 4 * time.Second

var Client = http.Default_Client

func format_query(vals url.Values) string {
   var b strings.Builder
   for key := range vals {
      val := vals.Get(key)
      b.WriteString(key)
      b.WriteByte('=')
      b.WriteString(val)
      b.WriteByte('\n')
   }
   return b.String()
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
      "droidguard_results": {"!"},
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

func (self Auth) Create(name string) error {
   query := format_query(self.Values)
   return os.WriteFile(name, []byte(query))
}

func (self *Auth) Exchange() error {
   // these values take from Android API 28
   body := url.Values{
      "Token": {self.Get_Token()},
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
   self.Values = parse_query(string(query))
   return nil
}

func (self Auth) Get_Auth() string {
   return self.Get("Auth")
}

func (self Auth) Get_Token() string {
   return self.Get("Token")
}

type Header struct {
   Auth Auth // Authorization
   Device Device // X-Dfe-Device-Id
   Single bool
}

func (h *Header) Open_Auth(name string) error {
   query, err := os.ReadFile(name)
   if err != nil {
      return err
   }
   h.Auth.Values = parse_query(string(query))
   return nil
}

// Purchase app. Only needs to be done once per Google account.
func (h Header) Purchase(app string) error {
   body := make(url.Values)
   body.Set("doc", app)
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/purchase",
      strings.NewReader(body.Encode()),
   )
   if err != nil {
      return err
   }
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   res, err := Client.Do(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}

func (h Header) Set_Agent(head http.Header) {
   var b []byte
   b = append(b, "Android-Finsky (sdk=9,versionCode="...)
   if h.Single {
      b = strconv.AppendInt(b, 8091_9999, 10)
   } else {
      b = strconv.AppendInt(b, 9999_9999, 10)
   }
   b = append(b, ')')
   head.Set("User-Agent", string(b))
}

func (h Header) Set_Auth(head http.Header) {
   head.Set("Authorization", "Bearer " + h.Auth.Get_Auth())
}

func (h Header) Set_Device(head http.Header) error {
   id, err := h.Device.ID()
   if err != nil {
      return err
   }
   head.Set("X-DFE-Device-ID", strconv.FormatUint(id, 16))
   return nil
}
