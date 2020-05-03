package cmd

import (
	"fmt"
	"github.com/alexruf/quicknote/common"
	"github.com/spf13/cobra"
	"runtime"
	"strings"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of " + common.ApplicationName + ".",
	Long:  `All software has versions. This is ` + common.ApplicationName + `'s.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s %s/%s\n", common.ApplicationName, printCurrentVersion(), runtime.GOOS, runtime.GOARCH)
	},
}

func printCurrentVersion() string {
	v := strings.TrimSpace(common.CurrentVersion)
	if len(v) == 0 || v == "latest" || v == "HEAD" {
		return v
	}
	if !strings.HasPrefix(v, "v") {
		v = "v" + v
	}
	return v
}
