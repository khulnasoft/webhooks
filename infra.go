package webhooks

import "time"

// Copy the Deploy-related types but make it about infrastructure provisioning instead of deploys.

// InfraChangeStatus describes the different infra change statuses.
type InfraChangeStatus string

const (
	// InfraChangePending means the infra change is pending, meaning it's created
	// but not yet queued.
	InfraChangePending InfraChangeStatus = "pending"

	// InfraChangeAwaitingApproval means the infra change is blocked awaiting approval.
	InfraChangeAwaitingApproval InfraChangeStatus = "awaiting_approval"

	// InfraChangeQueued means the infra change is queued.
	InfraChangeQueued InfraChangeStatus = "queued"

	// InfraChangeRunning means the infra change is running.
	InfraChangeRunning InfraChangeStatus = "running"

	// InfraChangeCompleted means the infra change is completed.
	// See InfraChangeConclusion to know if it succeeded or failed.
	InfraChangeCompleted InfraChangeStatus = "completed"
)

// InfraChangeConclusion defines the result of a infra change.
//
// Additional conclusions may be added in the future to distinguish
// between different types of failures.
type InfraChangeConclusion string

const (
	// InfraChangeSuccess means the infra change succeeded.
	InfraChangeSuccess InfraChangeConclusion = "success"

	// InfraChangeFailure means the infra change failed.
	InfraChangeFailure InfraChangeConclusion = "failure"

	// InfraChangeCanceled means the infra change was canceled.
	InfraChangeCanceled InfraChangeConclusion = "canceled"

	// InfraChangeRejected means the infra change was rejected by the user
	// at the manual approval stage.
	InfraChangeRejected InfraChangeConclusion = "rejected"
)

// InfraChange describes the infrastructure provisioning change phase of a rollout.
type InfraChange struct {
	// ID is the unique id of the infrastructure change.
	ID string `json:"id"`

	// Status describes the current status of the infra change.
	Status InfraChangeStatus `json:"status"`

	// Conclusion describes the conclusion of the infra change.
	// It is set only when the infra change is completed.
	Conclusion InfraChangeConclusion `json:"conclusion"`

	// QueuedAt defines when the infra change was queued.
	// It's nil if the change hasn't been queued yet.
	QueuedAt *time.Time `json:"queued_at"`

	// StartedAt defines when the infra change started.
	// It's nil if the infra change hasn't started yet.
	StartedAt *time.Time `json:"started_at"`

	// CompletedAt defines when the infra change completed.
	// It's nil if the infra change hasn't completed yet.
	CompletedAt *time.Time `json:"completed_at"`
}
