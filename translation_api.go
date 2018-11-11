package yandexapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type YandexLanguages struct {
	Translation []string          `json:"dirs"`
	Language    map[string]string `json:"langs"`
}

type LanguageDetection struct {
	RequestsCode   int    `json:"code"`
	DecodeLanguage string `json:"lang"`
}

type TextTranslation struct {
	Translatetext []string `json:"text"`
	Translation   string   `json:"lang"`
}

func Getyandexlanguages(api_key string, target_language string) *YandexLanguages {

	urlstring := "https://translate.yandex.net/api/v1.5/tr.json/getLangs"
	data := url.Values{}
	data.Add("key", api_key)
	data.Add("ui", target_language)

	r, _ := http.NewRequest("GET", urlstring, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded") // URL-encoded payload
	req, error := YandexhttpClient.Do(r)

	if error != nil {
		panic(error)
	}
	defer req.Body.Close()
	yandex_languages := new(YandexLanguages)
	if err := json.NewDecoder(req.Body).Decode(&yandex_languages); err != nil {
		panic("Error parsing api requests")
	}
	return yandex_languages

}

func Detectlanguage(api_key string, message string) *LanguageDetection {

	urlstring := "https://translate.yandex.net/api/v1.5/tr.json/detect"
	data := url.Values{}
	data.Add("key", api_key)
	data.Add("text", message)

	r, _ := http.NewRequest("GET", urlstring, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded") // URL-encoded payload
	req, error := YandexhttpClient.Do(r)

	if error != nil {
		panic(error)
	}
	defer req.Body.Close()
	language_detection := new(LanguageDetection)
	if err := json.NewDecoder(req.Body).Decode(&language_detection); err != nil {
		panic(err)
	}
	return language_detection

}

func Gettexttranslation(api_key string, message string, language string) *TextTranslation {

	language_detection := Detectlanguage(api_key, message)
	translation_code := fmt.Sprintf("%s-%s", language_detection.DecodeLanguage, language)

	urlstring := "https://translate.yandex.net/api/v1.5/tr.json/translate"
	data := url.Values{}
	data.Add("key", api_key)
	data.Add("text", message)
	data.Add("lang", translation_code)

	r, _ := http.NewRequest("POST", urlstring, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded") // URL-encoded payload
	resp, err := YandexhttpClient.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	text_translation := new(TextTranslation)
	if err := json.NewDecoder(resp.Body).Decode(&text_translation); err != nil {
		panic(err)
	}
	return text_translation
}
