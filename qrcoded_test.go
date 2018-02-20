package main

import (
	"image/png"
	"bytes"
	"testing"
)

// Call the Generator with specific data and store result for later
// inspection.
func TestGenerateQRCodeReturnValue(t *testing.T) {
	result := GenerateQRCode("555-5555")

	if result == nil {
		t.Errorf("Generated QRCode is nil")
	}
	if len(result) == 0 {
		t.Errorf("Generated QRCode has no data")
	}
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T)  {
	result := GenerateQRCode("555-2345")
	buffer := bytes.NewBuffer(result) 
	// Decode the byte array, discarding any positive results and focus
	// on the error.
	_, err := png.Decode(buffer) // `Decode` does not work on byte slices. 
	// It must satisfy the io.Reader interface  by wrapping in a buffer.

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG %s", err)
	}
}