/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/adityanagar10/too-doo/cmd"
	_ "github.com/adityanagar10/too-doo/cmd/add"
	_ "github.com/adityanagar10/too-doo/cmd/complete"
	_ "github.com/adityanagar10/too-doo/cmd/delete"
	_ "github.com/adityanagar10/too-doo/cmd/list"
)

func main() {
	cmd.Execute()
}
