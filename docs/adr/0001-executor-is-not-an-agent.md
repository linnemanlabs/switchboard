# ADR 0001: Executor is not an agent

- Status: accepted
- Date: 2026-03-27

## Context

This system will include components that investigate incidents, reason about possible actions, review proposed changes, and eventually execute approved runbooks.

One design option is to allow the same agentic system that reasons about incidents to also directly execute changes.

Another option is to separate reasoning from execution and make execution a deterministic subsystem with narrow responsibilities.

Because this system is intended to operate infrastructure, execution safety, predictability, and auditability are more important than flexibility at the point of action.

## Decision

The executor will not be an agent.

The executor will be a deterministic, policy-driven subsystem that:
- receives a valid execution request
- validates required approvals and integrity artifacts
- checks preconditions
- performs a bounded runbook action
- records execution results and verification data

The executor will not:
- decide what to do
- broaden its own scope
- reinterpret intent beyond the execution request
- bypass approval or integrity requirements

## Rationale

Separating execution from reasoning improves:

### Safety
A deterministic executor is easier to constrain than an agentic one.

### Auditability
It is easier to show exactly what was executed and why.

### Isolation of duties
Investigation, review, and execution are distinct responsibilities and should remain distinct in code and operation.

### Policy enforcement
Hard rules are easier to apply to a typed execution request than to an open-ended agent workflow.

### Replayability and debugging
Deterministic execution is easier to test, replay, and reason about after the fact.

## Consequences

### Positive
- reduced execution risk
- clearer trust boundaries
- simpler authorization model
- better alignment with transparency and signing requirements
- easier post-incident analysis

### Negative
- requires more explicit request shaping before execution
- less flexible than direct agent action
- may require more translation from proposal to execution request

## Alternatives considered

### Agent executes directly
Rejected because it collapses reasoning and action into one component and makes safe execution harder to guarantee.

### Hybrid executor with optional agentic behavior
Rejected for now because it muddies trust boundaries and increases operational risk early.

## Notes

Agents may still recommend, analyze, and review actions.
Only the executor performs infrastructure changes.