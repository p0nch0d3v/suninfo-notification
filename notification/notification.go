package notification

import (
	"encoding/json"
	"fmt"
	"sunrise-sunset-notification/settings"

	"github.com/twilio/twilio-go"
)

func SendNotification(message string, toPhoneNumber string) {
	accountSid, authToken, fromNumber := getTwilioSettings()

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(toPhoneNumber)
	params.SetFrom(fromNumber)
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}

func getTwilioSettings() (string, string, string) {
	accountSid := settings.GetSetting("TWILIO_ACCOUNT_SID")
	authToken := settings.GetSetting("TWILIO_AUTH_TOKEN")
	fromNumber := settings.GetSetting("TWILIO_AUTH_FROM_NUMBER")
	return accountSid, authToken, fromNumber
}
