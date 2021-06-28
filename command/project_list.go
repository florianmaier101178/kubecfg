package command

import (
	"fmt"
	"kubecfg/io"
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

	if io.IllegalConfigurationSetup() {
		fmt.Println("kubecfg is not properly configured")
		return 1
	}

	config, err := io.LoadConfigFromFileSystem()
	if err != nil {
		fmt.Println("could not load 'config.json'")
		return 1
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
