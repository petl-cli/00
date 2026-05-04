package commands

import "github.com/spf13/cobra"

var projectTemplatesCmd = &cobra.Command{
	Use:   "project-templates",
	Short: "",
}

func init() {
	rootCmd.AddCommand(projectTemplatesCmd)
}
