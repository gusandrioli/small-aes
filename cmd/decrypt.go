package cmd

import (
	"github.com/gusandrioli/small-aes/service"
	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts your message in 16, 24, or 32 bytes.",
	Long: `To effectively use this command, pass first your file that will be dencrypted
	then your key (which should be 16, 24, or 32 characters long). For example:

	small-aes decrypt your_text your_key`,
	Run: func(cmd *cobra.Command, args []string) {
		service.AESDecrypt(args)
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
