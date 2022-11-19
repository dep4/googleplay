package googleplay

import (
   "github.com/89z/rosso/crypto"
   "github.com/89z/rosso/http"
   "github.com/89z/rosso/protobuf"
   "io"
   "net/url"
   "os"
   "strconv"
   "strings"
   "time"
)

func (d Device) Create(name string) error {
   data := d.Marshal()
   return os.WriteFile(name, data, os.ModePerm)
}

func (n Native_Platform) String() string {
   b := []byte("nativePlatform")
   for key, val := range n {
      b = append(b, '\n')
      b = strconv.AppendInt(b, key, 10)
      b = append(b, ": "...)
      b = append(b, val...)
   }
   return string(b)
}

func (h *Header) Open_Device(name string) error {
   buf, err := os.ReadFile(name)
   if err != nil {
      return err
   }
   h.Device.Message, err = protobuf.Unmarshal(buf)
   if err != nil {
      return err
   }
   return nil
}

type Native_Platform map[int64]string

var Platforms = Native_Platform{
   // com.google.android.youtube
   0: "x86",
   // com.miui.weather2
   1: "armeabi-v7a",
   // com.kakaogames.twodin
   2: "arm64-v8a",
}

// These can use default values, but they must all be included
type Config struct {
   GL_ES_Version uint64
   GL_Extension []string
   Has_Five_Way_Navigation uint64
   Has_Hard_Keyboard uint64
   Keyboard uint64
   Navigation uint64
   New_System_Available_Feature []string
   Screen_Density uint64
   Screen_Layout uint64
   System_Shared_Library []string
   Touch_Screen uint64
}

func (d Device) ID() (uint64, error) {
   return d.Get_Fixed64(7)
}

type Device struct {
   protobuf.Message
}

var Phone = Config{
   New_System_Available_Feature: []string{
      // app.source.getcontact
      "android.hardware.location.gps",
      // br.com.rodrigokolb.realdrum
      "android.software.midi",
      // com.app.xt
      "android.hardware.camera.front",
      // com.clearchannel.iheartradio.controller
      "android.hardware.microphone",
      // com.google.android.apps.walletnfcrel
      "android.software.device_admin",
      // com.google.android.youtube
      "android.hardware.touchscreen",
      "android.hardware.wifi",
      // com.illumix.fnafar
      "android.hardware.sensor.gyroscope",
      // com.madhead.tos.zh
      "android.hardware.sensor.accelerometer",
      // com.miHoYo.GenshinImpact
      "android.hardware.opengles.aep",
      // com.pinterest
      "android.hardware.camera",
      "android.hardware.location",
      "android.hardware.screen.portrait",
      // com.supercell.brawlstars
      "android.hardware.touchscreen.multitouch",
      // com.sygic.aura
      "android.hardware.location.network",
      // com.xiaomi.smarthome
      "android.hardware.bluetooth",
      "android.hardware.bluetooth_le",
      "android.hardware.camera.autofocus",
      "android.hardware.usb.host",
      // kr.sira.metal
      "android.hardware.sensor.compass",
      // org.thoughtcrime.securesms
      "android.hardware.telephony",
      // org.videolan.vlc
      "android.hardware.screen.landscape",
   },
   System_Shared_Library: []string{
      // com.amctve.amcfullepisodes
      "org.apache.http.legacy",
      // com.binance.dev
      "android.test.runner",
      // com.miui.weather2
      "global-miui11-empty.jar",
   },
   GL_Extension: []string{
      // com.instagram.android
      "GL_OES_compressed_ETC1_RGB8_texture",
      // com.kakaogames.twodin
      "GL_KHR_texture_compression_astc_ldr",
   },
   // com.axis.drawingdesk.v3
   // valid range 0x3_0001 - 0x7FFF_FFFF
   GL_ES_Version: 0xF_FFFF,
}

func (h Header) Set_Agent(head http.Header) {
   // `sdk` is needed for `/fdfe/delivery`
   b := []byte("Android-Finsky (sdk=")
   // valid range 0 - 0x7FFF_FFFF
   b = strconv.AppendInt(b, 9, 10)
   // com.android.vending
   b = append(b, ",versionCode="...)
   if h.Single {
      // valid range 8032_0000 - 8091_9999
      b = strconv.AppendInt(b, 8091_9999, 10)
   } else {
      // valid range 8092_0000 - 0x7FFF_FFFF
      b = strconv.AppendInt(b, 9999_9999, 10)
   }
   b = append(b, ')')
   head.Set("User-Agent", string(b))
}

func (a Auth) Create(name string) error {
   query := format_query(a.Values)
   return os.WriteFile(name, []byte(query), os.ModePerm)
}

const Sleep = 4 * time.Second

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
