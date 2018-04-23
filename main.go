package main

import (
	"context"
	"fmt"
	"path"

	"./extLoggers/slack"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

)


func handler(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3

		sm := &slack.SlackMessage{
			Text: fmt.Sprintf("%s was uploaded into `%s` bucket", s3.Object.Key, s3.Bucket.Name),
		}

		attachment := slack.Attachment{
			Text:  s3.Object.Key,
			Color: slack.ExtensionToColor(path.Ext(s3.Object.Key)),
		}

		attachment.AddAction(slack.Action{
			Text:  "Show file",
			Type:  "button",
			URL:   fmt.Sprintf("https://s3.%s.amazonaws.com/%s/%s", record.AWSRegion, s3.Bucket.Name, s3.Object.Key),
			Style: "primary",
		})

		sm.Attachments = append(sm.Attachments, attachment)
		sm.Send()
	}

}


func main() {
	lambda.Start(handler)
}
