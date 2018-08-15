package main

import (
	"fmt"
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
	for id, url := range sc.GetUsersAvatars() {
		log.Printf("User %s has avatr image url %s", id, url)
		img := loadImageFromUrl(url)
		faces := finder.Detect(img)
		if lf := len(faces); lf == 1 {
			log.Printf("User %s has good avatar", id)
		} else {
			log.Printf("User %s: expected 1 face on avatar image, got %d", id, lf)
			sc.NotifyUser(id)
		}
	}
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
