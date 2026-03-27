# Architecture Documentation

This directory contains system-level architecture documents for the infrastructure operations assistant platform.

These documents describe the intended structure, trust model, data flow, and lifecycle of the system.

The goal of this documentation is to make it easier to understand:
- what the system is
- what each component is responsible for
- how decisions and actions flow through the system
- how trust, audit, review, and execution boundaries are enforced

## System purpose

This system is being built as a trustworthy machine collaborator for operating infrastructure.

It is not intended to be a generic multi-agent demo platform.
It is a practical operational system for:
- investigating incidents
- gathering evidence
- proposing bounded actions
- enforcing provenance and auditability
- reviewing potentially risky changes
- eventually executing approved runbooks safely
- later observing weak signals and referring suspicious patterns upstream

## Document guide

### `system-overview.md`
High-level description of the system, its planned components, principles, memory model, and V2 scope.

### `audit-model.md`
Defines the audit philosophy and the shared event-envelope model used across the system.

### `runbook-lifecycle.md`
Describes the lifecycle from detection through investigation, proposal, review, integrity recording, execution, verification, and closeout.

### `identity-and-auth.md`
Describes the intended internal trust model based on SPIFFE/SPIRE workload identity and SVID-based service authentication.

## Relationship to other directories

### `docs/adr/`
Contains Architecture Decision Records documenting why key technical decisions were made.

### `schemas/`
Contains formal machine-readable schemas for core system objects, including:
- audit events
- tool call requests
- tool results
- runbook proposals
- review decisions
- execution requests
- execution results

### `services/`
Contains component-specific documentation and code.
For example, `services/vigil/docs/` holds Vigil-specific design docs such as:
- V2 scope
- tool contracts

## Core model

The system is intentionally split into distinct responsibilities.

### Investigation
Handled by Vigil.
Responsible for alert triage, evidence gathering, conversational follow-up, and runbook proposal.

### Review
Handled by reviewer systems.
Includes both deterministic rule-based review and contextual agent-based review.

### Execution
Handled by a deterministic executor.
Execution is intentionally separate from agent reasoning.

### Observation
Handled by future observer components.
Observers may refer suspicious patterns upstream but do not execute actions.

### Audit and integrity
Handled through structured audit events and later transparency/integrity mechanisms.

## V2 boundary

The current major milestone is V2.

V2 includes:
- alert ingestion
- cleaned-up MCP-style tool boundary
- Wazuh as a first-class data source
- conversational follow-up
- one structured parameterized runbook proposal path
- audit event emission

V2 stops at proposal and does not execute changes.

## Recommended reading order

For someone new to the system, a good reading order is:

1. `system-overview.md`
2. `runbook-lifecycle.md`
3. `audit-model.md`
4. `identity-and-auth.md`
5. `../adr/README.md`
6. relevant schemas in `../../schemas/`

## Notes

These architecture docs describe intended design direction.
Some parts of the system are already live, while others are planned.
As implementation evolves, the docs should be updated to reflect stable decisions and changing scope.