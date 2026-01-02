package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func Execute(version string) {
	rootCmd = &cobra.Command{
		Use:   "rpt",
		Short: "rpt scaffolds RStudio development environments based on rocker images",
	}

	rootCmd.AddCommand(initCmd())
	rootCmd.AddCommand(versionCmd(version))
	rootCmd.AddCommand(completionCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func completionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate shell completion scripts",
	}
}
