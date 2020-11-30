package sha256

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

func Hash(args []string) {
	text, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	hasher := sha256.Sum256(text)

	err = ioutil.WriteFile(
		"hashed.txt",
		[]byte(fmt.Sprintf("%x", hasher[:])),
		0777,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x", hasher[:])
}
