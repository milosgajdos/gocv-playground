# Image arithmetics

This chapter deals with digital image airthmetics using `gocv`. We will look at image addition, subtraction, bitwise operations and image masking.

## Image addition and subtraction

It is possible to "add" two images together using `gocv.Add()` function; addition in the context of image processing means that the pixel intensity values of both images are added pixel by pixel in each channel. This requires that following conditions are met:
* both images must have the same size
* both images must have the same number of channels

We will demonstrate the image addition on the familiar picture of Lionel Messi. For the second image in the addition we will use the logo of [Wikimedia Commons](https://commons.wikimedia.org/wiki/Main_Page):
<img src="./commons.png" alt="Wikimedia Commons" width="200">

Now, as you figured, the image of messi and wiki commons logo are not of the same size. What we will do instead is, we will add the logo with the bottom left region (of the same size as the logo) of messi image. We will then display the resulting image.

The code is pretty simple; first we read in both images, then we select the bottom left region and finally we add the images together and write the resulting image to the filesystem (reading and writing image code is left out for brevity):

```go
// get image size
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

As you can see the left bottom part of the original image now contains wikimedia logo alas the resulting colors have pixel values that are a result of addition: note how orange-ish the centre of the logo turned as a result of red addition to green values. Similarly, the blue parts of logo turned green-ish:
<img src="./add_logo_messi.jpeg" alt="Messi add logo" width="200">

## Image blending

Image blending is an addition technique where the images have different weight coefficients in the resulting summation. This functionality is provided by `gocv.AddWeighted()` function. Using this function you are essentially creating a sort of libear combination of images. Let's demonstrate it on the messi image:

```go
gocv.AddWeighted(rio, 0.6, logo, 0.4, 0.0, rio)
```

The resulting image shows that the image of logo is a bit subdued due to our assigning it a lower coefficient (`0.4`). If both coefficients were set to `1.0` we would get the same result as when using `gocv.Add()`:
<img src="./add_weighted_logo_messi.jpeg" alt="Weighted Messi add logo" width="200">

## Image masking

Image masking is an incredibly powerful technique in digital image processing. It allows for fast processing when working with non-rectangular image regions of interest such as image overlaying, replacing objects in the image etc. In order to perform these functions we need to apply `bitwise` pixel operations such as AND, OR, NOT or XOR.

### Image thresholding

Before we proceed with image masking, we need to quickly mention an important topic: image thresholding. The concept is pretty simple: you define a threshold for a pixel intensity and if a pixel in the source image is higher than the chosen threshold it is assigned some value, else it's assigned different value. What those values are depends on the type of thresholding you decide to use. Let's demonstrate this on a simple example.

We will apply a *binary* threshold on the wikimedia commons logo First we need to turn the image into a grayscale image, then we can apply binary thresholding:

```go
// convert to grayscale
gocv.CvtColor(logo, logo, gocv.ColorBGRToGray)
// create binary image
gocv.Threshold(logo, logo, 10.0, 255.0, gocv.ThresholdBinary)
```

What the above code does is, it checks all the pixel values and sets the ones which are higher than 10 to 255 (white) and lower ones to 0 (black). The resulting binary thresholded image of wiki commons logo looks like this:
<img src="./wiki_commons_binary.jpeg" alt="Wiki commons binary" width="200">

#### Image bitwise operations

Imagine you wanted to mask out some non-rectangular parts of an image (certain pixels of certain colors or features), so that you can overlay something on top of the masked out region such as our wiki commons logo. What you need to do is to black-out a region of the image in areas where the logo would fit in.

`gocv` doesn't seem to have any convenient function that would let you apply a mask on colored images in at once, so what you need to do is a bit more elaborate:

```go
// create an inverse mask
maskInv := gocv.NewMat()
gocv.BitwiseNot(mask, maskInv)
// black-out the area of logo in roi i.e. in bottom left region
roiChans := gocv.Split(roi)
for i := 0; i < len(roiChans); i++ {
	gocv.BitwiseAnd(roiChans[i], maskInv, roiChans[i])
}
gocv.Merge(roiChans, roi)
```
What this code effectively does is it splits `roi` (region of interest where we want to overlay our logo) into separate channels and applies the *inverse mask* to each of them. It then merges all of the channels together. The net result is, the masked out pixels will have their intensity set to 0 (black) in the resulting image:
<img src="./wiki_commons_blackout_messi.jpeg" alt="Wiki commons blackout Messi" width="200">

## Image overlay

Now that we know how to apply masks and how to do basica airthmetic operations, we can put it all together and overlay the wiki commons log on top of the image of Messi.

What we need to do is apply the original mask we have obtained earlier to the wiki commons logo and then add together the resulting regions:

```go
// apply the mask on logo image
logoChans := gocv.Split(logo)
for i := 0; i < len(logoChans); i++ {
	gocv.BitwiseAnd(logoChans[i], mask, logoChans[i])
}
gocv.Merge(logoChans, logo)
// add logo to roi
gocv.Add(roi, logo, roi)
```
The resulting image looks like this:
<img src="./wiki_commons_messi.jpeg" alt="Wiki commons Messi" width="200">

You can find the full code and some commented out snippets in the `main.go` in this project directory


