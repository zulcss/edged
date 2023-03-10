package cmd

import (
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:	"server",
	Short:	"Server commands",
}

func init() {
	rootCmd.AddCommand(serverCmd)

	rootCmd.PersistentFlags().StringVarP(&ServerName, "server", "s", "", "Server name to register")
	rootCmd.PersistentFlags().StringVarP(&IPAddress, "ip-address", "i", "", "IP Address to register")
}
