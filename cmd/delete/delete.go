/*
Copyright © 2024 Aditya <includebitslabs@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/adityanagar10/too-doo/cmd"
	todo "github.com/adityanagar10/too-doo/utils"
	"github.com/spf13/cobra"
)

func DeleteCmd() *cobra.Command {
	todos := &todo.Todos{}
	if err := todos.Load(todo.FILENAME); err != nil {
    fmt.Fprintln(os.Stderr, err.Error())
    os.Exit(1)
  }
	var input int
	c := &cobra.Command{
		Use:   "delete [taskID]",
		Short: "Delete a task",
		Long:  `This command allows you to delete a task from your to-do list by providing its unique task ID. Once deleted, the task will be removed permanently and cannot be recovered. Use this command when you want to remove a specific task from your list of active tasks.`,		
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskID, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			input = taskID 
			todos.Delete(input)
			err = todos.Store(todo.FILENAME)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			todos.List()
		},
	}
	return c
}
func init() {
	deleteCmd := DeleteCmd()
	cmd.RootCmd.AddCommand(deleteCmd)
}
