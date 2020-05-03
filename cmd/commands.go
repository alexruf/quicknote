package cmd

import "github.com/spf13/cobra"

type commandsBuilder struct {
	commands []cmder
}

func newCommandsBuilder() *commandsBuilder {
	return &commandsBuilder{}
}

func (b *commandsBuilder) addCommands(commands ...cmder) *commandsBuilder {
	b.commands = append(b.commands, commands...)
	return b
}

func (b *commandsBuilder) addAll() *commandsBuilder {
	b.addCommands(
		b.newAddCmd(),
		b.newConfigureCmd(),
		b.newVersionCmd(),
	)
	return b
}

func (b *commandsBuilder) build() *quicknoteCmd {
	cc := b.newQuicknoteCmd()
	addChildCommands(cc.getCommand(), b.commands...)
	return cc
}

func addChildCommands(root *cobra.Command, commands ...cmder) {
	for _, command := range commands {
		cmd := command.getCommand()
		if cmd == nil {
			continue
		}
		root.AddCommand(cmd)
	}
}

type baseCmd struct {
	cmd *cobra.Command
}

func (c *baseCmd) getCommand() *cobra.Command {
	return c.cmd
}

func (b *commandsBuilder) newBaseCmd(cmd *cobra.Command) *baseCmd {
	return &baseCmd{cmd: cmd}
}

type baseBuilderCmd struct {
	*baseCmd
	*commandsBuilder
}

func (b *commandsBuilder) newBuilderCmd(cmd *cobra.Command) *baseBuilderCmd {
	return &baseBuilderCmd{commandsBuilder: b, baseCmd: b.newBaseCmd(cmd)}
}
