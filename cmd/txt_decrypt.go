package cmd

import (
	"log"

	"github.com/gusandrioli/small-aes/aes"
	"github.com/spf13/cobra"
)

// txtDecryptCmd represents the decrypt command
var txtDecryptCmd = &cobra.Command{
	Use:   "txtDecrypt",
	Short: "Decrypts your message with a key of 16, 24, or 32 bytes.",
	Long: `To effectively use this command, pass first your file that will be dencrypted
		then your key (which should be 16, 24, or 32 characters long). For example:

		small-aes decrypt your_text your_key`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			log.Fatal("Missing a 16, 24, or 32 bytes key")
			return
		}

		aes.DecryptTxt(args)
	},
}

func init() {
	rootCmd.AddCommand(txtDecryptCmd)
}
