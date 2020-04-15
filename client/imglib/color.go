package imglib

import (
	"fmt"
	"image"
	"log"

	"github.com/EdlinOrg/prominentcolor"
)

func AvgImgColor(img *image.RGBA) (prominentcolor.ColorRGB, error) {
	colors, err := prominentcolor.KmeansWithAll(3, img, prominentcolor.ArgumentDefault, prominentcolor.DefaultSize, []prominentcolor.ColorBackgroundMask{})
	if err != nil {
		log.Printf("Failed to process imagen\n", err)
		return prominentcolor.ColorRGB{R: 0, G: 0, B: 0}, fmt.Errorf("Error parsing avg color")
	}

	// for i := range colors {
	// 	fmt.Println("#" + colors[i].AsString())
	// }

	return colors[0].Color, nil
}

func GetColorHex(color prominentcolor.ColorRGB) string {
	ci := prominentcolor.ColorItem{Color: color, Cnt: 0}
	return ci.AsString()
}
