package cmd

import "github.com/spf13/cobra"

var nodeCmd = &cobra.Command{
        Use:    "node",
        Short:  "Node commands",
}

func init() {
	rootCmd.AddCommand(nodeCmd)
}
