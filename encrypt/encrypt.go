package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func aesEncrypt() {
	text := []byte("This is my secret message")

	key := []byte("passphrasewhichneedstobe32bytes!")

	//generate aes cipher
	c, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	// implement gcm or Galois/Galois Mode
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Fatal(err)
	}

	//new byte array the size of nonce
	nonce := make([]byte, gcm.NonceSize())

	//populate nonce with random sequence
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("myfile.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted into myfile.data")
}
