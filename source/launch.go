package source

import (
	"fmt"
	"go-script-launch-video-conversion/source/database"
	"go-script-launch-video-conversion/source/models"
	"os"
)

var config *models.Configuration

func Launch() {

	//for {
		err := readConfiguration()
		if err != nil {
			config.Logger.Error("%+v", err)
		}
		err = database.CompletScanOfFolders(config)
		if err != nil {
			fmt.Printf("Error Scanning Folders: %v", err)
		}
		fmt.Printf("Database %+v", config.Database)
		err = scanFolder()
		if err != nil {
			config.Logger.Info("No more files to process ")
			os.Exit(0)
		}
		// else {
		//	config.Logger.Info("Will process file : %s", foundFile)
		//
		//	err = startConversion(foundFile)
		//	if err != nil {
		//		config.Logger.Error("Error searching for file to process: %v ", err)
		//	}
		//}
	//}
}
