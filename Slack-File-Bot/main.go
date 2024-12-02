package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main()  {

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Readme.md"}
	for i := 0; i<len(fileArr); i++{
		params := slack.FileUploadParameters{
			Channels : channelArr,
			File : fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err!=nil{
			fmt.Printf("Error: %s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}