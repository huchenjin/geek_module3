package main

import (
	"github.com/huchenjin/geek_module3/service"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/healthz", service.Healthz)
	http.HandleFunc("/header", service.Header)
	http.HandleFunc("/clientip", service.ClientIP)
	http.HandleFunc("/version", service.Version)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
