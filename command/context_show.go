package command

import (
	"fmt"
	"kubecfg/arguments"
	"kubecfg/io"
	"strings"
)

type ContextShowCommand struct {
}

func (c *ContextShowCommand) Help() string {
	helpText := `
Usage: kubecfg context show NAME

  Show selected context of selected project.
  If this command is called parameterless, than the selected project (if a project is selected) is used.
  If an optional parameter with an project name is passed in, than for this project the selected context will be shown.
`
	return strings.TrimSpace(helpText)
}

func (c *ContextShowCommand) Run(args []string) int {
	optionalProjectNameArg, err := arguments.ParseOptionalProjectNameArgument(args)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		fmt.Println(c.Help())
		return 1
	}

	if io.IllegalConfigurationSetup() {
		fmt.Println("kubecfg is not properly configured")
		return 1
	}

	config, err := io.LoadConfigFromFileSystem()
	if err != nil {
		fmt.Println("could not load 'config.json'")
		return 1
	}

	hasSelectedProject := config.HasSelectedProject()
	if !hasSelectedProject {
		fmt.Println("no project is selected")
		return 1
	}

	var projectName string //from here on use projectName instead of contextRemoveArgs.ProjectName
	if optionalProjectNameArg.Available() {
		projectName = optionalProjectNameArg.Name()
	} else {
		projectName, err = determineProjectName(optionalProjectNameArg, config)
		if err != nil {
			fmt.Println(err)
			return 1
		}
	}

	project, err := config.GetProject(projectName)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Printf("selected context for project '%s':\n", project.Name)
	fmt.Printf("  %s\n", project.SelectedContext)
	fmt.Println()

	return 0
}

func (c *ContextShowCommand) Synopsis() string {
	return "Show selected context of selected project"
}
