package command

import (
	"fmt"
	"kubecfg/arguments"
	"strings"
)

type ContextListCommand struct {
}

func (c *ContextListCommand) Help() string {
	helpText := `
Usage: kubecfg context list [NAME]

  List contexts.
  If this command is called parameterless, than the selected project (if a project is selected) is used.
  If an optional parameter with an project name is passed in, than for this project the contexts will be listed.

Optional parameters:

NAME	projectName
`
	return strings.TrimSpace(helpText)
}

func (c *ContextListCommand) Run(args []string) int {
	optionalProjectNameArg, err := arguments.ParseOptionalProjectNameArgument(args)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		fmt.Println(c.Help())
		return 1
	}
	if len(args) > 1 {
		fmt.Println(c.Help())
		return 1
	}

	config, exitStatus := loadConfig()
	if exitStatus != 0 {
		return exitStatus
	}

	var projectName string //from here on use projectName instead of contextRemoveArgs.ProjectName
	if optionalProjectNameArg.Available() {
		projectName = optionalProjectNameArg.Name()
	} else {
		projectName, err = determineProjectName(optionalProjectNameArg, *config)
		if err != nil {
			fmt.Println(err)
			return 1
		}
	}

	existingProject, _ := config.ExistingProject(projectName)
	if !existingProject {
		fmt.Printf("No project with name: '%s' available for usage\n", projectName)
		return 1
	}

	fmt.Printf("available contexts for project: '%s'\n", projectName)
	project, err := config.GetProject(projectName)
	if err != nil {
		fmt.Println(err)
		return 1
	}
	for _, name := range project.GetContextNames() {
		fmt.Printf("  %s\n", name)
	}
	fmt.Println()

	return 0
}

func (c *ContextListCommand) Synopsis() string {
	return "List contexts"
}
