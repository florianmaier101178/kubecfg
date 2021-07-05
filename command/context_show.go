package command

import (
	"fmt"
	"kubecfg/arguments"
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

	config, exitStatus := loadConfig()
	if exitStatus != 0 {
		return exitStatus
	}

	hasSelectedProject := config.HasSelectedProject()
	if !hasSelectedProject {
		fmt.Println("no project is selected")
		return 1
	}

	projectName, exitStatus := provideProjectName(*optionalProjectNameArg, config)
	if exitStatus != 0 {
		return exitStatus
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
