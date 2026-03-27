package contracts

import "time"

type RunbookProposal struct {
	SchemaVersion        string         `json:"schema_version"`
	ProposalID           string         `json:"proposal_id"`
	IncidentID           string         `json:"incident_id"`
	CorrelationID        string         `json:"correlation_id"`
	RunbookID            string         `json:"runbook_id"`
	RunbookVersion       string         `json:"runbook_version"`
	Status               string         `json:"status"`
	ProposedBy           Actor          `json:"proposed_by"`
	ProposedAt           time.Time      `json:"proposed_at"`
	Parameters           map[string]any `json:"parameters"`
	TargetSelector       map[string]any `json:"target_selector"`
	ReasoningSummary     string         `json:"reasoning_summary"`
	Confidence           *float64       `json:"confidence,omitempty"`
	EvidenceRefs         []Reference    `json:"evidence_refs"`
	DetectionHints       []string       `json:"detection_hints,omitempty"`
	HardConstraints      []string       `json:"hard_constraints,omitempty"`
	VerificationChecks   []string       `json:"verification_checks,omitempty"`
	EstimatedBlastRadius *BlastRadius   `json:"estimated_blast_radius,omitempty"`
	BlockingRisks        []string       `json:"blocking_risks,omitempty"`
	RollbackSummary      *string        `json:"rollback_summary,omitempty"`
	RequiredReviews      []string       `json:"required_reviews"`
	Metadata             map[string]any `json:"metadata,omitempty"`
}
