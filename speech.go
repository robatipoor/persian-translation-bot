package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/evalphobia/google-tts-go/googletts"
)

// TextToSpeech convert text to speech audio
func TextToSpeech(text string) ([]byte, error) {
	
	url, err := googletts.GetTTSURL(text, "en")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return b, nil
}
