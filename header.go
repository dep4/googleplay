package googleplay

import (
   "github.com/89z/format/protobuf"
   "net/http"
   "net/url"
   "strconv"
   "strings"
)

type Header struct {
   http.Header
}

func (h Header) Category(cat string) (*Document, error) {
   return h.getCategory(cat, "")
}

func (h Header) CategoryNext(nextPage string) (*Document, error) {
   return h.getCategory("", nextPage)
}

func (h Header) Delivery(app string, ver int64) (*Delivery, error) {
   req, err := http.NewRequest(
      "GET", "https://play-fe.googleapis.com/fdfe/delivery", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = h.Header
   req.URL.RawQuery = url.Values{
      "doc": {app},
      "vc": {strconv.FormatInt(ver, 10)},
   }.Encode()
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   responseWrapper, err := protobuf.Decode(res.Body)
   if err != nil {
      return nil, err
   }
   status := responseWrapper.Get(1, "payload").
      Get(21, "deliveryResponse").
      GetVarint(1, "status")
   switch status {
   case 2:
      return nil, errorString("Geo-blocking")
   case 3:
      return nil, errorString("Purchase required")
   case 5:
      return nil, errorString("Invalid version")
   }
   appData := responseWrapper.Get(1, "payload").
      Get(21, "deliveryResponse").
      Get(2, "appDeliveryData")
   var del Delivery
   del.DownloadURL = appData.GetString(3, "downloadUrl")
   for _, data := range appData.GetMessages(15, "splitDeliveryData") {
      var split SplitDeliveryData
      split.ID = data.GetString(1, "id")
      split.DownloadURL = data.GetString(5, "downloadUrl")
      del.SplitDeliveryData = append(del.SplitDeliveryData, split)
   }
   return &del, nil
}

func (h Header) Details(app string) (*Details, error) {
   req, err := http.NewRequest(
      "GET", "https://android.clients.google.com/fdfe/details", nil,
   )
   if err != nil {
      return nil, err
   }
   req.Header = h.Header
   req.URL.RawQuery = "doc=" + url.QueryEscape(app)
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   responseWrapper, err := protobuf.Decode(res.Body)
   if err != nil {
      return nil, err
   }
   docV2 := responseWrapper.Get(1, "payload").
      Get(2, "detailsResponse").
      Get(4, "docV2")
   var det Details
   det.CurrencyCode = docV2.Get(8, "offer").GetString(2, "currencyCode")
   det.Micros = docV2.Get(8, "offer").GetVarint(1, "micros")
   det.NumDownloads = docV2.Get(13, "details").
      Get(1, "appDetails").
      GetVarint(70, "numDownloads")
   // The shorter path 13,1,9 returns wrong size for some packages:
   // com.riotgames.league.wildriftvn
   det.Size = docV2.Get(13, "details").
      Get(1, "appDetails").
      Get(34, "installDetails").
      GetVarint(2, "size")
   det.Title = docV2.GetString(5, "title")
   det.UploadDate = docV2.Get(13, "details").
      Get(1, "appDetails").
      GetString(16, "uploadDate")
   det.VersionCode = docV2.Get(13, "details").
      Get(1, "appDetails").
      GetVarint(3, "versionCode")
   det.VersionString = docV2.Get(13, "details").
      Get(1, "appDetails").
      GetString(4, "versionString")
   return &det, nil
}

// Purchase app. Only needs to be done once per Google account.
func (h Header) Purchase(app string) error {
   query := "doc=" + url.QueryEscape(app)
   req, err := http.NewRequest(
      "POST", "https://android.clients.google.com/fdfe/purchase",
      strings.NewReader(query),
   )
   if err != nil {
      return err
   }
   h.Set("Content-Type", "application/x-www-form-urlencoded")
   req.Header = h.Header
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return err
   }
   return res.Body.Close()
}

func (h Header) getCategory(cat, nextPage string) (*Document, error) {
   var buf strings.Builder
   buf.WriteString("https://android.clients.google.com/fdfe/")
   if cat != "" {
      buf.WriteString("list?")
      val := url.Values{
         "c": {"3"},
         "cat": {cat},
         "ctr": {"apps_topselling_free"},
      }.Encode()
      buf.WriteString(val)
   } else {
      buf.WriteString(nextPage)
   }
   req, err := http.NewRequest("GET", buf.String(), nil)
   if err != nil {
      return nil, err
   }
   req.Header = h.Header
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   responseWrapper, err := protobuf.Decode(res.Body)
   if err != nil {
      return nil, err
   }
   docV2 := responseWrapper.Get(1, "payload").
      Get(1, "listResponse").
      Get(2, "doc")
   var doc Document
   for _, elem := range docV2.GetMessages(11, "child") {
      var child Document
      child.ID = elem.GetString(1, "docID")
      child.Title = elem.GetString(5, "title")
      child.Creator = elem.GetString(6, "creator")
      doc.Child = append(doc.Child, child)
   }
   doc.NextPageURL = docV2.Get(12, "containerMetadata").
      GetString(2, "nextPageUrl")
   return &doc, nil
}

func (t Token) Header(dev *Device) (*Header, error) {
   return t.headerVersion(dev, 9999_9999)
}

func (t Token) SingleAPK(dev *Device) (*Header, error) {
   return t.headerVersion(dev, 8091_9999)
}

func (t Token) headerVersion(dev *Device, version int64) (*Header, error) {
   val := url.Values{
      "Token": {t.Token},
      "service": {"oauth2:https://www.googleapis.com/auth/googleplay"},
   }.Encode()
   req, err := http.NewRequest(
      "POST", "https://android.googleapis.com/auth", strings.NewReader(val),
   )
   if err != nil {
      return nil, err
   }
   req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
   LogLevel.Dump(req)
   res, err := new(http.Transport).RoundTrip(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return nil, errorString(res.Status)
   }
   var head Header
   head.Header = make(http.Header)
   auth := parseQuery(res.Body).Get("Auth")
   if auth != "" {
      head.Set("Authorization", "Bearer " + auth)
   }
   buf := []byte("Android-Finsky (sdk=9,versionCode=")
   buf = strconv.AppendInt(buf, version, 10)
   head.Set("User-Agent", string(buf))
   id := strconv.FormatUint(dev.AndroidID, 16)
   head.Set("X-DFE-Device-ID", id)
   return &head, nil
}