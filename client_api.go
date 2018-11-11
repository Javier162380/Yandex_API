package yandexapi

import (
	"net/http"
	"time"
)

var YandexhttpClient = &http.Client{
	Timeout: time.Second * 10,
}
