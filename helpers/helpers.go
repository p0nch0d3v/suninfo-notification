package helpers

import (
	"fmt"
	"suninfo-notification/settings"
	"time"
)

func FormatMessage(date string, sunset string, twilightEnd string) string {
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

func IsTimeInThreshold(dateTime time.Time, date string, sunset string, twilightEnd string) bool {
	minutesBefore := settings.GetNotifyBeforeInMinutes()
	var afterAssert time.Time
	var beforeAssert time.Time
	parsedSunset, _ := time.Parse("2006/01/02 3:04:05 PM", fmt.Sprintf(`%s %s`, date, sunset))
	parsedTwilightEnd, _ := time.Parse("2006/01/02 3:04:05 PM", fmt.Sprintf(`%s %s`, date, twilightEnd))

	if minutesBefore != 0 {
		parsedMinutesBefore, _ := time.ParseDuration(fmt.Sprintf("%dm", minutesBefore))
		afterAssert = parsedSunset.Add(parsedMinutesBefore)
		beforeAssert = parsedTwilightEnd
	} else {
		afterAssert = parsedSunset
		beforeAssert = parsedTwilightEnd
	}

	return dateTime.After(afterAssert) && dateTime.Before(beforeAssert)
}
