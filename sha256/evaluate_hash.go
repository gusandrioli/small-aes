package sha256

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

func EvaluateHash(args []string) {
	text, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	hash, err := ioutil.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	hasher := sha256.Sum256(text)

	hexHasher := []byte(fmt.Sprintf("%x", hasher[:]))

	if string(hexHasher) == string(hash) {
		fmt.Printf("True")
	} else {
		fmt.Printf("False")
	}

}
