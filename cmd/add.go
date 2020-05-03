package cmd

import (
	"github.com/spf13/cobra"
	jww "github.com/spf13/jwalterweatherman"
	"strings"
)

type addCmd struct {
	*baseBuilderCmd
	category string
}

func (b *commandsBuilder) newAddCmd() *addCmd {
	cc := &addCmd{}
	cc.baseBuilderCmd = b.newBuilderCmd(&cobra.Command{
		Use:     "add [note]",
		Short:   "Add a new note",
		Long:    `Add a new note.`,
		Example: "pn add \"A note\" -c \"meeting\"",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cc.Add(args)
		},
	})

	cc.getCommand().Flags().StringVarP(&cc.category, "category", "c", "", "Category name")

	return cc
}

func (a *addCmd) Add(args []string) error {
	jww.FEEDBACK.Println(a.getCommand().Short)
	jww.FEEDBACK.Println("category: " + a.category)
	jww.FEEDBACK.Println("note(s):\n" + strings.Join(args, "\n"))
	return nil
}
