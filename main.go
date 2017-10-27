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

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

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
	var list string
	var price string
	var stock string
	var food string
	url := "https://raw.githubusercontent.com/Yikaros/LineBotTemplate/master/images/"
    	for {
        	a, _, c := br.ReadLine()
        	if c == io.EOF {
            	break
        	}
		list = list + "&" + string(a)
    	}
	
	var list_array []string
	list_array = strings.Split(list, "&")
	
	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			var cow string
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			cow = url + "cow/" + fmt.Sprintf("%d", r.Intn(10))  + ".jpg"

			switch message := event.Message.(type) {
			case *linebot.TextMessage:     
				switch {
//賣菜的code
					case Contains(message.Text,"菜")||Contains(message.Text,"葉"):
						food = ""
						switch{
							case Contains(message.Text,"高麗菜"):
								food = "高麗菜"
							case Contains(message.Text,"小白菜"):
								food = "小白菜"
							
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
								case Contains(message.Text,"多少錢")||Contains(message.Text,"怎麼賣")||Contains(message.Text,"怎麼算"):
									bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "一斤" + price)).Do() 
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
							case Contains(message.Text,"石斑"):
								food = "石斑"
						}
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
							case Contains(message.Text,"多少錢")||Contains(message.Text,"怎麼賣")||Contains(message.Text,"怎麼算"):
								if food != "石斑"{								
									bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "一斤" + price)).Do() 
								}else{
									bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("拍謝啦! 我是笨笨的電腦，不知道您要問哪種石斑捏，我們有龍虎石斑、青斑、還有珍珠石斑")).Do()
								}
							case Contains(message.Text,"還有多少")||Contains(message.Text,"剩下多少")||Contains(message.Text,"庫存")||Contains(message.Text,"還有幾"):
								if food != "石斑"{
									bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(food + "大概還有" + stock + "尾可以買，賣完就沒了喔!! 趕快來電088953096/0939220743黃先生")).Do() 
								}else{
									bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("拍謝啦! 我是笨笨的電腦，不知道您要問哪種石斑捏，我們有龍虎石斑、青斑、還有珍珠石斑")).Do()
								}
						}
//以下是喇賽的code
					case Contains(message.Text,"87"):
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("87分，不能再高惹")).Do() 
					case Contains(message.Text,"母牛"):
						bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(cow,cow)).Do() 
					case Contains(message.Text,"洗眼"):
						bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(cow,cow)).Do() 
					case Contains(message.Text,"乳牛"):
						bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(cow,cow)).Do() 
					case Contains(message.Text,"淨化"):
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
