package main

import "testing"

func TestGetSHA1Hash(t *testing.T) {
	SHA1Hash := getSHA1Hash("password")
	if SHA1Hash != "5BAA61E4C9B93F3F0682250B6CF8331B7EE68FD8" {
		t.Errorf("Wrong hash, got \"%s\".", SHA1Hash)
	}
}
