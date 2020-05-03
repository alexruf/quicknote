package cmd

import (
	"github.com/alexruf/quicknote/common"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"runtime"
	"strings"
)

type versionCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newVersionCmd() *versionCmd {
	cc := &versionCmd{}
	cc.baseBuilderCmd = b.newBuilderCmd(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of " + common.ApplicationName + ".",
		Long:  `All software has versions. This is ` + common.ApplicationName + `'s.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cc.printCurrentVersion()
			return nil
		},
	})
	return cc
}

func (v *versionCmd) printCurrentVersion() {
	version := strings.TrimSpace(common.CurrentVersion)
	if len(version) == 0 {
		version = "latest"
	} else {
		if version != "latest" && !strings.HasPrefix(version, "v") {
			version = "v" + version
		}
	}
	jww.FEEDBACK.Printf("%s %s %s/%s\n", common.ApplicationName, version, runtime.GOOS, runtime.GOARCH)
}
