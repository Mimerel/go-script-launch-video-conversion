package source

import (
	"fmt"
	"go-utils"
)

func scanFolder() (err error) {
	for _ , value := range config.Database {
		if go_utils.StringInArray(value.Extension, config.OriginExtensions) {
				err = startConversion(value)
				if err != nil {
					config.Logger.Error("Error searching for file to process: %v ", err)
					return err
				}
		}
	}
	return fmt.Errorf("No file found. Errors %+v")
}


