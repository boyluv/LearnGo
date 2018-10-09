package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
	"image/jpeg"
	"log"
	"os"
	)
type Image struct{
	Height,Width int
}

func (m Image) ColorModel() color.Model  {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle  {
	return image.Rect(0,0,m.Height,m.Width)
}

func (m Image) At(x,y int) color.Color  {
	c := uint8(x^y)
	return color.RGBA{c,c,255,255}
}

func main() {
	m := Image{256,256}
	pic.ShowImage(m)

	file, err := os.Create("img/exercise-img.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	jpeg.Encode(file, m, &jpeg.Options{80})
}


