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
}
