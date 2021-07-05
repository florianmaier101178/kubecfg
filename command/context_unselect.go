package command

import (
	"fmt"
	"kubecfg/arguments"
	"kubecfg/io"
	"strings"
)

type ContextUnselectCommand struct {
}

func (c *ContextUnselectCommand) Help() string {
	helpText := `
Usage: kubecfg context unselect NAME

  Unselect context choice for a project.
  If this command is called parameterless, than the selected project (if a project is selected) is used.
  If an optional parameter with an project name is passed in, than for this project the context will be unselected.
`
	return strings.TrimSpace(helpText)
}

func (c *ContextUnselectCommand) Run(args []string) int {
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

	projectName, exitStatus := provideProjectName(*optionalProjectNameArg, config)
	if exitStatus != 0 {
		return exitStatus
	}

	existingProject, _ := config.ExistingProject(projectName)
	if !existingProject {
		fmt.Printf("No project with name: '%s' available for usage\n", projectName)
		return 1
	}

	project, err := config.GetProject(projectName)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	updatedConfig, err := config.UpdateProject(*project.UnselectContext())
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

	fmt.Printf("context selection unset for project '%s'\n", projectName)

	return 0
}

func (c *ContextUnselectCommand) Synopsis() string {
	return "Unselect context choice for a project"
}
