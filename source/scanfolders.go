package source

import (
	"fmt"
	"github.com/Mimerel/go-utils"
	"time"
)

func scanFolder() (file go_utils.Files, err error) {
	for index, _ := range config.Folders {
		sf := go_utils.NewParams()
		sf.Request.Path = config.Folders[index].Origin
		sf.Request.Ignore = []string{".grab"}
		sf.Request.Extensions = config.OriginExtensions
		sf.Request.MinAge = config.MinimumFileAge * time.Hour
		sf.ScanFolder()
		if len(sf.Result.Errors) == 0 && len(sf.Result.Files) > 0 {
			if config.FromEnd {
				return *sf.Result.Files[len(sf.Result.Files)-1], nil

			} else {
				return *sf.Result.Files[0], nil
			}
		}
	}
	return file, fmt.Errorf("No file found. Errors %+v")
}
