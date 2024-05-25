package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	file, err := os.Open("Torenia.jpg")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	colorModel := img.ColorModel()
	bounds := img.Bounds()
	fmt.Println("Image is a", colorModel)
	fmt.Println("Bounds:", bounds)
	fmt.Println("--------------------")

	resizedImg := resizeImage(img, 100, 100)

	asciiArt := convertToASCII(resizedImg)
	fmt.Print(asciiArt)
}

func resizeImage(img image.Image, newWidth, newHeight int) image.Image {
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	newImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	for y := 0; y < newHeight; y++ {
		for x := 0; x < newWidth; x++ {
			srcX := x * width / newWidth
			srcY := y * height / newHeight
			newImg.Set(x, y, img.At(srcX, srcY))
		}
	}

	return newImg
}

func RGBToANSI(r, g, b uint32) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r>>8, g>>8, b>>8)
}

func convertToASCII(img image.Image) string {
	asciiChars := "MND8OZ$7I?+=~:,.."
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var asciiString string

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			col := img.At(x, y)
			r, g, b, _ := col.RGBA()

			grayColor := color.GrayModel.Convert(col).(color.Gray)
			grayValue := grayColor.Y

			charIndex := int(grayValue) * (len(asciiChars) - 1) / 255
			asciiChar := string(asciiChars[charIndex])

			ansiColor := RGBToANSI(r, g, b)

			asciiString += ansiColor + asciiChar
		}
		asciiString += "\033[0m\n"
	}
	return asciiString
}
