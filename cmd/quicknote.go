package cmd

import (
	"fmt"
	"github.com/alexruf/quicknote/common"
	"github.com/alexruf/quicknote/config"
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
)

func init() {
	cobra.OnInitialize(config.InitConfig)
}

// The Response value from Execute.
type Response struct {
	// Err is set when the command failed to execute.
	Err error
	// The command that was executed.
	Cmd *cobra.Command
}

// IsUserError returns true is the Response error is a user error rather than a
// system error.
func (r Response) IsUserError() bool {
	return r.Err != nil && isUserError(r.Err)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// The args are usually filled with os.Args[1:].
func Execute(args []string) Response {
	quicknoteCmd := newCommandsBuilder().addAll().build()
	cmd := quicknoteCmd.getCommand()
	cmd.SetArgs(args)

	c, err := cmd.ExecuteC()

	var resp Response
	resp.Err = err
	resp.Cmd = c

	return resp
}

type quicknoteCmd struct {
	*baseBuilderCmd
	Verbose bool
	Debug   bool
	Trace   bool
}

func (b *commandsBuilder) newQuicknoteCmd() *quicknoteCmd {
	cc := &quicknoteCmd{}
	cc.baseBuilderCmd = b.newBuilderCmd(&cobra.Command{
		Use:   fmt.Sprintf("%s [command] [flags]", common.ApplicationShortName),
		Short: "A CLI tool that lets you easily save, access and organize your personal daily notes.",
		Long: `A CLI tool that lets you easily save, access and organize your personal daily notes.
Complete documentation is available at ` + common.Website,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Show help when users executes root command without any args.
			// This may not be the future behavior.
			if len(args) == 0 {
				return cmd.Help()
			}
			return nil
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cc.setStdoutThreshold()
		},
	})
	cc.cmd.SilenceUsage = true // Only show usage for user errors

	// Persistent flags
	cc.cmd.PersistentFlags().BoolVarP(&cc.Verbose, "verbose", "v", false, "verbose output")
	cc.cmd.PersistentFlags().BoolVarP(&cc.Debug, "debug", "", false, "debug output")
	_ = cc.cmd.PersistentFlags().MarkHidden("debug") // Hide in usage
	cc.cmd.PersistentFlags().BoolVarP(&cc.Trace, "trace", "", false, "trace output")
	_ = cc.cmd.PersistentFlags().MarkHidden("trace") // Hide in usage

	return cc
}

func (q quicknoteCmd) setStdoutThreshold() {
	if q.Verbose {
		jww.SetStdoutThreshold(jww.LevelInfo)
	}
	if q.Debug {
		jww.SetStdoutThreshold(jww.LevelDebug)
	}
	if q.Trace {
		jww.SetStdoutThreshold(jww.LevelTrace)
	}
}
