package servers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/FosteredGames/Odyssey/registry/internal/registry/data"
	"github.com/FosteredGames/Odyssey/registry/internal/registry/identity"
)

type Service struct {
	db *data.DB

	mux *http.ServeMux
}

func New(db *data.DB) *Service {
	mux := http.NewServeMux()

	service := &Service{
		db:  db,
		mux: mux,
	}

	mux.HandleFunc("/register", service.Register)

	return service
}

type RegistrationRequest struct {
	Name string `json:"name"`
}

func (service *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.InfoContext(r.Context(), "Servers Service", "path", r.URL.Path)
	service.mux.ServeHTTP(w, r)
}

func (service *Service) Register(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(identity.UserKeyContext).(string)
	if !ok {
		http.Error(w, "no user", http.StatusUnauthorized)
		return
	}

	var req RegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	Register(req.Name, userID)
}
