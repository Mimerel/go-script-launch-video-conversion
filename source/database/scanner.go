package database

import (
	"fmt"
	"github.com/Mimerel/go-utils"
	"go-script-launch-video-conversion/source/models"
	"strings"
)

func CompletScanOfFolders(config *models.Configuration) (err error) {
	config.Database = []models.Files{}
	for index, _ := range config.Folders {
		sf := go_utils.NewParams()
		sf.Request.Path = config.Folders[index].Origin
		sf.Request.Ignore = config.Ignore
		sf.ScanFolder()
		if len(sf.Result.Errors) == 0 && len(sf.Result.Files) > 0 {
			for _, file := range sf.Result.Files {
				newFile := new(models.Files)
				newFile.FileType = returnTypeOfFile(config, file)
				newFile.Path = file.Path
				newFile.FullPath = file.FullPath
				newFile.FullName = file.FullName
				newFile.Extension = file.Extension
				newFile.Name = file.Name
				config.Database = append(config.Database, *newFile)
			}
		}
	}
	return fmt.Errorf("No file found. Errors %+v")
}

func returnTypeOfFile(config *models.Configuration, file *go_utils.Files) (fileType string) {
	fileType = "Initial"
	if strings.Index(file.Path, config.ConvertedFileFolder) != -1 {
		for _, fileType := range config.ConversionParams {
			if strings.Index(file.Path, fileType.TypeName) != -1 {
				return fileType.TypeName
			}
		}
	}
	return fileType
}
