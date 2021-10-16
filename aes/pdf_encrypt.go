package aes

import (
	"fmt"
	"log"

	"github.com/gusandrioli/small-aes/utils"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func EncryptPDF(args []string) {
	var key string

	if len(args) == 0 {
		log.Fatal("No arguments provided. Please see small-aes -h")
	} else if len(args) == 2 {
		key = args[1]
	} else {
		key = utils.NewSequenceWithLength(127)
	}

	if err := api.EncryptFile(
		args[0],
		fmt.Sprintf("encrypted_%v", args[0]),
		pdfcpu.NewAESConfiguration(key, key, 256), // TODO differentiate between passwords
	); err != nil {
		log.Fatal(err)
	}

	fmt.Println("This is your password, save it!", key)
}
