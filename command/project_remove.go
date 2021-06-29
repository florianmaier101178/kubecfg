package command

import (
	"fmt"
	"kubecfg/io"
	"strings"
)

type ProjectRemoveCommand struct {
}

func (p *ProjectRemoveCommand) Help() string {
	helpText := `
Usage: kubecfg project remove NAME

  Remove a project.

  Important, only the entry from the 'config.json' will be removed.
  The project subdirectory with the actual context configurations will stay
  as they are.
`
	return strings.TrimSpace(helpText)
}

func (p *ProjectRemoveCommand) Run(args []string) int {
	if len(args) != 1 {
		fmt.Println(p.Help())
		return 1
	}
	projectName := args[0]
	fmt.Printf("remove project '%s' from kubecfg configuration\n", projectName)

	if io.IllegalConfigurationSetup() {
		fmt.Println("kubecfg is not properly configured")
		return 1
	}

	config, err := io.LoadConfigFromFileSystem()
	if err != nil {
		fmt.Println("could not load 'config.json'")
		return 1
	}

	existingProject, _ := config.ExistingProject(projectName)
	if !existingProject {
		fmt.Printf("No project with name: '%s' available for usage\n", projectName)
		return 1
	}

	selected := config.IsProjectSelected(projectName)
	if selected {
		exitStatus := io.UpdateConfigProjectSymlinkToUnselected()
		if exitStatus > 0 {
			return exitStatus
		}
	}

	updatedConfig, err := config.RemoveProject(projectName)
	if err != nil {
		fmt.Printf("could not remove project '%s' from config", projectName)
		return 1
	}

	exitStatus := io.WriteUpdatedConfigToFileSystem(updatedConfig)
	if exitStatus > 0 {
		return exitStatus
	}

	return 0
}

func (p *ProjectRemoveCommand) Synopsis() string {
	return "Remove a project"
}
