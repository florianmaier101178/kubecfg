package azure

import (
	"fmt"
	"os"
	"os/exec"
)

func CreateCredentialsForContext(resourceGroup string, clusterName string) {
	fmt.Println("If you get an error, try to switch the subscription first, unfortunately this is not possible from go code ...")

	command := exec.Command("az", "aks", "get-credentials",
		"--resource-group", fmt.Sprintf("%s", resourceGroup),
		"--name", fmt.Sprintf("%s", clusterName))

	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Run()
}
