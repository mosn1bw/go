package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	imageURL = "https://avatars1.githubusercontent.com/u/6083986"
)

type config struct {
	HostPort string `default:":8080"`
	IsDebug  bool   `env:"DEBUG"`
	Secret   string `default:"xx"`
	Token    string `default:"xx"`
	Tipshost string `default:"https://bowwow.tips"`
	Timeout  time.Duration
}

func random(min, mac int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(mac-min) + min
}

func main() {
	cfg := new(config)
	err := env.Fill(cfg)

	cfg.Secret = os.Getenv("LineSecret")
	cfg.Token = os.Getenv("LineToken")

	fmt.Println("Port:", cfg.HostPort)
	fmt.Println("Secret:", cfg.Secret)
	fmt.Println("Token:", cfg.Token)
	fmt.Println("Tipshost:", cfg.Tipshost)

	cl, err := linebot.New(
		cfg.Secret,
		cfg.Token,
	)
	if err != nil {
		log.Fatal(err)
	}

	template := linebot.NewCarouselTemplate(
		linebot.NewCarouselColumn(
			imageURL, "bowwow.tips", "bowwow lin",
			linebot.NewURIAction("bowwow", cfg.Tipshost),
		),
	)
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		res, err := cl.ParseRequest(req)
		if err != nil {
			log.Fatal(err)
		}
		for _, re := range res {
			if Type == linebot.EventTypeJoin {
				cl.ReplyMessage(ReplyToken, linebot.NewTextMessage("早上好！\n 請打「help」")).Do()
			}
			if Type == linebot.EventTypeFollow {
				n := time.Now()
				NowT := timeutil.Strftime(&n, "%Y年%m月%d日%H時%M分%S秒")
				p, _ := cl.GetProfile(Source.UserID).Do()
				cl.ReplyMessage(ReplyToken, linebot.NewTextMessage("謝謝你！\n"+p.DisplayName+"呃\n\n"+NowT)).Do()
				log.Println("DisplayName:" + p.DisplayName)
			}
			if Type == linebot.EventTypeMessage {
				switch msg := Message.(type) {
				case *linebot.TextMessage:
					if msg.Text == "test" {
						cl.ReplyMessage(ReplyToken, linebot.NewTextMessage("success")).Do()
					} else if msg.Text == "groupid" {
						cl.ReplyMessage(ReplyToken, linebot.NewTextMessage(string(Source.GroupID))).Do()
					} else if msg.Text == "byebye" {
						cl.ReplyMessage(ReplyToken, linebot.NewStickerMessage("3", "187")).Do()
						_, err := cl.LeaveGroup(Source.GroupID).Do()
						if err != nil {
							cl.LeaveRoom(Source.RoomID).Do()
						}
					} else if msg.Text == "help" {
						cl.ReplyMessage(ReplyToken, linebot.NewTextMessage("help\n・[image:画像url]=從圖片網址發送圖片\n・[speed]=測回話速度\n・[groupid]=發送GroupID\n・[roomid]=發送RoomID\n・[byebye]=取消訂閱\n・[about]=作者\n・[me]=發送發件人信息\n・[test]=test bowwow是否正常\n・[now]=現在時間\n・[mid]=mid\n・[sticker]=隨機圖片\n\n[其他機能]\n位置測試\n捉貼圖ID\n加入時發送消息")).Do()
					} else if msg.Text == "check" {
						fmt.Println(msg)
					} else if msg.Text == "now" {
						n := time.Now()
						NowT := timeutil.Strftime(&n, "%Y年%m月%d日%H時%M分%S秒")
						cl.ReplyMessage(ReplyToken, linebot.NewTextMessage(NowT)).Do()
					} else if msg.Text == "mid" {
						cl.ReplyMessage(ReplyToken, linebot.NewTextMessage(Source.UserID)).Do()
					} else if msg.Text == "roomid" {
						cl.ReplyMessage(ReplyToken, linebot.NewTextMessage(Source.RoomID)).Do()
					} else if msg.Text == "hidden" {
						cl.ReplyMessage(ReplyToken, linebot.NewTextMessage("hidden")).Do()
					} else if msg.Text == "bowwow" {
						_, err := cl.ReplyMessage(ReplyToken, linebot.NewImageMessage(imageURL, imageURL)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "sticker" {
						stid := random(180, 259)
						stidx := strconv.Itoa(stid)
						_, err := cl.ReplyMessage(ReplyToken, linebot.NewStickerMessage("3", stidx)).Do()
						if err != nil {
							log.Fatal(err)
						}
					} else if msg.Text == "1" {
					        bot.ReplyMessage(replyToken, linebot.NewTextMessage("💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9..0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.\n\n💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀"), linebot.NewTextMessage("7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7."), linebot.NewTextMessage("7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7."), linebot.NewTextMessage("7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7.7W0.G7.W0.G7W0.G7.W0.G7W0.G7.W0.G7W0.G7."), linebot.NewTextMessage("💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J\n\n─═≡ϻఠ_ఠsɛɳ≡═─\n\n.1.2.3.4.5.6.7.8.9.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9..0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.8.J.9.K.0.A.1.B.2.D.3.E.4.F.5.G.6.H.7.I.\n\n💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀💀")).Do()
						
					} else if msg.Text == "me" {
						mid := Source.UserID
						p, err := cl.GetProfile(mid).Do()
						if err != nil {
							cl.ReplyMessage(ReplyToken, linebot.NewTextMessage("新增同意"))
						}

						cl.ReplyMessage(ReplyToken, linebot.NewTextMessage("mid:"+mid+"\nname:"+p.DisplayName+"\nstatusMessage:"+p.StatusMessage)).Do()
					} else if msg.Text == "speed" {
						replytoken := ReplyToken
						start := time.Now()
						cl.ReplyMessage(replytoken, linebot.NewTextMessage("..")).Do()
						end := time.Now()
						result := fmt.Sprintf("%f [sec]", (end.Sub(start)).Seconds())
						_, err := cl.PushMessage(Source.GroupID, linebot.NewTextMessage(result)).Do()
						if err != nil {
							_, err := cl.PushMessage(Source.RoomID, linebot.NewTextMessage(result)).Do()
							if err != nil {
								_, err := cl.PushMessage(Source.UserID, linebot.NewTextMessage(result)).Do()
								if err != nil {
									log.Fatal(err)
								}
							}
						}
					} else if res := strings.Contains(msg.Text, "hello"); res == true {
						cl.ReplyMessage(ReplyToken, linebot.NewTextMessage("hello!"), linebot.NewTextMessage("my name is bowwow")).Do()
					} else if res := strings.Contains(msg.Text, "image:"); res == true {
						image_url := strings.Replace(msg.Text, "image:", "", -1)
						cl.ReplyMessage(ReplyToken, linebot.NewImageMessage(image_url, image_url)).Do()
					} else if msg.Text == "about" {
						_, err := cl.ReplyMessage(ReplyToken, linebot.NewTemplateMessage("hi", template)).Do()
						if err != nil {
							log.Println(err)
						}
					}
				case *linebot.StickerMessage:
					cl.ReplyMessage(ReplyToken, linebot.NewTextMessage("StickerId:"+msg.StickerID+"\nPackageId:"+msg.PackageID)).Do()
				case *linebot.LocationMessage:
					cl.ReplyMessage(ReplyToken, linebot.NewLocationMessage(
						msg.Title,
						msg.Address,
						msg.Latitude,
						msg.Longitude)).Do()
				}
			}
		}
	})
	if err := http.ListenAndServe(cfg.HostPort, nil); err != nil {
		log.Fatal(err)
	}
}
