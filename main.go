package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sunrise-sunset-notification/settings"
	"time"

	suninfo "sunrise-sunset-notification/sun_info"
)

func main() {
	fmt.Println("main")
	if len(os.Args) >= 4 {
		var latitude string = os.Args[1]
		var longitude string = os.Args[2]
		var phoneNumber string = os.Args[3]

		fmt.Println(latitude, longitude, phoneNumber)

		utcOffset, _ := strconv.ParseInt(settings.GetSetting("UTC_HOUR_OFFSET"), 10, 64)

		parsedUtcOffset, _ := time.ParseDuration(fmt.Sprintf("%dh", utcOffset))

		sunset, twilightEnd := suninfo.GetSunriseSunsetInfo(latitude, longitude)
		log.Println(sunset, twilightEnd)

		parsedSunset, _ := time.Parse("3:04:05 PM", sunset)
		fmt.Println(parsedSunset.Add(parsedUtcOffset))

		parsedTwilightEnd, _ := time.Parse("3:04:05 PM", twilightEnd)
		fmt.Println(parsedTwilightEnd.Add(parsedUtcOffset))
	}
}
