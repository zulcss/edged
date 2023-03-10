package cmd

import "github.com/spf13/cobra"

var (
	Name string
)

var siteCmd = &cobra.Command{
	Use:	"site",
	Short:	"site commands",
}

func init() {
	rootCmd.AddCommand(siteCmd)
	rootCmd.PersistentFlags().StringVarP(&Name, "site", "", "", "Site to add")
}
