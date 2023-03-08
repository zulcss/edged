package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/zulcss/edged/client/pkg/client"
)

var nodeHealthCmd = &cobra.Command{
	Use:	"health",
	Short:	"Query node health",
	Run: func(cmd *cobra.Command, args []string) {
		c := client.NewClient()
		node, err := c.GetStatus()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Server status is", node.Status)
	},
}

func init() {
	nodeCmd.AddCommand(nodeHealthCmd)
}
