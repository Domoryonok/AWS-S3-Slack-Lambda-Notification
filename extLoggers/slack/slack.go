package slack

import (
	"bytes"
	"os"

	"encoding/json"
	"net/http"

	"../../utils"
)

var SLACK_WEBHOOK_URL = os.Getenv("SLACK_WEBHOOK_URL")

type Action struct {
	Text  string `json:"text,omitempty"`
	Type  string `json:"type,omitempty"`
	URL   string `json:"url,omitempty"`
	Style string `json:"style,omitempty"`
}

/* AddAction(action Action) void */
type Attachment struct {
	Text    string   `json:"text,omitempty"`
	Actions []Action `json:"actions,omitempty"`
	Color   string   `json:"color,omitempty"`
}

func (at *Attachment) AddAction(action Action) {
	at.Actions = append(at.Actions, action)
}


/* Send() void */
type SlackMessage struct {
	Text        string       `json:"text,omitempty"`
	Attachments []Attachment `json:"attachments,omitempty"`
}

func (sm *SlackMessage) Send() {

	client := &http.Client{}

	body, err := json.Marshal(sm)
	utils.HandleError(err, "Marshalling")

	req, err := http.NewRequest("POST", SLACK_WEBHOOK_URL, bytes.NewReader(body))
	utils.HandleError(err, "Making new request")

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	utils.HandleError(err, "Sending slack request")
	defer resp.Body.Close()

}