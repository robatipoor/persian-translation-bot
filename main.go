package main

import (
	"bytes"
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

var port string
var token string
var appURL string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	appURL = os.Getenv("APP_URL")
	if appURL == "" {
		log.Fatal("Application URL not set")
	}
	token = os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatalln("TELEGRAM_TOKEN not set !")
	}
}

func main() {
	
	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: appURL},
	}
	setings := tb.Settings{
		Token:  token,
		Poller: webhook,
	}
	bot, err := tb.NewBot(setings)
	if err != nil {
		log.Fatal(err)
	}
	bot.Handle(tb.OnText, func(m *tb.Message) {
		if m.Text == "/start" {
			log.Println("show help message")
			bot.Send(m.Sender, "Please type in English or Farsi and send me !")
			return
		}
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
