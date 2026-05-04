package commands

import "github.com/spf13/cobra"

var portfoliosCmd = &cobra.Command{
	Use:   "portfolios",
	Short: "",
}

func init() {
	rootCmd.AddCommand(portfoliosCmd)
}
