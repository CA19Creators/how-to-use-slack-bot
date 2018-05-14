package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

type UPLOAD_FILE struct {
	ID string `json:"id"`
}

const (
	JSONFileName = "data.json"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	api := slack.New(os.Getenv("TOKEN"))
	found, _ := exists(JSONFileName)
	if found {
		bytes, err := ioutil.ReadFile(JSONFileName)
		if err != nil {
			panic(err)
		}
		var uf UPLOAD_FILE
		err = json.Unmarshal(bytes, &uf)
		if err != nil {
			panic(err)
		}
		err = api.DeleteFile(uf.ID)
		if err != nil {
			panic(err)
		}
		groups, err := api.GetChannels(false)
		if err != nil {
			panic(err)
		}
		var content string
		for _, group := range groups {
			content = content + fmt.Sprintf("#%s:\t%s\n", group.Name, group.Topic.Value)
		}
		params := slack.FileUploadParameters{
			Channels: []string{""},
			Title:    "slackの使い方",
			Content:  content,
		}
		file, err := api.UploadFile(params)
		if err != nil {
			panic(err)
		}
		var uploadFile UPLOAD_FILE
		uploadFile.ID = file.ID
		d, err := json.Marshal(uploadFile)
		if err != nil {
			panic(err)
		}
		ioutil.WriteFile(JSONFileName, d, os.ModePerm)
	}
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
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
