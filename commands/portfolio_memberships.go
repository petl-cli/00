package commands

import "github.com/spf13/cobra"

var portfolioMembershipsCmd = &cobra.Command{
	Use:   "portfolio-memberships",
	Short: "",
}

func init() {
	rootCmd.AddCommand(portfolioMembershipsCmd)
}
