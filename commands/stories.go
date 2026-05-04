package commands

import "github.com/spf13/cobra"

var storiesCmd = &cobra.Command{
	Use:   "stories",
	Short: "",
}

func init() {
	rootCmd.AddCommand(storiesCmd)
}
