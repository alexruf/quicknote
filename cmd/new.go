package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var category string
var project string

func init() {
	newCmd.Flags().StringVarP(&category, "category", "c", "", "Category name")
	newCmd.Flags().StringVarP(&project, "project", "p", "", "Project name")

	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:     "new [note]",
	Short:   "Create a new note",
	Long:    `Create a new note.`,
	Example: "pn new \"A note\" -c \"meeting\" -p \"my-project\"",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Short)
		fmt.Println("category: " + category)
		fmt.Println("project: " + project)
		fmt.Println("note(s):\n" + strings.Join(args, "\n"))
	},
}
