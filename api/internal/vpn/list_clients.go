package vpn

import (
	"bufio"
	"os"
	"strings"
)

type VPNClient struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func ListClients() ([]VPNClient, error) {
	file, err := os.Open(
		"/etc/openvpn/pki/index.txt",
	)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var clients []VPNClient

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "V") ||
			strings.HasPrefix(line, "R") {

			parts := strings.Split(line, "\t")

			if len(parts) < 6 {
				continue
			}

			status := "active"

			if strings.HasPrefix(line, "R") {
				status = "revoked"
			}

			namePart := parts[len(parts)-1]

			name := strings.TrimPrefix(
				namePart,
				"/CN=",
			)

			clients = append(
				clients,
				VPNClient{
					Name:   name,
					Status: status,
				},
			)
		}
	}

	return clients, nil
}
