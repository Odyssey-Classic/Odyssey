package servers

import "time"

type Server struct {
	Name string
	Host string
	Port int

	Registration Registration
}

type Registration struct {
	UserID      string
	APIKey      string `json:"-"`
	CreatedDate time.Time
}
