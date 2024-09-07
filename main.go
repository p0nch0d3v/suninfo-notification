package main

import (
	"os"
	"suninfo-notification/db"
	"suninfo-notification/helpers"

	"suninfo-notification/log"
	"suninfo-notification/notification"
	"suninfo-notification/settings"
	sunInfo "suninfo-notification/sun_info"
	"time"
)

func main() {
	log.Println("Init")
	settings.EnsureEnvValues()
	db.Init()

	if len(os.Args) == 2 {
		if os.Args[1] == "list" {
			db.PrintListAll()
		}
		os.Exit(0)
	}
	if len(os.Args) >= 4 {
		var currentDate time.Time = time.Now().UTC()
		var currentDateFormat string = currentDate.Format("2006/01/02")

		var isDateAlreadyAdded bool = db.IsDateAlreadyAdded(currentDateFormat)
		if !isDateAlreadyAdded {
			var latitude string = os.Args[1]
			var longitude string = os.Args[2]
			var phoneNumber string = os.Args[3]

			var sunset string
			var twilightEnd string

			sunset, twilightEnd = sunInfo.GetSunriseSunsetInfo(latitude, longitude)

			if helpers.IsTimeInThreshold(currentDate, currentDateFormat, sunset, twilightEnd) {

				var message string = helpers.FormatMessage(currentDateFormat, sunset, twilightEnd)

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
		log.FatalStr("Not enough parameters, [ latitude, longitude and phoneNumber are expected]")
		panic(1)
	}
}
