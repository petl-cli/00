package commands

import "github.com/spf13/cobra"

var tagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "",
}

func init() {
	rootCmd.AddCommand(tagsCmd)
}
