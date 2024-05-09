/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package complete

import (
	"fmt"
	"os"
	"strconv"

	"github.com/adityanagar10/too-doo/cmd"
	todo "github.com/adityanagar10/too-doo/utils"
	"github.com/spf13/cobra"
)

const (
  todoFile = ".todos.json"
)

func CompleteCmd() *cobra.Command {
	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
    fmt.Fprintln(os.Stderr, err.Error())
    os.Exit(1)
  }
	var input int
	c := &cobra.Command{
		Use:   "complete [taskID]",
		Short: "Mark a task as complete",
		Long:  `This command marks a task as complete by providing its ID.`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			taskID, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
			input = taskID 
			fmt.Printf("%d", input)
			todos.Complete(input)
			err = todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)
				os.Exit(1)
			}
		},
	}
	return c
}

func init() {
	completeCmd := CompleteCmd()
	cmd.RootCmd.AddCommand(completeCmd)
}
