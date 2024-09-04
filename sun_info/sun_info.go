package suninfo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sunrise-sunset-notification/models"
)

func Get_sunrise_sunset_info(latitude string, longitude string) (string, string) {
	url := fmt.Sprintf("https://api.sunrise-sunset.org/json?lat=%s&lng=%s", latitude, longitude)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var results models.SunrisSsunsetResult
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatalln(err)
	}

	return results.Results.Sunset, results.Results.CivilTwilightEnd
}
