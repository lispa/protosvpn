package vpn

import (
	"fmt"
	"os/exec"
)

func CreateClient(
	clientName string,
) error {
	command := exec.Command(
		"/usr/local/bin/easyrsa",
		"build-client-full",
		clientName,
		"nopass",
	)

	output, err := command.CombinedOutput()

	if err != nil {
		return fmt.Errorf(
			"failed to create client: %s",
			string(output),
		)
	}

	return nil
}
