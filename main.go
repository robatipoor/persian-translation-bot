package main

import (
	"log"
	"net/http"
	"os"
)


var address string

func init() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if os.Getenv("TELEGRAM_TOKEN") == "" {
		log.Fatalln("TELEGRAM_TOKEN not set !")
	}
	address = "0.0.0.0:" + port
}

func main() {
	go StartBot()
	// for live in heroku cloud
	go scheduledTask()
	http.HandleFunc("/", handler)
	log.Println("Start Server : ", address)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "Application/json")
	w.Write([]byte("Bot Start !"))
}
