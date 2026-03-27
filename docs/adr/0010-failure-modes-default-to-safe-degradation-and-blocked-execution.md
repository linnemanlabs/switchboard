# ADR 0010: Failure modes default to safe degradation and blocked execution

- Status: accepted
- Date: 2026-03-27

## Context

This system depends on multiple components and trust boundaries, including:
- alert ingestion
- audit event recording
- tool invocation
- LLM-based reasoning
- review systems
- integrity recording
- execution controls
- future observer components

These components will not always be available or healthy.

Examples include:
- LLM provider unavailable
- tool call timeout
- audit append failure
- reviewer unavailable
- transparency log unavailable
- timestamp authority unavailable
- executor precondition failure
- verification checks unavailable or inconclusive

Because this system is intended to operate infrastructure safely, failure handling must be explicit and consistent rather than ad hoc.

## Decision

Failure handling will follow this general rule:

- investigative and observational workflows may degrade in carefully bounded ways
- approval and execution workflows must fail closed

The system should prefer reduced capability over unsafe capability.

## Core policy

### 1. Investigation may degrade
If a non-critical dependency fails during investigation, the system may:
- continue with partial evidence
- return a degraded result
- ask for follow-up
- record the missing dependency explicitly
- avoid overstating confidence

Examples:
- LLM unavailable
- one tool unavailable
- one data source timeout
- partial evidence retrieval

### 2. Execution must fail closed
If a dependency required for review, integrity, approval, or execution is unavailable or incomplete, the system must not execute.

Examples:
- missing review decision
- missing required approval
- missing audit record required by policy
- missing transparency log reference
- missing timestamp proof
- failed precondition checks
- unavailable executor policy dependencies

### 3. Failures must be explicit
Failures should be represented as structured outcomes, not hidden or silently ignored.

The system should record:
- what failed
- when it failed
- whether it was retryable
- what was skipped or blocked because of it
- whether the system degraded or stopped

### 4. Confidence must not be inflated under degradation
If the system is operating with partial evidence or unavailable dependencies, it must reduce confidence accordingly and state the missing context explicitly.

## Failure handling by stage

## Alert ingestion

### If ingestion succeeds but downstream investigation fails
The alert should still be recorded if possible, and investigation should either:
- continue in degraded mode
- or produce an explicit failed/degraded triage result

### If ingestion itself fails
The failure should be surfaced immediately and recorded through the best available fallback path.

## Audit recording

### If ordinary audit event recording fails during investigation
The system may continue investigation only if policy allows degraded operation for that stage, but it must:
- surface that audit recording failed
- mark the result as degraded
- avoid treating the investigation as fully trustworthy

### If required audit or integrity recording fails for approval or execution
The system must block progression to execution.

Execution must not proceed without the required audit and integrity artifacts defined by policy.

## Tool invocation

### If a tool fails during investigation
The system may:
- continue with other evidence
- produce a partial result
- request more evidence
- return a degraded recommendation

The failure should be recorded as a structured tool failure.

### If a tool required for a hard constraint or verification check fails
The relevant review or execution step should block or return inconclusive status according to policy.

## LLM unavailability or failure

### During investigation
If the LLM is unavailable, the system may:
- return a degraded response
- fall back to non-LLM summaries if available
- record that reasoning capability was unavailable

### During review
If agent-based review is required by policy and unavailable, approval is incomplete and execution must not proceed.

### During execution
The executor must not depend on live LLM reasoning to perform a previously approved deterministic action.

Execution should remain possible only if all required approvals and artifacts already exist and executor policy allows it.

## Reviewer failures

### Rule-based reviewer unavailable
If rule-based review is required, execution must block.

### Agent-based reviewer unavailable
If agent-based review is required, execution must block.

### Human review unavailable
If human review is required by policy, execution must block.

## Integrity and transparency failures

### Transparency log unavailable
If transparency recording is required for execution, execution must block.

### Timestamp authority unavailable
If timestamp proof is required for execution, execution must block.

### Signature or integrity reference missing
Execution must block when required integrity artifacts are missing or invalid.

## Preconditions and verification failures

### Preconditions fail
Execution must not begin.

### Verification checks fail after execution
The system should:
- record execution result clearly
- mark resolution status accordingly
- recommend rollback or follow-up if appropriate
- avoid claiming success

### Verification cannot be completed
The result should be marked inconclusive or degraded, not successful.

## Observer failures

Observer components may fail without creating immediate execution risk because they do not execute actions.

If an observer dependency fails, the system should:
- record the failure
- degrade observation capability
- avoid fabricating weak-signal findings

## Rationale

### Safety first
Infrastructure actions should not continue when required trust, review, or integrity components are unavailable.

### Better operational clarity
Explicit degraded and blocked states are easier to reason about than silent failure or partial implicit behavior.

### Better auditability
Structured failure handling creates a more trustworthy record of what the system could and could not do.

### Better architecture
This decision reinforces the broader principle that action should have stricter requirements than investigation.

## Consequences

### Positive
- clearer operational behavior under failure
- safer execution posture
- more predictable handling of missing dependencies
- better trustworthiness and audit clarity

### Negative
- more explicit state handling required
- degraded modes must be carefully implemented
- some workflows will stop rather than “try anyway”

## Alternatives considered

### Best-effort behavior everywhere
Rejected because it risks unsafe action when required controls are unavailable.

### Fail closed for all failures including investigation
Rejected because it would make the system unnecessarily brittle and reduce practical usefulness during partial outages.

### Silent fallback behavior
Rejected because it weakens trust and makes post-incident reasoning harder.

## Notes

A useful shorthand for this ADR is:

- investigate with degradation when safe
- approve with explicit requirements
- execute only with complete required controls

This ADR should influence:
- audit event design
- tool result status handling
- review requirements
- execution policy
- user/operator-facing status reporting