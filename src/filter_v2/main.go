package main

import (
	"log"
	"net/http"
)

func main1 () {
	http.HandleFunc("/filter_task", RouterMultiple)
	err := http.ListenAndServe("8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

