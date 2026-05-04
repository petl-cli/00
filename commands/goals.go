package commands

import "github.com/spf13/cobra"

var goalsCmd = &cobra.Command{
	Use:   "goals",
	Short: "",
}

func init() {
	rootCmd.AddCommand(goalsCmd)
}
