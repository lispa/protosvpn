package vpn

import (
	"fmt"
	"os/exec"
)

func RevokeClient(
	clientName string,
) error {
	revokeCommand := exec.Command(
		"docker",
		"exec",
		"protosvpn-openvpn",
		"easyrsa",
		"--batch",
		"revoke",
		clientName,
	)

	revokeOutput, err := revokeCommand.CombinedOutput()

	if err != nil {
		return fmt.Errorf(
			"failed to revoke client: %s",
			string(revokeOutput),
		)
	}

	crlCommand := exec.Command(
		"docker",
		"exec",
		"protosvpn-openvpn",
		"easyrsa",
		"gen-crl",
	)

	crlOutput, err := crlCommand.CombinedOutput()

	if err != nil {
		return fmt.Errorf(
			"failed to generate CRL: %s",
			string(crlOutput),
		)
	}

	return nil
}
