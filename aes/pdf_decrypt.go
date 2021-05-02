package aes

import (
	"fmt"
	"log"

	"github.com/gusandrioli/small-aes/utils"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func DecryptPDF(args []string) {
	if len(args) == 0 {
		log.Fatal("No arguments provided. Please see small-aes -h")
	} else if len(args) == 1 {
		log.Fatal("You need to provide the password for decryption")
	}

	if err := api.DecryptFile(
		args[0],
		fmt.Sprintf("decrypted_%v", args[0]),
		pdfcpu.NewAESConfiguration(args[1], utils.NewSequence(), 256),
	); err != nil {
		log.Fatal(err)
	}
}
