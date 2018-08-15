package main

import (
	"github.com/nlopes/slack"
	"log"
)

type slackClient struct {
	api *slack.Client
}

type SlackClient interface {
	PrintTeamInfo()
	// user id -> image url
	GetUsersAvatars() map[string]string
	NotifyUser(id string)
}

func NewSlack(slackToken string) SlackClient {
	api := slack.New(slackToken)
	api.SetDebug(true)
	return &slackClient{
		api: api,
	}
}

func (sc *slackClient) PrintTeamInfo() {
	if teamInfo, err := sc.api.GetTeamInfo(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Team info received: %+v", *teamInfo)
	}
}

func (sc *slackClient) GetUsersAvatars() map[string]string {
	// todo: use pagination instead!
	if users, err := sc.api.GetUsers(); err != nil {
		panic(err)
	} else {
		m := make(map[string]string, len(users))
		for i, u := range users {
			log.Println(i, u)
			if !u.IsBot {
				m[u.ID] = u.Profile.Image192
			}
		}
		return m
	}
}

func (sc *slackClient) NotifyUser(id string) {
	params := slack.PostMessageParameters{}
	if _, _, err := sc.api.PostMessage(id, "Hi :) Please do not forget to set proper avatar picture!", params); err != nil {
		log.Printf("Fail to notify user %s: %+v", id, err)
	}
}
