package source

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"github.com/Mimerel/go-logger-client"
)

/**
Method that reads the configuration file.
If a environment variable is set, the program will read the configuration
file from the path provided otherwize it will use the path coded in hard
 */
func readConfiguration() (err error){
	pathToFile := os.Getenv("LOGGER_CONFIGURATION_FILE")
	if _, err := os.Stat("/root/go/src/go-script-launch-video-conversion/configuration.yaml"); !os.IsNotExist(err) {
		pathToFile = "/root/go/src/go-script-launch-video-conversion/configuration.yaml"
	} else {
		pathToFile = "./configuration.yaml"
	}
	yamlFile, err := ioutil.ReadFile(pathToFile)

	if err != nil {
		return (err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		config.Logger = logs.New("", "")
		return (err)
	} else {
		if config.Production {
			config.Logger = logs.New(config.Elasticsearch.Url, config.Host)
		} else {
			config.Logger = logs.New("", "")
		}
		config.Logger.Info("Configuration Loaded : %+v \n", config)
	}
	return nil
}
