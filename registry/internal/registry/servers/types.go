package servers

// ServerInfo holds metadata about a registered game server.
type ServerInfo struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	// Add more fields as needed (e.g., player count, region, etc.)
}
