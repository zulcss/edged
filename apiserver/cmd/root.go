package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

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
		db.InitDatabase()

		r := controller.Setup()
        	r.Run() // Listen and server on 0.0.0.0:8080
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {}
