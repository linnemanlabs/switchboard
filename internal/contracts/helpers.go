package contracts

func NewAuditEvent(eventType, eventID, correlationID string, actor Actor) AuditEvent {
	return AuditEvent{
		SchemaVersion: SchemaVersionV1,
		EventID:       eventID,
		EventType:     eventType,
		EventVersion:  1,
		CorrelationID: correlationID,
		Actor:         actor,
		Payload:       map[string]any{},
		EvidenceRefs:  []Reference{},
		PolicyRefs:    []Reference{},
		Tags:          []string{},
	}
}

func NewToolCallRequest(toolCallID, correlationID, toolName string, caller Actor) ToolCallRequest {
	return ToolCallRequest{
		SchemaVersion: SchemaVersionV1,
		ToolCallID:    toolCallID,
		CorrelationID: correlationID,
		ToolName:      toolName,
		Caller:        caller,
		Arguments:     map[string]any{},
		Priority:      "normal",
		ReadOnly:      true,
		Metadata:      map[string]any{},
	}
}

func NewToolResult(toolCallID, correlationID, toolName, status string, caller Actor) ToolResult {
	return ToolResult{
		SchemaVersion: SchemaVersionV1,
		ToolCallID:    toolCallID,
		CorrelationID: correlationID,
		ToolName:      toolName,
		Status:        status,
		Request: ToolRequest{
			Caller:    caller,
			Arguments: map[string]any{},
		},
		Result: ToolResponseResult{
			Data: map[string]any{},
		},
		EvidenceRefs: []Reference{},
		Tags:         []string{},
		Metadata:     map[string]any{},
	}
}

func NewRunbookProposal(proposalID, incidentID, correlationID, runbookID, runbookVersion, status string, proposedBy Actor) RunbookProposal {
	return RunbookProposal{
		SchemaVersion:      SchemaVersionV1,
		ProposalID:         proposalID,
		IncidentID:         incidentID,
		CorrelationID:      correlationID,
		RunbookID:          runbookID,
		RunbookVersion:     runbookVersion,
		Status:             status,
		ProposedBy:         proposedBy,
		Parameters:         map[string]any{},
		TargetSelector:     map[string]any{},
		EvidenceRefs:       []Reference{},
		RequiredReviews:    []string{},
		DetectionHints:     []string{},
		HardConstraints:    []string{},
		VerificationChecks: []string{},
		BlockingRisks:      []string{},
		Metadata:           map[string]any{},
	}
}
