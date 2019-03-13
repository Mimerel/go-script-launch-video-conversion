package source

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

/**
Method that reads the configuration file.
If a environment variable is set, the program will read the configuration
file from the path provided otherwize it will use the path coded in hard
 */
func readConfiguration() {
	pathToFile := os.Getenv("LOGGER_CONFIGURATION_FILE")
	if _, err := os.Stat("/root/go/src/go-script-launch-video-conversion/configuration.yaml"); !os.IsNotExist(err) {
		pathToFile = "/root/go/src/go-script-launch-video-conversion/configuration.yaml"
	} else {
		pathToFile = "./configuration.yaml"
	}
	yamlFile, err := ioutil.ReadFile(pathToFile)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Configuration Loaded : %+v \n", config)
	}
}
