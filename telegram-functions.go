package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
)

type message struct {
	Text string `json:"text"`
	Chat chat   `json:"chat"`
}

type chat struct {
	ID int64 `json:"id"`
}

type reply struct {
	ChatID    int64  `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

var token = os.Getenv("token")
var url = "https://api.telegram.org/bot" + token + "/"

func respond(chatID int64, text string) error {
	reqBody := &reply{
		ChatID:    chatID,
		Text:      text,
		ParseMode: "HTML",
	}

	reqBytes, err := json.Marshal(reqBody)

	if err != nil {
		return err
	}
	response, err := http.Post(url+"sendMessage", "application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
		return errors.New("Unexpected status code: " + response.Status)
	}

	return nil
}
