package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	List   bool
	Toggle int
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	// how to use: go run ./ --add="Buy milk"
	flag.StringVar(
		&cf.Add,                        // pointer to a string variable where the value will be stored
		"add",                          // the name of the flag (e.g., --add)
		"",                             // default value if the flag is not provided
		"Add a new todo specify title", // help message shown in `--help`
	)
	// how to use: go run ./ --toggle=1
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify todo by index to toggle complete true/false")
	// how to use: go run ./ --edit=1:"New title"
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	// how to use: go run ./ --del=1
	flag.IntVar(&cf.Del, "del", -1, "Specify todo by index to delete")
	// how to use: go run ./ --list
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid format for edit. Please use index:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index for edit.")
			os.Exit(1)

		}
		todos.edit(index, parts[1])
	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("Invalid command")
	}
}
