package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)
// Translate Fa to En and reverse
func Translate(text string,targetLang string) (string, error) {

	l := "https://translate.google.com/translate_a/single?&client=gtx&sl=auto"
	l = l + "&tl=" + string(targetLang)
	l = l + "&hl=" + string(targetLang)
	l = l + "&dt=t"
	l = l + "&text=" + url.QueryEscape(text)
	l = l + "&tk=" + tk(text)
	trans, err := get(l)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var sb strings.Builder
	var count int
	for {
		r := gjson.Get(trans, fmt.Sprintf("0.%d.0", count)).String()
		if r == "" {
			break
		}
		sb.WriteString(r)
		count++
	}

	return sb.String(), nil
}

func get(url string) (string, error) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(b), nil
}

func tk(s string) string {
	tkki, tkkf := gettkk()
	tki, tkf := gettk(int32(tkki), int32(tkkf), s)
	return (strconv.Itoa(int(tki)) + "." + strconv.Itoa(int(tkf)))
}

func gettkk() (int64, int64) {

	s, err := get("https://translate.google.com/")
	if err != nil {
		log.Panicln(err)
	}
	reg := regexp.MustCompile(`tkk:'.*?'`)
	tkkc := reg.FindString(s)
	tkkc = strings.Replace(tkkc, "tkk:", "", -1)
	tkkv := strings.Split(tkkc, ".")
	a, _ := strconv.ParseInt(tkkv[0], 10, 64)
	b, _ := strconv.ParseInt(tkkv[1], 10, 64)

	return a, b
}

func gettk(tkka int32, tkkb int32, str string) (int32, int32) {

	b := make([]byte, len(str))
	b = []byte(str)
	a := tkka
	for i := 0; i < len(b); i++ {
		a += int32((b[i]))
		a = xr(a, "+-a^+6")
	}
	a = xr(a, "+-3^+b+-f")
	a ^= tkkb
	if a < 0 {
		a = -a
	}
	a %= 1E6

	return a, a ^ tkka
}

func xr(a int32, b string) int32 {

	for c := 0; c < len(b)-2; c = c + 3 {
		var d int32
		d = int32(b[c+2])
		if d >= int32('a') {
			d = d - 87
		} else {
			d = d - 48
		}
		if int32('+') == int32(b[c+1]) {
			if a >= 0 {
				d = a >> uint32(d)
			} else {
				t := 4294967295 + int(a)
				t = t + 1
				d = int32(t >> uint(d))
			}
		} else {
			d = a << uint32(d)
		}

		if int32('+') == int32(b[c]) {
			a = a + d
		} else {
			a = a ^ d
		}
	}
	return a
}
