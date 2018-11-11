package yandexapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Lookup struct {
	Headers    map[string]string `json:"head,omitempty"`
	Definition `json:"def"`
}

type Definition []struct {
	Text             string `json:"text"`
	SemanticCategory string `json:"pos"`
	Translation      `json:"tr"`
}

type Translation []struct {
	Text             string `json:"text"`
	SemanticCategory string `json:"pos"`
	Synonyms         []struct {
		Text string `json:"text"`
	} `json:"syn"`
	Meaning []struct {
		Text string `json:"text"`
	} `json:"mean"`
	Examples []struct {
		Text string `json:"text"`
		Tr   []struct {
			Text string `json:"text"`
		} `json:"tr"`
	} `json:"ex"`
}

func DiccionarySearch(api_key string, message string, language string) *Lookup {
	language_detection := Detectlanguage(api_key, message)
	translation_code := fmt.Sprintf("%s-%s", language_detection.DecodeLanguage, language)

	client := &http.Client{}
	urlstring := "https://dictionary.yandex.net/api/v1/dicservice.json/lookup"
	data := url.Values{}
	data.Add("key", api_key)
	data.Add("lang", translation_code)
	data.Add("text", message)

	r, _ := http.NewRequest("POST", urlstring, strings.NewReader(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded") // URL-encoded payload
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	lookup_response := new(Lookup)
	if err := json.NewDecoder(resp.Body).Decode(&lookup_response); err != nil {
		panic(err)
	}
	return lookup_response

}
