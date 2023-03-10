package cmd

import (
        "github.com/spf13/cobra"
        "github.com/zulcss/edged/client/pkg/client"
)

var (
	ServerName string
	IPAddress string
)

var serverRegisterCmd = &cobra.Command{
	Use:	"create",
	Short:	"Create a server",
	Run:  func(cmd *cobra.Command, args []string) {
		c := client.NewClient(Endpoint)
		c.CreateServer(ServerName, IPAddress)
	},
}

func init() {
	serverCmd.AddCommand(serverRegisterCmd)

	rootCmd.PersistentFlags().StringVarP(&ServerName, "server", "s", "", "Server name to register")
	rootCmd.PersistentFlags().StringVarP(&IPAddress, "ip-address", "i", "", "IP Address to register")
}
