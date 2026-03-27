package contracts

import "time"

type ToolRequest struct {
	Caller          Actor          `json:"caller"`
	Arguments       map[string]any `json:"arguments"`
	ArgumentSummary *string        `json:"argument_summary,omitempty"`
}

type ToolResponseResult struct {
	Summary       string         `json:"summary"`
	Data          map[string]any `json:"data"`
	DataTruncated bool           `json:"data_truncated,omitempty"`
	RawRef        *string        `json:"raw_ref,omitempty"`
}

type ToolResult struct {
	SchemaVersion string             `json:"schema_version"`
	ToolCallID    string             `json:"tool_call_id"`
	IncidentID    *string            `json:"incident_id,omitempty"`
	CorrelationID string             `json:"correlation_id"`
	ToolName      string             `json:"tool_name"`
	ToolVersion   *string            `json:"tool_version,omitempty"`
	ToolKind      *string            `json:"tool_kind,omitempty"`
	Status        string             `json:"status"`
	StartedAt     time.Time          `json:"started_at"`
	CompletedAt   time.Time          `json:"completed_at"`
	DurationMS    int64              `json:"duration_ms"`
	Request       ToolRequest        `json:"request"`
	Result        ToolResponseResult `json:"result"`
	Error         *ToolError         `json:"error,omitempty"`
	EvidenceRefs  []Reference        `json:"evidence_refs,omitempty"`
	Tags          []string           `json:"tags,omitempty"`
	Metadata      map[string]any     `json:"metadata,omitempty"`
}
