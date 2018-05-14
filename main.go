package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	api := slack.New(os.Getenv("TOKEN"))

	// Delete previous　snippet
	{
		params := slack.GetFilesParameters{
			Types: "snippets",
		}
		files, _, _ := api.GetFiles(params)
		for _, file := range files {
			if file.Title == "slackの使い方" {
				err := api.DeleteFile(file.ID)
				if err != nil {
					panic(err)
				}
			}
		}
	}

	// Send a snippet with DM to a newly joined team
	{
		groups, err := api.GetChannels(false)
		if err != nil {
			panic(err)
		}
		var content string
		for _, group := range groups {
			content = content + fmt.Sprintf("#%s:\t%s\n", group.Name, group.Topic.Value)
		}
		params := slack.FileUploadParameters{
			Channels: []string{"CAPV96FJS"},
			Title:    "slackの使い方",
			Content:  content,
		}
		_, err = api.UploadFile(params)
		if err != nil {
			panic(err)
		}
	}
}

func StartSlackRTM(api *slack.Client) {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.TeamJoinEvent:
			// fmt.Println("join a new member!")
			fmt.Println(ev.User.Name)
		}
	}
}
