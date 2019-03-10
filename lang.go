package main

import (
	"errors"
	"log"

	"github.com/abadojack/whatlanggo"
)
// Lang = Language
type Lang string

const (
	// FA = Farsi
	FA Lang = "fa"
	// EN = English
	EN Lang = "en"
)
// DetectLanguage Function
func DetectLanguage(text string) (Lang, error) {

	info := whatlanggo.DetectScript(text)
	log.Println(whatlanggo.Scripts[info])
	switch whatlanggo.Scripts[info] {
	case "Arabic":
		return FA, nil
	case "Latin":
		return EN, nil
	default:
		return "", errors.New("language not supported")
	}
}
// TargetLang Function
func TargetLang(text string) (Lang, error) {

	lang, err := DetectLanguage(text)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var targetLang Lang
	if lang == EN {
		targetLang = FA
	} else {
		targetLang = EN
	}

	return targetLang, nil
}
