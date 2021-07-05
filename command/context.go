package command

import (
	"errors"
	"fmt"
	"kubecfg/arguments"
	"kubecfg/config"
	"strings"
)

type ContextCommand struct {
}

func (c *ContextCommand) Help() string {
	helpText := `
Usage: kubecfg context

  Add, delete, list, select and show contexts for the currently selected project.
`
	return strings.TrimSpace(helpText)
}

func (c *ContextCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(c.Help())
	}
	return 0
}

func (c *ContextCommand) Synopsis() string {
	return "Context management"
}

func determineProjectName(projectName arguments.OptionalProjectName, config config.Config) (string, error) {
	if projectName.Available() {
		return projectName.Name(), nil
	}
	if config.HasSelectedProject() {
		return config.SelectedProject, nil
	}
	return "", errors.New(
		fmt.Sprintf("neither projectName was passed in, nor config contains a selected project"))
}
