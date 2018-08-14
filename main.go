package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

func main() {
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
