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
	// params := slack.PostMessageParameters{}
	// // attachment := slack.Attachment{
	// // 	Pretext: "some pretext",
	// // 	Text:    "some text",
	// // }
	// // params.Attachments = []slack.Attachment{attachment}
	// channelID, timestamp, err := api.PostMessage("CAB37830E", "奥信先生まじパネェ", params)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.TeamJoinEvent:
			fmt.Println("join a new member!")
			fmt.Println(ev.User.Name)
		}
	}
}
