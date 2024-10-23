package webhooks

import "time"

// Application describes an KhulnaSoft application.
type Application struct {
	// ID is a unique id for this application.
	ID string `json:"id"`

	// Slug is the unique, human-readable string used to identify the application.
	Slug string `json:"slug"`

	// CreatedAt is when the application was created.
	CreatedAt time.Time `json:"created_at"`
}
