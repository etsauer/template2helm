package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "template2helm",
		Short: "Template2helm converts an OpenShift Template into a Helm Chart.",
		Long: `Template2helm converts an OpenShift Template into a Helm Chart.
      For more info, check out https://github.com/redhat-cop/template2helm`,
	}
)

// Execute - entrypoint for CLI tool
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
