package cmd

import (
	"github.com/spf13/cobra"
)

var (
	version string
	verbose bool

	rootCmd = &cobra.Command{
		Use:   "git-raycast",
		Short: "Automate git using Raycast AI",
		Long:  "Automate git using Raycast AI",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func SetVersion(v string) {
	version = v
	rootCmd.Version = version
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "V", false, "Send raycast url to output")
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
