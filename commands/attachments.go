package commands

import "github.com/spf13/cobra"

var attachmentsCmd = &cobra.Command{
	Use:   "attachments",
	Short: "",
}

func init() {
	rootCmd.AddCommand(attachmentsCmd)
}
