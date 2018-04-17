package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gocv.io/x/gocv"
)

func main() {
	// messi image
	messiPath := filepath.Join("messi.jpg")
	messi := gocv.IMRead(messiPath, gocv.IMReadColor)
	if messi.Empty() {
		fmt.Printf("Failed to read image: %s\n", messiPath)
		os.Exit(1)
	}
	// grayscale messi image
	grayMessi := gocv.NewMat()
	gocv.CvtColor(messi, &grayMessi, gocv.ColorBGRToGray)
	//rgbMessi := gocv.NewMat()
	//gocv.CvtColor(messi, &rgbMessi, gocv.ColorBGRToRGB)
	// write image to filesystem
	//outPath := filepath.Join("rgb_messi.jpeg")
	outPath := filepath.Join("gray_messi.jpeg")
	if ok := gocv.IMWrite(outPath, grayMessi); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
	// logo image
	logoPath := filepath.Join("commons.png")
	logo := gocv.IMRead(logoPath, gocv.IMReadColor)
	if logo.Empty() {
		fmt.Printf("Failed to read image: %s\n", logoPath)
		os.Exit(1)
	}
	// turn the logo to gray image
	grayLogo := gocv.NewMat()
	gocv.CvtColor(logo, &grayLogo, gocv.ColorBGRToGray)
	// binary threshold
	binLogo := gocv.NewMat()
	gocv.Threshold(grayLogo, &binLogo, 10.0, 255.0, gocv.ThresholdBinary)
	// inverse binary logo
	invBinLogo := gocv.NewMat()
	gocv.Threshold(grayLogo, &invBinLogo, 10.0, 255.0, gocv.ThresholdBinaryInv)
	// write logo to filesystem
	outBinLogo := "bin_logo.jpeg"
	if ok := gocv.IMWrite(outBinLogo, binLogo); !ok {
		fmt.Printf("Failed to write image: %s\n", outBinLogo)
		os.Exit(1)
	}
	outInvBinLogo := "inv_bin_logo.jpeg"
	if ok := gocv.IMWrite(outInvBinLogo, invBinLogo); !ok {
		fmt.Printf("Failed to write image: %s\n", outInvBinLogo)
		os.Exit(1)
	}
	// sudoku image
	sudokuPath := filepath.Join("sudoku.jpg")
	sudoku := gocv.IMRead(sudokuPath, gocv.IMReadGrayScale)
	if sudoku.Empty() {
		fmt.Printf("Failed to read image: %s\n", sudokuPath)
		os.Exit(1)
	}
	// adaptive mean
	adptMean := gocv.NewMat()
	gocv.AdaptiveThreshold(sudoku, &adptMean, 255.0, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, 5, 4.0)
	// adaptive gaussian
	adptGauss := gocv.NewMat()
	gocv.AdaptiveThreshold(sudoku, &adptGauss, 255.0, gocv.AdaptiveThresholdGaussian, gocv.ThresholdBinary, 5, 4.0)
	// write images to filesystem
	outAdptMean := "sudoku_adaptive_mean.jpeg"
	if ok := gocv.IMWrite(outAdptMean, adptMean); !ok {
		fmt.Printf("Failed to write image: %s\n", outAdptMean)
		os.Exit(1)
	}
	outAdptGauss := "sudoku_adaptive_gauss.jpeg"
	if ok := gocv.IMWrite(outAdptGauss, adptGauss); !ok {
		fmt.Printf("Failed to write image: %s\n", outAdptGauss)
		os.Exit(1)
	}
}
