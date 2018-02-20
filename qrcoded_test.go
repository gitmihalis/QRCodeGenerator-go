package main

import "testing"

// call the Generator with specific data and store result for later
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