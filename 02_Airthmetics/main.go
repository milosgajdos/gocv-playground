package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"gocv.io/x/gocv"
)

func main() {
	messiPath := filepath.Join("messi.jpg")
	logoPath := filepath.Join("commons.png")
	// read images
	messi := gocv.IMRead(messiPath, gocv.IMReadColor)
	if messi.Empty() {
		fmt.Printf("Failed to read image: %s\n", messiPath)
		os.Exit(1)
	}
	logo := gocv.IMRead(logoPath, gocv.IMReadColor)
	if logo.Empty() {
		fmt.Printf("Failed to read image: %s\n", messiPath)
		os.Exit(1)
	}
	// get image dimensions
	messiRows, messiCols := messi.Rows(), messi.Cols()
	logoRows, logoCols := logo.Rows(), logo.Cols()
	// number of channels
	fmt.Printf("%s channels: %d, size: %dx%d\n", messiPath, messi.Channels(), messiRows, messiCols)
	fmt.Printf("%s channels: %d, size: %dx%d\n", logoPath, logo.Channels(), logoRows, logoCols)
	// select bottom left region in messi picture
	minX, minY := 0, messiRows-logoRows
	maxX, maxY := logoCols, messiRows
	rec := image.Rectangle{Min: image.Point{minX, minY}, Max: image.Point{maxX, maxY}}
	roi := messi.Region(rec)
	// add two images
	//gocv.Add(roi, logo, roi)
	//gocv.AddWeighted(roi, 0.6, logo, 0.4, 0.0, roi)
	// turn color image into grayscale
	grayLogo := gocv.NewMat()
	gocv.CvtColor(logo, grayLogo, gocv.ColorBGRToGray)
	// create a binary mask
	mask := gocv.NewMat()
	gocv.Threshold(grayLogo, mask, 10.0, 255.0, gocv.ThresholdBinary)
	// create an inverse mask
	maskInv := gocv.NewMat()
	gocv.BitwiseNot(mask, maskInv)
	// black-out the area of logo in roi i.e. in bottom left region
	roiChans := gocv.Split(roi)
	for i := 0; i < len(roiChans); i++ {
		gocv.BitwiseAnd(roiChans[i], maskInv, roiChans[i])
	}
	gocv.Merge(roiChans, roi)
	// apply the mask on logo image
	logoChans := gocv.Split(logo)
	for i := 0; i < len(logoChans); i++ {
		gocv.BitwiseAnd(logoChans[i], mask, logoChans[i])
	}
	gocv.Merge(logoChans, logo)
	// add logo to roi
	gocv.Add(roi, logo, roi)
	// write new image to filesystem
	//outPath := filepath.Join("add_logo_messi.jpeg")
	//outPath := filepath.Join("add_weighted_logo_messi.jpeg")
	//outPath := filepath.Join("wiki_commons_blackout_messi.jpeg")
	outPath := filepath.Join("wiki_commons_messi.jpeg")
	if ok := gocv.IMWrite(outPath, messi); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
}
