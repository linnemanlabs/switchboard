# ADR 0008: Runbooks use detection hints, hard constraints, and verification checks

- Status: accepted
- Date: 2026-03-27

## Context

Runbooks in this system are not simple scripts attached directly to alerts. They are intended to be proposed, reviewed, and eventually executed in a controlled and auditable way.

A runbook recommendation must answer at least three different questions:

- why this runbook is relevant
- what conditions must prevent it from running
- how to confirm whether it succeeded safely

A naive design could mix all of these together as one undifferentiated set of conditions or notes.

However, those concerns serve different purposes and should remain distinct in both design and implementation.

## Decision

Runbooks will be modeled using three distinct categories of guidance:

### Detection hints
Signals that suggest a runbook may be relevant for a given situation.

Detection hints are used to guide proposal and investigation.
They increase confidence that a runbook is a good candidate.
They do not, by themselves, authorize execution.

Examples:
- specific alert classes
- specific log patterns
- particular symptom combinations
- known operational failure signatures

### Hard constraints
Conditions that should block review or execution if violated.

Hard constraints are safety boundaries.
They exist to prevent a runbook from being applied in the wrong context or at the wrong time.

Examples:
- target checked in too recently
- target is still healthy
- scope exceeds allowed size
- required integrity artifacts are missing
- required approvals are missing

### Verification checks
Conditions evaluated after execution to confirm whether the intended effect occurred safely.

Verification checks are used to determine:
- whether the action succeeded
- whether the original problem improved
- whether harmful side effects appeared

Examples:
- target removed from ring
- service health restored
- no unexpected redistribution spike
- no new critical alerts introduced

## Rationale

### Clear separation of purposes
Proposal guidance, execution blocking, and post-execution verification are different operational concerns and should not be collapsed into one list.

### Better auditability
It is useful to record separately:
- why a runbook was suggested
- why it was allowed or blocked
- how success or failure was verified

### Better safety
Hard constraints should remain explicit and visible so they can be enforced deterministically.

### Better review
Reviewers should be able to distinguish:
- recommendation signals
- hard safety boundaries
- outcome validation logic

### Better reuse
This structure should generalize well across many different runbook types.

## Consequences

### Positive
- clearer runbook design
- better separation between proposal and execution concerns
- stronger review and audit clarity
- easier policy enforcement
- easier post-execution evaluation

### Negative
- more explicit structure to maintain per runbook
- requires discipline to decide which guidance belongs in which category
- some signals may initially be ambiguous and require iteration

## Alternatives considered

### Single list of conditions and notes
Rejected because it mixes recommendation logic, blocking logic, and verification logic into one undifferentiated structure.

### Only hard constraints plus free-form notes
Rejected because it weakens proposal quality and post-execution verification clarity.

## Notes

This decision is intended to shape:
- runbook metadata
- proposal generation
- reviewer logic
- execution precondition checks
- post-execution verification design

It does not require every runbook to start with a large or perfect rule set.
The structure should exist from the beginning and improve over time.