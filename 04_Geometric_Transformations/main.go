package main

import (
	"fmt"
	"image"
	"math"
	"os"
	"path/filepath"

	"gocv.io/x/gocv"
)

const x = math.Pi / 180

func Rad(d float64) float64 { return d * x }
func Deg(r float64) float64 { return r / x }

func main() {
	messiPath := filepath.Join("messi.jpg")
	// read image
	messi := gocv.IMRead(messiPath, gocv.IMReadColor)
	if messi.Empty() {
		fmt.Printf("Failed to read image: %s\n", messiPath)
		os.Exit(1)
	}
	// size of the original image
	rows, cols := messi.Rows(), messi.Cols()
	fmt.Printf("Rows: %d, Cols: %d\n", rows, cols)
	// make messi bigger
	bigger := gocv.NewMat()
	gocv.Resize(messi, &bigger, image.Point{}, 2, 2, gocv.InterpolationCubic)
	fmt.Printf("Rows: %d, Cols: %d\n", bigger.Rows(), bigger.Cols())
	// bigger image of Messi
	biggerPath := filepath.Join("bigger_messi.jpeg")
	if ok := gocv.IMWrite(biggerPath, bigger); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}

	// make messi smaller
	smaller := gocv.NewMat()
	gocv.Resize(messi, &smaller, image.Point{}, 0.5, 0.5, gocv.InterpolationArea)
	fmt.Printf("Rows: %d, Cols: %d\n", smaller.Rows(), smaller.Cols())
	// smaller image of Messi
	smallerPath := filepath.Join("smaller_messi.jpeg")
	if ok := gocv.IMWrite(smallerPath, smaller); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}

	///////////////
	// ROTATTION //
	//////////////
	// rotate messi by 90 degrees clockwise
	rotated := gocv.NewMat()
	gocv.Rotate(messi, &rotated, gocv.Rotate90Clockwise)
	// rotate messig by 90 degrees clockwise
	rotatedPath := filepath.Join("rotated_messi.jpeg")
	if ok := gocv.IMWrite(rotatedPath, rotated); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}

	// arbitrary rotation by 45 degrees
	scale, angle := 1.0, 45.0
	radAngle := Rad(angle)
	center := image.Point{cols / 2, rows / 2}
	rotation := gocv.GetRotationMatrix2D(center, angle, scale)

	// Scale -> Rotate -> Translate
	// Scale
	scaleX, scaleY := float64(cols)*scale, float64(rows)*scale
	// Rotate
	newX := math.Abs(scaleX*math.Cos(radAngle)) + math.Abs(scaleY*math.Sin(radAngle))
	newY := math.Abs(scaleX*math.Sin(radAngle)) + math.Abs(scaleY*math.Cos(radAngle))
	// Translate
	tx := (newX - float64(cols)) / 2
	ty := (newY - float64(rows)) / 2
	rotation.SetDoubleAt(0, 2, rotation.GetDoubleAt(0, 2)+tx)
	rotation.SetDoubleAt(1, 2, rotation.GetDoubleAt(1, 2)+ty)

	rotated45 := gocv.NewMat()
	gocv.WarpAffine(messi, &rotated45, rotation, image.Point{int(newX), int(newY)})

	rotated45Path := filepath.Join("rotated_properly_messi.jpeg")
	if ok := gocv.IMWrite(rotated45Path, rotated); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}

	////////////////////////////////
	// PERSPECTIVE TRANSFORMATION //
	////////////////////////////////
	cardPath := filepath.Join("card.jpg")
	card := gocv.IMRead(cardPath, gocv.IMReadColor)
	if card.Empty() {
		fmt.Printf("Failed to read image: %s\n", cardPath)
		os.Exit(1)
	}

	// image coordinages corners of the selected business card object
	origImg := []image.Point{
		image.Point{128, 165}, // top-left
		image.Point{215, 275}, // bottom-left
		image.Point{385, 128}, // bottom-right
		image.Point{300, 40},  // top-right
	}

	// calculate height as a distance between (top-left, bottom-left) and (top-right, bottom-right)
	heightA := math.Sqrt(math.Pow(float64(origImg[0].X-origImg[1].X), 2) + math.Pow(float64(origImg[0].Y-origImg[1].Y), 2))
	heightB := math.Sqrt(math.Pow(float64(origImg[3].X-origImg[2].X), 2) + math.Pow(float64(origImg[3].Y-origImg[2].Y), 2))
	height := int(math.Max(heightA, heightB))

	// caluclate width as a distance between () and ()
	widthA := math.Sqrt(math.Pow(float64(origImg[0].X-origImg[3].X), 2) + math.Pow(float64(origImg[0].Y-origImg[3].Y), 2))
	widthB := math.Sqrt(math.Pow(float64(origImg[1].X-origImg[2].X), 2) + math.Pow(float64(origImg[1].Y-origImg[2].Y), 2))
	width := int(math.Max(widthA, widthB))

	// image coordinages corners of the target object which is a result of perspective transformation
	newImg := []image.Point{
		image.Point{0, 0},
		image.Point{0, height},
		image.Point{width, height},
		image.Point{width, 0},
	}

	transform := gocv.GetPerspectiveTransform(origImg, newImg)

	perspective := gocv.NewMat()
	gocv.WarpPerspective(card, &perspective, transform, image.Point{width, height})

	outPath := filepath.Join("card_perspective.jpg")
	if ok := gocv.IMWrite(outPath, perspective); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
}
