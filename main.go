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
	"fmt"
	"log"
    	"time"
	"net/http"
	"os"
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

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			var cow string
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			cow = "https://raw.githubusercontent.com/Yikaros/LineBotTemplate/master/images/" + fmt.Sprintf("%d", r.Intn(20))  + ".jpg"

			switch message := event.Message.(type) {
			case *linebot.TextMessage:     
				switch {
					case Contains(message.Text,"87"):
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("87分，不能再高惹")).Do() 
					case Contains(message.Text,"56"):
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("5566得第一")).Do() 
					case Contains(message.Text,"母牛"):
						bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(cow,cow)).Do() 
					case Contains(message.Text,"洗眼"):
						bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(cow,cow)).Do() 
					case Contains(message.Text,"乳牛"):
						bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(cow,cow)).Do() 
					case Contains(message.Text,"淨化"):
						bot.ReplyMessage(event.ReplyToken, linebot.NewImageMessage(cow,cow)).Do() 
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
