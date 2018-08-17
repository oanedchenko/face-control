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
	NotifyUser(id string, msg string)
	CheckChannelMembersAvatars(channel string, processProfileFn func(user *slack.User))
}

func NewSlack(slackToken string) SlackClient {
	api := slack.New(slackToken)
	//	api.SetDebug(true)
	//	api.GetChannels(true)
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

func (sc *slackClient) CheckChannelMembersAvatars(channel string, processProfileFn func(user *slack.User)) {
	// todo: use pagination instead!
	if channel, err := sc.api.GetChannelInfo(channel); err != nil {
		panic(err)
	} else {
		for i, mId := range channel.Members {
			log.Println(i, mId)
			if user, err := sc.api.GetUserInfo(mId); err != nil {
				panic(err)
			} else {
				processProfileFn(user)
			}
		}
	}
}

func (sc *slackClient) NotifyUser(id string, msg string) {
	params := slack.PostMessageParameters{}
	if _, _, err := sc.api.PostMessage(id, msg, params); err != nil {
		log.Printf("Fail to notify user %s: %+v", id, err)
	}
}
