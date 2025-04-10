package main

import (
	"flag"
	"fmt"
)

type CmdFlags struct {
	Add string
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
	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.add(cf.Add)
	default:
		fmt.Println("Invalid command")
	}
}
