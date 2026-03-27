# ADR 0007: Review includes both rule-based and agent-based analysis

- Status: accepted
- Date: 2026-03-27

## Context

This system will eventually support runbook proposals that need evaluation before execution.

A review stage could be implemented in several ways:
- purely rule-based checks
- purely agent-based analysis
- human approval only
- a combination of methods

Rule-based review is strong at enforcing explicit hard constraints and policy boundaries.
Agent-based review is strong at contextual reasoning, blast radius estimation, evidence synthesis, and spotting issues not captured in static rules.

Because the system is intended to operate infrastructure safely, relying exclusively on either approach would create avoidable weaknesses.

## Decision

The review phase will include both:
- rule-based analysis
- agent-based analysis

These reviewer classes are complementary and should remain distinct.

### Rule-based reviewer
The rule-based reviewer is responsible for deterministic, hard-edged checks such as:
- explicit safety thresholds
- forbidden target conditions
- scope constraints
- required artifacts present or missing
- policy violations

### Agent-based reviewer
The agent-based reviewer is responsible for contextual analysis such as:
- blast radius estimation
- reasoning about surrounding evidence
- detecting suspicious context combinations
- identifying when scope should be narrowed
- identifying when more evidence is needed

Human review may also be added later as an additional layer where appropriate.

## Rationale

### Defense in depth
Rule-based and agent-based review catch different classes of problems.

### Deterministic safety plus contextual judgment
Some risks should always be blocked by hard rules.
Others require broader situational reasoning.

### Better practical safety
A purely agentic reviewer may occasionally miss hard boundaries.
A purely rule-based reviewer may miss contextual nuance.

### Better system evolution
The system can accumulate hard rules reactively over time while still benefiting from broader reasoning early.

### Better auditability
It is useful to separately record:
- what hard rules were evaluated
- what contextual concerns were raised
- how each reviewer arrived at its decision

## Consequences

### Positive
- stronger safety posture
- better blast radius evaluation
- clearer separation between hard constraints and contextual judgment
- easier iterative improvement of both reviewer classes
- better audit trail for execution approval decisions

### Negative
- more components and state to manage
- potential disagreement between reviewer classes
- more explicit decision-merging logic required

## Alternatives considered

### Rule-based review only
Rejected because it is too brittle and cannot capture all contextual risks.

### Agent-based review only
Rejected because it provides weaker guarantees for hard safety boundaries.

### Human review only
Rejected as the sole mechanism because it does not scale well and does not encode reusable automated policy.

## Notes

Where reviewers disagree, the system should fail safely.

As a general principle:
- hard rule violations should block execution
- contextual concerns may approve, reject, request more evidence, or request narrower scope depending on severity

This decision does not define the final merge logic, only the requirement that both reviewer classes exist as distinct parts of the review model.