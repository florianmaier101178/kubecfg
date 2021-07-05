package command

import (
	"fmt"
	"strings"
)

type ProjectShowCommand struct {
}

func (p *ProjectShowCommand) Help() string {
	helpText := `
Usage: kubecfg project show

  Show selected project.
`
	return strings.TrimSpace(helpText)
}

func (p *ProjectShowCommand) Run(args []string) int {
	if len(args) > 0 {
		fmt.Println(p.Help())
		return 1
	}

	config, exitStatus := loadConfig()
	if exitStatus != 0 {
		return exitStatus
	}

	fmt.Println("selected project:")
	fmt.Printf("  %s\n", config.SelectedProject)
	fmt.Println()

	return 0
}

func (p *ProjectShowCommand) Synopsis() string {
	return "Show selected project"
}
