package commands

import "github.com/spf13/cobra"

var customFieldsCmd = &cobra.Command{
	Use:   "custom-fields",
	Short: "",
}

func init() {
	rootCmd.AddCommand(customFieldsCmd)
}
