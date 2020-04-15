package imglib

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"time"

	"github.com/EdlinOrg/prominentcolor"
)

func SaveImg(img *image.RGBA) {
	f, err := os.Create(fmt.Sprintf("./screenshot-%d.png", time.Now().Unix()))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func TestGreenImageErr() {
	f, err := os.Open("./colors.jpg")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	imgRaw, _, err := image.Decode(f)
	b := imgRaw.Bounds()
	m := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(m, m.Bounds(), imgRaw, b.Min, draw.Src)
	color, _ := AvgImgColor(m)
	ci := prominentcolor.ColorItem{Color: color, Cnt: 0}
	fmt.Println("#" + ci.AsString())
}
