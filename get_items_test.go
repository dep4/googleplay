package googleplay

import (
   "os"
   "testing"
)

func Test_Get_Items(t *testing.T) {
   home, err := os.UserHomeDir()
   if err != nil {
      t.Fatal(err)
   }
   var head Header
   if err := head.Open_Device(home + "/googleplay/x86.bin"); err != nil {
      t.Fatal(err)
   }
   res, err := head.Get_Items("com.google.android.youtube")
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   res.Create("ignore.txt")
   /*
   cat, err := item.Category()
   if err != nil {
      t.Fatal(err)
   }
   if cat != "Video Players & Editors" {
      t.Fatal(cat)
   }
   */
}
