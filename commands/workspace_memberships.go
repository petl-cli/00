package commands

import "github.com/spf13/cobra"

var workspaceMembershipsCmd = &cobra.Command{
	Use:   "workspace-memberships",
	Short: "",
}

func init() {
	rootCmd.AddCommand(workspaceMembershipsCmd)
}
