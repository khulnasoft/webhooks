package webhooks

import "time"

// DeployStatus defines the different deploy statuses.
type DeployStatus string

const (
	// DeployPending means the deploy is pending, meaning it's created
	// but not yet queued.
	DeployPending DeployStatus = "pending"

	// DeployQueued means the deploy is queued.
	DeployQueued DeployStatus = "queued"

	// DeployRunning means the deploy is running.
	DeployRunning DeployStatus = "running"

	// DeployCompleted means the deploy is completed.
	// See DeployConclusion to know if it succeeded or failed.
	DeployCompleted DeployStatus = "completed"
)

// DeployConclusion defines the result of a deploy.
//
// Additional conclusions may be added in the future to distinguish
// between different types of failures.
type DeployConclusion string

const (
	// DeploySuccess means the deploy succeeded.
	DeploySuccess DeployConclusion = "success"

	// DeployFailure means the deploy failed.
	DeployFailure DeployConclusion = "failure"

	// DeployCanceled means the deploy was canceled.
	DeployCanceled DeployConclusion = "canceled"
)

// Deploy describes the deploy phase of a rollout.
type Deploy struct {
	// ID is the unique id of the deploy.
	ID string `json:"id"`

	// Status describes the current status of the deploy.
	Status DeployStatus `json:"status"`

	// Conclusion describes the conclusion of the deploy.
	// It is set only when the deploy is completed.
	Conclusion DeployConclusion `json:"conclusion"`

	// QueuedAt defines when the deploy was queued.
	// It's nil if the status is pending.
	QueuedAt *time.Time `json:"queued_at"`

	// StartedAt defines when the deploy started.
	// It's nil if the deploy hasn't started yet.
	StartedAt *time.Time `json:"started_at"`

	// CompletedAt defines when the deploy completed.
	// It's nil if the deploy hasn't completed yet.
	CompletedAt *time.Time `json:"completed_at"`
}
