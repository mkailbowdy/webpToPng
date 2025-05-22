package main

import (
	"flag"
	"fmt"
	"golang.org/x/image/webp"
	"image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	mime := flag.String("mime", "jpeg", "MIME type to convert to")
	q := flag.Int("q", 100, "Quality of image")
	flag.Parse()
	fmt.Printf("Mime type: %s\n", *mime)
	options := jpeg.Options{Quality: *q}
	// Get directory user's system
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	subdir := os.Args[1:]
	fmt.Printf("subdir path: %s\n", subdir)

	fullPath := filepath.Join(home, subdir[1])

	// Get all webp files in the directory
	files, err := filepath.Glob(fullPath + "/*.webp")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	for i, file := range files {
		// Get the current file's name, remove the .webp suffix, and replace with .png
		name := strings.TrimSuffix(filepath.Base(file), ".webp")
		img, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		// Decode the webp data, then create a new file
		webpImg, err := webp.Decode(img)

		var newFile *os.File
		if *mime == "jpeg" {
			log.Printf("jpeg")
			newFile, err = os.Create(fullPath + "/" + name + strconv.Itoa(i) + ".jpg")
			err = jpeg.Encode(newFile, webpImg, &options)
			if err != nil {
				fmt.Println(err)

				log.Fatal(err)
			}
		} else {
			log.Printf("png")
			newFile, err = os.Create(fullPath + "/" + name + strconv.Itoa(i) + ".png")
			if err != nil {
				fmt.Println(err)

				log.Fatal(err)
			}

			// encode the data into png and write it to the new file
			err = png.Encode(newFile, webpImg)
			if err != nil {
				fmt.Println(err)

				log.Fatal(err)
			}

		}
		img.Close()
	}
}
