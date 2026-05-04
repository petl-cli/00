package commands

import "github.com/spf13/cobra"

var workspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "",
}

func init() {
	rootCmd.AddCommand(workspacesCmd)
}
