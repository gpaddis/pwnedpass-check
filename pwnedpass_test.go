package main

import "testing"

func TestGetMD5Hash(t *testing.T) {
	MD5Hash := getMD5Hash("password")
	if MD5Hash != "5f4dcc3b5aa765d61d8327deb882cf99" {
		t.Errorf("Wrong hash, got \"%s\".", MD5Hash)
	}
}
