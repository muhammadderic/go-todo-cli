package main

import (
	"flag"
	"fmt"
)

type CmdFlags struct {
	Add  string
	List bool
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
	default:
		fmt.Println("Invalid command")
	}
}
