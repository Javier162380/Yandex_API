package main

import (
	"fmt"
	"javier162380/yandexapi"
	"os"
)

func main() {
	translation_api_key := os.Args[1] //<<INSERT YOUR YANDEX API TRANSLATION KEY HERE>>
	dictionary_api_key := os.Args[2]  //<<INSERT YOUR YANDEX API DICTIONARY KEY HERE>>

	languagetranslation := yandexapi.Getyandexlanguages(translation_api_key, "en")
	fmt.Printf("%s", languagetranslation.Language)

	detectlanguage := yandexapi.Detectlanguage(translation_api_key, "españa")
	fmt.Printf("%s", detectlanguage.DecodeLanguage)

	translatemessage := yandexapi.Gettexttranslation(translation_api_key,
		"mensaje de prueba",
		"en")
	fmt.Printf("%s", translatemessage.Translatetext)

	lookupmethod := yandexapi.DiccionarySearch(dictionary_api_key,
		"lawyer",
		"es")
	fmt.Printf("%s", lookupmethod)

}
