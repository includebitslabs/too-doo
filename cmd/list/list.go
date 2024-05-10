/*
Copyright © 2024 Aditya <includebitslabs@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/adityanagar10/too-doo/cmd"
	todo "github.com/adityanagar10/too-doo/utils"
	"github.com/spf13/cobra"
)

func ListCmd() *cobra.Command {
	todos := &todo.Todos{}
	if err := todos.Load(todo.FILENAME); err != nil {
    fmt.Fprintln(os.Stderr, err.Error())
    os.Exit(1)
  }
	c := &cobra.Command{
		Use:   "list",
		Short: "List all tasks",
		Long:  `This command allows you to list all the available to-dos`,
		Run: func(cmd *cobra.Command, args []string) {
			todos.List()
		},
	}
	return c
}
func init() {
	listCmd := ListCmd()
	cmd.RootCmd.AddCommand(listCmd)
}
