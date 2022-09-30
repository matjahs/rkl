package cmd

import (
  "github.com/spf13/cobra"
)

func initialize() *cobra.Command {
  init := &cobra.Command{
    Use:     "initialize",
    Short:   "init the rkl config.",
    Long:    "init provisions the rkl configuration file",
    Example: "rkl init",
    Aliases: []string{"i", "init"},
    RunE: func(cmd *cobra.Command, args []string) error {
      return nil
    },
  }

  return init
}
