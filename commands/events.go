package commands

import "github.com/spf13/cobra"

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "",
}

func init() {
	rootCmd.AddCommand(eventsCmd)
}
