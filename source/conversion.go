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
	fileOut := strings.Replace(fileIn, ".ts", config.DestinationExtension, -1)

	config.Logger.Info("Output filename : %s", fileOut)

	args := []string{"-i",fileIn,"-o",fileOut}
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
	if cmd.Stdout != nil {
		config.Logger.Info("Output HandBrakeCLI %v", cmd.Stdout)
	}
	fmt.Printf("out %+v %+v", cmd.Stderr, cmd.Stdout)
	config.Logger.Info("finished converting : %s", fileIn)
	config.Logger.Info("removing file : %s", fileIn)
	os.Remove(fileIn)
	SendProwlNotification("End", fileIn)
	return nil
}

func SendProwlNotification(action string, fileIn string) {
	var params *go_utils.HttpRequestParams
	params.Url = config.Prowl + "/Plex_Transcode/" + action +"/" + fileIn
	params.Method = "POST"
	_, _ = go_utils.HttpExecuteRequest(params)
}