package main

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New("TOKEN")
	groups, err := api.GetGroups(false)
	if err != nil {
		panic(err)
	}

	for _, g := range groups {
		fmt.Printf("ID: %s, Name: %s\n", g.ID, g.Name)
	}
}
