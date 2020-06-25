package main

import (
	"os"
	"log"
	//"fmt"
	"time"
	"strconv"
	"net/http"
	//"io/ioutil"
	//"encoding/json"
	tb "gopkg.in/tucnak/telebot.v2"
)

func MainHandler(resp http.ResponseWriter, _ *http.Request) {
    resp.Write([]byte("Hi there! I'm PopravkiBot!"))
}

func main() {
	bot_token := os.Getenv("BOT_TOKEN")
	b, err := tb.NewBot(tb.Settings{
		Token: bot_token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		return
	}

	log.Printf("Authorized on account popravki_bot")

	zab := tb.InlineButton{
		Unique: "ZA",
		Text:   "✔️ За",
	}
	
	protivb := tb.InlineButton{
		Unique: "P",
		Text:   "❌ Против",
	}


	mainInline := [][]tb.InlineButton{
		[]tb.InlineButton{protivb, zab},
	}


	http.HandleFunc("/", MainHandler)
    go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	b.Handle("/start", func(m *tb.Message) {
			var za int
			var protiv int
			uresp := "Привет! 🗳️\nДавай голосовать по поправкам!\n\nЗа: " + strconv.Itoa(za) + "\nПротив: " + strconv.Itoa(protiv)
			b.Send(m.Sender, uresp, &tb.ReplyMarkup{
				InlineKeyboard: mainInline,
			})


			b.Handle(&zab, func(c *tb.Callback) {

				log.Println(m.Sender.Username, ": za")

				za += 1

				uresp := "Привет! 🗳️\nДавай голосовать по поправкам!\n\nЗа: " + strconv.Itoa(za) + "\nПротив: " + strconv.Itoa(protiv)
				b.Edit(c.Message, uresp, &tb.ReplyMarkup{
					InlineKeyboard: mainInline,
				})

				b.Respond(c, &tb.CallbackResponse{})
			})

			b.Handle(&protivb, func(c *tb.Callback) {

				log.Println(m.Sender.Username, ": protiv")

				protiv += 1

				uresp := "Привет! 🗳️\nДавай голосовать по поправкам!\n\nЗа: " + strconv.Itoa(za) + "\nПротив: " + strconv.Itoa(protiv)
				b.Edit(c.Message, uresp, &tb.ReplyMarkup{
					InlineKeyboard: mainInline,
				})

				b.Respond(c, &tb.CallbackResponse{})
			})
	})

	b.Start()
}