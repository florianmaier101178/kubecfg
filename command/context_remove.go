package command

import (
	"fmt"
	"kubecfg/io"
	"strings"
)

type ContextRemoveCommand struct {
}

func (c *ContextRemoveCommand) Help() string {
	helpText := `
Usage: kubecfg context remove NAME [options]

  Remove a context from a project.
  If this command is called parameterless, than the selected project (if a project is selected) is used.
  If an optional parameter with an project name is passed in, than for this project the context will be removed.

Options:

-project=string			Name of project, e.g. "business"
`
	return strings.TrimSpace(helpText)
}

func (c *ContextRemoveCommand) Run(args []string) int {
	projectNameAndContextArgs, err := ParseProjectNameAndContextArguments(args)
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

	projectName, exitStatus := provideProjectName(*projectNameAndContextArgs, config)
	if exitStatus != 0 {
		return exitStatus
	}

	project, err := config.GetProject(projectName)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	updatedProject, err := project.RemoveContext(projectNameAndContextArgs.ContextName)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	updatedConfig, err := config.UpdateProject(*updatedProject)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	exitStatus = io.WriteUpdatedConfigToFileSystem(updatedConfig)
	if exitStatus > 0 {
		return exitStatus
	}

	exitStatus = io.UpdateProjectContextSymlinkToUnselected(projectName)
	if exitStatus > 0 {
		return exitStatus
	}

	fmt.Printf("removed context '%s' from project '%s'\n", projectNameAndContextArgs.ContextName, projectNameAndContextArgs.ProjectName())

	return 0
}

func (c *ContextRemoveCommand) Synopsis() string {
	return "Remove a context from a project"
}
