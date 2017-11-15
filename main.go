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
						case message.Text=="運動名單":
							d1 := []byte("hello\ngo\n")
							err := ioutil.WriteFile("buffer/test", d1, 0644)
						    	if err != nil {
								fmt.Printf("Error: %s\n", err)
								return
							}
						case Contains(message.Text,"艾魯"):
							template := linebot.NewConfirmTemplate("大家覺得艾魯是不是該刪掉刀塔?", linebot.NewMessageTemplateAction("是", "是"),linebot.NewMessageTemplateAction("4", "4"))
							messgage := linebot.NewTemplateMessage("Sorry :(, please update your app.", template)
							bot.ReplyMessage(event.ReplyToken, messgage).Do()
						case Contains(message.Text,"尻槍"):
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
						case Contains(message.Text,"什麼菜")||Contains(message.Text,"菜單")||message.Text=="我想買菜":
							
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
						case Contains(message.Text,"什麼魚")||message.Text=="我想買魚":
							template := linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									url + "/fish/bluef.jpg", "青斑", "一斤330元，物美價廉，最適合一般家庭",
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
						case message.Text=="我要加入":
							if len(profile) > 0{
								bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(profile[2] + "你已經是菜市場的會員囉，不用再申請加入啦")).Do()
							}else{
								ppljoin = uid
								bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要加入Protoss菜市場嗎? 請問您叫什麼名字呢?")).Do() 
							}
	//賣菜的code
						case Contains(message.Text,"菜")||Contains(message.Text,"葉"):						
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
									case Contains(message.Text,"一斤多少")||Contains(message.Text,"多少錢")||Contains(message.Text,"怎麼賣")||Contains(message.Text,"怎麼算"):
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "一斤" + price)).Do() 
									case Contains(message.Text,"要買"):
										/**if len(profile) > 0{
											bot.PushMessage(admin,linebot.NewTextMessage(profile[1] + profile[2] + "要買" + food)).Do() 
											bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "嗎? 我已經幫你聯絡老闆了，晚點他就會主動跟你聯繫，請耐心等一下喔")).Do()
										}else{
											bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買菜嗎? 可是你好像還不是我們菜市場的會員捏，麻煩跟管理員聯繫幫你加入菜市場會員，會員才有特別優惠喔!!")).Do() 	
										}**/
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買" + food + "嗎? 趕快撥打 0909 022 890 跟鄉下小孩-張耀東買菜喔!!")).Do()
								}
							}else{
								if Contains(message.Text,"要買"){
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
						case Contains(message.Text,"斑")||Contains(message.Text,"班"):
							food = ""
							switch{
								case Contains(message.Text,"龍虎"):
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
									case Contains(message.Text,"一斤多少")||Contains(message.Text,"多少錢")||Contains(message.Text,"怎麼賣")||Contains(message.Text,"怎麼算"):
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "一斤" + price)).Do() 
									case Contains(message.Text,"還有多少")||Contains(message.Text,"剩下多少")||Contains(message.Text,"庫存")||Contains(message.Text,"還有幾"):
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "大概還有" + stock + "尾可以買，賣完就沒了喔!! 趕快來電088953096/0939220743黃先生")).Do() 
									case Contains(message.Text,"要買"):/**
										if len(profile) > 0{
											bot.PushMessage(admin,linebot.NewTextMessage(profile[1] + profile[2] + "要買" + food)).Do() 
											bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "嗎? 我已經幫你聯絡老闆了，晚點他就會主動跟你聯繫，請耐心等一下喔")).Do()
										}else{
											bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買魚嗎? 可是你好像還不是我們菜市場的會員捏，麻煩跟管理員聯繫幫你加入菜市場會員，會員才有特別優惠喔!!")).Do() 	
										}**/
										bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("你要買石斑嗎? 趕快撥打 0939 220 743 跟石斑膠膠哥-黃大哥買魚喔!!")).Do()
								}
							}else{
								if Contains(message.Text,"要買"){/**
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
						case Contains(message.Text,"母牛")||Contains(message.Text,"洗眼")||Contains(message.Text,"乳牛")||Contains(message.Text,"淨化"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(cow,cow)).Do() 
						case Contains(message.Text,"刀塔"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(url + "6569950-1490833625.jpg", url + "6569950-1490833625.jpg")).Do() 
						case Contains(message.Text,"黑人問號"):
							bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(url + "blackman.jpg", url + "blackman.jpg")).Do() 					
					}
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
