package commands

import "github.com/spf13/cobra"

var batchApiCmd = &cobra.Command{
	Use:   "batch-api",
	Short: "",
}

func init() {
	rootCmd.AddCommand(batchApiCmd)
}
