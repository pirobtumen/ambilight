package serial

import (
	"fmt"
	"io"
	"log"

	"github.com/EdlinOrg/prominentcolor"
	"github.com/jacobsa/go-serial/serial"
)

type OpenOptions = serial.OpenOptions

func Connect(options OpenOptions) io.ReadWriteCloser {
	connection, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	return connection
}

func SendColor(port io.ReadWriteCloser, color prominentcolor.ColorRGB) {
	b := []byte(fmt.Sprintf("%d,%d,%d;", color.R, color.G, color.B))
	n, err := port.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}

	fmt.Println("Wrote", n, "bytes.")
}
