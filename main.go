package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joho/godotenv"
)

type Slack struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel"`
}

func HandleRequest(ctx context.Context) (string, error) {

	godotenv.Load()

	id := os.Getenv("APP_ID")
	url := fmt.Sprintf("https://itunes.apple.com/jp/app/%s?mt=8", id)

	doc, _ := goquery.NewDocument(url)

	title := doc.Find("header > h1").Children().Remove().End().Text()
	title = strings.TrimSpace(title)

	rank := doc.Find(".inline-list__item").First().Text()
	rank = strings.TrimSpace(rank)

	text := fmt.Sprintf("%s は %s です\n", title, rank)
	post(text)

	return "", nil
}

func main() {
	lambda.Start(HandleRequest)
}

func post(s string) {

	params := Slack{
		Text:      s,
		Username:  os.Getenv("SLACK_USERNAME"),
		IconEmoji: os.Getenv("SLACK_ICON_EMOJI"),
		Channel:   os.Getenv("SLACK_CHANNEL"),
	}
	jsonparams, _ := json.Marshal(params)
	resp, _ := http.PostForm(
		os.Getenv("INCOMING_WEBHOOK_URL"),
		url.Values{"payload": {string(jsonparams)}},
	)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	println(string(body))
}
