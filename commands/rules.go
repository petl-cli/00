package commands

import "github.com/spf13/cobra"

var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "",
}

func init() {
	rootCmd.AddCommand(rulesCmd)
}
