# Runbook Lifecycle

## Purpose

This document defines the lifecycle of a runbook as it moves from idea to possible execution.

The lifecycle is intentionally split into distinct stages so reasoning, review, integrity, and execution remain separable.

## Key Objects

The lifecycle involves several different objects.

### Incident
The operational context containing alerts, findings, evidence, and conversation.

### Runbook Proposal
A structured recommendation that a specific runbook should be considered, with parameters and evidence.

### Review Decision
A reviewer response evaluating the proposal.

### Transparency Record
A record proving a request was submitted to the transparency layer and timestamped.

### Execution Request
A structured request to actually perform the runbook.

### Execution Result
The result of a runbook attempt, including verification data and failures if any.

## Lifecycle Stages

## 1. Detection
An alert or referral creates or updates an incident.

Sources may include:
- alertmanager
- future observer referrals
- operator request
- security event streams

## 2. Investigation
Vigil gathers evidence using tools and produces a working understanding of the problem.

Outputs may include:
- likely explanation
- additional questions
- relevant evidence references
- candidate next actions

## 3. Proposal
If appropriate, Vigil produces a runbook proposal.

A proposal should include:
- runbook identity and version
- parameters
- target selector
- reasoning summary
- evidence references
- constraints
- blast radius estimate
- required reviews
- rollback summary

At this point, nothing has been executed.

## 4. Review
One or more reviewer systems evaluate the proposal.

Planned reviewer classes:
- rule-based reviewer
- agent-based reviewer
- optional human approval

A review may:
- approve
- reject
- request more evidence
- request narrower scope
- request parameter changes

## 5. Integrity Recording
Before execution, the execution request is recorded in the transparency/integrity layer.

This stage is expected to provide artifacts such as:
- transparency log reference
- timestamp authority proof
- signature material or equivalent references

## 6. Execution Request
A separate execution request is created from an approved proposal.

This object should be narrower than the proposal and contain only what the executor needs.

## 7. Execution
The executor performs the allowed runbook action using deterministic logic and strong policy checks.

The executor should refuse execution if required approvals or integrity artifacts are missing.

## 8. Verification
Post-execution checks confirm whether the intended effect occurred and whether harmful side effects appeared.

## 9. Closeout
The incident is updated with:
- what was proposed
- what was approved
- what was executed
- what changed
- whether the issue resolved
- whether rollback or follow-up is required

## V2 Boundary

V2 stops at the proposal stage.

V2 includes:
- detection
- investigation
- conversational follow-up
- proposal
- audit event emission

Review, integrity recording, and execution are planned next stages.

## Guidance Model

Runbooks should be designed using three distinct categories:

- detection hints
- hard constraints
- verification checks

These categories are defined as a core design decision in `docs/adr/0008-runbooks-use-detection-hints-hard-constraints-and-verification-checks.md`.

This model should remain consistent across future runbooks, reviewer logic, execution preconditions, and post-execution verification.

### Detection hints
Signals that suggest a runbook may be relevant.

### Hard constraints
Conditions that should block review or execution.

### Verification checks
Conditions that confirm whether the action succeeded safely.

This model should remain consistent across future runbooks.