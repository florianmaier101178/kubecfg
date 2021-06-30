package command

import (
	"fmt"
	"strings"
)

type ProjectCommand struct {
}

func (p *ProjectCommand) Help() string {
	helpText := `
Usage: kubecfg project

  Add, delete, list, select and show projects.
`
	return strings.TrimSpace(helpText)
}

func (p *ProjectCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println(p.Help())
	}
	return 0
}

func (p *ProjectCommand) Synopsis() string {
	return "Project management"
}
