# Image arithmetics

This chapter covers the topic of digital image airthmetics. We will look at image addition, subtraction and bitwise operations. We will demonstrate all concepts on practical examples.

## Image addition and subtraction

It is possible to "add" two digital images together. `gocv` provides `Add()` function which calculates the addition; in context of digital image processing, "addition" means that the pixel intensity values of both images are added together pixel by pixel in each channel. This requires that the following conditions are met:
* both images must have the same size (rows x columns)
* both images must have the same number of channels

We will demonstrate image addition on the familiar picture of Lionel Messi. We will add the image of him to the logo of [Wikimedia Commons](https://commons.wikimedia.org/wiki/Main_Page):

<img src="./commons.png" alt="Wikimedia Commons" width="200">

As you may have figured, the image of Messi and wiki commons logo are not of the same size and therefore the first of the conditions we listed earlier can't be fulfilled. What we will do instead is, we will add the logo to the bottom left region (of the same size as the logo) of Messi image.

The code to do this is pretty simple; first we read in both images, then we select the bottom left region of interest (`roi`) and finally we add the two images together (reading and writing image code is left out for brevity):

```go
// get image sizes
messiRows, messiCols := messi.Rows(), messi.Cols()
logoRows, logoCols := logo.Rows(), logo.Cols()
// select bottom left region in messi picture
minX, minY := 0, messiRows-logoRows
maxX, maxY := logoCols, messiRows
rec := image.Rectangle{Min: image.Point{minX, minY}, Max: image.Point{maxX, maxY}}
rio := messi.Region(rec)
// add two images
gocv.Add(rio, logo, rio)
```

As you can see in the resulting image below, the left bottom part of the original image now contains the logo, alas the resulting colors have pixel values that are a result of addition to the messi image colors: note how orange-ish the centre of the logo is. Similarly, the blue parts of logo turned green-ish:

<img src="./add_logo_messi.jpeg" alt="Messi add logo" width="200">

## Image blending

Image blending is an addition technique where the two images have different weight coefficients in the resulting summation. This functionality is provided by `AddWeighted()` function. By calling this function you are essentially creating a linear combination of images. Let's demonstrate it on some example. We will add the messi and logo images together, but add more weight to the messi image and less weight to the image of logo:

```go
gocv.AddWeighted(rio, 0.6, logo, 0.4, 0.0, rio)
```

The resulting image shows that the image of logo is a bit "subdued" due to our assigning it a lower coefficient (`0.4`). If both coefficients were set to `1.0` we would get the same result as when using `gocv.Add()`:

<img src="./add_weighted_logo_messi.jpeg" alt="Weighted Messi add logo" width="200">

## Image masking

Image masking is an incredibly powerful technique in digital image processing. It allows for fast processing when working with non-rectangular image regions of interest such as image overlaying, replacing different kinds of objects in the image etc. In order to perform these functions we need to apply `bitwise` pixel operations such as AND, OR, NOT or XOR.

### Image thresholding

Before we proceed with image masking, we need to quickly discuss an important topic: image thresholding. The concept is pretty simple; you define a threshold for a pixel intensity and if a pixel in the source image is higher than the threshold it is assigned some value, else it's assigned different value. What those values are depends on the type of thresholding you decide to use.

Let's demonstrate this on a simple example. We will apply a *binary* threshold on the logo. First we need to turn the image into a grayscale image, then we can perform binary thresholding:

```go
// convert to grayscale
gocv.CvtColor(logo, logo, gocv.ColorBGRToGray)
// create binary image
gocv.Threshold(logo, &logo, 10.0, 255.0, gocv.ThresholdBinary)
```

What the above code does is, it checks all the pixel values and sets the ones which are higher than 10 to 255 (white) and the lower ones to 0 (black). The resulting image of logo looks like this:

<img src="./wiki_commons_binary.jpeg" alt="Wiki commons binary" width="200">

#### Image bitwise operations

Imagine you wanted to mask out some non-rectangular parts of an image (certain pixels of certain colors or features), so that you can overlay something on top of the masked out region such as the wiki commons logo. What you need to do is to black-out a region of interest in the image in areas where the logo would fit in.

`gocv` doesn't seem to have any convenient function that would let you apply a mask on colored images in one function call, so what you need to do is a bit more elaborate:

```go
// create an inverse mask
maskInv := gocv.NewMat()
gocv.BitwiseNot(mask, &maskInv)
// black-out the area of logo in roi i.e. in bottom left region
roiMask := gocv.NewMat()
gocv.Merge([]gocv.Mat{maskInv, maskInv, maskInv}, &roiMask)
gocv.BitwiseAnd(roi, roiMask, &roi)
```
What this code effectively does is it creates an image mask by assembling the binary thresholded image into 3 channels image mask - if were masking grayscale image we wouldn't have to do this, however we must do this now because `BitwiseAnd()` expect the mask and the source image to have both the exact size and the same number of channels. The net result of this operation is the masked out pixels will have their intensity set to 0 (black) in the resulting image:

<img src="./wiki_commons_blackout_messi.jpeg" alt="Wiki commons blackout Messi" width="200">

## Image overlay

Now that we know how to apply masks and how to do basica airthmetic operations, we can put it all together and overlay the wiki commons log on top of the image of Messi.

What we need to do is to apply the original mask we have obtained earlier to the wiki commons logo and then `Add()` the resulting region to the masked logo image:

```go
// apply the mask on logo image
logoMask := gocv.NewMat()
gocv.Merge([]gocv.Mat{mask, mask, mask}, &logoMask)
gocv.BitwiseAnd(logo, logoMask, &logo)
// add logo to roi
gocv.Add(roi, logo, &roi)
```
The resulting image looks like this:

<img src="./wiki_commons_messi.jpeg" alt="Wiki commons Messi" width="200">

You can find the full code and some commented out snippets in the `main.go` in this project directory
