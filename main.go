package main

import (
	"log"
	"net/http"
	"os"
)


var addr string

func init() {
	var port string
	if os.Getenv("PORT") == "" {
		port = "8080"
	} else {
		port = os.Getenv("PORT")
	}
	if os.Getenv("TELEGRAM_TOKEN") == "" {
		log.Fatalln("TELEGRAM_TOKEN not set !")
	}
	addr = "0.0.0.0:" + port
}

func main() {
	go StartBot()
	http.HandleFunc("/", handler)
	log.Printf("Start Server %s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "Application/json")
	w.Write([]byte("Bot Start !"))
}
