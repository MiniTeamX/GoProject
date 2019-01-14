package main

import (
	"net/http"
)

func init () {

}

func RouterMultiple (w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filterType := r.Form.Get("type")
	if filterType == "normal" {
		HandleNormalTasks(w, r)
	} else if filterType == "bt" {
		HandleBtTasks(w, r)
	}
}







