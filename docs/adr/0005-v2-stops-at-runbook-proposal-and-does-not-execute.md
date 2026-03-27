# ADR 0005: V2 stops at runbook proposal and does not execute

- Status: accepted
- Date: 2026-03-27

## Context

The long-term system vision includes:
- investigation
- conversational follow-up
- runbook proposal
- review
- transparency/integrity recording
- guarded execution
- future proactive observation and referral

A natural temptation is to extend V2 all the way through execution because the path from proposal to execution is conceptually adjacent.

However, execution introduces materially different requirements:
- approval workflow
- hard safety constraints
- transparency recording
- timestamping
- deterministic executor behavior
- post-execution verification
- failure handling and rollback concerns

V2 is intended to be a strong, publishable milestone that improves investigation structure and auditability without prematurely collapsing into full action-taking autonomy.

## Decision

V2 stops at structured runbook proposal and does not perform infrastructure execution.

V2 includes:
- alert ingestion
- MCP-style tool boundary
- Wazuh as first-class data source
- conversational follow-up
- one structured parameterized runbook proposal path
- audit event emission

V2 does not include:
- autonomous execution
- execution request handling
- transparency log enforcement for execution
- timestamp authority requirements for execution
- guarded runbook execution
- automated rollback behavior

## Rationale

### Clear milestone boundary
Stopping at proposal creates a real, shippable milestone with meaningful value and clear scope.

### Better sequencing
Execution safety deserves its own design and implementation phase rather than being rushed into the same milestone as tool and proposal cleanup.

### Reduced risk
Proposal generation can be exercised and evaluated without exposing the system to execution mistakes.

### Better learning
The proposal phase will teach what review, policy, and execution controls are actually needed before those parts are built.

### Stronger public artifact
A well-scoped V2 is easier to document, explain, and publish than an overextended version trying to do too much.

## Consequences

### Positive
- cleaner V2 scope
- lower delivery risk
- faster path to a stable milestone
- better separation between reasoning and action
- more time to design review and execution properly

### Negative
- V2 cannot directly remediate issues
- proposal-to-action remains manual or deferred until later phases
- some users may expect execution sooner because the runbook path exists conceptually

## Alternatives considered

### Include direct execution in V2
Rejected because it would expand scope significantly and introduce safety requirements that deserve dedicated treatment.

### Include partial execution behind ad hoc safeguards
Rejected because it would blur the milestone boundary and risk producing a weak or inconsistent execution model.

## Notes

This decision is about sequencing, not capability.

Execution remains part of the intended system roadmap.
It is simply not part of V2.

The next phase after V2 is expected to add:
- review decisions
- integrity recording
- execution requests
- deterministic guarded execution
- execution result recording