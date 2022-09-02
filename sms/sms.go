package sms

import (
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func Notify(msg string) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo("+12533185980")
	params.SetFrom("+12062222873")
	params.SetBody(msg)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return err
	}
	return nil
}
