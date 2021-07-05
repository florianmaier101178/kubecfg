package main

import (
	"kubecfg/command"
	"log"
	"os"

	"github.com/mitchellh/cli"
)

const version = "1.0.0"

func main() {
	c := cli.NewCLI("kubecfg", version)
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"context": func() (cli.Command, error) {
			return &command.ContextCommand{}, nil
		},
		"context add": func() (cli.Command, error) {
			return &command.ContextAddCommand{}, nil
		},
		"context list": func() (cli.Command, error) {
			return &command.ContextListCommand{}, nil
		},
		"context remove": func() (cli.Command, error) {
			return &command.ContextRemoveCommand{}, nil
		},
		"context show": func() (cli.Command, error) {
			return &command.ContextShowCommand{}, nil
		},
		"context select": func() (cli.Command, error) {
			return &command.ContextSelectCommand{}, nil
		},
		"context unselect": func() (cli.Command, error) {
			return &command.ContextUnselectCommand{}, nil
		},
		"dump": func() (cli.Command, error) {
			return &command.DumpCommand{}, nil
		},
		"info": func() (cli.Command, error) {
			return &command.InfoCommand{}, nil
		},
		"init": func() (cli.Command, error) {
			return &command.InitCommand{}, nil
		},
		"project": func() (cli.Command, error) {
			return &command.ProjectCommand{}, nil
		},
		"project add": func() (cli.Command, error) {
			return &command.ProjectAddCommand{}, nil
		},
		"project list": func() (cli.Command, error) {
			return &command.ProjectListCommand{}, nil
		},
		"project remove": func() (cli.Command, error) {
			return &command.ProjectRemoveCommand{}, nil
		},
		"project select": func() (cli.Command, error) {
			return &command.ProjectSelectCommand{}, nil
		},
		"project show": func() (cli.Command, error) {
			return &command.ProjectShowCommand{}, nil
		},
		"project unselect": func() (cli.Command, error) {
			return &command.ProjectUnselectCommand{}, nil
		},
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Version: version,
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
