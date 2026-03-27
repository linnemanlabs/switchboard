package contracts

import "time"

type AuditEvent struct {
	SchemaVersion string         `json:"schema_version"`
	EventID       string         `json:"event_id"`
	EventType     string         `json:"event_type"`
	EventVersion  int            `json:"event_version"`
	OccurredAt    time.Time      `json:"occurred_at"`
	CorrelationID string         `json:"correlation_id"`
	IncidentID    *string        `json:"incident_id,omitempty"`
	Actor         Actor          `json:"actor"`
	Subject       *Subject       `json:"subject,omitempty"`
	Payload       map[string]any `json:"payload"`
	EvidenceRefs  []Reference    `json:"evidence_refs,omitempty"`
	PolicyRefs    []Reference    `json:"policy_refs,omitempty"`
	Integrity     *Integrity     `json:"integrity,omitempty"`
	Tags          []string       `json:"tags,omitempty"`
}
