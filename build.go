package webhooks

import "time"

// BuildStatus defines the different build statuses.
type BuildStatus string

const (
	// BuildQueued means the build is queued.
	BuildQueued BuildStatus = "queued"

	// BuildRunning means the build is running.
	BuildRunning BuildStatus = "running"

	// BuildCompleted means the build is completed.
	// See BuildConclusion to know if it succeeded or failed.
	BuildCompleted BuildStatus = "completed"
)

// BuildConclusion defines the result of a build.
//
// Additional conclusions may be added in the future to distinguish
// between different types of failures.
type BuildConclusion string

const (
	// BuildSuccess means the build succeeded.
	BuildSuccess BuildConclusion = "success"

	// BuildFailure means the build failed.
	BuildFailure BuildConclusion = "failure"

	// BuildCanceled means the build was canceled.
	BuildCanceled BuildConclusion = "canceled"
)

type Build struct {
	// ID is the unique id of the build.
	ID string `json:"id"`

	// Status describes the current status of the build.
	Status BuildStatus `json:"status"`

	// Conclusion describes the conclusion of the build.
	// It is set only when the build is completed.
	Conclusion BuildConclusion `json:"conclusion"`

	// CommitHash is the commit hash being built.
	CommitHash string `json:"commit_hash"`

	// QueuedAt defines when the build was queued.
	QueuedAt time.Time `json:"queued_at"`

	// StartedAt defines when the build started.
	// It's nil if the build hasn't started yet.
	StartedAt *time.Time `json:"started_at"`

	// CompletedAt defines when the build completed.
	// It's nil if the build hasn't completed yet.
	CompletedAt *time.Time `json:"completed_at"`
}
