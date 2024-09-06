package notification

import (
	"encoding/json"
	"fmt"
	"suninfo-notification/settings"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

func SendNotification(message string, toPhoneNumber string) bool {
	accountSid, authToken, fromNumber := settings.GetTwilioSettings()

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
		return false
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
		return true
	}
}
