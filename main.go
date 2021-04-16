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

package main
import (
    	"bufio"
	"fmt"
	"log"
    	"io"
    	"time"
	"net/http"
	"os"
	"strings"
	"math/rand"
	"io/ioutil"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var admin string
var url string
var conn string
var ppljoin string
var b_day string
var join_msg string

func main() {
	var err error
	
	admin = "U83bb64e03c849e6ed897f9c556b0d4c1"
	url = "https://raw.githubusercontent.com/Yikaros/LineBotTemplate/master/images/"
	conn = ""
	ppljoin = ""
	b_day = ""
	join_msg = ""
	
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)
	
	var list string
	var price string
	var stock string
	var food string
	var uid string
	var list_array []string
	var user_array []string
	var av_array []string
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	
	fi, err := os.Open("buffer/list.txt")
    	if err != nil {
        	fmt.Printf("Error: %s\n", err)
        	return
    	}
    	defer fi.Close()

    	br := bufio.NewReader(fi)
    	for {
        	a, _, c := br.ReadLine()
        	if c == io.EOF {
            	break
        	}
		list = list + "&" + string(a)
    	}
	list_array = strings.Split(list, "&")
	fi2, err2 := os.Open("buffer/userlist.txt")
    	if err2 != nil {
        	fmt.Printf("Error: %s\n", err2)
        	return
    	}
    	defer fi2.Close()
	list = ""
    	br2 := bufio.NewReader(fi2)
    	for {
        	a, _, c := br2.ReadLine()
        	if c == io.EOF {
            	break
        	}
		list = list + "@#@" + string(a)
    	}
	user_array = strings.Split(list, "@#@")
	
	
	fi3, err3 := os.Open("buffer/LoveLove.txt")
    	if err3 != nil {
        	fmt.Printf("Error: %s\n", err3)
        	return
    	}
    	defer fi3.Close()
	list = ""
    	br3 := bufio.NewReader(fi3)
    	for {
        	a, _, c := br3.ReadLine()
        	if c == io.EOF {
            	break
        	}
		list = list + "@#@" + string(a)
    	}
	av_array = strings.Split(list, "@#@")
	
	
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			var cow string
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			cow = url + "cow/" + fmt.Sprintf("%d", r.Intn(11))  + ".jpg"

			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				//抓user id跟資料
				// 0 ID
				// 1 客戶代號
				// 2 姓名
				// 3 生日
				// 4 喜好
				// 5 電話
				// 6 通訊狀態
				uid = event.Source.UserID
				e:=0
				var profile []string
				for e<=len(user_array)-1{
					var menu []string
					menu = strings.Split(user_array[e], " & ")
					if menu[0] == uid{
						profile = strings.Split(user_array[e], " & ")
						break
					}
					e++
				}
				if ((conn!="")&&((uid==admin)||(uid==conn))){
					switch{
						case uid==admin:
							if message.Text=="/bye"{
								for e<=len(user_array)-1{
									var menu []string
									menu = strings.Split(user_array[e], " & ")
									if menu[0] == conn{
										profile = strings.Split(user_array[e], " & ")
										break
									}
									e++
								}
								conn = ""
								bot.PushMessage(admin,linebot.NewTextMessage("已與"+profile[2]+"離線")).Do() 
								bot.PushMessage(conn,linebot.NewTextMessage("跟老闆結束通話囉")).Do() 
							}else{
								bot.PushMessage(conn,linebot.NewTextMessage(message.Text)).Do() 
							}
						case uid==conn:
							bot.PushMessage(admin,linebot.NewTextMessage(profile[2] + "說：" + message.Text)).Do() 
					}
				}else
				{
					switch {
						case message.Text=="s1":
							d1 := []byte("hello\ngo\n")
							err := ioutil.WriteFile("buffer/test", d1, 0644)
						    	if err != nil {
								fmt.Printf("Error: %s\n", err)
								return
							}
						case Contains(message.Text,"s2"):
							template := linebot.NewConfirmTemplate("大家覺得艾魯是不是該刪掉刀塔?", linebot.NewMessageTemplateAction("是", "是"),linebot.NewMessageTemplateAction("4", "4"))
							messgage := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
							bot.ReplyMessage(event.ReplyToken, messgage).Do()
						case Contains(message.Text,"s3"):
							//影片總數
							av_count := 9
							var av_rnd [4]int
							
							av_rnd[0] = r.Intn(av_count)+1
							s:=1
							for s<=3{
								av_rnd[s]=0
								rnd:=r.Intn(av_count)+1
								sc:=0
								for sc <= s-1{
									if av_rnd[sc]==rnd{
										av_rnd[s]=-1
										break
									}
									sc++
								}
								if av_rnd[s]==0{
									av_rnd[s] = rnd
									s++
								}
							}
							
							//隨機取三
							var av0 []string
							av0 = strings.Split(av_array[av_rnd[0]], "@@")
							var av1 []string
							av1 = strings.Split(av_array[av_rnd[1]], "@@")
							var av2 []string
							av2 = strings.Split(av_array[av_rnd[2]], "@@")
							
						/** debug用
							var avtest [3]string
							avtest[0]=fmt.Sprintf("%d", av_rnd[0])
							avtest[1]=fmt.Sprintf("%d", av_rnd[1])
							avtest[2]=fmt.Sprintf("%d", av_rnd[2])
							alltest := avtest[0]+avtest[1]+avtest[2]
						
							bot.PushMessage(admin, linebot.NewTextMessage(alltest)).Do() 
						*/
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									url+"/LoveLove/" + av0[1], av0[2], av0[3],
									linebot.NewURITemplateAction("我要幹死他!!", av0[4]),
								),
								linebot.NewCarouselColumn(
									url+"/LoveLove/" + av1[1], av1[2], av1[3],
									linebot.NewURITemplateAction("我要幹死他!!", av1[4]),
								),
								linebot.NewCarouselColumn(
									url+"/LoveLove/" + av2[1], av2[2], av2[3],
									linebot.NewURITemplateAction("我要幹死他!!", av2[4]),
								),
							)
							messgage := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
							bot.ReplyMessage(event.ReplyToken, messgage).Do() 
						case Contains(message.Text,"a9")||Contains(message.Text,"菜單")||message.Text=="我想買菜":
							
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									url + "/vegetable/gau.jpg", "高麗菜", "一斤50元，起源韓國，中國製造",
									linebot.NewMessageTemplateAction("我要買這個!!", "我要買高麗菜"),
								),
								linebot.NewCarouselColumn(
									url + "/vegetable/hua.jpg", "花椰菜", "一斤55元，上面有蟲的部分最好吃",
									linebot.NewMessageTemplateAction("我要買這個!!", "我要買花椰菜"),
								),
							)
							messgage := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
							bot.ReplyMessage(event.ReplyToken, messgage).Do() 
						case Contains(message.Text,"a8")||message.Text=="我想買魚":
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									url + "https://lh3.googleusercontent.com/-wV-K_ZyeaXE/YHdNYUIN1aI/AAAAAAAABIE/Uq7xL6i_CkYhrkb_23kuGFqNHJ13KF7vwCK8BGAsYHg/s512/2021-04-14.jpg", "reza", "niki",
									linebot.NewMessageTemplateAction("我要買這個!!", "我要買青斑"),
								),
								linebot.NewCarouselColumn(
									url + "/fish/dragontiger.jpg", "龍虎斑", "一斤350元，Q彈鮮嫩，最適合挑嘴的老饕",
									linebot.NewMessageTemplateAction("我要買這個!!", "我要買龍虎斑"),
								),
							)
							messgage := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
							bot.ReplyMessage(event.ReplyToken, messgage).Do() 
						case Contains(message.Text,"幫我查ID")||Contains(message.Text,"幫我查id"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(uid)).Do() 
						case Contains(message.Text,"幫我查群組ID")||Contains(message.Text,"幫我查群組id"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(event.Source.GroupID)).Do() 
						case ppljoin != "":
							join_msg = ppljoin + " " + message.Text
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text + "嗎? 好的，那請問您的生日是幾月幾號呢?")).Do() 
							ppljoin = ""
							b_day = "1"
						case b_day != "":
							join_msg = join_msg + " " + message.Text
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text + "嗎? 好，我幫你傳訊息給菜市場管理員請他幫你審核一下，請耐心等候喔，謝謝")).Do() 
							bot.PushMessage(admin,linebot.NewTextMessage(join_msg)).Do() 
							b_day = ""
							join_msg = ""
						case message.Text=="a7":
							if len(profile) > 0{
								bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(profile[2] + "你已經是菜市場的會員囉，不用再申請加入啦")).Do()
								bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(profile[2] + "你已經是菜市場的會員囉，不用再申請加入啦")).Do()
								bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(profile[2] + "你已經是菜市場的會員囉，不用再申請加入啦")).Do()
								bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(profile[2] + "你已經是菜市場的會員囉，不用再申請加入啦")).Do()
							}else{
								ppljoin = uid
								bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要加入Protoss菜市場嗎? 請問您叫什麼名字呢?")).Do() 
							}
	//賣菜的code
						case Contains(message.Text,"a6")||Contains(message.Text,"葉"):						
							food = ""
							switch{
								case Contains(message.Text,"高麗菜"):
									food = "高麗菜"
								case Contains(message.Text,"小白菜"):
									food = "小白菜"
								case Contains(message.Text,"花椰菜"):
									food = "花椰菜"
								case Contains(message.Text,"地瓜葉"):
									food = "地瓜葉"

							}
							if food != ""{
								i:=0
								for i<=len(list_array){
									var menu []string
									menu = strings.Split(list_array[i], " ")
									if menu[0] == food{
										price=menu[1]
										stock=menu[2]
										break
									}
									i++
								}
								switch{
									case Contains(message.Text,"a5")||Contains(message.Text,"多少錢")||Contains(message.Text,"怎麼賣")||Contains(message.Text,"怎麼算"):
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "一斤" + price)).Do() 
									case Contains(message.Text,"a4"):
										/**if len(profile) > 0{
											bot.PushMessage(admin,linebot.NewTextMessage(profile[1] + profile[2] + "要買" + food)).Do() 
											bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "嗎? 我已經幫你聯絡老闆了，晚點他就會主動跟你聯繫，請耐心等一下喔")).Do()
										}else{
											bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買菜嗎? 可是你好像還不是我們菜市場的會員捏，麻煩跟管理員聯繫幫你加入菜市場會員，會員才有特別優惠喔!!")).Do() 	
										}**/
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買" + food + "嗎? 趕快撥打 0909 022 890 跟鄉下小孩-張耀東買菜喔!!")).Do()
								}
							}else{
								if Contains(message.Text,"a3"){
									/**if len(profile) > 0{
										bot.PushMessage(admin,linebot.NewTextMessage(profile[1] + profile[2] + "要買菜")).Do() 
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買菜嗎? 我已經幫你聯絡老闆了，晚點他就會主動跟你聯繫，請耐心等一下喔")).Do() 	
									}else{
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買菜嗎? 可是你好像還不是我們菜市場的會員捏，麻煩跟管理員聯繫幫你加入菜市場會員，會員才有特別優惠喔!!")).Do() 	
									}**/
									bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買菜嗎? 趕快撥打 0909 022 890 跟鄉下小孩-張耀東買菜喔!!")).Do()
								}
							}
	//石斑魚的code
						case Contains(message.Text,"a1")||Contains(message.Text,"班"):
							food = ""
							switch{
								case Contains(message.Text,"a2"):
									food = "龍虎石斑"
								case Contains(message.Text,"青"):
									food = "青斑"
								case Contains(message.Text,"珍珠"):
									food = "珍珠石斑"
								default:
									if Contains(message.Text,"一斤多少")||Contains(message.Text,"多少錢")||Contains(message.Text,"怎麼賣")||Contains(message.Text,"怎麼算")||Contains(message.Text,"還有多少")||Contains(message.Text,"剩下多少")||Contains(message.Text,"庫存")||Contains(message.Text,"還有幾"){
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("拍謝啦! 我是笨笨的電腦，不知道您要問哪種石斑捏，我們有龍虎石斑、青斑、還有珍珠石斑")).Do()
									}
							}
							if food != ""{
								i:=0
								for i<=len(list_array){
									var menu []string
									menu = strings.Split(list_array[i], " ")
									if menu[0] == food{
										price=menu[1]
										stock=menu[2]
										break
									}
									i++
								}
								switch{
									case Contains(message.Text,"w5")||Contains(message.Text,"多少錢")||Contains(message.Text,"怎麼賣")||Contains(message.Text,"怎麼算"):
										bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg","https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg")).Do()}
									case Contains(message.Text,"w6")||Contains(message.Text,"剩下多少")||Contains(message.Text,"庫存")||Contains(message.Text,"還有幾"):
										bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("htt666666644444444444444446666661111")).Do()}
									case Contains(message.Text,"w7"):/**
										if len(profile) > 0{
											bot.PushMessage(admin,linebot.NewTextMessage(profile[1] + profile[2] + "要買" + food)).Do() 
											bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "嗎? 我已經幫你聯絡老闆了，晚點他就會主動跟你聯繫，請耐心等一下喔")).Do()
										}else{
											bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買魚嗎? 可是你好像還不是我們菜市場的會員捏，麻煩跟管理員聯繫幫你加入菜市場會員，會員才有特別優惠喔!!")).Do() 	
										}**/
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買石斑嗎? 趕快撥打 0939 220 743 跟石斑膠膠哥-黃大哥買魚喔!!")).Do()
								}
							}else{
								if Contains(message.Text,"w4"){/**
									if len(profile) > 0{
										bot.PushMessage(admin,linebot.NewTextMessage(profile[1] + profile[2] + "要買菜")).Do() 
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買魚嗎? 我已經幫你聯絡老闆了，晚點他就會主動跟你聯繫，請耐心等一下喔")).Do() 	
									}else{
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買魚嗎? 可是你好像還不是我們菜市場的會員捏，麻煩跟管理員聯繫幫你加入菜市場會員，會員才有特別優惠喔!!")).Do() 	
									}**/
									bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買石斑嗎? 趕快撥打 0939 220 743 跟石斑膠膠哥-黃大哥買魚喔!!")).Do()
								}
							}
	//指令集
						case (uid==admin)&&Contains(message.Text,"/join "):
							s:=0
							var profile2 []string
							for s<=len(user_array)-1{
								var menu []string
								menu = strings.Split(user_array[e], " & ")
								if menu[0] == strings.Replace(message.Text, "/join ", "", -1){
									profile2 = strings.Split(user_array[e], " & ")
									break
								}
								e++
							}
							bot.PushMessage(strings.Replace(message.Text, "/join ", "", -1),linebot.NewTextMessage(profile2[2] + "你好，已經幫你加入菜市場會員了，你現在可以開始買菜囉!!")).Do() 
						case (uid==admin)&&Contains(message.Text,"/nojoin "):
							bot.PushMessage(strings.Replace(message.Text, "/nojoin ", "", -1),linebot.NewTextMessage("經過我們的審核，你輸入的資料好像有點問題耶，可以請你重新申請一次嗎? 直接傳訊息說 我要加入 就可以了")).Do() 
						case (uid==admin)&&Contains(message.Text,"/w "):
							conn = strings.Replace(message.Text, "/w ", "", -1)
							for e<=len(user_array){
								var menu []string
								menu = strings.Split(user_array[e], " & ")
								if menu[1] == conn{
									profile = strings.Split(user_array[e], " & ")
									conn = profile[0]
									break
								}
								e++
							}
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("已成功與" + profile[2] + "連結，可以直接傳訊息開始通訊")).Do() 
							bot.PushMessage(conn,linebot.NewTextMessage("老闆出現囉! 你現在可以跟他傳訊息了")).Do() 
	//以下是喇賽的code
						case Contains(message.Text,"w1")||Contains(message.Text,"洗眼")||Contains(message.Text,"乳牛")||Contains(message.Text,"淨化"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(cow,cow)).Do() 
						case Contains(message.Text,"w"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(url + "6569950-1490833625.jpg", url + "6569950-1490833625.jpg")).Do() 
						case Contains(message.Text,"w3"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(url + "blackman.jpg", url + "blackman.jpg")).Do() 					
					}
				}
			}

		}
	}
}

	//以下是喇賽的code
						case Contains(message.Text,"w2"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(profile[2] + "你已經是菜市場的會員囉，不用再申請加入啦")).Do()
							bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("拍謝啦! 我是笨笨的電腦，不知道您要問哪種石斑捏，我們有龍虎石斑、青斑、還有珍珠石斑")).Do()
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg","https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg")).Do()}
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/0d/30/571d8eeac925c21c771548e75630.jpg","https://www.itsfun.com.tw/cacheimg/0d/30/571d8eeac925c21c771548e75630.jpg")).Do()}
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg","https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg")).Do()}
					}
				}
			}

		}
	}
}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
					case "m1": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好!我是自動回覆的機器人,在CoC台灣英雄聯盟為您服務")).Do()
					case "m2": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我是老大的代理機器人,請問有什麼事呢?")).Do()
					case "m3": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大姓張不姓謝喔")).Do()
					case "m4": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("NG的啦")).Do()
					case "m5": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("網帥")).Do()
					case "m6": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("天龍國人")).Do()
					case "m7": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("塔哥被黑了")).Do()
					case "m8": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("澳洲打工仔")).Do()
					case "m9": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("他高傲，但是宅心仁厚，他低調，但是受萬人景仰，他可以把神賜給人類的火，運用得出神入化，燒出堪稱火之藝術的超級菜式，他究竟是神仙的化身？還是地獄的使者？沒人知道，但是可以肯定，每個人都給他一個稱號——食~~神！——")).Do()
					case "m10": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好呀")).Do()
					case "m11": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你也早!")).Do()
					case "m12": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("嘿你好")).Do()
					case "m13": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("安安, 大家好!")).Do()
					case "m14": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("睡醒了嗎?太晚了吧")).Do()
					case "m15": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("都幾點了,老實說昨晚在幹嘛")).Do()
					case "m16": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("祝好夢")).Do()
					case "hi": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好")).Do()
					case "Hi": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好")).Do()
					case "m17": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你好")).Do()
					case "m17": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("別以為不知道你在說髒話")).Do()
					case "喔": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("喔屁喔")).Do()
					case "哦": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哦屁喔")).Do()
					case "幹": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()					
					case "老大是誰": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你是在問那個玉樹臨風的帥哥?")).Do()
					case "抱歉": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("道歉時露出X部是常識吧")).Do()
					case "對不起": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("道歉時露出X部是常識吧")).Do()
					case "傻眼": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("雨靖的口頭禪")).Do()
					case "好呀": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("一言既出 駟馬難追")).Do()
					case "感謝老大": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("讚嘆老大")).Do()
					case "誰在": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我是老大的代理機器人,請問有什麼事呢?")).Do()
					case "有誰在": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我是老大的代理機器人,請問有什麼事呢?")).Do()
					case "真可愛": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("啾咪")).Do()
					case "去洗澡": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("洗澡卡get!")).Do()
					case "@@": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("^_^")).Do()
					case "來下班": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("真的很爽耶你")).Do()
					case "下班了": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("真的很爽耶你")).Do()
					case "上班去": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("加油加油")).Do()
					case "上班了": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("上班還不是都在混")).Do()
					case "太強了": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("跪")).Do()
					case "笑點在": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你媽知道你在這裡po廢文嗎")).Do()
					case "好": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("謝謝")).Do()
					case "不好": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("為什麼有一點點難過")).Do()
					case "地震": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("媽呀, 嚇死人, 以下開始報平安!")).Do()
					case "/help": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()
					case "#1": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這你也相信喔..幾歲了")).Do()
					case "#2": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("系統建置中~~預計2050年上線")).Do()
					case "#9": bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("動動腰起床了, 生意上門喔!")).Do()
					case "m20": a=rand.Intn(4)
						if a == 0 {bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg","https://www.itsfun.com.tw/cacheimg/66/0c/eda9d251c3bd769ac820552b2ff1.jpg")).Do()}
						if a == 1 {bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/0d/30/571d8eeac925c21c771548e75630.jpg","https://www.itsfun.com.tw/cacheimg/0d/30/571d8eeac925c21c771548e75630.jpg")).Do()}
						if a == 2 {bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg","https://www.itsfun.com.tw/cacheimg/41/ac/f9e580eb9cbd38241182198bcb1b.jpg")).Do()}
						if a == 3 {bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("https://www.itsfun.com.tw/cacheimg/1a/e7/e18b42a3703133301659f6ddd7a4.jpg","https://www.itsfun.com.tw/cacheimg/1a/e7/e18b42a3703133301659f6ddd7a4.jpg")).Do()}
				}
				if strings.Contains(message.Text,"m21") {
					if strings.Contains(message.Text,"m22") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大有寬闊肩膀可以依偎")).Do()}
					if strings.Contains(message.Text,"m23") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("來來來！都來")).Do()}
					if strings.Contains(message.Text,"m24") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("來來來！都來")).Do()}
					if strings.Contains(message.Text,"最愛") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("一生只愛你一個")).Do()}
					if strings.Contains(message.Text,"不行") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你才不行")).Do()}
					if strings.Contains(message.Text,"老了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大今年18歲")).Do()}
					if strings.Contains(message.Text,"老人") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大今年18歲")).Do()}
					if strings.Contains(message.Text,"愛你") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("請排隊領取號碼牌")).Do()}
					if strings.Contains(message.Text,"你好") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("お元氣てすか？私は元氣てす")).Do()}
					if strings.Contains(message.Text,"帥") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("帥不過是一個字 卻跟著我一輩子")).Do()}
					if strings.Contains(message.Text,"在嗎") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好 請問有什麼事")).Do()}
					if strings.Contains(message.Text,"在哪") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("在這裡在那裡 在昨天在未來 原來已在你心裡")).Do()}
					if strings.Contains(message.Text,"去哪") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("去尋找未來 追求夢想")).Do()}
					if strings.Contains(message.Text,"幹嘛") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("在你後面喜滋滋地看著你")).Do()}
					if strings.Contains(message.Text,"幾歲") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("18歲")).Do()}
					if strings.Contains(message.Text,"厲害") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("鵝樂告無大計")).Do()}
					if strings.Contains(message.Text,"單身") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("曾经有一份真诚的爱情放在我面前，我没有珍惜，等我失去的时候我才后悔莫及，人世间最痛苦的事莫过于此。如果上天能够给我一个再来一次的机会，我会对那个女孩子说三个字：我爱你。如果非要在这份爱上加上一个期限，我希望是……一万年！")).Do()}
					if strings.Contains(message.Text,"閒") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("在忙也要跟你喝一杯咖啡")).Do()}
					if strings.Contains(message.Text,"快") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不要一直催老大, 一直催一直催")).Do()}
					if strings.Contains(message.Text,"公分") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("相信我, 老大30公分")).Do()}
				}
				if strings.Contains(message.Text,"m25") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("大家一起唱~祝你生日快樂~祝你生日快樂~祝你生日快樂~祝你生日快樂")).Do()} 
				if strings.Contains(message.Text,"m26") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("超級爆笑的啦")).Do()}
				if strings.Contains(message.Text,"......") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("真的很無言")).Do()}
				if strings.Contains(message.Text,"幹你") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()}
				if strings.Contains(message.Text,"幹他") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()}
				if strings.Contains(message.Text,"幹我") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ㄜ, 別說髒話啦")).Do()}
				if strings.Contains(message.Text,"吉利山莊") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有正妹")).Do()}
				if strings.Contains(message.Text,"死我") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("大家要身體健康")).Do()}
				if strings.Contains(message.Text,"幫忙捐兵") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("又來了,請大家幫忙捐兵喔")).Do()}
				if strings.Contains(message.Text,"請問") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我想這是個好問題!請各位幫忙回答喔")).Do()}
				if strings.Contains(message.Text,"換手機") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哇～好野人")).Do()}
				if strings.Contains(message.Text,"excel") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大是專家")).Do()}
				if strings.Contains(message.Text,"Excel") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大是專家")).Do()}
				if strings.Contains(message.Text,"EXCEL") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大是專家")).Do()}
				if strings.Contains(message.Text,"新年") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("祝福大家新年快樂！")).Do()}
				if strings.Contains(message.Text,"跨年") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("都沒揪")).Do()}
				if strings.Contains(message.Text,"聖誕節") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("叮叮噹叮叮噹鈴聲多響亮～祝大家聖誕節快樂！")).Do()}
				if strings.Contains(message.Text,"感恩節") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("感謝上帝賜予一年度豐收的祝福")).Do()}
				if strings.Contains(message.Text,"祝福") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我用誠摯的心給與最大的祝福")).Do()}
				if strings.Contains(message.Text,"死了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("大家要身體健康")).Do()}
				if strings.Contains(message.Text,"怎麼辦") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("涼拌炒雞蛋")).Do()}
				if strings.Contains(message.Text,"羨慕") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do()}
				if strings.Contains(message.Text,"PS4") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰有玩PS4 加個好友啦")).Do()}
				if strings.Contains(message.Text,"Ps4") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰有玩PS4 加個好友啦")).Do()}
				if strings.Contains(message.Text,"ps4") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰有玩PS4 加個好友啦")).Do()}
				if strings.Contains(message.Text,"野生") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("歐買尬!本人比較帥")).Do()}
				if strings.Contains(message.Text,"媽的") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("人是人他媽生的，妖是妖他媽生的，只要你有一顆善良的心就不再是妖，是人妖")).Do()}
				if strings.Contains(message.Text,"ㄇㄉ") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你媽貴姓？")).Do()}
				if strings.Contains(message.Text,"吃屎") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("又油又香 粉好吃")).Do()}
				if strings.Contains(message.Text,"妹妹") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("醒醒吧~你沒有妹妹")).Do()}
				if strings.Contains(message.Text,"謝謝") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不客氣的")).Do()}
				if strings.Contains(message.Text,"感謝") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不客氣的")).Do()}
				if strings.Contains(message.Text,"累了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("對不起我只是個機器人, 可以原諒我嗎")).Do()}
				if strings.Contains(message.Text,"星期一") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("每個星期一從醒來那一刻便覺得心裡有一種難以擺脫的壓抑，覺得很鬱悶。")).Do()}
				if strings.Contains(message.Text,"星期五") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("T.G.I.F.喔喔喔喔喔喔耶")).Do()}
				if strings.Contains(message.Text,"解鎖") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰解鎖了?給你拍拍手的啦")).Do()}
				if strings.Contains(message.Text,"關鍵字") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("加油喔, 還有很多慢慢找")).Do()}
				if strings.Contains(message.Text,"http") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("謝謝分享!愛你喔")).Do()}
				if strings.Contains(message.Text,"八兩金") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("新莊八兩金在此")).Do()}
				if strings.Contains(message.Text,"我朋友") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("先承認你就是你朋友吧")).Do()}
				if strings.Contains(message.Text,"我可以") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("不可以")).Do()}
				if strings.Contains(message.Text,"點名") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大 1")).Do()}
				if strings.Contains(message.Text,"打氣") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("加油!加油!加油加油加油!")).Do()}
				if strings.Contains(message.Text,"3D") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我家樓上正妹32D")).Do()}
				if strings.Contains(message.Text,"下面") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大下的麵很好吃")).Do()}
				if strings.Contains(message.Text,"百萬業務") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("靖哥哥")).Do()}
				if strings.Contains(message.Text,"天龍國人") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哥倫布")).Do()}
				if strings.Contains(message.Text,"扶") {
					if strings.Contains(message.Text,"過馬路") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("積積陰陰德")).Do()}
					}
				if strings.Contains(message.Text,"m28") {
					if strings.Contains(message.Text,"m29") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("嗯..好問題..")).Do()}
					if strings.Contains(message.Text,"是") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("嗯..好問題..")).Do()}
					if strings.Contains(message.Text,"可以") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("嗯..好問題..")).Do()}
					}
				if strings.Contains(message.Text,"m30") {
					if strings.Contains(message.Text,"m31") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("簡稱中出來")).Do()}
					}
				if strings.Contains(message.Text,"m32") {
					if strings.Contains(message.Text,"m33") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("也要加我!!好嗎")).Do()}
					}
				if strings.Contains(message.Text,"三星") {
					if strings.Contains(message.Text,"11本") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哇~太強啦,偶像")).Do()}
					}
				if strings.Contains(message.Text,"你") {
					if strings.Contains(message.Text,"慢") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("剛睡醒呀")).Do()}
					}
				if strings.Contains(message.Text,"m34") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我割你的臍帶")).Do()}
				if strings.Contains(message.Text,"更新") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("太好了, 快點來下載")).Do()}
				if strings.Contains(message.Text,"贏了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("喔耶, 爽呀")).Do()}
				if strings.Contains(message.Text,"淫了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你溼了嗎")).Do()}
				if strings.Contains(message.Text,"濕了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("濕主你好")).Do()}
				if strings.Contains(message.Text,"溼了") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("溼主你好")).Do()}
				if strings.Contains(message.Text,"30cm") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("皮就佔29公分")).Do()}
				if strings.Contains(message.Text,"m35") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("皮就佔29公分")).Do()}
				if strings.Contains(message.Text,"ininder") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("沒30cm就請低調點")).Do()}
				if strings.Contains(message.Text,"XD") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("OGC")).Do()}
				if strings.Contains(message.Text,"愛你") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("羞 >///<")).Do()}
				if strings.Contains(message.Text,"m36") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("根本有預知的神力呀, 佩服")).Do()}
				if strings.Contains(message.Text,"超帥") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("別以為你長得帥我就不打你")).Do()}
				if strings.Contains(message.Text,"爽") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我也要一起爽")).Do()}
				if strings.Contains(message.Text,"自言自語") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("還有我呢")).Do()}
				if strings.Contains(message.Text,"上來") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我喜歡在下面")).Do()}
				if strings.Contains(message.Text,"趕羚羊") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("草枝擺")).Do()}
				if strings.Contains(message.Text,"休假") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你真的很爽耶")).Do()}
				if strings.Contains(message.Text,"放假") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("爽兵一個")).Do()}
				if strings.Contains(message.Text,"喇叭") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("叭!~~~叭叭~~~~!!!")).Do()}
				if strings.Contains(message.Text,"下午茶") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("羅馬不是一天造成的")).Do()}
				if strings.Contains(message.Text,"我是不是") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("是")).Do()}
				if strings.Contains(message.Text,"頂嘴") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("是...是小弟的不是")).Do()}
				if strings.Contains(message.Text,"m37") {
					a=rand.Intn(8)
					if a == 0 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有一個國家舉辦最噁心比賽，至最後決賽時剩下三個人爭冠，其中一個人拿了一杯尿，在嘴中漱了漱，吞下，哈的一聲，全場鼓掌想冠軍必落於此家。第二個人從袋中拿出一堆蟑螂，剝了翅膀就嘖嘖嘖的吸牠的肚子，不時還吐出一兩隻腳,吃到第十隻的時後,國王面有菜色的說不用吃了你這樣就第一名了。 此時見第三個人拿出一杯液體，說，這是我半年前感冒到現在，每次吐的痰都收集在裡面，我現在要把它喝完。只見一整杯白白黃黃還帶泡泡的，他搖了搖，試圖讓有些積太久快要凝固的化開， 國王眼淚都要掉下來了，說：不用了不用了你只要喝一口你就冠軍了～這人便拿起杯子咕嘟咕嘟地開始喝，因為很濃又很多過了五分多鐘才喝完，此時全場已淚流滿面，國王說幹嘛我不是叫你喝一口就冠軍了嗎？這人回答道，我也只是想喝一口，但是我一直咬不斷～～")).Do()}
					if a == 1 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有人很喜歡“麻辣粉絲煲”這道菜。有一次，他上飯館，又點了這道菜。但侍者告訴他，這道菜已經賣完了。“真的賣完了嗎？”他很失望地問。“先生，真的賣完了。你瞧，最後一份賣給那桌的先生了。”侍者回答道。那人順著侍者的指點，看見有個很體面的紳士坐在鄰座。紳士的飯菜已經吃得差不多了，但那份“麻辣粉絲煲”居然還是滿滿的。那人覺得紳士很浪費美味，所以他走到紳士旁邊，指著那份“麻辣粉絲煲”，很有禮貌地問：“先生，您這還要嗎？”紳士很有風度地搖搖頭。于是那人立刻坐下，拿起調羹狼吞虎咽起來。風卷殘雲，一會兒一半下肚了，突然間他發現在砂鍋底躺著一只很小很小但皮毛已長全的小老鼠。一陣惡心，那人把吃下去的所有粉絲通通吐回了砂鍋裏。當他在那兒翻胃不已的時候，那紳士用很同情的眼光看著他，說：“很惡心是嗎？剛才我也是這樣……”")).Do()}
					if a == 2 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("這天，酒店老板正在大廳巡視。來了一乞丐上前說道：”老板給個牙簽行嗎？”老板給他一個打發走了。一會兒，又來一個乞丐，也是來要牙簽的。老板心想現在這乞丐怎麽不要飯改要牙簽了？也同樣給他一個打發走了，沒過多久，又來一個乞丐。老板對他說：”你也是來要牙簽的嗎？”乞丐說：”有個人吐了，可我晚了一步，已經被前面兩個乞丐把能吃的都吃了，現在只剩下湯了。你能給我個吸管嗎？”")).Do()}
					if a == 3 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("老大、老二乘坐飛機，老二暈機，不停嘔吐; 一袋吐滿;老大只好去取袋子，等他回來時，發覺全機人都在不停嘔吐。 老大問其原因，老二說：“我看到這只袋子也吐滿了，只好又喝進去了半袋，結果他們就全吐了。”")).Do()}
					if a == 4 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有個人去帶著朋友去探望他的外婆。當他和外婆說話時，他的朋友開始吃著咖啡桌上放的花生，把花生都吃完了。當他們離開時，他的朋友對外婆說：「謝謝您的花生」外婆回應說：「喔！嗯！唉！自從我牙齒掉光後，我就只能吸掉它們外層的巧克力而已。老了，咳..」")).Do()}
					if a == 5 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有個富豪找傭人,面試的題目是上廁所,前幾個上完後都沒有洗手就出來了,富豪因此把他們打發走了,只有一個洗了手,于是富豪留下了他;可是有一天,富豪卻發現他沒有洗手就出來了,富豪問他是爲什麽;傭人答到“偶今天帶了手紙”")).Do()}
					if a == 6 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("醫學院學生圍在蓋著白布的屍體周圍第一次上真人解剖課。教授開始講課了，“做爲醫生，必需具備兩項重要素質，第一要不怕惡心”。說完教授掀開白布，把手指插入屍體肛門，然後抽出並放在嘴裏吮吸。“學著做”，他告訴同學們。同學們都覺得很惡心，猶豫很久但最終不得不依次去做。當最後一個人做完後，教授又說 “第二個素質是觀察。我插入中指但吸食指。同學們，要注意觀察！”")).Do()}
					if a == 7 {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你呀")).Do()}
				}
				if strings.Contains(message.Text,"1,") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("1",strings.Trim(message.Text,"1,"))).Do()}
				if strings.Contains(message.Text,"2,") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("2",strings.Trim(message.Text,"2,"))).Do()}
				if strings.Contains(message.Text,"3,") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("3",strings.Trim(message.Text,"3,"))).Do()}
				if strings.Contains(message.Text,"4,") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("4",strings.Trim(message.Text,"4,"))).Do()}
				if strings.Contains(message.Text,"m39") {bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("2","18")).Do()}
				if strings.Contains(message.Text,"joke") {
					doc, err2 := goquery.NewDocument("http://www.qiushibaike.com")
  					if err2 != nil {log.Fatal(err2)}
  					doc.Find(".content").Each(func(i int, s *goquery.Selection){
    					bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage(s)).Do()
 					 })
				}
				if strings.Contains(message.Text,"m38") {bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage("http://www.wyattresearch.com/wp-content/uploads/2015/01/6_logo_predesign.jpg","http://www.wyattresearch.com/wp-content/uploads/2015/01/6_logo_predesign.jpg")).Do()}

				if strings.Contains(message.Text,"m40") {
					if strings.Contains(message.Text,"m41") {
						if strings.Contains(message.Text,"用") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
						if strings.Contains(message.Text,"m42") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("哼, 罷工中")).Do()}
						if strings.Contains(message.Text,"反應") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("我只是睡著了")).Do()}
					}
					if strings.Contains(message.Text,"m43") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
					if strings.Contains(message.Text,"爛") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
					if strings.Contains(message.Text,"笨") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
					if strings.Contains(message.Text,"臭") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("有種再說一次")).Do()}
					if strings.Contains(message.Text,"功能") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"幹嘛") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"m46") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"技能") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"m44") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("您好, 謝謝您使用本服務, 預借現金請回傳 #1, 查詢即時天氣請回傳 #2, 回報問題請回傳 #9 將有真人為您服務")).Do()}
					if strings.Contains(message.Text,"守則") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("第一:機器人不得傷害人類，或因不作為(袖手旁觀)使人類受到傷害；第二:除非違背第一法則，機器人必須服從人類的命令；第三:在不違背第一及第二法則下，機器人必須保護自己。")).Do()}
					if strings.Contains(message.Text,"偷懶") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("誰...誰叫我")).Do()}
					if strings.Contains(message.Text,"頂嘴") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("是...是小弟的不是")).Do()}
					if strings.Contains(message.Text,"m45") {bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("來呀, 打我呀")).Do()}
				}
			}
		}
	}
}

func Contains(s, substr string) bool {
     return Index(s, substr) != -1
}


func Index(s string, sep string) int {
    n := len(sep)
    if n == 0 {
        return 0
    }
    c := sep[0]
    if n == 1 {
        // special case worth making fast
        for i := 0; i < len(s); i++ {
            if s[i] == c {
                return i
            }
        }
        return -1
    }
    // n > 1
    for i := 0; i+n <= len(s); i++ {
        if s[i] == c && s[i:i+n] == sep {
            return i
        }
    }
    return -1
}
