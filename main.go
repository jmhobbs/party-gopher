package main

import (
	"image/color"
	"image/gif"
	"log"
	"os"
	"strconv"
)

var (
	ParrotColors     []color.Color
	DarkParrotColors []color.Color
	LightGopherBlue  color.Color
	DarkGopherBlue   color.Color
)

func init() {
	var err error

	for _, s := range []string{
		"FF6B6B",
		"FF6BB5",
		"FF81FF",
		"FF81FF",
		"D081FF",
		"81ACFF",
		"81FFFF",
		"81FF81",
		"FFD081",
		"FF8181",
	} {
		c, err := hexToColor(s)
		if err != nil {
			log.Fatal(err)
		}
		ParrotColors = append(ParrotColors, c)
		DarkParrotColors = append(DarkParrotColors, darken(c))
	}

	LightGopherBlue, err = hexToColor("8BD0FF")
	if err != nil {
		log.Fatal(err)
	}
	DarkGopherBlue, err = hexToColor("82C2EE")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Indexes for light and dark blue in palettes.
	var (
		lbi int
		dbi int
	)

	// Open the dancing gopher gif
	f, err := os.Open("dancing-gopher.gif")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Decode the gif so we can edit it
	gopher, err := gif.DecodeAll(f)
	if err != nil {
		log.Fatal(err)
	}

	// Instead of changing pixels with Set(x, y int, c color.Color)
	// let's just tweak the palette to replace those colors.
	for i, frame := range gopher.Image {
		lbi = frame.Palette.Index(LightGopherBlue)
		dbi = frame.Palette.Index(DarkGopherBlue)

		frame.Palette[lbi] = ParrotColors[i%len(ParrotColors)]
		frame.Palette[dbi] = DarkParrotColors[i%len(DarkParrotColors)]
	}

	// Save it out!
	o, _ := os.OpenFile("party-gopher.gif", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	defer o.Close()
	gif.EncodeAll(o, gopher)
}

// The colors I extracted were in HTML style hex format.
// Instead of manually converting them, I just wrote this.
func hexToColor(hex string) (color.Color, error) {
	c := color.RGBA{0, 0, 0, 255}

	r, err := strconv.ParseInt(hex[0:2], 16, 16)
	if err != nil {
		return c, err
	}

	g, err := strconv.ParseInt(hex[2:4], 16, 16)
	if err != nil {
		return c, err
	}

	b, err := strconv.ParseInt(hex[4:6], 16, 16)
	if err != nil {
		return c, err
	}

	c.R = uint8(r)
	c.G = uint8(g)
	c.B = uint8(b)

	return c, nil
}

// To make the shadow shades I darken all the channels by 15/255
func darken(c color.Color) color.Color {
	r, g, b, a := c.RGBA()
	r = r - 15
	g = g - 15
	b = b - 15
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
