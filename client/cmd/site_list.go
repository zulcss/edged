package cmd

import (
	"github.com/spf13/cobra"

	"github.com/zulcss/edged/client/pkg/client"
)

var siteListCmd = &cobra.Command{
	Use:	"list",
	Short:	"List sites available",
	Run:  func(cmd *cobra.Command, args []string) {
		c := client.NewClient()
		c.ListSites()
	},
}

func init() {
	siteCmd.AddCommand(siteListCmd)
}
