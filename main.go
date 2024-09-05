package main

import (
	"fmt"
	"log"
	"os"
	"sunrise-sunset-notification/settings"
	sunInfo "sunrise-sunset-notification/sun_info"
	"time"
)

func main() {
	fmt.Println("main")
	if len(os.Args) >= 4 {
		var latitude string = os.Args[1]
		var longitude string = os.Args[2]
		var phoneNumber string = os.Args[3]

		fmt.Println(latitude, longitude, phoneNumber)

		sunset, twilightEnd := sunInfo.GetSunriseSunsetInfo(latitude, longitude)
		message := formatMessage(sunset, twilightEnd)
		log.Println(message)
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
