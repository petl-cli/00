package commands

import "github.com/spf13/cobra"

var projectMembershipsCmd = &cobra.Command{
	Use:   "project-memberships",
	Short: "",
}

func init() {
	rootCmd.AddCommand(projectMembershipsCmd)
}
