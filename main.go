package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

type SlackRequestBody struct {
	Text string `json:"text"`
}

func main() {

	currentTime := time.Now()
	file, err := os.Open("new.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	webhookUrl := ""
	fault := SendSlackNotification(webhookUrl, currentTime.String())
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		err = SendSlackNotification(webhookUrl, scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if fault != nil {
		log.Fatal(fault)
	}
}

func SendSlackNotification(webhookUrl string, msg string) error {

	slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
	req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}
	return nil
}
