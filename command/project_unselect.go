package command

import (
	"fmt"
	"kubecfg/io"
	"strings"
)

type ProjectUnselectCommand struct {
}

func (p *ProjectUnselectCommand) Help() string {
	helpText := `
Usage: kubecfg project unselect

  Unselect project choice.
`
	return strings.TrimSpace(helpText)
}

func (p *ProjectUnselectCommand) Run(args []string) int {
	if len(args) > 0 {
		fmt.Println(p.Help())
		return 1
	}
	fmt.Println("unselect project choice")

	if io.IllegalConfigurationSetup() {
		fmt.Println("kubecfg is not properly configured")
		return 1
	}

	config, err := io.LoadConfigFromFileSystem()
	if err != nil {
		fmt.Println("could not load 'config.json'")
		return 1
	}

	exitStatus := io.WriteUpdatedConfigToFileSystem(config.UnselectProject())
	if exitStatus > 0 {
		return exitStatus
	}

	exitStatus = io.UpdateConfigProjectSymlinkToUnselected()
	if exitStatus > 0 {
		return exitStatus
	}

	return 0
}

func (p *ProjectUnselectCommand) Synopsis() string {
	return "Unselect project choice"
}
