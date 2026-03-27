# Architecture Decision Records

This directory contains Architecture Decision Records (ADRs) for the system.

ADRs capture important technical decisions that are expected to shape the design, trust model, or operational behavior of the system over time.

They are not general design documents.
They exist to answer questions such as:
- what decision was made
- why it was made
- what alternatives were considered
- what consequences follow from that decision

## Status values

ADRs should include one of the following status values:

- proposed
- accepted
- superseded
- deprecated

## File naming

ADRs are numbered in creation order:

- `0001-...`
- `0002-...`
- `0003-...`

The number should remain stable once assigned.

## When to write an ADR

Write an ADR when a decision:
- affects trust boundaries
- affects system architecture
- changes how responsibilities are split
- changes how data is modeled or audited
- changes safety or execution behavior
- is likely to be revisited later and worth documenting now

## Relationship to other docs

- `docs/architecture/` explains how the system is intended to work
- `docs/adr/` explains why specific important decisions were made
- `schemas/` defines formal machine-readable contracts

## Current ADRs

- `0001-executor-is-not-an-agent.md`
- `0002-runbook-proposal-and-execution-request-are-separate.md`
- `0003-audit-events-use-a-common-envelope-with-event-specific-payloads.md`
- `0004-internal-service-authentication-uses-spiffe-svids.md`
- `0005-v2-stops-at-runbook-proposal-and-does-not-execute.md`
- `0006-observer-components-may-refer-but-not-execute.md`
- `0007-review-includes-both-rule-based-and-agent-based-analysis.md`

## Notes

ADRs are expected to evolve.

If a later decision replaces an earlier one, the older ADR should usually remain in place and be marked as superseded rather than deleted.