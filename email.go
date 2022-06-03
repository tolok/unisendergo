package unisendergo

import (
	"context"
)

type EmailSendResponse struct {
	Status       string            `json:"status"`
	JobId        string            `json:"job_id"`
	Emails       []string          `json:"emails"`
	FailedEmails map[string]string `json:"failed_emails"`
}

func (u *Unisender) EmailSend(ctx context.Context, message Message) (*EmailSendResponse, error) {
	var response EmailSendResponse

	if err := request(ctx, u.client, sendEmail, message, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
