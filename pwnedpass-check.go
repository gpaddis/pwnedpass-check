package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// https://haveibeenpwned.com/API/v2#PwnedPasswords
// GET https://api.pwnedpasswords.com/range/{first 5 SHA1 hash chars}

func getSHA1Hash(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return strings.ToUpper(hash)
}

func main() {
	pass := os.Args[1]
	hash := getSHA1Hash(pass)
	res, err := http.Get("https://api.pwnedpasswords.com/range/" + hash[0:5])
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	result := strings.Fields(string(body))

	suffix := strings.ToUpper(hash[5:])
	for _, v := range result {
		if strings.Contains(v, suffix) {
			count := strings.Split(v, ":")[1]
			fmt.Println("The password", pass, "was found in", count, "breaches.")
		}
	}
}
