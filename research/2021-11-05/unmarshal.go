package main

import (
   "fmt"
   "github.com/segmentio/encoding/proto"
)

/*
number:1 type:2 value:
   number:2 type:2 value:
      number:4 type:2 value:
         number:13 type:2 value:
            number:1 type:2 value:
               number:1 type:2 value:Google LLC
               number:3 type:0 value:1524094400
               number:4 type:2 value:16.43.34
*/
type details struct {
   N1 struct {
      N2 struct {
         N4 struct {
            N13 struct {
               N1 struct {
                  N4 string `protobuf:"bytes,4"`
               } `protobuf:"bytes,1"`
            } `protobuf:"bytes,13"`
         } `protobuf:"bytes,4"`
      } `protobuf:"bytes,2"`
   } `protobuf:"bytes,1"`
}

var buf = []byte("\n\xe0D\x12\xddD\"\xc7B\n\x1acom.google.android.youtube\x12\x1acom.google.android.youtube\x18\x01 \x03*\aYouTube2\nGoogle LLC:\xdb\x11Get the official YouTube app on Android phones and tablets. See what the world is watching -- from the hottest music videos to what’s popular in gaming, fashion, beauty, news, learning and more. Subscribe to channels you love, create content of your own, share with friends, and watch on any device.<br><br>Watch and subscribe <br>● Browse personal recommendations on Home<br>● See the latest from your favorite channels in Subscriptions <br>● Look up videos you’ve watched, liked, and saved for later in Library <br><br>Explore different topics, what’s popular, and on the rise (available in select countries)<br>● Stay up to date on what’s popular in music, gaming, beauty, news, learning and more<br>● See what’s trending on YouTube and around the world on Explore<br>● Learn about the coolest Creators, Gamers, and Artists on the Rise (available in select countries)<br><br>Connect with the YouTube community<br>● Keep up with your favorites creators with Posts, Stories, Premieres, and Live streams<br>● Join the conversation with comments and interact with creators and other community members<br><br>Create content from your mobile device<br>● Create or upload your own videos directly in the app <br>● Engage with your audience in real time with live streaming right from the app<br><br>Find the experience that fits you and your family (available in select countries)<br>● Every family has their own approach to online video. Learn about your options: the YouTube Kids app or a new parent supervised experience on YouTube at youtube.com/myfamily<br><br>Support creators you love with channel memberships (available in select countries)<br>● Join channels that offer paid monthly memberships and support their work<br>● Get access to exclusive perks from the channel &amp; become part of their members community<br>● Stand out in comments and live chats with a loyalty badge next to your username<br><br>Upgrade to YouTube Premium (available in select countries)<br>● Watch videos uninterrupted by ads, while using other apps, or when the screen is locked<br>● Save videos for when you really need them – like when you’re on a plane or commuting<br>● Get access to YouTube Music Premium as part of your benefitsB\x13\b\x00\x12\x03USD\x1a\x00(\x00@\x01\xd2\x01\x03BuyJ\x02(\x01Rx\b\x02*ihttps://play-lh.googleusercontent.com/vA4tG0v4aasE7oIvRIvTkOYTwom07DfqHdUPr6k7jmrDwy_qA_SonqZkw6KX0OXKAdkH\x01z\a#363636Rr\b\x04*ihttps://play-lh.googleusercontent.com/lMoItBgdPPVDJsNOVtP26EKHePkwBg-PkuY9NOrc-fumRtTFP4XhpUNk_22syN4DatcH\x01\xc8\x01\x01Rw\b\x01\x13\x18\xb8\b \xf0\x10\x14*ihttps://play-lh.googleusercontent.com/195bIEw9tXffhbtw_sgG68AkuVRslTPMAV1qvizuzyNzibGHZNv7frfUOHXMNk4efhAH\x01Rw\b\x01\x13\x18\xb8\b \xf0\x10\x14*ihttps://play-lh.googleusercontent.com/ueBLRqpvC7wa-Gzy_u_J9Q96VKTqZ87GJgnBFdejMNaB0ib9iFCk0YJ8mJLN2GSvCPoH\x01Rx\b\x01\x13\x18\xb8\b \xf0\x10\x14*jhttps://play-lh.googleusercontent.com/tCydHDF9rUgVuHh9VVU-DmBPEhtqQJtz8XXgdFuckTikHinDMKy-aAPnJj8DBSvH91xqH\x01Rv\b\x01\x13\x18\xb8\b \xf0\x10\x14*hhttps://play-lh.googleusercontent.com/1QyKb-0B05Tzhmiua6vugeDaAbxi9UYWKmqSjM5Vpu8o9xgobyilE67haQxR2jCiSgH\x01j\xed\x13\n\xea\x13\n\nGoogle LLC\x18\xc0\xab\xdf\xd6\x05\"\b16.43.34H\x95\xaf\xc4\x10R)android.permission.ACCESS_COARSE_LOCATIONR'android.permission.ACCESS_FINE_LOCATIONR'android.permission.ACCESS_NETWORK_STATER$android.permission.ACCESS_WIFI_STATER\x19android.permission.CAMERAR%android.permission.FOREGROUND_SERVICER\x1fandroid.permission.GET_ACCOUNTSR#android.permission.GET_PACKAGE_SIZER\x1bandroid.permission.INTERNETR\"android.permission.MANAGE_ACCOUNTSR#android.permission.MANAGE_DOCUMENTSR\x16android.permission.NFCR android.permission.READ_CONTACTSR(android.permission.READ_EXTERNAL_STORAGER#android.permission.READ_PHONE_STATER)android.permission.RECEIVE_BOOT_COMPLETEDR\x1fandroid.permission.RECORD_AUDIOR&android.permission.SYSTEM_ALERT_WINDOWR android.permission.USE_BIOMETRICR\"android.permission.USE_CREDENTIALSR\"android.permission.USE_FINGERPRINTR\x1aandroid.permission.VIBRATER\x1candroid.permission.WAKE_LOCKR)android.permission.WRITE_EXTERNAL_STORAGER*com.google.android.c2dm.permission.RECEIVER4com.google.android.gms.permission.AD_ID_NOTIFICATIONR:com.google.android.providers.gsf.permission.READ_GSERVICESR1com.google.android.youtube.permission.C2D_MESSAGER)com.htc.launcher.permission.READ_SETTINGSR+com.htc.launcher.permission.UPDATE_SHORTCUTR.com.sec.android.provider.badge.permission.READR/com.sec.android.provider.badge.permission.WRITER0com.sonyericsson.home.permission.BROADCAST_BADGER4com.sonymobile.home.permission.PROVIDER_INSERT_BADGEZ\x1cytandroid-support@google.comb5https://support.google.com/youtube/topic/2422554?rd=1j\x0f10,000,000,000+r\x1acom.google.android.youtubezjFor new features, look for in-product education &amp; notifications sharing the feature and how to use it!\x82\x01\fOct 27, 2021\x8a\x01\r\b\x00\x10\xc0\xab\xdf\xd6\x05\x18\xac֧\t\x8a\x01\x19\b\x00\x10\xc0\xab\xdf\xd6\x05\x18\x9e\x86D\"\vconfig.ldpi\x8a\x01\x17\b\x00\x10\xc0\xab\xdf\xd6\x05\x18\x99\xe3\x0f\"\tconfig.en\x8a\x01!\b\x00\x10\xc0\xab\xdf\xd6\x05\x18\xb2\xef\xc8\x06\"\x12config.armeabi_v7a\x9a\x01\x1bJLskwF5H4K76aKWKdmF52bYTpgA\xa8\x01\x01\xf2\x01\fContains ads\x80\x02\x1f\x92\x02\xc7\x02\x10\x95\xaf\xc4\x10\x1a#\n\x16com.google.android.gms\x10\xe0ۜf \x01(\x000\x01 \x1f(\x028\x03P\x04Z\vconfig.ldpiZ\tconfig.enZ\x12config.armeabi_v7a\x8a\x01\xe5\x01\n\xe2\x01AB-xQnpz23vL298UqsJb7FaqV55RQTC16dGRSrvvYpcDnpGTYvmYqPkkVj798bHYqA6fDIAbtYFEG2zdeYoG7XHRhUfiPGSMANlsPxX8LiDxYJRcQYQHNiiaJy9VS8UHsI7KRYgZYyUEXDSjMpsBDsccb3f-nero5Gsc6xWjEdYt_u2Vvo1Jri9uaZN4VFmRbd_iuglmv0m0LEMtm2EBIA6GZo3A9OuWNg\x9a\x02\x91\x01\b\x01\x10\x00\x18\x002~\b\x01*xhttps://play-lh.googleusercontent.com/DX5LZpc4SUuvYv2kLA31fi05liLQwizh4DbHYwQthmNP15rul46zG-aquf2Vaf3M9EeBIihWmCsyIWFSLQH\x01:\aYouTube@\x02\xea\x02.1600 Amphitheatre Parkway, Mountain View 94043\xf2\x02;\x129\x127browseDeveloperPage?docid=developer-5700313618786177705\x82\x03\x17Video Players & Editors\xa8\x03\x80ȯ\xa0%\xea\x03\x0410B+\x82\x04\x1c\n\fOct 20, 2010\x12\f\b\x9c\xfe\xfc\xe5\x04\x10\x80\x98\xff\xcd\x01\xa2\x04\f\n\nGoogle LLC\xb0\x04\x9c\x87\x85\xde'\xca\x04\b\n\x06\b\aB\x02\b\x03\xd2\x04\b\n\x06\b\aB\x02\b\x0e\xe2\x04\b\n\x06\b\aB\x02\b\x13\xea\x04\x0310B\xf2\x04\n10 billionrE\b\x02\x15pև@\x18\xe5\x90\xe6; \xa7\xe6\xdc\x06(\xba\xfc\xf1\x010۵\xb6\x038\xa2\xbd\xdd\x05@\x85\xbb\x83*X壖\x14\x8a\x01\x034.2\x92\x01\x04125M\x9a\x01\v125 millionz\xce\x0eJ\x8b\x03\n\fContains ads\x12\xf9\x01\b\x06*xhttps://play-lh.googleusercontent.com/fIlJSil0lKMCEw_jC8R18ZsZZyf4Zh-4-lE1JhlMg981NqIPTuKEcXQpFXYm2wtolaWizL6CFDGQZfqTVAH\x01\xe2\x01xhttps://play-lh.googleusercontent.com/0Yzbn6OLLsP6GXofukTOoL5cruQFHrDmZAbjbsojrKSLKqwoP1N4jNToMSZA6A8lA86KPeLiCdfyARlmVA\"\u007fAds are placed by the app developer. <br/> <a href=\"https://support.google.com/googleplay/?p=ads_badge&hl=en_US\">Learn more</a>\x92\x01&http://www.google.com/policies/privacy\xea\x01\x8b\x02\n\x04Teen\x12\x84\x01\b\x06\x13\x18\x80\x04 \x80\x04\x14*vhttps://play-lh.googleusercontent.com/mw_NfsvKM8m6RPv8Fz2GQawCOsqWv010saMnc7zbWalMxuaA9IY8h7E0VMieLxSxAFB98NFeYqbFrXXqH\x01\"vUsers Interact, Digital Purchases<br/><a href=\"https://support.google.com/googleplay?p=appgame_ratings\">Learn more</a>*\x04Teen\xba\x04C\n\x17Video Players & Editors\x12 \x12\x1e\x12\x1chomeV2?cat=VIDEO_PLAYERS&c=3\x1a\x06\b\aB\x02\b\x05\xc2\x04\xaf\x03\n\x034.2\x12\xf9\x01\b\x06*xhttps://play-lh.googleusercontent.com/B8PBzYOiqekDl0G2cVEpyywdw6gUqK-3s-uY_xsjhknCDK8nZi6bI1qexaKof-oNmdLqx9d9K5cawnW9IAH\x01\xe2\x01xhttps://play-lh.googleusercontent.com/Ig2Xk7RjPW6it1-nypnuxw8N3At656wHwtVWcZSFTba5pKzuLD5PQotaZqH2v6Paw2MT_HK7E5vA0zWOSw\x1a\f125M reviews\"\x06\b\aB\x02\b\x04:e\n'rev?doc=com.google.android.youtube&n=20\x12:\x128\x9a\x035allReviews?doc=com.google.android.youtube&sfilter=ALLB/Average rating 4.2 stars in 125 million reviews\xc2\x04;\n\x0410B+\x1a\tDownloads\"\x06\b\aB\x02\b\x03B Downloaded 10 billion plus times\xc2\x04\x94\x03\x12\x84\x01\b\x06\x13\x18\x80\x04 \x80\x04\x14*vhttps://play-lh.googleusercontent.com/mw_NfsvKM8m6RPv8Fz2GQawCOsqWv010saMnc7zbWalMxuaA9IY8h7E0VMieLxSxAFB98NFeYqbFrXXqH\x01\x1a\x04Teen\"\x06\b\aB\x02\b\x0f2\xe7\x01\n\x86\x01\b\x06\x13\x18\x90\x02 \x90\x02\x14*xhttps://play-lh.googleusercontent.com/xl04GurQer374yogi24hRkvUcXtMStMOcOo86Sprmrid9j6wQHF0d9xMh73CwUEg57lbsGIwpOXgrJ927QH\x01\x1a!Users Interact, Digital Purchases*9\n7https://support.google.com/googleplay?p=appgame_ratingsB\x13Content rating Teen\x82\x058apps/detailsLiveOpsStream?doc=com.google.android.youtube\xaa\x05=promotion/detailsPagePromotion?doc=com.google.android.youtube\x8a\x06 \x12\x1e\x12\x1chomeV2?cat=VIDEO_PLAYERS&c=3\xb2\x06\b\n\x06\b\aB\x02\b\x05\xe2\x06\rVIDEO_PLAYERS\x82\x01&details?doc=com.google.android.youtube\x8a\x01Hhttps://play.google.com/store/apps/details?id=com.google.android.youtube\x92\x01'rev?doc=com.google.android.youtube&n=20\xa8\x01\x01\xc2\x01\xa1\x01\b\x02\x1aK\n\"\x12 \n\x1acom.google.android.youtube\x10\x01\x18\x03\x10\x002\x13\b\xbd\xc0\x88\x9bā\xf4\x02\x15\xea\x12\xc9\n\x1d\xea\x1d\n\xae\x8a\x01\r\b\x00\x12\t\n\x05en-US\x10\x00\xaa\x02O\x1aM\b\x00\x12\"\n \n\x1acom.google.android.youtube\x10\x01\x18\x03J\x13\b\xbd\xc0\x88\x9bā\xf4\x02\x15\xea\x12\xc9\n\x1d\xea\x1d\n\xae\xfa\x01\x0f\n\r\b\x00\x12\t\n\x05en-US\x10\x00\xca\x01;\x12\x1a\n\nOffered by\x1a\f\x12\nGoogle LLC\x12\x1d\n\vReleased on\x1a\x0e\x12\fOct 20, 2010\xd0\x01\x00\xda\x01FEnjoy your favorite videos and channels with the official YouTube app.\xe8\x01\x00\x80\x02\x00\xb2\x02\n\b\x01\x12\x03lpi\x1a\x011\xba\x02,reviewSummary?doc=com.google.android.youtube\xc2\x02:\x128\x9a\x035allReviews?doc=com.google.android.youtube&sfilter=ALL\xd8\x02\x03\xe8\x02\x01\x8a\x03\f\b\x01\x12\b\n\x06\b\aB\x02\b\x12\x92\x03\xb2\x02\n\x04Teen\x12vUsers Interact, Digital Purchases<br/><a href=\"https://support.google.com/googleplay?p=appgame_ratings\">Learn more</a>\x1a\x84\x01\x10\x02\x1a\x06\b\x80\x04\x10\x80\x042x\nvhttps://play-lh.googleusercontent.com/mw_NfsvKM8m6RPv8Fz2GQawCOsqWv010saMnc7zbWalMxuaA9IY8h7E0VMieLxSxAFB98NFeYqbFrXXq2!Users Interact, Digital PurchasesB\b\n\x06\b\aB\x02\b\x0f\x9a\x03\b\n\x06\b\aB\x02\b\x042\x97\x01\b\x01\x12K\b\x05\"\"\x12 \n\x1acom.google.android.youtube\x10\x01\x18\x03Z\x13\b\xbd\xc0\x88\x9bā\xf4\x02\x15\xea\x12\xc9\n\x1d\xea\x1d\n\xae\x82\x01\r\b\x00\x12\t\n\x05en-US\x10\x00\xaa\x02E\nC\n\fDETAILS_PAGE\x123\n \n\x1acom.google.android.youtube\x10\x01\x18\x03\x1a\x0f\n\r\b\x00\x12\t\n\x05en-US\x10\x00@\x01j/getDetailsStream?doc=com.google.android.youtube\x8a\x01:getPostAcquireDetailsStream?doc=com.google.android.youtube\xaa\x01\x06popups*\x03\b\xd3\x01J\x1c\b\x12\x9a\x01\x17\n\x13\b\xbd\xc0\x88\x9bā\xf4\x02\x15\xea\x12\xc9\n\x1d\xea\x1d\n\xae\x10\x04")

func main() {
   var det details
   err := proto.Unmarshal(buf, &det)
   if err != nil {
      panic(err)
   }
   fmt.Printf("%+v\n", det)
}
