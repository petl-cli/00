package commands

import "github.com/spf13/cobra"

var timePeriodsCmd = &cobra.Command{
	Use:   "time-periods",
	Short: "",
}

func init() {
	rootCmd.AddCommand(timePeriodsCmd)
}
