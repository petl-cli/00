package commands

import "github.com/spf13/cobra"

var teamMembershipsCmd = &cobra.Command{
	Use:   "team-memberships",
	Short: "",
}

func init() {
	rootCmd.AddCommand(teamMembershipsCmd)
}
