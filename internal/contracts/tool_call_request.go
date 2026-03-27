package contracts

import "time"

type ToolCallRequest struct {
	SchemaVersion   string         `json:"schema_version"`
	ToolCallID      string         `json:"tool_call_id"`
	IncidentID      *string        `json:"incident_id,omitempty"`
	CorrelationID   string         `json:"correlation_id"`
	ToolName        string         `json:"tool_name"`
	ToolVersion     *string        `json:"tool_version,omitempty"`
	ToolKind        *string        `json:"tool_kind,omitempty"`
	RequestedAt     time.Time      `json:"requested_at"`
	Caller          Actor          `json:"caller"`
	Arguments       map[string]any `json:"arguments"`
	ArgumentSummary *string        `json:"argument_summary,omitempty"`
	TimeoutMS       *int64         `json:"timeout_ms,omitempty"`
	Priority        string         `json:"priority,omitempty"`
	ReadOnly        bool           `json:"read_only"`
	IdempotencyKey  *string        `json:"idempotency_key,omitempty"`
	Metadata        map[string]any `json:"metadata,omitempty"`
}
