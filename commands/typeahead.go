package commands

import "github.com/spf13/cobra"

var typeaheadCmd = &cobra.Command{
	Use:   "typeahead",
	Short: "",
}

func init() {
	rootCmd.AddCommand(typeaheadCmd)
}
