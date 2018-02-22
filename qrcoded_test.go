package main

import (
	"bytes"
	"errors"
	"image/png"
	"testing"
)

// Call the Generator with specific data and store result for later
// inspection.

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
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

// Test  an implementation of io.Writer that generates an error.
type ErrorWriter struct{}

// This function turns the ErrorWriter into a stub satisfying io.Writer.
func (e *ErrorWriter) Write(b []byte) (int, error) {
	return 0, errors.New("Expected error") // Its implementation is straightforward: every call to Write will return an error.
}

func TestGenerateQRCodePropogatesErrors(t *testing.T) {
	w := new(ErrorWriter)
	err := GenerateQRCode(w, "555-2334")

	if err == nil || err.Error() != "Expected error" {
		t.Errorf("Error not propogated correctly, got %v", err)
	}
}
