package openvpn

import (
	"bufio"
	"os"
	"strings"
)

type Client struct {
	Name          string `json:"name"`
	RealAddress   string `json:"real_address"`
	BytesReceived string `json:"bytes_received"`
	BytesSent     string `json:"bytes_sent"`
}

func ParseStatusFile(path string) ([]Client, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var clients []Client

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "CLIENT_LIST") {
			parts := strings.Fields(line)

			if len(parts) < 6 {
				continue
			}

			client := Client{
				Name:          parts[1],
				RealAddress:   parts[2],
				BytesReceived: parts[4],
				BytesSent:     parts[5],
			}

			clients = append(clients, client)
		}
	}

	return clients, scanner.Err()
}
