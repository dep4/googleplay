package main

import (
   "io"
   "net/http"
   "net/http/httputil"
   "net/url"
   "os"
   "strings"
)

func main() {
   var req http.Request
   req.Header["Authorization"] = []string{authorization}
   req.Header["X-Dfe-Device-Id"] = []string{device}
   req.Body = io.NopCloser(body)
   req.Header = make(http.Header)
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
   req.Method = "POST"
   req.URL = new(url.URL)
   req.URL.Host = "play-fe.googleapis.com"
   req.URL.Path = "/fdfe/sync"
   req.URL.RawPath = ""
   val := make(url.Values)
   req.URL.RawQuery = val.Encode()
   req.URL.Scheme = "https"
   res, err := new(http.Transport).RoundTrip(&req)
   if err != nil {
      panic(err)
   }
   defer res.Body.Close()
   buf, err := httputil.DumpResponse(res, true)
   if err != nil {
      panic(err)
   }
   os.Stdout.Write(buf)
}

var body = strings.NewReader("\n1:/\n-\n+rh_dzz-uEzvauarv0XpP9CbIiw6UIOlt_57Izz3THCw\n\x1fJ\x1d\n\x06310270\x12\x13\n\x11\b\x80؉\xe9\x85\xc6F\x12\aAndroid\n1B/\n-\n+rh_dzz-uEzvauarv0XpP9CbIiw6UIOlt_57Izz3THCw\n\xad\"R\xaa\"\n#\n!android.hardware.sensor.proximity\n'\n%android.hardware.sensor.accelerometer\n\x1c\n\x1aandroid.hardware.faketouch\n \n\x1eandroid.hardware.usb.accessory\n\x19\n\x17android.software.backup\n\x1e\n\x1candroid.hardware.touchscreen\n)\n'android.hardware.touchscreen.multitouch\n\x18\n\x16android.software.print\n2\n0com.google.android.feature.PIXEL_2017_EXPERIENCE\n$\n\"android.software.voice_recognizers\n2\n0com.google.android.feature.PIXEL_2018_EXPERIENCE\n+\n)android.hardware.sensor.relative_humidity\n#\n!android.hardware.camera.autofocus\n)\n'com.google.android.feature.GOOGLE_BUILD\n \n\x1eandroid.hardware.telephony.gsm\n\x1b\n\x19android.software.sip.voip\n\x1b\n\x19android.hardware.usb.host\n\x1f\n\x1dandroid.hardware.audio.output\n\x1f\n\x1dandroid.hardware.camera.front\n\"\n android.hardware.screen.portrait\n*\n(com.google.android.feature.TURBO_PRELOAD\n-\n+android.hardware.sensor.ambient_temperature\n\x1e\n\x1candroid.software.home_screen\n\x1d\n\x1bandroid.hardware.microphone\n-\n+com.google.android.feature.PIXEL_EXPERIENCE\n2\n0android.hardware.touchscreen.multitouch.jazzhand\n#\n!android.hardware.sensor.barometer\n\x1e\n\x1candroid.software.app_widgets\n \n\x1eandroid.software.input_methods\n\x1f\n\x1dandroid.hardware.sensor.light\n\x1f\n\x1dandroid.software.device_admin\n&\n$com.google.android.feature.WELLBEING\n\x19\n\x17android.hardware.camera\n#\n!android.hardware.screen.landscape\n \n\x1eandroid.software.managed_users\n\x1a\n\x18android.software.webview\n\x1d\n\x1bandroid.hardware.camera.any\n$\n\"android.software.connectionservice\n2\n0android.hardware.touchscreen.multitouch.distinct\n#\n!android.hardware.location.network\n\x16\n\x14android.software.sip\n*\n(com.google.android.apps.dialer.SUPPORTED\n!\n\x1fandroid.software.live_wallpaper\n.\n,com.google.android.feature.GOOGLE_EXPERIENCE\n)\n'com.google.android.feature.EXCHANGE_6_2\n\x1f\n\x1dandroid.hardware.location.gps\n\x17\n\x15android.hardware.wifi\n\x1b\n\x19android.hardware.location\n\x1c\n\x1aandroid.hardware.telephony\x12 com.google.android.media.effects\x12\x1dcom.android.location.provider\x12 com.android.future.usb.accessory\x12\x12android.ext.shared\x12\njavax.obex\x12\x16com.google.android.gms\x12\x14android.ext.services\x12\x13android.test.runner\x12!com.google.android.dialer.support\x12\x17com.google.android.maps\x12\x16org.apache.http.legacy\x12\x1fcom.android.media.remotedisplay\x12\x1bcom.android.mediadrm.signer\x1a\x02af\x1a\x05af-ZA\x1a\x02am\x1a\x05am-ET\x1a\x02ar\x1a\x05ar-EG\x1a\x05ar-XB\x1a\x02az\x1a\x05az-AZ\x1a\x02be\x1a\x05be-BY\x1a\x02bg\x1a\x05bg-BG\x1a\x02bn\x1a\x05bn-BD\x1a\x05bs-BA\x1a\x02ca\x1a\x05ca-ES\x1a\x02cs\x1a\x05cs-CZ\x1a\x02da\x1a\x05da-DK\x1a\x02de\x1a\x05de-AT\x1a\x05de-CH\x1a\x05de-DE\x1a\x05de-LI\x1a\x02el\x1a\x05el-GR\x1a\x02en\x1a\x05en-AU\x1a\x05en-CA\x1a\x05en-GB\x1a\x05en-IN\x1a\x05en-NZ\x1a\x05en-SG\x1a\x05en-US\x1a\x05en-XA\x1a\x02es\x1a\x05es-ES\x1a\x05es-US\x1a\x02et\x1a\x05et-EE\x1a\x02eu\x1a\x05eu-ES\x1a\x02fa\x1a\x05fa-IR\x1a\x02fi\x1a\x05fi-FI\x1a\x02fr\x1a\x05fr-BE\x1a\x05fr-CA\x1a\x05fr-CH\x1a\x05fr-FR\x1a\x02gl\x1a\x05gl-ES\x1a\x02gu\x1a\x05gu-IN\x1a\x02hi\x1a\x05hi-IN\x1a\x02hr\x1a\x05hr-HR\x1a\x02hu\x1a\x05hu-HU\x1a\x02hy\x1a\x05hy-AM\x1a\x02id\x1a\x02in\x1a\x05in-ID\x1a\x02is\x1a\x05is-IS\x1a\x02it\x1a\x05it-CH\x1a\x05it-IT\x1a\x02iw\x1a\x05iw-IL\x1a\x02ja\x1a\x05ja-JP\x1a\x02ka\x1a\x05ka-GE\x1a\x02kk\x1a\x05kk-KZ\x1a\x02km\x1a\x05km-KH\x1a\x02kn\x1a\x05kn-IN\x1a\x02ko\x1a\x05ko-KR\x1a\x02ky\x1a\x05ky-KG\x1a\x02lo\x1a\x05lo-LA\x1a\x02lt\x1a\x05lt-LT\x1a\x02lv\x1a\x05lv-LV\x1a\x02mk\x1a\x05mk-MK\x1a\x02ml\x1a\x05ml-IN\x1a\x02mn\x1a\x05mn-MN\x1a\x02mr\x1a\x05mr-IN\x1a\x02ms\x1a\x05ms-MY\x1a\x02my\x1a\x05my-MM\x1a\x02nb\x1a\x05nb-NO\x1a\x02ne\x1a\x05ne-NP\x1a\x02nl\x1a\x05nl-BE\x1a\x05nl-NL\x1a\x02pa\x1a\x05pa-IN\x1a\x02pl\x1a\x05pl-PL\x1a\x02pt\x1a\x05pt-BR\x1a\x05pt-PT\x1a\x02ro\x1a\x05ro-RO\x1a\x02ru\x1a\x05ru-RU\x1a\x02si\x1a\x05si-LK\x1a\x02sk\x1a\x05sk-SK\x1a\x02sl\x1a\x05sl-SI\x1a\x02sq\x1a\x05sq-AL\x1a\x02sr\x1a\asr-Latn\x1a\x05sr-RS\x1a\x02sv\x1a\x05sv-SE\x1a\x02sw\x1a\x05sw-TZ\x1a\x02ta\x1a\x05ta-IN\x1a\x02te\x1a\x05te-IN\x1a\x02th\x1a\x05th-TH\x1a\x03fil\x1a\x06fil-PH\x1a\x02tr\x1a\x05tr-TR\x1a\x02uk\x1a\x05uk-UA\x1a\x02ur\x1a\x05ur-PK\x1a\x02uz\x1a\x05uz-UZ\x1a\x02vi\x1a\x05vi-VN\x1a\x05zh-CN\x1a\x05zh-HK\x1a\x05zh-TW\x1a\x02zu\x1a\x05zu-ZA\" ANDROID_EMU_async_frame_commands\"\x1eANDROID_EMU_async_unmap_buffer\" ANDROID_EMU_gles_max_version_3_1\"\x1fANDROID_EMU_host_composition_v1\"\x1fANDROID_EMU_host_composition_v2\"\x1dANDROID_EMU_host_side_tracing\"\x1cANDROID_EMU_sync_buffer_data\" GL_APPLE_texture_format_BGRA8888\"\x19GL_EXT_color_buffer_float\"\x1eGL_EXT_color_buffer_half_float\"\x13GL_EXT_debug_marker\"\x11GL_EXT_robustness\"\x1eGL_EXT_texture_format_BGRA8888\"#GL_KHR_texture_compression_astc_ldr\"\x10GL_OES_EGL_image\"\x19GL_OES_EGL_image_external\"\x1fGL_OES_EGL_image_external_essl3\"\x0fGL_OES_EGL_sync\"\x1eGL_OES_blend_equation_separate\"\x1aGL_OES_blend_func_separate\"\x15GL_OES_blend_subtract\"\x17GL_OES_byte_coordinates\"#GL_OES_compressed_ETC1_RGB8_texture\"\"GL_OES_compressed_paletted_texture\"\x0eGL_OES_depth24\"\x0eGL_OES_depth32\"\x14GL_OES_depth_texture\"\x13GL_OES_draw_texture\"\x19GL_OES_element_index_uint\"\x18GL_OES_fbo_render_mipmap\"\x19GL_OES_framebuffer_object\"\x1bGL_OES_packed_depth_stencil\"\x17GL_OES_point_size_array\"\x13GL_OES_point_sprite\"\x11GL_OES_rgb8_rgba8\"\x17GL_OES_single_precision\"\x0fGL_OES_stencil1\"\x0fGL_OES_stencil4\"\x0fGL_OES_stencil8\"\x13GL_OES_stencil_wrap\"\x17GL_OES_texture_cube_map\"\x1bGL_OES_texture_env_crossbar\"\x14GL_OES_texture_float\"\x1bGL_OES_texture_float_linear\"\x19GL_OES_texture_half_float\" GL_OES_texture_half_float_linear\"\x1dGL_OES_texture_mirored_repeat\"\x13GL_OES_texture_npot\"\x1aGL_OES_vertex_array_object\"\x18GL_OES_vertex_half_float\n\bZ\x06\b\x02\x10\x01\x18\x01\n.b,\n\aunknown\x12\aNexus 4\x1a\avbox86p\"\avbox86p*\x06Google\n\xcf\x01j\xcc\x01\x12d\n\x16com.google.android.gms\x12\x1bOJGKRT0HGZNU-LGa8F7GViztV4g\x1a+8P1sW0EPJcslw7UzRsiXL64w-O50Ed-RBICtay1g24M \x04\x12d\n\x16com.google.android.gms\x12\x1bOJGKRT0HGZNU-LGa8F7GViztV4g\x1a+8P1sW0EPJcslw7UzRsiXL64w-O50Ed-RBICtay1g24M \x04\n\vr\t\t\x00\x00\x00\x00\x00\x00\x00\x10\n\x11z\x0f\b\x00\x10\x80\x80\xe1\xe9\a\x18\x04\"\x03x86\n\x0e\x82\x01\v\n\tGMT-05:00\n\xaa\x01\x8a\x01\xa6\x01\n\xa3\x01dKUsArA7S_yvxPkY_UlMlk:APA91bGwkDvn0eW3hxikpj99SVR9BegBzKLr3fa6VkeUr4N-6_391RmH75g3DSBmq9AiuZDmuLs9CsCO7rw-gQURLlFWPRe1exrFFkMx1SpP0vyaow8k2dVIS6QSonsHH4dIu852a6tB\nB\x92\x01?\n\nam-unknown\x12\x16play-ms-android-google\x1a\x19play-ad-ms-android-google\n\b\x9a\x01\x05\x10\xae\xf0\xcb'\n\x10\xa2\x01\r\b\x03\x10\x80\x06\x18\x80\n \x02(\xc0\x02\nL\xaa\x01I\n;google/vbox86p/vbox86p:7.1.1/NMF26Q/433:userdebug/test-keys\x10\x19\x1a\x00\"\x00*\x000\x80\x80\b")
