package azure

import (
	"fmt"
	"kubecfg/arguments"
	"os"
	"os/exec"
)

func CreateCredentialsForContext(args *arguments.ContextAddArguments) {
	fmt.Println("If you get an error, try to switch the subscription first, unfortunately this is not possible from go code ...")

	command := exec.Command("az", "aks", "get-credentials",
		"--resource-group", fmt.Sprintf("%s", args.ResourceGroup),
		"--name", fmt.Sprintf("%s", args.ClusterName))

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}
