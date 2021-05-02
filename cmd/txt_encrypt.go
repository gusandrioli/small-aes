package cmd

import (
	"github.com/gusandrioli/small-aes/aes"
	"github.com/spf13/cobra"
)

// txtEncryptCmd represents the aesEncrypt command
var txtEncryptCmd = &cobra.Command{
	Use:   "txtEncrypt",
	Short: "Encrypts your message with a key 16, 24, or 32 bytes.",
	Long: `To effectively use this command, pass first your text file that will be encrypted,
		then your key (which should be 16, 24, or 32 characters long). For example:

		small-aes encrypt your_text your_key`,
	Run: func(cmd *cobra.Command, args []string) {
		aes.EncryptTxt(args)
	},
}

func init() {
	rootCmd.AddCommand(txtEncryptCmd)
}
