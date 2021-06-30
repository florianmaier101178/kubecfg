package command

import (
	"errors"
	"fmt"
	"kubecfg/config"
	"kubecfg/io"
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
	if len(args) > 1 {
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

	projectName, err := determineProjectName(args, config)
	if err != nil {
		fmt.Println(err)
		return 1
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

func determineProjectName(args []string, config config.Config) (string, error) {
	if len(args) == 1 {
		return args[0], nil
	}
	if config.HasSelectedProject() {
		return config.SelectedProject, nil
	}
	return "", errors.New(
		fmt.Sprintf("neither projectName was passed in, nor config contains a selected project"))
}

func (c *ContextListCommand) Synopsis() string {
	return "List contexts"
}
