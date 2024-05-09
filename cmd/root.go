/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	todo "github.com/adityanagar10/too-doo/utils"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "too-doo",
	Short: "Manage your to-do list from the command line",
	Long:  `Too-Doo is a command-line application for managing your to-do list efficiently. With Too-Doo, you can add, delete, complete, and view tasks directly from your terminal, making it easy to stay organized and productive. This tool is built using Cobra, a CLI library for Go, and offers a seamless experience for managing your tasks with ease. Use Too-Doo to streamline your task management process and stay on top of your daily agenda.`,	
Run: func(cmd *cobra.Command, args []string) {
	todos := &todo.Todos{}
	
	if err := todos.Load(todo.FILENAME); err != nil {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
	}
	
	todos.List()
},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


