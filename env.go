package env

import (
	"bufio"
	"os"
	"strings"
)

/**
The map stores configurations.
*/
var config map[string]string

/**
The config file's location.
*/
var configFilePath string

/**
Reading configuration.
*/
func Env(key string) string {

	if config == nil {
		readConfigToMap()
	}
	return config[key]
}

/**
Set the configuration file path to init.
*/
func SetConfigFilePath(path string) {
	config = nil
	configFilePath = path
}

/**
Reading configurations from the given file path.
*/
func readConfigToMap() {
	var confFile *os.File
	var err error
	if len(configFilePath) == 0 {
		confFile, err = os.Open("./.env")
	} else {
		confFile, err = os.Open(configFilePath)
	}
	if err != nil {
		panic(err)
	}
	fileInput := bufio.NewScanner(confFile)

	config = make(map[string]string)

	for fileInput.Scan() {
		line := fileInput.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}
		keyValuePair := strings.Split(line, "=")
		config[strings.TrimSpace(keyValuePair[0])] = strings.TrimSpace(keyValuePair[1])
	}
}
