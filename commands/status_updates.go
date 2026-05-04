package commands

import "github.com/spf13/cobra"

var statusUpdatesCmd = &cobra.Command{
	Use:   "status-updates",
	Short: "",
}

func init() {
	rootCmd.AddCommand(statusUpdatesCmd)
}
