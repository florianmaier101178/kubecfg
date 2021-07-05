package command

import (
	"fmt"
	"kubecfg/io"
	"strings"
)

type ProjectSelectCommand struct {
}

func (p *ProjectSelectCommand) Help() string {
	helpText := `
Usage: kubecfg project select NAME

  Select a project.
`
	return strings.TrimSpace(helpText)
}

func (p *ProjectSelectCommand) Run(args []string) int {
	if len(args) != 1 {
		fmt.Println(p.Help())
		return 1
	}
	projectName := args[0]
	fmt.Printf("select project '%s' for usage\n", projectName)

	config, exitStatus := loadConfig()
	if exitStatus != 0 {
		return exitStatus
	}

	existingProject, _ := config.ExistingProject(projectName)
	if !existingProject {
		fmt.Printf("No project with name: '%s' available for usage\n", projectName)
		return 1
	}

	updatedConfig, err := config.SelectProject(projectName)
	if err != nil {
		fmt.Printf("could not add project '%s' to config\n", projectName)
		return 1
	}

	exitStatus = io.WriteUpdatedConfigToFileSystem(updatedConfig)
	if exitStatus > 0 {
		return exitStatus
	}

	exitStatus = io.UpdateConfigProjectSymlink(projectName)
	if exitStatus > 0 {
		return exitStatus
	}

	return 0
}

func (p *ProjectSelectCommand) Synopsis() string {
	return "Select a project"
}
