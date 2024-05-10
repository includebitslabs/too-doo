/*
Copyright © 2024 Aditya <includebitslabs@gmail.comS>
*/
package add

import (
	"fmt"
	"os"

	"github.com/adityanagar10/too-doo/cmd"
	todo "github.com/adityanagar10/too-doo/utils"
	"github.com/spf13/cobra"
)

func AddCmd() *cobra.Command {
	todos := &todo.Todos{}
	if err := todos.Load(todo.FILENAME); err != nil {
    fmt.Fprintln(os.Stderr, err.Error())
    os.Exit(1)
  }
	var input string
	c := &cobra.Command{
		Use:   "add [message]",
		Short: "Add a task to the to-do list",
		Long:  `This command allows you to add a new task to your to-do list by providing a brief message describing the task. Once added, the task will become part of your list of active tasks. Use this command when you want to create a new task and include it in your to-do list.`,		
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			input = args[0]
			todos.Add(input)
			err := todos.Store(todo.FILENAME)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			todos.List()
		},
	}
	return c
}

func init() {
	addCmd := AddCmd()
	cmd.RootCmd.AddCommand(addCmd)
}
