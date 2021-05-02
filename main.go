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
var user_mosen = "ub5e4ae027d8d4a82736222b2a8dc77df"
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
		"人被殺，就會死。",
		"當別人贏過你時，你就輸了！",
		"研究指出日本人的母語是日語",
		"你知道嗎 當你背對太陽 你就看不見金星",
		"當你失眠的時候，你就會睡不著",
		"今天是昨天的明天。",
		"吃得苦中苦，那一口特別苦",
	}
var answers_ImageMessage = []string{
		"傳這甚麼廢圖？你有認真在分享嗎？",
	}
var answers_StickerMessage = []string{
		"腳踏實地打字好嗎？傳這甚麼貼圖？",
	}
var answers_VideoMessage = []string{
		"看甚麼影片，不知道我的流量快用光了嗎？",
	}
var answers_AudioMessage = []string{
		"說的比唱的好聽，唱得鬼哭神號，是要嚇唬誰？",
	}
var answers_LocationMessage = []string{
		"這是哪裡啊？火星嗎？",
	}
var answers_ReplyCurseMessage = []string{
		"真的無恥",
		"有夠無恥",
		"超級無恥",
		"就是無恥",
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
		log.Println("現在時間(台北): " + nowString)
		bot.ReplyMessage(replyToken, linebot.NewTextMessage("現在時間(台北): " + nowString)).Do()
	} else if silent != true {
		log.Println("自動報時(台北): " + nowString)
		bot.PushMessage(replyToken, linebot.NewTextMessage("自動報時(台北): " + nowString)).Do()
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
				
				if source.UserID != "" && source.UserID != user_zchien {
					profile, err := bot.GetProfile(source.UserID).Do()
					if err != nil {
						log.Print(err)
					} else if _, err := bot.PushMessage(user_zchien, linebot.NewTextMessage(profile.DisplayName + ": "+message.Text)).Do(); err != nil {
							log.Print(err)
					}
				}
				
    elif text == '1':	
        line_bot_api.reply_message(
        event.reply_token, [
        TextSendMessage(text='ᖼ.O.ᗱ.ᗴ.ℕ.♡.~.☆.💖.💔.💙.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.ᖼ.O.ᗱ.ᗴ.ℕ.♡.~.☆.💖.💔.💙'),
        TextSendMessage(text='ᖼ.O.ᗱ.ᗴ.ℕ.♡.~.☆.💖.💔.💙.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.ᖼ.O.ᗱ.ᗴ.ℕ.♡.~.☆.💖.💔.💙'),
        TextSendMessage(text='ᖼ.O.ᗱ.ᗴ.ℕ.♡.~.☆.💖.💔.💙.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.ᖼ.O.ᗱ.ᗴ.ℕ.♡.~.☆.💖.💔.💙'),
        TextSendMessage(text='ᖼ.O.ᗱ.ᗴ.ℕ.♡.~.☆.💖.💔.💙.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.ᖼ.O.ᗱ.ᗴ.ℕ.♡.~.☆.💖.💔.💙'),
        TextSendMessage(text='ᖼ.O.ᗱ.ᗴ.ℕ.♡.~.☆.💖.💔.💙.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.2.2.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.2.4.1.6.0.5.4.1.4.0.2.0.3.4.1.6.0.3.4.1.4.3.0.3.4.1.6.0.3.4.1.4.3.0.2.4.1.6.0.5.4.1.4.3.0.ᖼ.O.ᗱ.ᗴ.ℕ.♡.~.☆.💖.💔.💙')
        ])

				a=rand.Intn(8)
				if a == 0 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有一個國家舉辦最噁心比賽，至最後決賽時剩下三個人爭冠，其中一個人拿了一杯尿，在嘴中漱了漱，吞下，哈的一聲，全場鼓掌想冠軍必落於此家。第二個人從袋中拿出一堆蟑螂，剝了翅膀就嘖嘖嘖的吸牠的肚子，不時還吐出一兩隻腳,吃到第十隻的時後,國王面有菜色的說不用吃了你這樣就第一名了。 此時見第三個人拿出一杯液體，說，這是我半年前感冒到現在，每次吐的痰都收集在裡面，我現在要把它喝完。只見一整杯白白黃黃還帶泡泡的，他搖了搖，試圖讓有些積太久快要凝固的化開， 國王眼淚都要掉下來了，說：不用了不用了你只要喝一口你就冠軍了～這人便拿起杯子咕嘟咕嘟地開始喝，因為很濃又很多過了五分多鐘才喝完，此時全場已淚流滿面，國王說幹嘛我不是叫你喝一口就冠軍了嗎？這人回答道，我也只是想喝一口，但是我一直咬不斷～～")).Do()}
				if a == 1 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有人很喜歡“麻辣粉絲煲”這道菜。有一次，他上飯館，又點了這道菜。但侍者告訴他，這道菜已經賣完了。“真的賣完了嗎？”他很失望地問。“先生，真的賣完了。你瞧，最後一份賣給那桌的先生了。”侍者回答道。那人順著侍者的指點，看見有個很體面的紳士坐在鄰座。紳士的飯菜已經吃得差不多了，但那份“麻辣粉絲煲”居然還是滿滿的。那人覺得紳士很浪費美味，所以他走到紳士旁邊，指著那份“麻辣粉絲煲”，很有禮貌地問：“先生，您這還要嗎？”紳士很有風度地搖搖頭。于是那人立刻坐下，拿起調羹狼吞虎咽起來。風卷殘雲，一會兒一半下肚了，突然間他發現在砂鍋底躺著一只很小很小但皮毛已長全的小老鼠。一陣惡心，那人把吃下去的所有粉絲通通吐回了砂鍋裏。當他在那兒翻胃不已的時候，那紳士用很同情的眼光看著他，說：“很惡心是嗎？剛才我也是這樣……”")).Do()}
				if a == 2 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這天，酒店老板正在大廳巡視。來了一乞丐上前說道：”老板給個牙簽行嗎？”老板給他一個打發走了。一會兒，又來一個乞丐，也是來要牙簽的。老板心想現在這乞丐怎麽不要飯改要牙簽了？也同樣給他一個打發走了，沒過多久，又來一個乞丐。老板對他說：”你也是來要牙簽的嗎？”乞丐說：”有個人吐了，可我晚了一步，已經被前面兩個乞丐把能吃的都吃了，現在只剩下湯了。你能給我個吸管嗎？”")).Do()}
				if a == 7 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你呀")).Do()}
				}
				if strings.Contains(message.Text,"1,") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("1",strings.Trim(message.Text,"1,"))).Do()}
				if strings.Contains(message.Text,"4,") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("4",strings.Trim(message.Text,"4,"))).Do()}
				if strings.Contains(message.Text,"測試") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("2","18")).Do()}

				if strings.Contains(message.Text, "a1") {
					tellTime(replyToken, true)
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("QQ")).Do()
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("QQ")).Do()
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("QQ")).Do()
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("QQ")).Do()
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("QQ")).Do()
				} else if strings.Contains(message.Text, "a2") {
					tellTime(replyToken, true)
				} else if "a3" == message.Text {
					tellTime(replyToken, true)
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("請神容易送神難, 我偏不要, 嘿嘿")).Do()
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("麥克風測試，1、2、3... OK")).Do()
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("麥克風測試，1、2、3... OK")).Do()
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("麥克風測試，1、2、3... OK")).Do()
					bot.ReplyMessage(replyToken, linebot.NewTextMessage("Bot can't use profile API without user ID")).Do()
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
							"Imagemap alt text",
							linebot.ImagemapBaseSize{1040, 1040},
							linebot.NewURIImagemapAction("https://store.line.me/family/manga/en", linebot.ImagemapArea{0, 0, 520, 520}),
							linebot.NewURIImagemapAction("https://store.line.me/family/music/en", linebot.ImagemapArea{520, 0, 520, 520}),
							linebot.NewURIImagemapAction("https://store.line.me/family/play/en", linebot.ImagemapArea{0, 520, 520, 520}),
							linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{520, 520, 520, 520}),
						),
					).Do(); err != nil {
						log.Print(err)
					}
				} else if "a5" == message.Text {
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
				} else if "a4" == message.Text {
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
