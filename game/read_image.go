package game

import (
	"image"
	_ "image/png"
	"log"
	"os"
)

func readImage(path string) image.Image {
	reader, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return m
}
