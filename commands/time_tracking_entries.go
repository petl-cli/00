package commands

import "github.com/spf13/cobra"

var timeTrackingEntriesCmd = &cobra.Command{
	Use:   "time-tracking-entries",
	Short: "",
}

func init() {
	rootCmd.AddCommand(timeTrackingEntriesCmd)
}
