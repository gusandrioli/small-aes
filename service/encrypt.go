package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func AESEncrypt(args []string) {
	key := []byte(args[1])

	if len(key) != 32 && len(key) != 24 && len(key) != 16 {
		log.Fatal("Wrong size for key")
	}

	text, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

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
