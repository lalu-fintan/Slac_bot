package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("Slack_bot_token", "xoxb-3868005067767-3906384339216-5ItMBQj50D58Kh7Ecanld9S0")
	os.Setenv("Channel_ID", "C03RRUCMYGN")
	api := slack.New(os.Getenv("Slack_bot_token"))
	channelArr := []string{os.Getenv("Channel_ID")}
	fileArr := []string{"gobook.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name:%s,URL:%s\n", file.Name, file.URL)
	}

}
