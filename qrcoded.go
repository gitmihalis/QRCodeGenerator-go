package main

import (
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	log.Println("Hello QR Code")
	// Create png file
	file, err := os.Create("qrcode.png")
	// Error handling
	if err != nil {
		log.Fatal(err) // automatically calls os.Exit(1)
	}
	// Close file when func is finished running
	defer file.Close()
	// Generate the QR code and catch any error
	err = GenerateQRCode(file, "555-3556", Version(1))
	if err != nil {
		log.Fatal(err)
	}
}

// GenerateQRCode Generates a QR Code from a string
func GenerateQRCode(w io.Writer, code string, version Version) error {
	size := 4*int(version) + 17
	img := image.NewNRGBA(image.Rect(0, 0, size, size)) // NRGBA is an in-memory image whose At method returns color.NRGBA values.


	return png.Encode(w, img) 	// Save the png to file and return any error
}

type Version int8
