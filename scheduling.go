package main

import (
	"github.com/jasonlvhit/gocron"
	"log"
)


func scheduledTask() {

	s := gocron.NewScheduler()
	s.Every(30).Minutes().Do(func() {
		b, err := get("https://persiantranslationbot.herokuapp.com/")
		if err != nil {
			log.Println(err)
		}
		log.Println(string(b))
	})
	log.Println("Start Scheduled Task")
	<-s.Start()
}