package main

// https://haveibeenpwned.com/API/v2#PwnedPasswords
// GET https://api.pwnedpasswords.com/range/{first 5 SHA1 hash chars}

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func getSHA1Hash(password string) string {
	hasher := sha1.New()
	hasher.Write([]byte(password))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return strings.ToUpper(hash)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	password := os.Args[1]
	hash := getSHA1Hash(password)
	res, err := http.Get("https://api.pwnedpasswords.com/range/" + hash[0:5])
	checkErr(err)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	checkErr(err)

	hashList := strings.Fields(string(body))
	suffix := hash[5:]
	for _, h := range hashList {
		if strings.Contains(h, suffix) {
			count := strings.Split(h, ":")[1]
			fmt.Println("The password", password, "was pwned and appears", count, "times in the database.")
			return
		}
	}

	fmt.Println("The password", password, "was not found in any breaches.")
}
