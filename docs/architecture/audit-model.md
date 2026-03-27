# Audit Model

## Purpose

The audit model provides a durable, append-only record of important system activity.

It exists to answer questions such as:
- what happened
- when it happened
- who or what did it
- what evidence existed at the time
- what decision was made
- what policy or constraints applied
- what happened next

## Design Goals

The audit model should be:
- append-only
- structured
- machine-readable
- easy to correlate across services
- easy to sign, hash, and timestamp later
- usable both for debugging and for trust verification

## Common Event Envelope

All audit events should share one common outer structure.

The common envelope allows every service to emit auditable events in a uniform shape while still using event-specific payloads.

The common envelope includes:
- event identity
- event type
- event version
- timestamp
- correlation identifiers
- actor identity
- subject identity
- payload
- evidence references
- policy references
- integrity metadata

## Event Types

Examples of event types include:
- incident.created
- alert.received
- tool.called
- tool.completed
- tool.failed
- investigation.updated
- runbook.proposed
- review.requested
- review.completed
- transparency.recorded
- execution.requested
- execution.started
- execution.completed
- execution.failed

Event types should describe facts that occurred, not vague intentions.

## Correlation

At minimum, the system should support:
- event_id
- correlation_id
- incident_id

### event_id
Unique identifier for this specific event.

### correlation_id
Identifier used to tie together related activity across components and time.

### incident_id
Identifier for the operational incident or investigation context.

## Actor Model

Each event should identify the actor that caused it.

Examples:
- user
- service
- agent
- reviewer
- executor
- tool
- system

For internal services, actor identity should map cleanly to SPIFFE identity where possible.

## Subject Model

The subject identifies what the event was about.

Examples:
- incident
- alert
- tool call
- runbook proposal
- review
- execution request
- execution result

## Evidence References

Events should not always embed all raw evidence inline.

Instead, events can carry references to evidence such as:
- logs
- metrics
- traces
- Wazuh findings
- AWS query results
- prior events
- stored investigation artifacts

This keeps the audit model flexible while preserving traceability.

## Policy References

When policy affects a decision, the event should reference the policy source that applied.

Examples:
- rule-based reviewer policy
- execution guardrails
- runbook constraints
- trust requirements

## Integrity Metadata

The event envelope reserves space for later integrity controls such as:
- event hash
- previous event hash
- signature reference
- timestamp authority reference
- transparency log reference

Not every early event must use every integrity feature, but the envelope should make those additions natural.

## Proposal vs Execution

A proposal is not an execution request.

These are intentionally separate objects and should produce separate events.

This separation improves:
- auditability
- reviewability
- role isolation
- signature boundaries
- reasoning clarity

## Storage Model

The audit event schema defines the wire and event contract.

Storage may include relational tables, event logs, or other systems optimized for query and retention.

Storage design should not be allowed to blur or redefine the event contract.

## Rule of Thumb

If an action, decision, recommendation, or state transition would matter during debugging, review, or trust verification, it should probably be represented as a structured audit event.