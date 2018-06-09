# gocv basics

This chapter covers the basics of working with [gocv](https://gocv.io/). We will look at how to read and write digital images using `gocv` and how they're represented in `gocv`'s API. Armed with this knowledge, we will proceed to slightly more advanced topics such as image region selection, image blurring and image borders.

# Read image

Reading digital images using `gocv` is super easy. Let's show how to do this on an image of my favorite football player: [Lionel Messi](https://en.wikipedia.org/wiki/Lionel_Messi) [1](https://commons.wikimedia.org/wiki/File:Lionel_Messi_Player_of_the_Year_2011.jpg):

<img src="./messi.jpg" alt="Colored image of Messi" width="200">

Read color image:

```go
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
	if img.Empty() {
		fmt.Printf("Failed to read image: %s\n", imgPath)
		os.Exit(1)
	}
}
```

Format of the digital image is determined by the content of the file it is stored in, in particular the format is determined by reading first few bytes of the image file. By default `gocv` reads an image in **BGR** color scheme, **not RGB**. Why? Because of some historical [OpenCV](https://docs.opencv.org/trunk/index.html) [reasons](https://stackoverflow.com/questions/14556545/why-opencv-using-bgr-colour-space-instead-of-rgb). So, don't forget to convert the image into a color scheme you are interested in working with!

Notice that `gocv.IMRead()` function accepts `gocv.IMReadFlag` as the second parameter. There is a wide range of `gocv.IMReadFlag`s which can simplify reading different types of images. The above code shows an example of how to read a colored image by passing in `gocv.IMReadColor` flag to `gocv.IMRead()`.

If you want to read a grayscale image, or a colored image as grayscale, simply pass in `gocv.IMReadGrayScale` flag to `gocv.IMRead()`. This will read the image and automatically convert it to grayscale color map:

```go
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
	// read in a colored image as grayscale
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)
	if img.Empty() {
		fmt.Printf("Failed to read image: %s\n", imgPath)
		os.Exit(1)
	}
}
```

Here is a grayscale image of Messi read in using the code snippet showed above:

<img src="./gray_messi.jpg" alt="Grayscale image of Messi" width="200">

# Write image

Writing an image to filesystem is almost as easy as reading an image:

```go
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
	// read in a colored image as grayscale
	img := gocv.IMRead(imgPath, gocv.IMReadGrayScale)
	if img.Empty() {
		fmt.Printf("Failed to read image: %s\n", imgPath)
		os.Exit(1)
	}
	// write an image into filesystem
	outPath := filepath.Join("new_messi.jpg")
	if ok := gocv.IMWrite(outPath, img); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
}
```

The format of the image written to the filesystem is determined by the provided file extension i.e. `messi.jpg` or `messi.jpeg` would save the image in the filesystem as [JPEG](https://en.wikipedia.org/wiki/JPEG) image.

Note, there is also `gocv.IMWriteWithParams` function provided by `gocv` which allows you to pass various image write [parameters](https://docs.opencv.org/master/d4/da8/group__imgcodecs.html#ga292d81be8d76901bff7988d18d2b42ac) that can modify the image before it's written to the filesystem.

# Image matrix

`gocv.IMRead()` returns an image as `gocv.Mat` object which is a multidimensional array whose elements store values of pixel intensities of the image. `gocv.Mat` type provides a lot of useful functions to work with the image. Let's look at the most basic functions that allow you to obtain various kinds of useful information about read images.

## Image size

You can find out image size by using `Rows()` and `Cols()` functions:

```go
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
	if img.Empty() {
		fmt.Printf("Failed to read image: %s\n", imgPath)
		os.Exit(1)
	}
	// size of an image
	fmt.Printf("%s size: %d x %d\n", imgPath, img.Rows(), img.Cols())
}
```

This will print out the following output:

```console
messi.jpg size: 480 x 388
```

## Image channels

Every digital image has a certain number of [channels](https://en.wikipedia.org/wiki/Channel_(digital_image)). Colored images have R (Red), G (Green), B (Blue) and often alpha (opacity) channel. A grayscale image, as you would expec, has only one channel. You can obtain the number of image channels using `gocv.Channels()` function.

Color image:
```go
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
	if img.Empty() {
		fmt.Printf("Failed to read image: %s\n", imgPath)
		os.Exit(1)
	}
	// image channels
	fmt.Printf("%s channels: %d\n", imgPath, img.Channels())
}
```

This will print out the following output:

```console
messi.jpg channels: 3
```

`gocv` provides `gocv.Split()` function which allows you to extract particular channel from the image so you can work with it separately from the other channels. You can merge the split channels  back together in any order chosen by you using `gocv.Merge()` function. The following code converts `BGR` image to `RGB` one:

```go
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
	if img.Empty() {
		fmt.Printf("Failed to read image: %s\n", imgPath)
		os.Exit(1)
	}
	// split image channels: remember IMRead reads images in BGR scheme
	bgr = gocv.Split(img)
	// reorder image channels from BGR -> RGB
	gocv.Merge(bgr[2], bgr[1], bgr[0], &img)
}
```

## Image type

Every image read using OpenCV has a certain *image type*. Image type is not the same thing as image *format* (JPEG, PNG etc.). The type of an image is more related to things like color scheme etc. It's important to know the type of an image when accessing and manipulating image pixels using `gocv`.

Finding out the image type is simple:

```go
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
	// image type
	fmt.Printf("image type: %d\n", img.Type())
}
```

This will print out the following output:

```console
image type: 16
```

The returned value above is an `integer`, which unfortunately does not implement `Stringer()` so you need to dig into the documentation and source code to find out what does the value actually mean. Luckily it's not that hard to figure it out.

In the case of our color image, the image type would be `gocv.MatTypeCV8UC3` which basically means we have an unsigned integer 3 channels image. You can read more about it [here](https://stackoverflow.com/a/27184054/569763).

If we tried to find out the type of a grayscale image we would find out the type to be `gocv.MatTypeCV8UC1` i.e. an unsigned integer 1 channel image.

### Image pixels

Just like OpenCV, `gocv` provides a dedicated type to store pixel intensity values: `gocv.Scalar`. However, at the moment of this writing, it falls short in providing a conveniet function to return the pixel intensity values for all available color channels in a single `go` slice; when using `python opencv` you can conveniently request pixel values across all image channels by simply indexing an image matrix i.e. `img[100,100]` would give you a list of all pixel intensity values on the requested pixel position in the image.

Getting the pixel intensity values for each channel using `gocv` requires a bit more work. Let's say we would like to find out a specific pixel values across all image channels. We need to do something like this:

```go
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
	// split image channels
	bgr := gocv.Split(img)
	// pixel values for each channel - we know this is a BGR image
	fmt.Printf("Pixel B: %d\n", bgr[0].GetUCharAt(100, 100))
	fmt.Printf("Pixel G: %d\n", bgr[1].GetUCharAt(100, 100))
	fmt.Printf("Pixel R: %d\n", bgr[2].GetUCharAt(100, 100))
}
```

This will print out the following output:

```consol
Pixel B: 37
Pixel G: 44
Pixel R: 41
```

Notice a few things:
* We know that the image has three channels, ordered as BGR
* We also know, that the image type is `gocv.MatTypeCV8UC3`

This basically means we can access particular pixel intensity value using `GetUCharAt()` function of type `gocv.Mat`. Equivalently, we can also modify particular pixel intensity values using `SetUCharAt()` function.

If we simply try to use `img.GetUCharAt(100, 100)` i.e. if we didn't separate the color channels, `gocv` would return only the last channel's value: in this case that would be R channel since the image is read in BGR color scheme.

### Image region

Often we need to select a small region in the image and do something with it eg. when you detect an object you might want to blur it or apply some transformation on it or whatnot. `gocv.Mat` provides a conveniently named function called `Region()`. You can select a rectangle region using the standard library `image.Rectangle`.

Let's demonstrate this on a simple example using our colored image. We will select a region of the image that contains football and we will blur it. Don't worry about the clurring part of the code for now, we will cover it in later tutorials. The code to select the region and blur it is very simple:

```go
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
	// rectangle region
	rec := image.Rectangle{Min: image.Point{214, 383}, Max: image.Point{292, 460}}
	ball := img.Region(rec)
	// apply Gaussian blur
	gocv.GaussianBlur(ball, &ball, image.Pt(35, 35), 0, 0, gocv.BorderDefault)
	// write image to filesystem
	outPath := filepath.Join("blur_messi.jpg")
	if ok := gocv.IMWrite(outPath, img); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
}
```

Note that the `image.Rectangle` values were figured out manually. I used my favorite macOS tool for this: pixelmator. Once you have the pixel coordinates of the football region, you select it in the original image and then apply `gocv.GaussianBlur()` on it. That will give you the following result:

<img src="./blur_messi.jpg" alt="Blurred football" width="200">

Notice that the selected region points to a region of the original image so any modifications you will make on it will also be applied to the original image.

# Image border

Finally, to conclude this chapter you can draw a border around an image using `gocv.CopyMakeBorder()` function. There are several types of borders you can choose. We will demonstrate the functionality by using `gocv.BorderConstant` border type i.e. basic line border:

```go
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
	// rectangle region
	rec := image.Rectangle{Min: image.Point{214, 383}, Max: image.Point{292, 460}}
	ball := img.Region(rec)
	// apply Gaussian blur
	gocv.GaussianBlur(ball, &ball, image.Pt(35, 35), 0, 0, gocv.BorderDefault)
	// select BLUE color
	blue := color.RGBA{B: 255}
	gocv.CopyMakeBorder(img, &img, 10, 10, 10, 10, gocv.BorderConstant, blue)
	// write image to filesystem
	outPath := filepath.Join("border_blur_messi.jpg")
	if ok := gocv.IMWrite(outPath, img); !ok {
		fmt.Printf("Failed to write image: %s\n")
		os.Exit(1)
	}
}
```

The resulting image looks like this:

<img src="./border_blur_messi.jpg" alt="Blurred football borders" width="200">

[[1]](https://commons.wikimedia.org/wiki/File:Lionel_Messi_Player_of_the_Year_2011.jpg) Lionel Messi Player of the Year 2011
