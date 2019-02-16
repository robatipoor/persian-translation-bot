package main

import (
	"errors"
	"log"

	"github.com/abadojack/whatlanggo"
)

type Lang string

const (
	FA Lang = "fa"
	EN Lang = "en"
)

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
