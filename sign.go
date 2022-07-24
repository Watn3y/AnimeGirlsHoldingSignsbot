package main

import (
	"bytes"
	"github.com/fogleman/gg"
	"image"
	"image/png"
	"regexp"
	"strings"
)

var maxWidth int
var maxHeight int
var anchorX int
var anchorY int
var colour string
var fontsize float64
var currentImage image.Image
var wordcount int

func isSign(cmd string) bool {
	sign, err := regexp.MatchString("^sign[0-9]+$", cmd)
	if cmd == "yannik" {
		sign = true
	}
	if err != nil {
		PrintErr("Failed to regex command", err, false)
	}
	return sign
}

func selectImage(cmd string, text string) []byte {

	if text == "" {
		text = "Brudi"
	}

	wordcount := countWords(text)
	println("Wordcount: ", wordcount)

	switch cmd {
	case "sign1":
		sign1()
	case "sign2":
		sign2()
	case "sign3":
		sign3()
	case "sign4":
		sign4()
	case "sign5":
		sign5()
	case "sign6":
		sign6()
	case "sign7":
		sign7()
	case "yannik":
		signYannik()

	default:
		sign1()
	}

	textImage := addText(text)

	imageBytes := drawImage(textImage)

	return imageBytes

}

func sign1() {
	maxWidth = 347
	maxHeight = 514
	anchorX = 180
	anchorY = 133
	colour = "#000000"
	fontsize = 18
	currentImage = image1
}
func sign2() {
	maxWidth = 160
	maxHeight = 192
	anchorX = 293
	anchorY = 536
	colour = "#000000"
	fontsize = 13
	currentImage = image2
}
func sign3() {
	maxWidth = 1100 - 390
	maxHeight = 1570 - 1130
	anchorX = 390
	anchorY = 1130
	colour = "#000000"
	fontsize = 13
	currentImage = image3

}
func sign4() {
	maxWidth = 700 - 190
	maxHeight = 1090 - 810
	anchorX = 190
	anchorY = 810
	colour = "#000000"
	fontsize = 13
	currentImage = image4
}
func sign5() {
	maxWidth = 235
	maxHeight = 80
	anchorX = 280
	anchorY = 200
	colour = "#000000"
	fontsize = 13
	currentImage = image5
}
func sign6() {
	maxWidth = 260 - 65
	maxHeight = 490 - 385
	anchorX = 65
	anchorY = 385
	colour = "#000000"
	fontsize = 9
	currentImage = image6
}
func sign7() {
	maxWidth = 440
	maxHeight = 40
	anchorX = 270
	anchorY = 20
	colour = "#000000"
	fontsize = 13
	currentImage = image7
}
func signYannik() {
	maxWidth = 720 - 25
	maxHeight = 115 - 5
	anchorX = 25
	anchorY = 5
	colour = "#000000"
	fontsize = 15
	currentImage = yannikImage
}

func countWords(s string) int {
	wordcount = len(strings.Fields(s))
	return wordcount
}

func addText(text string) image.Image {
	textImage := gg.NewContext(maxWidth, maxHeight)

	textImage.SetHexColor(colour)

	err := textImage.LoadFontFace("assets/fonts/arial-bold.ttf", fontsize)
	if err != nil {
		PrintErr("Failed to load font", err, false)
	}

	textImage.DrawStringWrapped(text, 0, 0, 0, 0, float64(maxWidth), 1.5, gg.AlignLeft)

	return textImage.Image()

}

func drawImage(textImage image.Image) []byte {

	girlImage := gg.NewContextForImage(currentImage)

	girlImage.DrawImage(textImage, anchorX, anchorY)

	buf := new(bytes.Buffer)

	err := png.Encode(buf, girlImage.Image())
	if err != nil {
		PrintErr("Failed to encode png", err, false)
	}

	return buf.Bytes()

}
