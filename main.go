package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
	"os"
)

func main() {

	finder := NewFinder()
	sc := NewSlack(os.Getenv("SLACK_TOKEN"))
	//sc.PrintTeamInfo()
	// #garage-go channel:
	sc.CheckChannelMembersAvatars(os.Getenv("CHANNEL_ID"), func(user *slack.User) {
		log.Printf("Processing user %s(%s %s)", user.ID, user.Name, user.RealName)
		img := loadImageFromUrl(user.Profile.Image192)
		faces := finder.Detect(img)
		if lf := len(faces); lf == 1 {
			log.Printf("User %s has good avatar", user.Name)
		} else {
			log.Printf("User %s: expected 1 face on avatar image, got %d", user.Name, lf)
			sc.NotifyUser(user.ID, fmt.Sprintf("Hi %s :) Please do not forget to set proper avatar picture!", user.Name))
		}
	})
}

func imageRecognitionTest() {
	finder := NewFinder()
	baseImage := loadImage()

	faces := finder.Detect(baseImage)
	for i, f := range faces {
		fmt.Printf("Face %d detected: %+v\r\n", i, f)
	}
}

func loadImage() image.Image {
	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		log.Fatalf("error loading image from stdin: %+v", err)
	}
	return img
}

func loadImageFromUrl(url string) image.Image {
	if resp, err := http.Get(url); err != nil {
		panic(err)
	} else if img, _, err := image.Decode(resp.Body); err != nil {
		panic(err)
	} else {
		return img
	}
}
