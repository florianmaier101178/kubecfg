package command

import (
	"fmt"
	"kubecfg/io"
	"strings"
)

type ContextSelectCommand struct {
}

func (c *ContextSelectCommand) Help() string {
	helpText := `
Usage: kubecfg context select NAME [options]

  Select a context for a project.
  If this command is called parameterless, than the selected project (if a project is selected) is used.
  If an optional parameter with an project name is passed in, than for this project the context will be selected.

Options:

-project=string			Name of project, e.g. "business"
`
	return strings.TrimSpace(helpText)
}

func (c *ContextSelectCommand) Run(args []string) int {
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

	updatedProject, err := project.SelectContext(projectNameAndContextArgs.ContextName)
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

	exitStatus = io.UpdateProjectContextSymlink(projectName, projectNameAndContextArgs.ContextName)
	if exitStatus > 0 {
		return exitStatus
	}

	fmt.Printf("selected context '%s' for project '%s'\n", projectNameAndContextArgs.ContextName, projectName)

	return 0
}

func (c *ContextSelectCommand) Synopsis() string {
	return "Select a context for a project"
}
