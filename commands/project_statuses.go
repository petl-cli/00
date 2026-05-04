package commands

import "github.com/spf13/cobra"

var projectStatusesCmd = &cobra.Command{
	Use:   "project-statuses",
	Short: "",
}

func init() {
	rootCmd.AddCommand(projectStatusesCmd)
}
