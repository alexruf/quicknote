package cmd

import (
	"github.com/alexruf/quicknote/common"
	"github.com/alexruf/quicknote/config"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

type configureCmd struct {
	*baseBuilderCmd
}

func (b *commandsBuilder) newConfigureCmd() *configureCmd {
	cc := &configureCmd{}
	cc.baseBuilderCmd = b.newBuilderCmd(&cobra.Command{
		Use:   "configure",
		Short: "Configure " + common.ApplicationName + ".",
		Long:  `Configure ` + common.ApplicationName + ` with your preferred settings.`,
		Run: func(cmd *cobra.Command, args []string) {
			jww.FEEDBACK.Println(cmd.Short)
			if config.HasConfig() {
				jww.FEEDBACK.Println(common.ApplicationName + " already detected a configuration. Do you want to overwrite the existing configuration?")
			} else {
				jww.FEEDBACK.Println(common.ApplicationName + " Needs to know a few things before we can start...")
			}
		},
	})
	return cc
}
