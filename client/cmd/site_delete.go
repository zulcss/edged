package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zulcss/edged/client/pkg/client"
)

var siteRemoveCmd = &cobra.Command{
	Use:	"remove",
	Short:	"Remove a site",
	Run:  func(cmd *cobra.Command, args []string) {
		c := client.NewClient()
		c.DeleteSite(Name)
	},
}

func init() {
	siteCmd.AddCommand(siteRemoveCmd)
}
