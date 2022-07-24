package main

import (
	"github.com/fogleman/gg"
	"image"
)

var image1 image.Image
var image2 image.Image
var image3 image.Image
var image4 image.Image
var image5 image.Image
var image6 image.Image
var image7 image.Image
var yannikImage image.Image

func preloadImages() {
	var err error
	imgdir := "assets/images/"
	image1, err = gg.LoadImage(imgdir + "1.png")
	if err != nil {
		PrintErr("Failed to load image file", err, true)
	}
	image2, err = gg.LoadImage(imgdir + "2.png")
	if err != nil {
		PrintErr("Failed to load image file", err, true)
	}
	image3, err = gg.LoadImage(imgdir + "3.png")
	if err != nil {
		PrintErr("Failed to load image file", err, true)
	}
	image4, err = gg.LoadImage(imgdir + "4.png")
	if err != nil {
		PrintErr("Failed to load image file", err, true)
	}
	image5, err = gg.LoadImage(imgdir + "5.png")
	if err != nil {
		PrintErr("Failed to load image file", err, true)
	}
	image6, err = gg.LoadImage(imgdir + "6.png")
	if err != nil {
		PrintErr("Failed to load image file", err, true)
	}
	image7, err = gg.LoadImage(imgdir + "7.png")
	if err != nil {
		PrintErr("Failed to load image file", err, true)
	}
	yannikImage, err = gg.LoadImage(imgdir + "yannik.png")
	if err != nil {
		PrintErr("Failed to load image file", err, true)
	}

}
