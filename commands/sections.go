package commands

import "github.com/spf13/cobra"

var sectionsCmd = &cobra.Command{
	Use:   "sections",
	Short: "",
}

func init() {
	rootCmd.AddCommand(sectionsCmd)
}
