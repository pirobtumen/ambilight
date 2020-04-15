package main

import (
	"fmt"
	"image"
	"io"
	"log"
	"time"

	"./imglib"
	"./screenshot"
	"./serial"
)

func capture(c *screenshot.Connection, s io.ReadWriteCloser) *image.RGBA {
	start := time.Now()
	img, err := screenshot.CaptureScreen(c)
	if err != nil {
		panic(err)
	}

	color, err := imglib.AvgImgColor(img)
	if err == nil {
		fmt.Println("#" + imglib.GetColorHex(color))
		serial.SendColor(s, color)
		elapsed := time.Since(start)
		log.Printf("Elapsed %g seconds", elapsed.Seconds())
		time.Sleep(time.Duration(250 * time.Millisecond))
	}

	return img
}

func main() {
	options := serial.OpenOptions{
		PortName:        "/dev/ttyACM0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	c := screenshot.Connect()
	s := serial.Connect(options)
	defer c.Close()
	defer s.Close()

	go func() {
		for {
			capture(c, s)
		}
	}()

	fmt.Println("Press enter to exit")
	var input string
	fmt.Scanf("%s", &input)
}
