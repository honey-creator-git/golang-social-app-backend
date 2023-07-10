package controllers

import (
	"gopkg.in/mailgun/mailgun-go.v1"
)

type MailRequest struct {
	from             string
	title            string
	subject          string
	htmlMessage      string
	plainTextMessage string
	to               []string
}

// Create new mail request
func NewMailRequest(from string, title string, htmlMessage string, textMessage string, receivers []string) *MailRequest {
	return &MailRequest{
		from:             from,
		title:            title,
		htmlMessage:      htmlMessage,
		plainTextMessage: textMessage,
		to:               receivers,
	}
}

//SendMail is used to send message, it will ask user about title, htmlMessage, textMessage and list of recipient
func (mailRequest *MailRequest) SendMail() (bool, error) {
	// NewMailGun creates a new client instance
	mg := mailgun.NewMailgun("replies.ourserendipityapp.com", "165a8415a40f7ad0dc954e63596e6022-1b3a03f6-a09659ba", "pubkey-092b2b4e5898ff87053853600e1eea9c")

	// Create message
	message := mg.NewMessage(
		mailRequest.from,
		mailRequest.title,
		mailRequest.plainTextMessage,
		mailRequest.to...,
	)

	message.SetHtml(mailRequest.htmlMessage)

	// Send message and get result
	if _, _, err := mg.Send(message); err != nil {
		return false, err
	}

	return true, nil
}
