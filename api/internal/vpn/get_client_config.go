package vpn

import (
	"fmt"
	"os/exec"
)

func GetClientConfig(
	clientName string,
) ([]byte, error) {
	command := exec.Command(
		"docker",
		"exec",
		"protosvpn-openvpn",
		"ovpn_getclient",
		clientName,
	)

	output, err := command.CombinedOutput()

	if err != nil {
		return nil, fmt.Errorf(
			"failed to get client config: %s",
			string(output),
		)
	}

	return output, nil
}
