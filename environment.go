package webhooks

import "time"

type EnvironmentType string

const (
	// EnvironmentTypeProduction represents production environments.
	EnvironmentTypeProduction EnvironmentType = "production"
	// EnvironmentTypeDevelopment represents development environments,
	// meaning a cloud-hosted, persistent environment that is not production.
	EnvironmentTypeDevelopment EnvironmentType = "development"
	// EnvironmentTypePreview represents preview environments,
	// meaning an ephemeral development environment for a specific Pull Request.
	EnvironmentTypePreview EnvironmentType = "preview"
)

// Environment describes an KhulnaSoft application environment.
type Environment struct {
	// ID is a unique id for this environment.
	ID string `json:"id"`

	// Name is the name of the environment.
	Name string `json:"name"`

	// Type is the type of the environment.
	Type EnvironmentType `json:"type"`

	// APIBaseURL is the base URL for making requests to this environment.
	APIBaseURL string `json:"api_base_url"`

	// CreatedAt is the time the environment was created.
	CreatedAt time.Time `json:"created_at"`
}
