package main

import (
	"fmt"
	"javier162380/yandexapi"
	"os"
)

func main() {
	yandex_api_key := os.Args[1]
	fmt.Printf("%s", yandex_api_key)
	language_translation := yandexapi.Get_yandex_languages(
		yandex_api_key,
		"en")
	fmt.Printf("%s", language_translation.Translation)
}
