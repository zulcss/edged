package cmd

import (
        "github.com/spf13/cobra"
        "github.com/zulcss/edged/client/pkg/client"
)

var (
	ServerName string
	IPAddress string
)

var serverCreateCmd = &cobra.Command{
	Use:	"create",
	Short:	"Create a server",
	Run:  func(cmd *cobra.Command, args []string) {
		c := client.NewClient(Endpoint)
		c.CreateServer(ServerName, IPAddress)
	},
}

func init() {
	serverCmd.AddCommand(serverCreateCmd)
}
