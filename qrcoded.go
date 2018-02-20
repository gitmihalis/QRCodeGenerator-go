package main

import (
	"image/png"
	"bytes"
	"image"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello QR Code")

	qrcode := GenerateQRCode("555-5555")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

func GenerateQRCode(code string) []byte {
	// Create an image and encode it using `png.Encode`.
	img := image.NewNRGBA(image.Rect(0, 0, 21, 21))
	buf := new(bytes.Buffer)
	_ = png.Encode(buf, img)

	return buf.Bytes()
}
