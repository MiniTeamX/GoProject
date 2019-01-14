package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type NormalTaskRequest struct {
	Url        string
	FileName   string
	FileSize   uint64
	Gcid       string
	Md5        string
	Sha1       string
}

type NormalTaskSliceRequest struct {
	NormalTasks     []NormalTaskRequest
}


type NormalTaskResponse struct {
	Hit       int64
	Md5       string
	ErrInfo   string
	Url       string
	Sha1      string
	RetCode   int64
	Gcid      string
}

type NormalTaskSliceResponse struct {
	RetCode    int64
	Result     []NormalTaskResponse
	ErrInfo    string
}

func HandleNormalTasks(w http.ResponseWriter, r *http.Request) {
	normalTasks := NormalTaskSliceRequest{}
	if r.Method == "post" {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal("Read request json failed:", err)
		}
		defer r.Body.Close()
		err = json.Unmarshal(b, &normalTasks)
		if err != nil {
			log.Fatal("Json format error:", err)
		}
	}
	log.Println("Request tasks is:", normalTasks)

    responseResult := NormalTaskSliceResponse {RetCode:0, ErrInfo:"1"}

	for _,v := range normalTasks.NormalTasks {
		response := handleOneTask(v)
		responseResult.Result = append(responseResult.Result, *response)
	}
}

func handleOneTask(taskRequest *NormalTaskRequest) (response *NormalTaskResponse){


}
