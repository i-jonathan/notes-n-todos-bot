package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type webHookReqBody struct {
	Message message `json:"message"`
}

var db = initDB()

func main() {
	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, http.HandlerFunc(messageHandler))
	log.Println(err)
}

func messageHandler(resp http.ResponseWriter, req *http.Request) {
	body := &webHookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		log.Println(err)
		return
	}

	go processRequest(body)
}
