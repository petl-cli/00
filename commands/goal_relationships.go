package commands

import "github.com/spf13/cobra"

var goalRelationshipsCmd = &cobra.Command{
	Use:   "goal-relationships",
	Short: "",
}

func init() {
	rootCmd.AddCommand(goalRelationshipsCmd)
}
