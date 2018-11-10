package main

import (
	"fmt"
	"javier162380/yandexapi"
	"os"
)

func main() {
	yandex_api_key := os.Args[1] //<<INSERT YOUR YANDEX API KEY HERE>>
	fmt.Printf("%s", yandex_api_key)
	languagetranslation := yandexapi.Getyandexlanguages(yandex_api_key, "en")
	fmt.Printf("%s", languagetranslation.Language)

	detectlanguage := yandexapi.Detectlanguage(yandex_api_key, "españa")
	fmt.Printf("%s", detectlanguage.DecodeLanguage)

	translatemessage := yandexapi.Gettexttranslation(yandex_api_key,
		"estoy temblando",
		"en")
	fmt.Printf("%s", translatemessage.Translatetext)

}
