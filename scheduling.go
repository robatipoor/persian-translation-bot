package main

import (
	"github.com/jasonlvhit/gocron"
	"log"
)


func scheduledTask() {

	s := gocron.NewScheduler()
	s.Every(15).Minutes().Do(func() {
		s, err := get("https://persiantranslationbot.herokuapp.com/")
		if err != nil {
			log.Println(err)
		}
		log.Println(s)
	})
	log.Println("Start Scheduled Task")
	<-s.Start()
}