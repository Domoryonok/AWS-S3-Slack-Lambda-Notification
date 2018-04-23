package main

import (
	"fmt"
	"net/http"
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"bytes"
	"log"
	"encoding/json"
	"os"
)

var SLACK_WEBHOOK_URL = os.Getenv("SLACK_WEBHOOK_URL")

type slackMessage struct {
	Text string `json:"text"`
}

func handleError(err error, reason string) {
	if err != nil {
		log.Fatalf("%s error: %v", reason, err)
	}
}

func (sm *slackMessage) Send() {

	client := &http.Client{}

	body, err := json.Marshal(sm)
	handleError(err, "Marshalling")

	req, err := http.NewRequest("POST", SLACK_WEBHOOK_URL, bytes.NewReader(body))
	handleError(err, "Making new request")

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	handleError(err, "Sending slack request")
	defer resp.Body.Close()

}

func handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		sm := &slackMessage{
			fmt.Sprintf("%s was uploaded into `%s` bucket", s3.Object.Key, s3.Bucket.Name),
		}
		sm.Send()
	}

}

func main() {
	lambda.Start(handler)
}