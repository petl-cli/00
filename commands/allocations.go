package commands

import "github.com/spf13/cobra"

var allocationsCmd = &cobra.Command{
	Use:   "allocations",
	Short: "",
}

func init() {
	rootCmd.AddCommand(allocationsCmd)
}
