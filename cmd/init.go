package cmd

import (
	"fmt"
	"github.com/alexruf/quicknote/common"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize " + common.ApplicationName + ".",
	Long:  `Initialize ` + common.ApplicationName + ` with it's default configuration'.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(viper.ConfigFileUsed())
	},
}
