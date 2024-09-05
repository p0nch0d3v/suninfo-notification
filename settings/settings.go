package settings

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sunrise-sunset-notification/models"
)

func getSetting(key string) string {
	readedConfigs, err := readEnvFile("./.env")
	if err != nil {
		log.Println("")
	}
	value := getConfigValue(readedConfigs, key)
	return value
}

func readEnvFile(filePath string) ([]models.EnvConfigItem, error) {
	readFile, err := os.Open(filePath)

	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	configs := []models.EnvConfigItem{}
	var re = regexp.MustCompile(`(?mi)^(?P<key>\w+)(\=+)(?P<value>.*)$`)

	for _, line := range fileLines {
		matches := re.FindAllStringSubmatch(line, -1)
		if len(matches) == 1 && len(matches[0]) > 0 {
			configs = append(configs, models.EnvConfigItem{Key: matches[0][1], Value: matches[0][3]})
		}
	}
	return configs, err
}

func getConfigValue(configs []models.EnvConfigItem, key string) string {
	idx := slices.IndexFunc(configs, func(c models.EnvConfigItem) bool {
		return strings.EqualFold(c.Key, key)
	})
	if idx > -1 {
		return configs[idx].Value
	}
	return ""
}

func GetUtcHourOffset() int64 {
	utcHourOffsetEnvValue := getSetting("UTC_HOUR_OFFSET")
	utcOffset, _ := strconv.ParseInt(utcHourOffsetEnvValue, 10, 64)
	return utcOffset
}

func GetTwilioSettings() (string, string, string) {
	accountSid := getSetting("TWILIO_ACCOUNT_SID")
	authToken := getSetting("TWILIO_AUTH_TOKEN")
	fromNumber := getSetting("TWILIO_AUTH_FROM_NUMBER")
	return accountSid, authToken, fromNumber
}
