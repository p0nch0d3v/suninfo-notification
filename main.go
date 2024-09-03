package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type SunrisSsunsetInfo struct {
	Date        string
	Sunrise     string
	Sunset      string
	FirstLight  string
	LastLight   string
	Dawn        string
	Dusk        string
	SolarNoon  string
	GoldenHour string
	DayLength  string
	Timezone    string
	UtcOffset  int
}

type SunrisSsunsetResult struct {
	Results SunrisSsunsetInfo
	Status  string
}

func main() {
	fmt.Println("main")
	if len(os.Args) >= 4 {
		var latitude string = os.Args[1]
		var longitude string = os.Args[2]
		var phoneNumber string = os.Args[3]

		fmt.Println(latitude, longitude, phoneNumber)

		get_sunrise_sunset_info(latitude, longitude)
	}
}

func get_sunrise_sunset_info(latitude string, longitude string) {
	resp, err := http.Get(fmt.Sprintf("https://api.sunrisesunset.io/json?lat=%s&lng=%s", latitude, longitude))
	if err != nil {
		log.Fatalln(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

  var results SunrisSsunsetResult
	err = json.Unmarshal(body, &results)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(results.Results.Sunset)
  log.Println(results.Results.Dusk)
}
