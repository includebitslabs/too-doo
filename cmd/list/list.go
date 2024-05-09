/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/adityanagar10/too-doo/cmd"
	todo "github.com/adityanagar10/too-doo/utils"
	"github.com/spf13/cobra"
)

const (
  todoFile = ".todos.json"
)

func ListCmd() *cobra.Command {
	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
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
