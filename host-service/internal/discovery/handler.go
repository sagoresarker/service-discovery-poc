package discovery

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterAgent(w http.ResponseWriter, r *http.Request) {
	var data struct {
		AgentID   string `json:"agent_id"`
		IPAddress string `json:"ip_address"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.service.RegisterAgent(data.AgentID, data.IPAddress); err != nil {
		http.Error(w, "Failed to register agent", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Add this new handler method
func (h *Handler) GetAllAgents(w http.ResponseWriter, r *http.Request) {
	agents := h.service.GetAllAgents()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(agents)
}
