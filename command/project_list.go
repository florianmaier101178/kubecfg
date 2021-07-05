package command

import (
	"fmt"
	"strings"
)

type ProjectListCommand struct {
}

func (p *ProjectListCommand) Help() string {
	helpText := `
Usage: kubecfg project list

  List projects.
`
	return strings.TrimSpace(helpText)
}

func (p *ProjectListCommand) Run(args []string) int {
	if len(args) > 0 {
		fmt.Println(p.Help())
		return 1
	}

	config, exitStatus := loadConfig()
	if exitStatus != 0 {
		return exitStatus
	}

	fmt.Println("available projects:")
	for _, name := range config.GetProjectNames() {
		fmt.Printf("  %s\n", name)
	}
	fmt.Println()

	return 0
}

func (p *ProjectListCommand) Synopsis() string {
	return "List projects"
}
