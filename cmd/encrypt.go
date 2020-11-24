package cmd

import (
	"github.com/gusandrioli/small-aes/service"
	"github.com/spf13/cobra"
)

// aesEncryptCmd represents the aesEncrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts your message in 16, 24, or 32 bytes.",
	Long: `To effectively use this command, pass first your text that will be encrypted
		then your key (which should be 16, 24, or 32 characters long). For example:

		small-aes encrypt your_text your_key`,
	Run: func(cmd *cobra.Command, args []string) {
		service.AESEncrypt(args)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("24-key-size", "k-24", false, "Size of the key message is 24 bytes.")
}
