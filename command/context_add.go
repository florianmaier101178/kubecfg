package command

import (
	"fmt"
	"kubecfg/azure"
	"kubecfg/io"
	"strings"
)

type ContextAddCommand struct {
}

func (c *ContextAddCommand) Help() string {
	helpText := `
Usage: kubecfg context add NAME [options] 
  
  Add a new context for the currently selected project.

Options:

-project=string			Name of project, e.g. "business"
-rg=string			Resource group name, e.g. "hau-bus-dev-aks-rg". 
-context=string			Name of context, e.g. "haubusdevaks".
`
	return strings.TrimSpace(helpText)
}

/*
	Unfortunately it is not possible to set the azure subscription
	via a exec.Command call.
	If the the subscription name contains upper case letters, like
	in "HAU Business Dev" than this name gets somehow passed to the
	shell as "hau business dev". This leads to an error in the
	az account set command.
*/
func (c *ContextAddCommand) Run(args []string) int {
	contextAddArgs, err := ParseContextAddArguments(args)
	if err != nil {
		fmt.Println(err)
		fmt.Println()
		fmt.Println(c.Help())
		return 1
	}

	config, exitStatus := loadConfig()
	if exitStatus != 0 {
		return exitStatus
	}

	existingProject, _ := config.ExistingProject(contextAddArgs.ProjectName)
	if !existingProject {
		fmt.Printf("No project with name: '%s' available for usage\n", contextAddArgs.ProjectName)
		return 1
	}

	io.BackupKubeConfigFile()
	azure.CreateCredentialsForContext(contextAddArgs.ResourceGroup, contextAddArgs.ClusterName)
	io.SaveCreatedContextFile(contextAddArgs.ProjectName, string(contextAddArgs.ContextName))
	io.RestoreBackedUpKubeConfigFile()

	project, err := config.GetProject(contextAddArgs.ProjectName)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	updatedProject, err := project.AddContext(contextAddArgs.ContextName)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	updatedConfig, err := config.UpdateProject(*updatedProject)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	exitStatus = io.WriteUpdatedConfigToFileSystem(updatedConfig)
	if exitStatus > 0 {
		return exitStatus
	}

	fmt.Printf("Added context: '%s' to project: '%s'\n", contextAddArgs.ContextName, contextAddArgs.ProjectName)
	return 0
}

func (c *ContextAddCommand) Synopsis() string {
	return "Add a new context for the currently selected project"
}
