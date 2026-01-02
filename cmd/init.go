package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tomyukn/rocker-project-template/internal/generator"
)

func initCmd() *cobra.Command {
	var rVersion string
	var serviceName string

	cmd := &cobra.Command{
		Use:   "init <project-name>",
		Short: "Initialize a rocker/rstudio based R project",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			projectName := args[0]

			if rVersion == "" {
				rVersion = "latest"
			}

			if serviceName == "" {
				serviceName = "rstudio"
			}

			cfg := generator.ProjectConfig{
				ProjectName: projectName,
				RVersion:    rVersion,
				ServiceName: serviceName,
			}

			if err := generator.Generate(cfg); err != nil {
				return err
			}

			fmt.Printf("Project %q has been created successfully.\n", projectName)
			return nil
		},
	}

	cmd.Flags().StringVar(&rVersion, "r-version", "", "R version (default: latest)")
	cmd.Flags().StringVar(&serviceName, "name", "rstudio", "Docker Compose service name")

	return cmd
}
