package yandexapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type YandexLanguages struct {
	Translation []string    `json:"dirs"`
	Language    interface{} `json:"langs"`
}

type LanguageDetection struct {
	RequestsCode   int    `json:"code"`
	DecodeLanguage string `json:"lang"`
}

func Get_yandex_languages(api_key string, target_language string) *YandexLanguages {
	url_req := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/"+
		"getLangs?key=%s&ui=%s", api_key, target_language)
	req, error := http.Get(url_req)
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

func Detect_language(api_key string, message string) *LanguageDetection {
	url_req := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/detect"+
		"?key=%s&text=%s", api_key, message)
	req, error := http.Get(url_req)
	if error != nil {
		panic(error)
	}
	defer req.Body.Close()
	language_detection := new(LanguageDetection)
	if err := json.NewDecoder(req.Body).Decode(&language_detection); err != nil {
		panic("Error parsing api requests")
	}
	return language_detection

}

func Get_text_translation(api_key string, message string, language string) map[string]interface{} {

	language_detection := Detect_language(api_key, message)
	translation_code := fmt.Sprintf("%s-%s", language_detection.DecodeLanguage, language)
	fmt.Printf(translation_code)
	url_req := fmt.Sprintf("https://translate.yandex.net/api/v1.5/tr.json/detect"+
		"?key=%s&text=%s=&lang=%s", api_key, message, translation_code)
	req, error := http.Get(url_req)
	if error != nil {
		panic(error)
	}
	defer req.Body.Close()
	text_translation := make(map[string]interface{})
	if err := json.NewDecoder(req.Body).Decode(&text_translation); err != nil {
		panic("Error parsing api requests")
	}
	return text_translation
}
