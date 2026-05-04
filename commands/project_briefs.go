package commands

import "github.com/spf13/cobra"

var projectBriefsCmd = &cobra.Command{
	Use:   "project-briefs",
	Short: "",
}

func init() {
	rootCmd.AddCommand(projectBriefsCmd)
}
