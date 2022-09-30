package cmd

import (
  "context"

  "github.com/spf13/cobra"
)

func Execute() error {
  rootCmd := &cobra.Command{
    Version: "0.0.1",
    Use:     "rkl",
    Long:    "Rekall (rkl) is a CLI tool for remembering things.",
    Example: "rkl",
    RunE: func(cmd *cobra.Command, args []string) error {
      return nil
    },
  }

  rootCmd.AddCommand(initialize())

  return rootCmd.ExecuteContext(context.Background())
}
