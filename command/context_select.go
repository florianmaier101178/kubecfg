package command

import (
	"fmt"
	"kubecfg/arguments"
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
	contextSelectArgs, err := arguments.ParseContextSelectArguments(args)
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

	var projectName string //from here on use projectName instead of contextRemoveArgs.ProjectName
	if contextSelectArgs.ProjectNameAvailable{
		projectName = contextSelectArgs.ProjectName
	} else {
		projectName, err = determineProjectName(contextSelectArgs, config)
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

	project, err := config.GetProject(projectName)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	updatedProject, err := project.SelectContext(contextSelectArgs.ContextName)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	updatedConfig, err := config.UpdateProject(*updatedProject)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	exitStatus := io.WriteUpdatedConfigToFileSystem(updatedConfig)
	if exitStatus > 0 {
		return exitStatus
	}

	exitStatus = io.UpdateProjectContextSymlink(projectName, contextSelectArgs.ContextName)
	if exitStatus > 0 {
		return exitStatus
	}

	fmt.Printf("selected context '%s' for project '%s'\n", contextSelectArgs.ContextName, projectName)

	return 0
}

func (c *ContextSelectCommand) Synopsis() string {
	return "Select a context for a project"
}
