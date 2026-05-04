package commands

import "github.com/spf13/cobra"

var organizationExportsCmd = &cobra.Command{
	Use:   "organization-exports",
	Short: "",
}

func init() {
	rootCmd.AddCommand(organizationExportsCmd)
}
