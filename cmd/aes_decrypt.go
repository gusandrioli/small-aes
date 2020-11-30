package cmd

import (
	"github.com/gusandrioli/small-aes/aes"
	"github.com/spf13/cobra"
)

// aesDecryptCmd represents the decrypt command
var aesDecryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts your message in 16, 24, or 32 bytes.",
	Long: `To effectively use this command, pass first your file that will be dencrypted
		then your key (which should be 16, 24, or 32 characters long). For example:

		small-aes decrypt your_text your_key`,
	Run: func(cmd *cobra.Command, args []string) {
		aes.Decrypt(args)
	},
}

func init() {
	rootCmd.AddCommand(aesDecryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// aesDecryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// aesDecryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
