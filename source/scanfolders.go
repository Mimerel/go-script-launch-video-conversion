package source

import (
	"fmt"
	"github.com/Mimerel/go-utils"
)

func scanFolder() (path string, err error) {
	for index, _ := range config.Folders {
		_, files, scanErr := go_utils.ScanDirectory(config.Folders[index].Origin, config.OriginExtensions, []string{})
		if len(scanErr) == 0 && len(files) > 0 {
			return files[0], nil
		}
	}
	return "", fmt.Errorf("No file found. Errors %+v")
}
