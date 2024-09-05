package models

type SunrisSsunsetInfo struct {
	Sunrise                   string `json:"sunrise"`
	Sunset                    string `json:"sunset"`
	SolarMoon                 string `json:"solar_noon"`
	DayLength                 string `json:"day_length"`
	CivilTwilightBegin        string `json:"civil_twilight_begin"`
	CivilTwilightEnd          string `json:"civil_twilight_end"`
	NauticalTwilightBegin     string `json:"nautical_twilight_begin"`
	NauticalTwilightEnd       string `json:"nautical_twilight_end"`
	AstronomicalTwilightBegin string `json:"astronomical_twilight_begin"`
	AstronomicalTwilightEnd   string `json:"astronomical_twilight_end"`
}

type SunrisSsunsetResult struct {
	Results SunrisSsunsetInfo
	Status  string
	TzId    string
}

type EnvConfigItem struct {
	Key   string
	Value string
}

type LogItem struct {
	Id          string
	Sunset      string
	TwilightEnd string
	Message     string
}
