package screenshot

import (
	"image"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

// Connection type that wraps the underlying library type
type Connection = xgb.Conn

// Connect to the window manager system
func Connect() *Connection {
	c, err := xgb.NewConn()
	if err != nil {
		panic(err)
	}

	return c
}

// CaptureScreen capture the current connection screen
func CaptureScreen(c *Connection) (*image.RGBA, error) {
	screen := xproto.Setup(c).DefaultScreen(c)
	x := screen.WidthInPixels
	y := screen.HeightInPixels

	rect := image.Rect(0, 0, int(x), int(y))

	xImg, err := xproto.GetImage(c, xproto.ImageFormatZPixmap, xproto.Drawable(screen.Root), int16(rect.Min.X), int16(rect.Min.Y), uint16(rect.Max.X), uint16(rect.Max.Y), 0xffffffff).Reply()
	if err != nil {
		return nil, err
	}

	data := xImg.Data
	for i := 0; i < len(data); i += 4 {
		data[i], data[i+2], data[i+3] = data[i+2], data[i], 255
	}

	img := &image.RGBA{data, 4 * rect.Max.X, rect}
	return img, nil
}
