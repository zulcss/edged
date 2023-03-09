package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/zulcss/edged/apiserver/pkg/config"
	"github.com/zulcss/edged/apiserver/pkg/controller"
        "github.com/zulcss/edged/shared/db"
)

var (
	ConfigFile string
)

var rootCmd = &cobra.Command{
	Use:	"edged",
	Short:  "Edged rest-api server",
	Run: func(cmd *cobra.Command, args []string) {
		config.ReadConfig(ConfigFile)
		log.Printf("Initializing server for %s", viper.GetString("api.host"))
		db.InitDatabase()

		r := controller.Setup()
        	r.Run(fmt.Sprintf("%s", viper.GetString("api.host")))
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
