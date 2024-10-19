package discovery

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	hostServiceURL string
}

func NewClient(hostServiceURL string) *Client {
	return &Client{
		hostServiceURL: hostServiceURL,
	}
}

func (c *Client) RegisterAgent(agentID, ipAddress string) error {
	data := map[string]string{
		"agent_id":   agentID,
		"ip_address": ipAddress,
	}

	log.Printf("Registering agent with data: %v", data)

	log.Printf("Sending registration request to: %s", fmt.Sprintf("%s/register", c.hostServiceURL))

	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON data: %v", err)
	}

	log.Printf("Request payload: %s", string(jsonData))

	log.Printf("Registering agent with URL: %s/register", c.hostServiceURL)

	resp, err := http.Post(fmt.Sprintf("%s/register", c.hostServiceURL), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send registration request: %v", err)
	}
	defer resp.Body.Close()

	log.Printf("Received response with status code: %d", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("registration failed with status code: %d", resp.StatusCode)
	}

	return nil
}
