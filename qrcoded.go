package main

import "fmt"
import "io/ioutil"

func main() {
	fmt.Println("Hello QR Code")

	qrcode := GenerateQRCode("555-5555")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

func GenerateQRCode(code string) []byte {
	return nil
}