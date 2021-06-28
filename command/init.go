package command

import (
	"fmt"
	"kubecfg/config"
	"kubecfg/io"
	"strings"
)

type InitCommand struct {
}

func (i *InitCommand) Help() string {
	helpText := `
Usage: kubecfg init

  Initializes the kubecfg tool. It creates a ~/.kubecfg directory
  to store all configuration needed for kubecfg.
`
	return strings.TrimSpace(helpText)
}

func (i *InitCommand) Run(args []string) int {
	fmt.Println("initialize kubecfg into ~/.kubecfg")

	exitStatus := io.SetupConfigDir()
	if exitStatus > 0 {
		return exitStatus
	}

	exitStatus = io.CreateUnselectedProjectMarkerFile()
	if exitStatus > 0 {
		return exitStatus
	}

	exitStatus = io.CreateProjectSymlink()
	if exitStatus > 0 {
		return exitStatus
	}

	exitStatus = io.WriteInitialConfigToFileSystem(config.NewConfig())
	if exitStatus > 0 {
		return exitStatus
	}

	return 0
}

func (i *InitCommand) Synopsis() string {
	return "Initialize kubecfg tool"
}
