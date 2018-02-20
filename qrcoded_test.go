package main

import (
	"image/png"
	"bytes"
	"testing"
)

// Call the Generator with specific data and store result for later
// inspection.




func TestGenerateQRCodeGeneratesPNG(t *testing.T)  {
	buffer := new(bytes.Buffer)
	GenerateQRCode(buffer, "555-2345")
	
	if buffer.Len() == 0 {
		t.Errorf("No QRCode generated")
	}
	// Decode the byte array, discarding any positive results and focus
	// on the error.
	_, err := png.Decode(buffer) // `Decode` does not work on byte slices. 
	// It must satisfy the io.Reader interface  by wrapping in a buffer.

	if err != nil {
		t.Errorf("Generated QRCode is not a PNG %s", err)
	}
}