package cmd

import (
	"fmt"
	"log"
	"os"
	"github.com/spf13/cobra"

	"github.com/zulcss/edged/apiserver/pkg/config"
	"github.com/zulcss/edged/agent/pkg/service"
)

var (
	ConfigFile string
)


var rootCmd = &cobra.Command{
	Use:	"agent",
	Short: 	"edged agent",
	Run: func(cmd *cobra.Command, args []string) {
		log.Printf("Starting agent...")
		config.ReadConfig(ConfigFile)

		service.Run()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	 rootCmd.PersistentFlags().StringVar(&ConfigFile, "config", "", "config file to load")
}
