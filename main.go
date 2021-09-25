// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// https://github.com/line/line-bot-sdk-go/tree/master/linebot

package main

import (
	"strconv"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Constants
var timeFormat = "01/02 PM03:04:05"
var user_zchien = "U696bcb700dfc9254b27605374b86968b"
var user_yaoming = "U3aaab6c6248bb38f194134948c60f757"
var user_jackal = "U3effab06ddf5bcf0b46c1c60bcd39ef5"
var user_shane = "U2ade7ac4456cb3ca99ffdf9d7257329a"

// Global Settings
var channelSecret = os.Getenv("CHANNEL_SECRET")
var channelToken = os.Getenv("CHANNEL_TOKEN")
//var baseURL = os.Getenv("APP_BASE_URL")
var baseURL = "https://line-talking-bot-go.herokuapp.com"
var endpointBase = os.Getenv("ENDPOINT_BASE")
var tellTimeInterval int = 15
var answers_TextMessage = []string{
		"",
	}
var answers_ImageMessage = []string{
		"",
	}
var answers_StickerMessage = []string{
		"",
	}
var answers_VideoMessage = []string{
		"",
	}
var answers_AudioMessage = []string{
		"",
	}
var answers_LocationMessage = []string{
		"",
	}
var answers_ReplyCurseMessage = []string{
		"",
	}

var silentMap = make(map[string]bool) // [UserID/GroupID/RoomID]:bool

//var echoMap = make(map[string]bool)

var loc, _ = time.LoadLocation("Asia/Tehran")
var bot *linebot.Client


func tellTime(replyToken string, doTell bool){
	var silent = false
	now := time.Now().In(loc)
	nowString := now.Format(timeFormat)
	
	if doTell {
		log.Println("ساعت بوقت تهران:  " + nowString)
		bot.ReplyMessage(replyToken, linebot.NewTextMessage("ساعت بوقت تهران: " + nowString)).Do()
	} else if silent != true {
		log.Println("ساعت بوقت تهران:  " + nowString)
		bot.PushMessage(replyToken, linebot.NewTextMessage("ساعت بوقت تهران:  " + nowString)).Do()
	} else {
		log.Println("tell time misfired")
	}
}

func tellTimeJob(sourceId string) {
	for {
		time.Sleep(time.Duration(tellTimeInterval) * time.Minute)
		now := time.Now().In(loc)
		log.Println("time to tell time to : " + sourceId + ", " + now.Format(timeFormat))
		tellTime(sourceId, false)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	/*
	go func() {
		tellTimeJob(user_zchien);
	}()
	go func() {
		for {
			now := time.Now().In(loc)
			log.Println("keep alive at : " + now.Format(timeFormat))
			//http.Get("https://line-talking-bot-go.herokuapp.com")
			time.Sleep(time.Duration(rand.Int31n(29)) * time.Minute)
		}
	}()
	*/

	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	log.Print("URL:"  + r.URL.String())
	
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		var replyToken = event.ReplyToken

		var source = event.Source //EventSource		
		var userId = source.UserID
		var groupId = source.GroupID
		var roomId = source.RoomID
		log.Print("callbackHandler to source UserID/GroupID/RoomID: " + userId + "/" + groupId + "/" + roomId)
		
		var sourceId = roomId
		if sourceId == "" {
			sourceId = groupId
			if sourceId == "" {
				sourceId = userId
			}
		}
		
		if event.Type == linebot.EventTypeMessage {
			_, silent := silentMap[sourceId]
			
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				log.Print("ReplyToken[" + replyToken + "] TextMessage: ID(" + message.ID + "), Text(" + message.Text  + "), current silent status=" + strconv.FormatBool(silent) )
				//if _, err = bot.ReplyMessage(replyToken, linebot.NewTextMessage(message.ID+":"+message.Text+" OK!")).Do(); err != nil {
				//	log.Print(err)
				//}
				
				if strings.Contains(message.Text, "test") {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("success")).Do()
				} else if "groupid"  == message.Text {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(string(source.GroupID))).Do()
				} else if "help" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("help\n・[image:画像url]=從圖片網址發送圖片\n・[speed]=測回話速度\n・[groupid]=發送GroupID\n・[roomid]=發送RoomID\n・[byebye]=取消訂閱\n・[about]=作者\n・[me]=發送發件人信息\n・[test]=test bowwow是否正常\n・[now]=現在時間\n・[mid]=mid\n・[sticker]=隨機圖片\n\n[其他機能]\n位置測試\n捉貼圖ID\n加入時發送消息")).Do()
				} else if "mid" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(source.UserID)).Do()
				} else if "mee6" == message.Text  {
					bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg","https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg")).Do()
				} else if "mee5" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg","https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg")).Do()
					bot.ReplyMessage(replyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/1a/e7/e18b42a3703133301659f6ddd7a4.jpg","https://www.itsfun.com.tw/cacheimg/1a/e7/e18b42a3703133301659f6ddd7a4.jpg")).Do()
				} else if "roomid" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(source.RoomID)).Do()
				} else if "hidden" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("hidden")).Do()
				} else if "bowwow" == message.Text  {
					bot.ReplyMessage(replyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg","https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg")).Do()}
					if err != nil {
						log.Fatal(err)
				} else if "me" == message.Text  {
					mid := source.UserID
					p, err := bot.GetProfile(mid).Do()
					if err != nil {
						bot.ReplyMessage(replyToken, linebot.NewTextMessage("新增同意"))
					}

					bot.ReplyMessage(replyToken, linebot.NewTextMessage("mid:"+mid+"\nname:"+p.DisplayName+"\nstatusMessage:"+p.StatusMessage)).Do()
				} else if res := strings.Contains(message.Text, "hello"); res == true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("hello!"), linebot.NewTextMessage("my name is bowwow")).Do()
				} else if res := strings.Contains(message.Text, "image:"); res == true {
					image_url := strings.Replace(message.Text, "image:", "", -1)
					bot.ReplyMessage(replyToken, linebot.NewImageMessage(image_url, image_url)).Do()
				} else if "/1" == message.Text {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9..0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.\n\n💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀"), linebot.NewTextMessage("7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7."), linebot.NewTextMessage("7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7."), linebot.NewTextMessage("7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7."), linebot.NewTextMessage("💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9..0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.\n\n💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀")).Do()
				} else if  "time" == message.Text {
					tellTime(replyToken, true)
				} else if  "/4" == message.Text {
					silentMap[sourceId] = false
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H"),linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H"),linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H"), linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H"), linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H") ,linebot.NewTextMessage(".1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H")).Do()
				} else if "profile" == message.Text {
					if source.UserID != "" {
						profile, err := bot.GetProfile(source.UserID).Do()
						if err != nil {
							log.Print(err)
						} else if _, err := bot.ReplyMessage(
							replyToken,
							linebot.NewTextMessage("Display name: "+profile.DisplayName + ", Status message: "+profile.StatusMessage)).Do(); err != nil {
								log.Print(err)
						}
					} else {
						bot.ReplyMessage(replyToken, linebot.NewTextMessage("Bot can't use profile API without user ID")).Do()
					}
				} else if "buttons" == message.Text {
					imageURL := baseURL + "/static/buttons/1040.jpg"
					//log.Print(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> "+imageURL)
					template := linebot.NewButtonsTemplate(
						imageURL, "My button sample", "Hello, my button",
						linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
						linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", ""),
						linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
						linebot.NewMessageTemplateAction("Say message", "Rice=米"),
					)
					if _, err := bot.ReplyMessage(
						replyToken,
						linebot.NewTemplateMessage("Buttons alt text", template),
					).Do(); err != nil {
						log.Print(err)
					}
				} else if "confirm" == message.Text {
					template := linebot.NewConfirmTemplate(
						"Do it?",
						linebot.NewMessageTemplateAction("Yes", "Yes!"),
						linebot.NewMessageTemplateAction("No", "No!"),
					)
					if _, err := bot.ReplyMessage(
						replyToken,
						linebot.NewTemplateMessage("Confirm alt text", template),
					).Do(); err != nil {
						log.Print(err)
					}
				} else if "LL"== message.Text {
						// imageURL = SystemImageURL
						template := linebot.NewConfirmTemplate(
							"驗證成功！\n要現在進入管理介面嗎？",
							linebot.NewPostbackTemplateAction("是","開啟管理者選單", ""),
							linebot.NewPostbackTemplateAction("否","登出管理者", ""),
						)
						obj_message := linebot.NewTemplateMessage("這功能只支援最新版本 APP 使用。\n請用 APP 端查看下一步。", template)
						if _, err = bot.ReplyMessage(event.ReplyToken,obj_message).Do(); err != nil {
							log.Print(2162)
							log.Print(err)
					    	}
					    }
			    	return
				} else if "s2"== message.Text {
			    		imageURL = SystemImageURL
			    		LineTemplate_test := linebot.NewCarouselTemplate(
			    			linebot.NewCarouselColumn(
			    				imageURL, "test", "For test mode.",
			    				linebot.NewPostbackTemplateAction("測試","test", ""),
			    				linebot.NewPostbackTemplateAction("管理模式","admin", ""),
			    				linebot.NewPostbackTemplateAction("申請使用管理者","開發者", "開發者"),
			    			),
			    			// LineTemplate_other_example,
			    			// LineTemplate_other,
			    			//LineTemplate_CarouselColumn_feedback,
			    		)
			    		no_temp_msg := "你已觸發測試模式，請更新最新版本的 LINE 查看內容 。"
			    		obj_message := linebot.NewTemplateMessage(no_temp_msg, LineTemplate_test)
			    		if _, err = bot.ReplyMessage(event.ReplyToken, obj_message).Do(); err != nil {
			    				log.Print(847)
			    				log.Print(err)
			    		}
			    	return
				} else if "s3"== message.Text {
		    		imageURL = "https://lh3.googleusercontent.com/-HH0XJGbc8rg/XvGESH1WZ5I/AAAAAAAAFro/hQKPpJ5OVM80ZsBNIR7Ou7aoPh4Rz6DAACK8BGAsYHg/s512/2020-06-22.jpg"
		    		LineTemplate_test := linebot.NewCarouselTemplate(
		    			linebot.NewCarouselColumn(
		    				imageURL, "💠ϻఠ_ఠsɛɳ💠вσт💠", "ℓ๐νә",
		    				linebot.NewURITemplateAction("Google Map","https://goo.gl/maps/h6s5ccdXrL52"),
		    				linebot.NewURITemplateAction("add me","line://ti/p/~M_BW"),
		    				linebot.NewURITemplateAction("click","https://lh3.googleusercontent.com/-pOOr24kfmm0/XzzpnO3Br8I/AAAAAAAAIVQ/NOaSZuWwV0QhaBW9h7wWsiCukigu8XOkACK8BGAsYHg/s512/2020-08-19.gif"),
		    			),
		    			// LineTemplate_other_example,
		    			// LineTemplate_other,
		    			linebot.NewCarouselColumn(
		    				imageURL, "ϻఠ_ఠsɛɳ", "вσт",
		    				linebot.NewMessageTemplateAction("ϻఠ_ఠsɛɳ", "вσт"),
		    				linebot.NewMessageTemplateAction("『ᖼOᗱᗴℕ』", "ℓ๐νә"),
		    				linebot.NewMessageTemplateAction("ϻఠ_ఠsɛɳ", "вσт"),
		    			),
		    			LineTemplate_CarouselColumn_feedback,
		    		)
		    		no_temp_msg := "ℓ๐νә"
		    		obj_message := linebot.NewTemplateMessage(no_temp_msg, LineTemplate_test)
		    
		    		obj_message_map := linebot.NewLocationMessage("ℓ๐νә", "ℓ๐νә", 25.007408,121.537688) //台北市信義區富陽街46號
		    		if _, err = bot.ReplyMessage(event.ReplyToken, obj_message, obj_message_map).Do(); err != nil {
		    			log.Print(1876)
		    			log.Print(err)
		    		}
		    		return
				} else if "s5"== message.Text {
					bot_msg = "你找我的製造者？OK！\n我跟你講我的夥伴喵在哪，你去加他。\n他跟製造者很親近的，跟他說的話製造者都會看到。\nhttps://line.me/R/ti/p/%40uwk0684z\n\n\n你也可以從下面這個連結直接去找開發者線上對話。\n\n如果他不在線上一樣可以留言給他，\n他會收到的！\n這跟手機、電腦桌面軟體都有同步連線。" +
					"\n\nhttp://www.smartsuppchat.com/widget?key=77b943aeaffa11a51bb483a816f552c70e322417&vid=" + target_id_code +
					"&lang=tw&pageTitle=%E9%80%99%E6%98%AF%E4%BE%86%E8%87%AA%20LINE%40%20%E9%80%B2%E4%BE%86%E7%9A%84%E5%8D%B3%E6%99%82%E9%80%9A%E8%A8%8A"
					log.Print("觸發找製造者")
					template := linebot.NewCarouselTemplate(
						linebot.NewCarouselColumn(
							SystemImageURL, "開發者相關資訊", "你可以透過此功能\n聯絡 開發者",
							LineTemplate_addme,
							LineTemplate_chat,
							linebot.NewPostbackTemplateAction("聯絡 LINE 機器人開發者", "開發者", "開發者"),
						),
					)
					obj_message := linebot.NewTemplateMessage("上面這些都是聯絡開發者的相關方法。", template)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(bot_msg),obj_message).Do(); err != nil {
						log.Print(1672)
						log.Print(err)
					}
					//HttpPost_JANDI(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "yellow" , "LINE 同步：執行找開發者",target_id_code)
					//HttpPost_IFTTT(target_item + " " + user_talk + "：" + message.Text + `\n<br>` + userImageUrl + `\n<br>` + userStatus, "LINE 同步：執行找開發者",target_id_code)
					HttpPost_Zapier(target_item + " [" + user_talk + "](" + userImageUrl + ")：" + message.Text + `\n` + userStatus, "LINE 同步：執行找開發者",target_id_code,user_talk)
					return
				} else if "carousel" == message.Text {
					imageURL := baseURL + "/static/buttons/1040.jpg"
					template := linebot.NewCarouselTemplate(
						linebot.NewCarouselColumn(
							imageURL, "hoge", "fuga",
							linebot.NewURITemplateAction("Go to line.me", "https://line.me"),
							linebot.NewPostbackTemplateAction("Say hello1", "hello こんにちは", ""),
						),
						linebot.NewCarouselColumn(
							imageURL, "hoge", "fuga",
							linebot.NewPostbackTemplateAction("言 hello2", "hello こんにちは", "hello こんにちは"),
							linebot.NewMessageTemplateAction("Say message", "Rice=米"),
						),
					)
					if _, err := bot.ReplyMessage(
						replyToken,
						linebot.NewTemplateMessage("Carousel alt text", template),
					).Do(); err != nil {
						log.Print(err)
					}
				} else if "imagemap" == message.Text {
					if _, err := bot.ReplyMessage(
						replyToken,
						linebot.NewImagemapMessage(
							baseURL + "/static/rich",
							"Imagemap alt text",{
							linebot.ImagemapBaseSize{1040, 1040},
							linebot.NewURIImagemapAction("https://store.line.me/family/manga/en", linebot.ImagemapArea{0, 0, 520, 520}),
							linebot.NewURIImagemapAction("https://store.line.me/family/music/en", linebot.ImagemapArea{520, 0, 520, 520}),
							linebot.NewURIImagemapAction("https://store.line.me/family/play/en", linebot.ImagemapArea{0, 520, 520, 520}),
							linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{520, 520, 520, 520}),
						),
					).Do(); err != nil {
						log.Print(err)
					}
				} else if "/bye" == message.Text {
					if rand.Intn(100) > 70 {
						bot.ReplyMessage(replyToken, linebot.NewTextMessage("請神容易送神難, 我偏不要, 嘿嘿")).Do()
					} else {
						switch source.Type {
						case linebot.EventSourceTypeUser:
							bot.ReplyMessage(replyToken, linebot.NewTextMessage("我想走, 但是我走不了...")).Do()
						case linebot.EventSourceTypeGroup:
							bot.ReplyMessage(replyToken, linebot.NewTextMessage("我揮一揮衣袖 不帶走一片雲彩")).Do()
							bot.LeaveGroup(source.GroupID).Do()
						case linebot.EventSourceTypeRoom:
							bot.ReplyMessage(replyToken, linebot.NewTextMessage("我揮一揮衣袖 不帶走一片雲彩")).Do()
							bot.LeaveRoom(source.RoomID).Do()
						}
					}
				} else if "無恥" == message.Text {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_ReplyCurseMessage[rand.Intn(len(answers_ReplyCurseMessage))])).Do()
				} else if silentMap[sourceId] != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_TextMessage[rand.Intn(len(answers_TextMessage))])).Do()
				}
			case *linebot.ImageMessage :
				log.Print("ReplyToken[" + replyToken + "] ImageMessage[" + message.ID + "] OriginalContentURL(" + message.OriginalContentURL + "), PreviewImageURL(" + message.PreviewImageURL + ")" )
				if silent != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_ImageMessage[rand.Intn(len(answers_ImageMessage))])).Do()
				}
			case *linebot.VideoMessage :
				log.Print("ReplyToken[" + replyToken + "] VideoMessage[" + message.ID + "] OriginalContentURL(" + message.OriginalContentURL + "), PreviewImageURL(" + message.PreviewImageURL + ")" )
				if silent != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_VideoMessage[rand.Intn(len(answers_VideoMessage))])).Do()
				}
			case *linebot.AudioMessage :
				log.Print("ReplyToken[" + replyToken + "] AudioMessage[" + message.ID + "] OriginalContentURL(" + message.OriginalContentURL + "), Duration(" + strconv.Itoa(message.Duration) + ")" )
				if silent != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_AudioMessage[rand.Intn(len(answers_AudioMessage))])).Do()
				}
			case *linebot.LocationMessage:
				log.Print("ReplyToken[" + replyToken + "] LocationMessage[" + message.ID + "] Title(" + message.Title  + "), Address(" + message.Address + "), Latitude(" + strconv.FormatFloat(message.Latitude, 'f', -1, 64) + "), Longitude(" + strconv.FormatFloat(message.Longitude, 'f', -1, 64) + ")" )
				if silent != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_LocationMessage[rand.Intn(len(answers_LocationMessage))])).Do()
				}
			case *linebot.StickerMessage :
				log.Print("ReplyToken[" + replyToken + "] StickerMessage[" + message.ID + "] PackageID(" + message.PackageID + "), StickerID(" + message.StickerID + ")" )
				if silent != true {
					bot.ReplyMessage(replyToken, linebot.NewTextMessage(answers_StickerMessage[rand.Intn(len(answers_StickerMessage))])).Do()
				}
			}
		} else if event.Type == linebot.EventTypePostback {
		} else if event.Type == linebot.EventTypeBeacon {
		}
	}
	
}
