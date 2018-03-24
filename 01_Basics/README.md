# gocv basics

This directory contains source files which demonstrate basic usage of [gocv](https://gocv.io/) `Go` package.

# Read image

We will demonstrate reading images using an image of my favorite football player, [Lionel Messi](https://en.wikipedia.org/wiki/Lionel_Messi):

<img src="./messi.jpg" alt="Colored image of Messi" width="200">

Read in color image:

```go
	imgPath := filepath.Join("messi.jpg")
	// read in a color image
	img := gocv.IMRead(imgPath, gocv.IMReadColor)
	if img.Empty() {
		fmt.Printf("Failed to read image %s\n", imgPath)
		os.Exit(1)
	}
```

`gocv` by default reads in image in *BGR* color scheme. Why? Because of historical [OpenCV](https://docs.opencv.org/trunk/index.html) reasons. So, don't forget to convert the read image into particular color scheme you are interested in working with!

Notice that `gocv.IMRead` accepts `gocv.IMReadFlag` as a second argument. There is a wide range of `gocv.IMReadFlag`s which can simplify reading different types of images. The above code shows how to read a colored image. `gocv.IMReadColor` flag won't read in alpha channel; you need to pass in `gocv.IMReadUnchanged` flag to read in all image channels.

If you want to read a grayscale image or a colored image as grayscale, simply pass in `gocv.IMReadGrayScale` flag.Here is a grayscale image of Messi:

<img src="./gray_messi.jpg" alt="Grayscale image of Messi" width="200">

# Write image

Write image to filesystem:

```go
	// write an image into filesystem
	outPath := filepath.Join("new_messi.jpg")
	if ok := gocv.IMWrite(outPath, img); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
```

Note there is also `gocv.IMWriteWithParams` function which allows you to pass various image write [parameters](https://docs.opencv.org/master/d4/da8/group__imgcodecs.html#ga292d81be8d76901bff7988d18d2b42ac)
