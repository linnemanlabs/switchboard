# ADR 0002: Runbook proposal and execution request are separate objects

- Status: accepted
- Date: 2026-03-27

## Context

When Vigil identifies a candidate action, it needs to express both:
- a recommendation based on evidence and reasoning
- a later, narrower request to actually perform the action

A naive design could use one object for both recommendation and execution.

However, this system requires review, approval, integrity recording, and strong execution constraints.
Those concerns benefit from separating reasoning-oriented objects from execution-oriented objects.

## Decision

Runbook proposals and execution requests will be separate objects with separate schemas, lifecycle stages, and audit records.

### Runbook proposal
A proposal represents a recommendation.
It may include:
- reasoning summary
- evidence references
- confidence
- detection hints
- hard constraints
- verification ideas
- estimated blast radius
- required reviews

### Execution request
An execution request represents a request to actually perform a bounded action.
It should include:
- final parameters
- final target scope
- approval references
- integrity references
- preconditions
- verification checks
- expiration
- status for execution lifecycle

## Rationale

### Separation of concerns
A proposal is about reasoning.
An execution request is about controlled action.

### Better auditability
The system should clearly record:
- what was proposed
- what was reviewed
- what was approved
- what was actually requested for execution

### Stronger trust boundaries
Different actors may create, review, sign, and execute these objects.

### Cleaner policy enforcement
Execution policy should act on a narrow object designed for execution, not on a broad reasoning object.

### Easier signing and integrity controls
Each step can be independently recorded, signed, timestamped, and verified.

## Consequences

### Positive
- clearer lifecycle
- better review flow
- easier policy checks
- stronger audit trail
- cleaner schema design

### Negative
- more objects to manage
- more explicit state transitions
- some duplication between proposal and execution request

## Alternatives considered

### Single object for proposal and execution
Rejected because it mixes reasoning, review, and action into one lifecycle artifact.

### Proposal mutates into execution request
Rejected because it weakens audit boundaries and makes it harder to distinguish recommendation from action authorization.

## Notes

This decision is foundational to later reviewer, transparency, and executor components.