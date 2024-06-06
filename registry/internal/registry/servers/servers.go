package servers

import (
	"net/http"
)

type ServersServer struct {
}

func (s *ServersServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
