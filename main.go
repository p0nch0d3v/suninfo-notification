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

		utc_offset, _ := strconv.ParseInt(settings.GetSetting("UTC_HOUR_OFFSET"), 10, 64)

		parsed_utc_offset, _ := time.ParseDuration(fmt.Sprintf("%dh", utc_offset))

		sunset, twilight_end := suninfo.Get_sunrise_sunset_info(latitude, longitude)
		log.Println(sunset, twilight_end)

		parsed_sunset, _ := time.Parse("3:04:05 PM", sunset)
		fmt.Println(parsed_sunset.Add(parsed_utc_offset))

		parsed_twilight_end, _ := time.Parse("3:04:05 PM", twilight_end)
		fmt.Println(parsed_twilight_end.Add(parsed_utc_offset))
	}
}
