package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zulcss/edged/client/pkg/client"
)

var siteCreateCmd = &cobra.Command{
	Use:	"create",
	Short:	"create a new site",
	Run:  func(cmd *cobra.Command, args []string) {
		c := client.NewClient(Endpoint)
		c.CreateSite(Name)
	},
}

func init() {
	siteCmd.AddCommand(siteCreateCmd)
}
