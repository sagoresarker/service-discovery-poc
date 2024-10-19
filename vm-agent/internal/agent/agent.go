package agent

import (
	"fmt"
	"net"

	"github.com/google/uuid"
	"github.com/sagoresarker/svc-discovery-vm-poc/internal/config"
	"github.com/sagoresarker/svc-discovery-vm-poc/internal/discovery"
)

type Agent struct {
	cfg     *config.Config
	client  *discovery.Client
	agentID string
}

func NewAgent(cfg *config.Config) *Agent {
	return &Agent{
		cfg:    cfg,
		client: discovery.NewClient(cfg.HostServiceURL),
	}
}

func (a *Agent) Run() error {
	// Generate a unique agent ID
	a.agentID = generateAgentID()

	// Get the VM's IP address
	ipAddress, err := getIPAddress()
	if err != nil {
		return fmt.Errorf("failed to get IP address: %v", err)
	}

	// Register the agent with the host service
	if err := a.client.RegisterAgent(a.agentID, ipAddress); err != nil {
		return fmt.Errorf("failed to register agent: %v", err)
	}

	fmt.Printf("Agent registered successfully. ID: %s, IP: %s\n", a.agentID, ipAddress)
	return nil
}

func generateAgentID() string {
	id := uuid.New()
	return id.String()
}

func getIPAddress() (string, error) {
	// Get the VM's IP address dynamically
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("failed to get network interfaces: %v", err)
	}

	for _, iface := range interfaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			addrs, err := iface.Addrs()
			if err != nil {
				return "", fmt.Errorf("failed to get addresses for interface %s: %v", iface.Name, err)
			}
			for _, addr := range addrs {
				if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil {
					return ipNet.IP.String(), nil // Return the first non-loopback IPv4 address
				}
			}
		}
	}
	return "", fmt.Errorf("no valid IP address found")
}
