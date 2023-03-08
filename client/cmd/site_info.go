package cmd

import (
	"github.com/spf13/cobra"

	"github.com/zulcss/edged/client/pkg/client"
)

var siteInfoCmd = &cobra.Command{
	Use:	"info",
	Short:	"site info",
	Run:  func(cmd *cobra.Command, args []string) {
		c := client.NewClient()
		c.GetSite(Name)
	},
}

func init() {
	siteCmd.AddCommand(siteInfoCmd)
}
