package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://haveibeenpwned.com/API/v2#PwnedPasswords
// GET https://api.pwnedpasswords.com/range/{first 5 hash chars}

func getMD5Hash(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	res, err := http.Get("https://api.pwnedpasswords.com/range/28675")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
