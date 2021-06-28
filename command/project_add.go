package command

import (
	"fmt"
	cfg "kubecfg/config"
	"kubecfg/io"
	"strings"
)

type ProjectAddCommand struct {
}

func (p *ProjectAddCommand) Help() string {
	helpText := `
Usage: kubecfg project add NAME

  Add a new project.
`
	return strings.TrimSpace(helpText)
}

func (p *ProjectAddCommand) Run(args []string) int {
	if len(args) != 1 {
		fmt.Println(p.Help())
		return 1
	}
	projectName := args[0]
	fmt.Printf("add project '%s' to kubecfg configuration\n", projectName)

	if io.IllegalConfigurationSetup() {
		fmt.Println("kubecfg is not properly configured")
		return 1
	}

	config, err := io.LoadConfigFromFileSystem()
	if err != nil {
		fmt.Println("could not load 'config.json'")
		return 1
	}

	exitStatus := io.CreateProjectDir(projectName)
	if exitStatus > 0 {
		return exitStatus
	}

	exitStatus = io.CreateUnselectedContextMarkerFile(projectName)
	if exitStatus > 0 {
		return exitStatus
	}

	exitStatus = io.CreateProjectContextSymlink(projectName)
	if exitStatus > 0 {
		return exitStatus
	}

	updatedConfig, err := config.AddProject(*cfg.NewProject(projectName))
	if err != nil {
		fmt.Printf("could not add project '%s' to config", projectName)
		return 1
	}

	exitStatus = io.WriteConfigToFileSystem(updatedConfig)
	if exitStatus > 0 {
		return exitStatus
	}

	return 0
}

func (p *ProjectAddCommand) Synopsis() string {
	return "Add a new project"
}
