package googleplay

import (
   "bytes"
   "github.com/89z/rosso/protobuf"
   "io"
   "net/http"
)

// A Sleep is needed after this.
func (c Config) Checkin(native_platform string) (*Device, error) {
   checkin := protobuf.Message{
      4: protobuf.Message{ // checkin
         1: protobuf.Message{ // build
            // sdkVersion
            // multiple APK valid range 14 - 0x7FFF_FFFF
            // single APK valid range 14 - 28
            10: protobuf.Varint(28),
         },
         18: protobuf.Varint(1), // voiceCapable
      },
      // version
      // valid range 2 - 3
      14: protobuf.Varint(3),
      18: protobuf.Message{ // deviceConfiguration
         1: protobuf.Varint(c.Touch_Screen),
         2: protobuf.Varint(c.Keyboard),
         3: protobuf.Varint(c.Navigation),
         4: protobuf.Varint(c.Screen_Layout),
         5: protobuf.Varint(c.Has_Hard_Keyboard),
         6: protobuf.Varint(c.Has_Five_Way_Navigation),
         7: protobuf.Varint(c.Screen_Density),
         8: protobuf.Varint(c.GL_ES_Version),
         11: protobuf.String(native_platform),
      },
   }
   for _, library := range c.System_Shared_Library {
      // .deviceConfiguration.systemSharedLibrary
      checkin.Get(18).Add_String(9, library)
   }
   for _, extension := range c.GL_Extension {
      // .deviceConfiguration.glExtension
      checkin.Get(18).Add_String(15, extension)
   }
   for _, name := range c.New_System_Available_Feature {
      // .deviceConfiguration.newSystemAvailableFeature
      checkin.Get(18).Add(26, protobuf.Message{
         1: protobuf.String(name),
      })
   }
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/checkin",
      bytes.NewReader(checkin.Marshal()),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-protobuffer")
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   buf, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   var dev Device
   dev.Message, err = protobuf.Unmarshal(buf)
   if err != nil {
      return nil, err
   }
   return &dev, nil
}
