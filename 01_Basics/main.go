package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gocv.io/x/gocv"
)

func main() {
	imgPath := filepath.Join("messi.jpg")
	// read in a color image
	//img := gocv.IMRead(imgPath, gocv.IMReadColor)
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)
	if img.Empty() {
		fmt.Printf("Failed to read image: %s\n", imgPath)
		os.Exit(1)
	}
	// write an image into filesystem
	outPath := filepath.Join("gray_messi.jpg")
	if ok := gocv.IMWrite(outPath, img); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
}
