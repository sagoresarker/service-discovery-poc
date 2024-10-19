package agent

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

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
	a.agentID = generateAgentID()
	ipAddress, err := getIPAddress()
	if err != nil {
		return fmt.Errorf("failed to get IP address: %v", err)
	}

	if err := a.client.RegisterAgent(a.agentID, ipAddress); err != nil {
		return fmt.Errorf("failed to register agent: %v", err)
	}

	fmt.Printf("Agent registered successfully. ID: %s, IP: %s\n", a.agentID, ipAddress)

	// Start the health check endpoint
	http.HandleFunc("/health", a.healthHandler)
	go func() {
		log.Printf("Starting health check server on port %d", a.cfg.AgentPort)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", a.cfg.AgentPort), nil); err != nil {
			log.Fatalf("Failed to start health check server: %v", err)
		}
	}()

	// Keep the agent running
	for {
		time.Sleep(30 * time.Second)
		log.Println("Agent is still running...")
	}
}

func (a *Agent) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Agent is healthy"))
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
