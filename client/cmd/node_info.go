package cmd

import (
        "fmt"
        "github.com/spf13/cobra"

        "github.com/zulcss/edged/client/pkg/client"
)

var nodeInfoCmd = &cobra.Command{
	Use:	"info",
	Short:	"Show rest-api information",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClient()
		node, err := c.GetInfo()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(node.Hostname)
	},
}

func init() {
	nodeCmd.AddCommand(nodeInfoCmd)
}
