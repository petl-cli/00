package commands

import "github.com/spf13/cobra"

var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "",
}

func init() {
	rootCmd.AddCommand(teamsCmd)
}
