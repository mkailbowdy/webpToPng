package main

import (
	"golang.org/x/image/webp"
	"image/png"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	// Get directory user's system
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	subdir := os.Args[1:]
	fullPath := filepath.Join(home, subdir[0])

	// Get all webp files in the directory
	files, err := filepath.Glob(fullPath + "/*.webp")
	if err != nil {
		log.Fatal(err)
	}

	for i, file := range files {
		// Get the current file's name, remove the .webp suffix, and replace with .png
		name := strings.TrimSuffix(filepath.Base(file), ".webp")
		img, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		// Decode the webp data, then create a new file
		webpImg, err := webp.Decode(img)
		newFile, err := os.Create(fullPath + "/" + name + strconv.Itoa(i) + ".png")
		if err != nil {
			log.Fatal(err)
		}

		// encode the data into png and write it to the new file
		err = png.Encode(newFile, webpImg)
		if err != nil {
			log.Fatal(err)
		}

		img.Close()
	}
}
