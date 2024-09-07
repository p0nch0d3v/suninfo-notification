package sunInfo

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"suninfo-notification/log"
	"suninfo-notification/models"
)

func GetSunriseSunsetInfo(latitude string, longitude string) (string, string) {
	url := fmt.Sprintf("https://api.sunrise-sunset.org/json?lat=%s&lng=%s", latitude, longitude)
	log.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		log.FatalErr(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.FatalErr(err)
	}

	var results models.SunrisSsunsetResult
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.FatalErr(err)
	}

	return results.Results.Sunset, results.Results.CivilTwilightEnd
}
