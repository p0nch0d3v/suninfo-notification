package main

import (
	"fmt"
	"log"
	"os"
	"sunrise-sunset-notification/db"
	"sunrise-sunset-notification/notification"
	"sunrise-sunset-notification/settings"
	sunInfo "sunrise-sunset-notification/sun_info"
	"time"
)

func main() {
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

			var message string = formatMessage(sunset, twilightEnd)

			notification.SendNotification(message, phoneNumber)

			db.AddSunInfo(currentDateFormat, sunset, twilightEnd)

			os.Exit(0)
		} else {
			log.Println(fmt.Sprintf(`The date [%s] is already processed`, currentDateFormat))
		}
	} else {
		log.Fatal("Not enough parameters, [ latitude, longitude and phoneNumber are expected]")
		panic(1)
	}
}

func formatMessage(sunset string, twilightEnd string) string {
	var timeFormat string = "3:04:05 PM"
	utcOffset := settings.GetUtcHourOffset()
	parsedUtcOffset, _ := time.ParseDuration(fmt.Sprintf("%dh", utcOffset))

	parsedSunset, _ := time.Parse(timeFormat, sunset)
	parsedSunset = parsedSunset.Add(parsedUtcOffset)

	parsedTwilightEnd, _ := time.Parse(timeFormat, twilightEnd)
	parsedTwilightEnd = parsedTwilightEnd.Add(parsedUtcOffset)

	message := fmt.Sprintf("%s -> %s", parsedSunset.Format(timeFormat), parsedTwilightEnd.Format(timeFormat))

	return message
}
