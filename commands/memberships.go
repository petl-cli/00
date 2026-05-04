package commands

import "github.com/spf13/cobra"

var membershipsCmd = &cobra.Command{
	Use:   "memberships",
	Short: "",
}

func init() {
	rootCmd.AddCommand(membershipsCmd)
}
