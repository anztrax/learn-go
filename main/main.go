package main

import (
	"github.com/ungerik/go-cairo"
	"image"
	_ "image/jpeg"
	"net/http"
)

func decodeBase64() image.Image{
	res, err := http.Get("https://tvlk.imgix.net/imageResource/2018/10/15/1539606255527-bc6318215478e9e8a66ba763a7605bea.png?auto=compress%2Cformat&cs=srgb&fm=png&ixlib=java-1.1.12&q=75");
	if err != nil || res.StatusCode != 200 {
		// handle errors
		panic("can't get image")
	}
	defer res.Body.Close()
	m, _, err := image.Decode(res.Body)
	if err != nil {
		// handle error

		panic("can't get image")
	}
	return m
}

func createImage(){
	surface := cairo.NewSurface(cairo.FORMAT_ARGB32, 6000, 6000)
	surface.SelectFontFace("serif", cairo.FONT_SLANT_NORMAL, cairo.FONT_WEIGHT_BOLD)
	surface.SetFontSize(32.0)
	surface.SetSourceRGB(0.0, 0.0, 1.0)
	surface.MoveTo(10.0, 50.0)
	surface.ShowText("Hello World")
	surface.SetImage(decodeBase64());
	surface.WriteToPNG("hello.png")
	surface.Finish()
}

func main() {
	createImage()
}