package source

import (
	"bytes"
	"fmt"
	"github.com/Mimerel/go-utils"
	"os"
	"os/exec"
	"strings"
)

func startConversion(fileIn string) (err error) {
	SendProwlNotification("Start", fileIn)
	config.Logger.Info("Copying file to local /tmp/file.ts" )
	err = go_utils.CopyFileContents(fileIn, config.TemporaryFile + ".ts")
	if err != nil {
		return err
	}
	config.Logger.Info("Copying to local finished" )

	fileOut := strings.Replace(fileIn, ".ts", config.DestinationExtension, -1)

	config.Logger.Info("Output filename : %s", fileOut)

	args := []string{"-i",config.TemporaryFile + ".ts","-o", config.TemporaryFile + ".mp4"}
	args = append(args, config.Params...)
	config.Logger.Info("Running command : HandBrakeCLI %s", args)

	cmd := exec.Command("HandBrakeCLI", args...)
	var out bytes.Buffer
	var errcmd bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &errcmd
	err = cmd.Start()
	if err != nil {
		fmt.Printf("start\n")
		return err
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("wait\n")

		return err
	}
	config.Logger.Info("finished converting : %s", fileIn)
	config.Logger.Info("removing file : %s", fileIn)
	os.Remove(fileIn)

	config.Logger.Info("Copying file to Nas " )
	err = go_utils.CopyFileContents(config.TemporaryFile + ".mp4", fileOut)
	if err != nil {
		return err
	}
	config.Logger.Info("Copying to Nas finished " )
	os.Remove(config.TemporaryFile + ".mp4")

	SendProwlNotification("End", fileIn)
	return nil
}

func SendProwlNotification(action string, fileIn string) {
	var params go_utils.HttpRequestParams
	filename := strings.Replace(fileIn, "/", "-", -1)
	filename = strings.Replace(filename, " ", ".", -1)
	params.Url = config.Prowl + "/Plex_Transcode/" + action +"/" + filename
	params.Method = "POST"
	_, _ = go_utils.HttpExecuteRequest(&params)
}