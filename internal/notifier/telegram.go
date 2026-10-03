package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Telegram struct {
	token  string
	chatID string
}

func NewTelegram(token, chatID string) *Telegram {
	return &Telegram{
		token:  token,
		chatID: chatID,
	}
}

func (tele *Telegram) SendMessage(name, email, message string) error {
	text := fmt.Sprintf(
		"Name: %s\n"+
			"Email: %s\n"+
			"%s",
		name,
		email,
		message,
	)

	payload := map[string]string{
		"chat_id": tele.chatID,
		"text":    text,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil
	}

	endpoint := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage",
		tele.token,
	)

	resp, err := http.Post(
		endpoint,
		"application/json",
		bytes.NewReader(body),
	)
	if err != nil {
		return fmt.Errorf("send Telegram message %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("Telegram returned %s", resp.Status)
	}
	return nil
}
