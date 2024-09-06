package main

import (
	"fmt"
	"log"
	"os"
	"suninfo-notification/db"
	"suninfo-notification/notification"
	"suninfo-notification/settings"
	sunInfo "suninfo-notification/sun_info"
	"time"
)

func main() {
	settings.EnsureEnvValues()
	if len(os.Args) >= 4 {
		var currentDate time.Time = time.Now().UTC()
		var currentDateFormat string = currentDate.Format("2006/01/02")

		db.Init()

		var isDateAlreadyAdded bool = db.IsDateAlreadyAdded(currentDateFormat)
		if !isDateAlreadyAdded {
			var latitude string = os.Args[1]
			var longitude string = os.Args[2]
			var phoneNumber string = os.Args[3]

			var sunset string
			var twilightEnd string

			sunset, twilightEnd = sunInfo.GetSunriseSunsetInfo(latitude, longitude)

			if isTimeToSend(currentDate, currentDateFormat, sunset, twilightEnd) {

				var message string = formatMessage(currentDateFormat, sunset, twilightEnd)

				if notification.SendNotification(message, phoneNumber) {
					db.AddSunInfo(currentDateFormat, sunset, twilightEnd, message)
				}
				os.Exit(0)
			} else {
				log.Println("Not able to send yet")
			}

		} else {
			log.Printf(`The date [%s] is already processed`, currentDateFormat)
		}
	} else {
		log.Fatal("Not enough parameters, [ latitude, longitude and phoneNumber are expected]")
		panic(1)
	}
}

func formatMessage(date string, sunset string, twilightEnd string) string {
	var timeFormat string = "3:04:05 PM"
	utcOffset := settings.GetUtcHourOffset()
	parsedUtcOffset, _ := time.ParseDuration(fmt.Sprintf("%dh", utcOffset))

	parsedSunset, _ := time.Parse(timeFormat, sunset)
	parsedSunset = parsedSunset.Add(parsedUtcOffset)

	parsedTwilightEnd, _ := time.Parse(timeFormat, twilightEnd)
	parsedTwilightEnd = parsedTwilightEnd.Add(parsedUtcOffset)

	message := fmt.Sprintf("%s - %s -> %s", date, parsedSunset.Format(timeFormat), parsedTwilightEnd.Format(timeFormat))

	return message
}

func isTimeToSend(dateTime time.Time, date string, sunset string, twilightEnd string) bool {
	parsedSunset, _ := time.Parse("2006/01/02 3:04:05 PM", fmt.Sprintf(`%s %s`, date, sunset))
	parsedTwilightEnd, _ := time.Parse("2006/01/02 3:04:05 PM", fmt.Sprintf(`%s %s`, date, twilightEnd))

	return dateTime.After(parsedSunset) && dateTime.Before(parsedTwilightEnd)
}
