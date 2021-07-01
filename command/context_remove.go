package command

import (
	"fmt"
	"kubecfg/arguments"
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
  If an optional parameter with an project name is passed in, than for this project the contexts will be listed.

Options:

-project=string			Name of project, e.g. "business"
`
	return strings.TrimSpace(helpText)
}

func (c *ContextRemoveCommand) Run(args []string) int {
	contextRemoveArgs, err := arguments.ParseContextRemoveArguments(args)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		fmt.Println(c.Help())
		return 1
	}

	fmt.Printf("remove context '%s' from project '%s'\n", contextRemoveArgs.ContextName, contextRemoveArgs.ProjectName)

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
	if contextRemoveArgs.ProjectNameAvailable{
		projectName = contextRemoveArgs.ProjectName
	} else {
		projectName, err = determineProjectName(contextRemoveArgs, config)
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

	updatedProject, err := project.RemoveContext(contextRemoveArgs.ContextName)
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

	//TODO handle context link

	return 0
}

func (c *ContextRemoveCommand) Synopsis() string {
	return "Remove a context from a project"
}
