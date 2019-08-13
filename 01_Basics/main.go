package main

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"

	"gocv.io/x/gocv"
)

func main() {
	imgPath := filepath.Join("messi.jpg")
	// read in a color image
	img := gocv.IMRead(imgPath, gocv.IMReadColor)
	//img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)
	if img.Empty() {
		fmt.Printf("Failed to read image: %s\n", imgPath)
		os.Exit(1)
	}
	// size of an image
	fmt.Printf("%s size: %d x %d\n", imgPath, img.Rows(), img.Cols())
	fmt.Printf("%s channels: %d\n", imgPath, img.Channels())
	// split image channels
	bgr := gocv.Split(img)
	// gocv.Merge(bgr[2], bgr[1], bgr[0], &img)
	// image type; for colored image: gocv.MatTypeCV8UC3
	fmt.Printf("Image type: %v\n", img.Type())
	// BGR pixel values
	fmt.Printf("Pixel B: %d\n", bgr[0].GetUCharAt(100, 100))
	fmt.Printf("Pixel G: %d\n", bgr[1].GetUCharAt(100, 100))
	fmt.Printf("Pixel R: %d\n", bgr[2].GetUCharAt(100, 100))
	// Same as the last channel value i.e. R
	fmt.Printf("Pixel (same as R): %d\n", img.GetUCharAt(100, 100))
	// ball image region: image coordinates have been selected manually
	rec := image.Rectangle{Min: image.Point{214, 383}, Max: image.Point{292, 460}}
	ball := img.Region(rec)
	// This will blur the selected ball region
	gocv.GaussianBlur(ball, &ball, image.Pt(35, 35), 0, 0, gocv.BorderDefault)
	// draw a border around image
	blue := color.RGBA{B: 255}
	gocv.CopyMakeBorder(img, &img, 10, 10, 10, 10, gocv.BorderConstant, blue)
	// write an image into filesystem
	filename := "border_blur_messi.jpg"
	//outPath := filepath.Join("new_messi.jpg")
	//outPath := filepath.Join("gray_messi.jpg")
	//outPath := filepath.Join("blur_messi.jpg")
	outPath := filepath.Join(filename)
	if ok := gocv.IMWrite(outPath, img); !ok {
		fmt.Printf("Failed to write image: %s\n", filename)
		os.Exit(1)
	} else {
		fmt.Printf("%s is stored", filename)
	}
}
