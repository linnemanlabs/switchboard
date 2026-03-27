# ADR 0003: Audit events use a common envelope with event-specific payloads

- Status: accepted
- Date: 2026-03-27

## Context

This system will generate many different event types, including:
- alert receipt
- tool calls
- tool results
- investigation updates
- runbook proposals
- review decisions
- transparency records
- execution requests
- execution results

One design option is for each event type to define a fully separate structure.

Another option is to use a common outer event envelope and allow event-specific payloads inside it.

Because this system needs consistent auditability, correlation, integrity handling, and future signing/timestamping, event uniformity is important.

## Decision

All audit events will use a shared common envelope with event-specific payloads.

The common envelope includes fields such as:
- schema version
- event id
- event type
- event version
- timestamp
- correlation id
- incident id
- actor
- subject
- payload
- evidence refs
- policy refs
- integrity metadata

The `payload` field is event-specific.

## Rationale

### Consistency
Every event can be processed, indexed, and validated in a predictable way.

### Easier correlation
Shared identifiers and actor/subject structure make it easier to connect activity across services.

### Integrity support
A common envelope makes it straightforward to add:
- hashes
- chain references
- signatures
- timestamp references
- transparency log references

### Flexibility
Different event types still retain their own useful detail inside the payload.

### Lower operational complexity
Consumers of audit data can rely on one top-level model rather than many unrelated shapes.

## Consequences

### Positive
- simpler event ingestion
- simpler event storage mapping
- easier validation
- easier future integrity controls
- better consistency across components

### Negative
- requires discipline to keep payloads well-structured
- some event-specific schemas may still be needed for strong validation
- envelope design becomes an important long-term contract

## Alternatives considered

### Separate schema per event type with no shared envelope
Rejected because it would increase event handling complexity and weaken consistency.

### Fully generic untyped event payloads
Rejected because it would make audit data harder to validate and reason about.

## Notes

The shared envelope standardizes the audit layer without forcing every event payload to look identical.
This balances consistency with flexibility.