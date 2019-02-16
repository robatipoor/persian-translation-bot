package main

import (
	"bytes"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

func StartBot() {
	bot, err := tb.NewBot(tb.Settings{
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle(tb.OnText, func(m *tb.Message) {
		targetLang, err := TargetLang(m.Text)
		if err != nil {
			log.Println(err)
			bot.Send(m.Sender, "Detect Language Faild !")
			return
		}
		tr, err := Translate(m.Text, string(targetLang))
		if err != nil {
			log.Println(err)
			bot.Send(m.Sender, "Translate Problem !")
			return
		}
		log.Println(tr)
		if targetLang == FA {
			speech, err := TextToSpeech(m.Text)
			if err != nil {
				log.Println(err)
				bot.Send(m.Sender, "Text To Speech Problem !")
				return
			}
			audio := &tb.Audio{File: tb.FromReader(bytes.NewBuffer(speech))}
			audio.Title = m.Text
			audio.Caption = tr
			bot.Send(m.Sender, audio)
		} else {
			bot.Send(m.Sender, tr)
		}
	})

	log.Println("Start Bot ...")
	bot.Start()
}
