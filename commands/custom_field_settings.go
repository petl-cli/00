package commands

import "github.com/spf13/cobra"

var customFieldSettingsCmd = &cobra.Command{
	Use:   "custom-field-settings",
	Short: "",
}

func init() {
	rootCmd.AddCommand(customFieldSettingsCmd)
}
