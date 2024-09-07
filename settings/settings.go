package settings

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"suninfo-notification/models"
)

func getSetting(key string) string {
	value := os.Getenv(key)
	// No environment value
	if len(strings.Trim(value, " ")) == 0 {
		// Try with local env file
		readedConfigs, err := readEnvFile(".env.local")

		if err == nil || len(readedConfigs) > 0 {
			value = getConfigValue(readedConfigs, key)
		}

		if err != nil || len(readedConfigs) == 0 {
			// Try with env file
			readedConfigs, err = readEnvFile(".env")
			if err == nil || len(readedConfigs) > 0 {
				value = getConfigValue(readedConfigs, key)
			}
		}

	}
	return value
}

func readEnvFile(filePath string) ([]models.EnvConfigItem, error) {
	readFile, err := os.Open(filePath)

	if err != nil {
		log.Println(err)
		return []models.EnvConfigItem{}, err
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
	return os.Getenv(key)
}

func GetUtcHourOffset() int64 {
	utcHourOffsetEnvValue := getSetting("UTC_HOUR_OFFSET")
	utcOffset, _ := strconv.ParseInt(utcHourOffsetEnvValue, 10, 64)
	return utcOffset
}

func GetTwilioSettings() (string, string, string, bool) {
	accountSid := getSetting("TWILIO_ACCOUNT_SID")
	authToken := getSetting("TWILIO_AUTH_TOKEN")
	fromNumber := getSetting("TWILIO_AUTH_FROM_NUMBER")
	byPassTwilio, _ := strconv.ParseBool(getSetting("TWILIO_BYPASS"))

	return accountSid, authToken, fromNumber, byPassTwilio
}

func EnsureEnvValues() {
	if !ensureEnvValue("UTC_HOUR_OFFSET") || !ensureEnvValue("TWILIO_ACCOUNT_SID") || !ensureEnvValue("TWILIO_AUTH_TOKEN") || !ensureEnvValue("TWILIO_AUTH_FROM_NUMBER") {
		os.Exit(1)
	}
}

func ensureEnvValue(key string) bool {
	value := getSetting(key)
	if len(value) == 0 {
		log.Printf("Missing Env Value [%s]", key)
		return false
	}
	return true
}
