package contracts

import "time"

const SchemaVersionV1 = "1.0.0"

type ActorType string

const (
	ActorTypeUser     ActorType = "user"
	ActorTypeService  ActorType = "service"
	ActorTypeAgent    ActorType = "agent"
	ActorTypeReviewer ActorType = "reviewer"
	ActorTypeExecutor ActorType = "executor"
	ActorTypeTool     ActorType = "tool"
	ActorTypeSystem   ActorType = "system"
)

type SubjectType string

const (
	SubjectTypeIncident           SubjectType = "incident"
	SubjectTypeAlert              SubjectType = "alert"
	SubjectTypeToolCall           SubjectType = "tool_call"
	SubjectTypeToolResult         SubjectType = "tool_result"
	SubjectTypeInvestigation      SubjectType = "investigation"
	SubjectTypeRunbookProposal    SubjectType = "runbook_proposal"
	SubjectTypeReview             SubjectType = "review"
	SubjectTypeTransparencyRecord SubjectType = "transparency_record"
	SubjectTypeExecutionRequest   SubjectType = "execution_request"
	SubjectTypeExecutionResult    SubjectType = "execution_result"
)

type Reference struct {
	Ref     string  `json:"ref"`
	Kind    *string `json:"kind,omitempty"`
	Summary *string `json:"summary,omitempty"`
}

type Actor struct {
	Type        string  `json:"type"`
	ID          string  `json:"id"`
	DisplayName *string `json:"display_name,omitempty"`
}

type Subject struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type Integrity struct {
	Hash               *string `json:"hash,omitempty"`
	PrevEventHash      *string `json:"prev_event_hash,omitempty"`
	SignatureRef       *string `json:"signature_ref,omitempty"`
	TSARef             *string `json:"tsa_ref,omitempty"`
	TransparencyLogRef *string `json:"transparency_log_ref,omitempty"`
}

type BlastRadius struct {
	Level        string  `json:"level"`
	ScopeSummary *string `json:"scope_summary,omitempty"`
	Rationale    *string `json:"rationale,omitempty"`
}

type ToolError struct {
	Code      string         `json:"code"`
	Message   string         `json:"message"`
	Retryable bool           `json:"retryable"`
	Details   map[string]any `json:"details,omitempty"`
}

type TimeRange struct {
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
	DurationMS  int64     `json:"duration_ms"`
}
