package webhooks

import "time"

// RolloutStatus defines the different rollout statuses.
type RolloutStatus string

const (
	// RolloutPending means the rollout is pending, meaning it's created
	// but not yet queued.
	RolloutPending RolloutStatus = "pending"

	// RolloutQueued means the rollout is queued.
	RolloutQueued RolloutStatus = "queued"

	// RolloutRunning means the rollout is running.
	RolloutRunning RolloutStatus = "running"

	// RolloutCompleted means the rollout is completed.
	// See RolloutConclusion to know if it succeeded or failed.
	RolloutCompleted RolloutStatus = "completed"
)

// RolloutConclusion defines the result of a rollout.
//
// Additional conclusions may be added in the future to distinguish
// between different types of failures.
type RolloutConclusion string

const (
	// RolloutSuccess means the rollout succeeded.
	RolloutSuccess RolloutConclusion = "success"

	// RolloutFailure means the rollout failed.
	RolloutFailure RolloutConclusion = "failure"

	// RolloutCanceled means the rollout was canceled.
	RolloutCanceled RolloutConclusion = "canceled"
)

type Rollout struct {
	// ID is the unique id of the rollout.
	ID string `json:"id"`

	// Status describes the current status of the rollout.
	Status RolloutStatus `json:"status"`

	// Conclusion describes the conclusion of the rollout.
	// It is set only when the rollout is completed.
	Conclusion RolloutConclusion `json:"conclusion"`

	// Build is the build used by the rollout.
	// It's always non-nil.
	//
	// KhulnaSoft re-uses builds when possible, so the build may
	// be shared by multiple rollouts, and may be already underway or
	// completed when a rollout is created.
	Build *Build `json:"build"`

	// InfraProvision describes the infrastructure provisioning phase of the rollout.
	// It's nil until the infrastructure provisioning phase starts.
	InfraProvision *InfraChange `json:"infra_provision"`

	// Deploy describes the deploy phase of the rollout.
	// It's nil until the deploy phase starts.
	Deploy *Deploy `json:"deploy"`

	// QueuedAt defines when the rollout was queued.
	// It's nil if the status is pending.
	QueuedAt *time.Time `json:"queued_at"`

	// StartedAt defines when the rollout started.
	// It's nil if the rollout hasn't started yet.
	StartedAt *time.Time `json:"started_at"`

	// CompletedAt defines when the rollout completed.
	// It's nil if the rollout hasn't completed yet.
	CompletedAt *time.Time `json:"completed_at"`
}
