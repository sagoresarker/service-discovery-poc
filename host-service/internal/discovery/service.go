package discovery

import "fmt"

type Service struct {
	agents map[string]string
}

func NewService() *Service {
	return &Service{
		agents: make(map[string]string),
	}
}

func (s *Service) RegisterAgent(agentID, ipAddress string) error {
	if _, exists := s.agents[agentID]; exists {
		return fmt.Errorf("agent with ID %s already registered", agentID)
	}

	s.agents[agentID] = ipAddress
	return nil
}

// Add this new method
func (s *Service) GetAllAgents() map[string]string {
	return s.agents
}
