package settings

import (
  "log"
  "os"
  "regexp"
  "bufio"
  "slices"
  "strings"
)

type EnvConfigItem struct {
	Key   string
	Value string
}

func GetSetting(key string) string {
  readedConfigs, err := readEnvFile("./.env")
  if err != nil {
    log.Println("")
   }

  value := getConfigValue(readedConfigs, key)

  return value
}

func readEnvFile(filePath string) ([]EnvConfigItem, error) {
	readFile, err := os.Open(filePath)
	// common.CheckError(err, "readEnvFile")
  if err != nil {
    log.Println(err)
    os.Exit(0)
  }

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	configs := []EnvConfigItem{}

	var re = regexp.MustCompile(`(?mi)^(?P<key>\w+)(\=+)(?P<value>.*)$`)

	for _, line := range fileLines {
		matches := re.FindAllStringSubmatch(line, -1)
		if len(matches) == 1 && len(matches[0]) > 0 {
			configs = append(configs, EnvConfigItem{Key: matches[0][1], Value: matches[0][3]})
		}
	}
	return configs, err
}

func getConfigValue(configs []EnvConfigItem, key string) string {
	idx := slices.IndexFunc(configs, func(c EnvConfigItem) bool { return strings.ToLower(c.Key) == strings.ToLower(key) })
	if idx > -1 {
		return configs[idx].Value
	}
	return ""
}
