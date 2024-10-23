package webhooks

// RolloutEvent is the interface implemented by all rollout-related events.
type RolloutEvent interface {
	GetRollout() *Rollout
	GetEnv() *Environment
	GetApp() *Application
}

// RolloutCreatedEvent describes the "rollout.created" webhook event.
type RolloutCreatedEvent struct {
	// Rollout describes the current rollout state.
	// It's always non-nil.
	Rollout *Rollout `json:"rollout"`

	// Env is the environment the rollout is targeting.
	// It's always non-nil.
	Env *Environment `json:"env"`

	// App is the application the event is for.
	// It's always non-nil.
	App *Application `json:"app"`
}

func (e *RolloutCreatedEvent) GetRollout() *Rollout { return e.Rollout }
func (e *RolloutCreatedEvent) GetEnv() *Environment { return e.Env }
func (e *RolloutCreatedEvent) GetApp() *Application { return e.App }

// RolloutAwaitingInfraApprovalEvent describes the "rollout.awaiting_infra_approval" webhook event.
type RolloutAwaitingInfraApprovalEvent struct {
	// Rollout describes the current rollout state.
	// It's always non-nil.
	Rollout *Rollout `json:"rollout"`

	// Env is the environment the rollout is targeting.
	// It's always non-nil.
	Env *Environment `json:"env"`

	// App is the application the event is for.
	// It's always non-nil.
	App *Application `json:"app"`
}

func (e *RolloutAwaitingInfraApprovalEvent) GetRollout() *Rollout { return e.Rollout }
func (e *RolloutAwaitingInfraApprovalEvent) GetEnv() *Environment { return e.Env }
func (e *RolloutAwaitingInfraApprovalEvent) GetApp() *Application { return e.App }

// RolloutCompletedEvent describes the "rollout.completed" webhook event.
type RolloutCompletedEvent struct {
	// Rollout describes the current rollout state.
	// It's always non-nil.
	Rollout *Rollout `json:"rollout"`

	// Env is the environment the rollout is targeting.
	// It's always non-nil.
	Env *Environment `json:"env"`

	// App is the application the event is for.
	// It's always non-nil.
	App *Application `json:"app"`
}

func (e *RolloutCompletedEvent) GetRollout() *Rollout { return e.Rollout }
func (e *RolloutCompletedEvent) GetEnv() *Environment { return e.Env }
func (e *RolloutCompletedEvent) GetApp() *Application { return e.App }

func parseRolloutEvent(typ string) (any, error) {
	switch typ {
	case "rollout.created":
		return &RolloutCreatedEvent{}, nil
	case "rollout.awaiting_infra_approval":
		return &RolloutAwaitingInfraApprovalEvent{}, nil
	case "rollout.completed":
		return &RolloutCompletedEvent{}, nil
	default:
		return nil, errUnknownEvent
	}
}
