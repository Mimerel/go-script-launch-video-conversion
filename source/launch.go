package source

import (
	"time"
	"github.com/Mimerel/go-logger-client"
)

type Elasticsearch struct {
	Url string `yaml:"url,omitempty"`
}

type Folders struct {
	Origin      string `yaml:"origin,omitempty"`
	Destination string `yaml:"destination,omitempty"`
}

type Configuration struct {
	Elasticsearch        Elasticsearch `yaml:"elasticSearch,omitempty"`
	Host                 string        `yaml:"host,omitempty"`
	Folders              []Folders     `yaml:"folders,omitempty"`
	OriginExtensions     []string      `yaml:"originExtensions,omitempty"`
	DestinationExtension string        `yaml:"destinationExtensions,omitempty"`
	Params               []string      `yaml:"params,omitempty"`
	Logger               logs.LogParams
}

var config *Configuration

func Launch() {

	for {
		err:= readConfiguration()
		if err != nil {
			config.Logger.Error("%+v",err)
		}


		foundFile, err := scanFolder()
		if err != nil {
			config.Logger.Info("No more files to process ")
			time.Sleep(1 * time.Minute)
		} else {
			config.Logger.Info("Will process file : %s", foundFile)

			err = startConversion(foundFile)
			if err != nil {
				config.Logger.Error("Error searching for file to process: %v ", err)
			}
		}
	}
}
