# Identity and Authentication

## Purpose

This document describes the intended trust model for internal service identity and authentication.

The goal is to avoid shared static secrets and instead rely on strong workload identity for service-to-service trust.

## Identity Model

Internal services identify themselves using SPIFFE identities issued through SPIRE.

Examples of service identities may look like:
- spiffe://linnemanlabs.com/switchboard/vigil
- spiffe://linnemanlabs.com/switchboard/reviewer
- spiffe://linnemanlabs.com/switchboard/executor
- spiffe://linnemanlabs.com/switchboard/transparency-log

Exact path format may evolve, but the principle is stable:
service identity should be explicit, machine-verifiable, and short-lived.

## Authentication Model

Internal service-to-service authentication uses SVID-based mutual authentication.

This applies to calls such as:
- Vigil to reviewer
- Vigil to transparency service
- reviewer to executor
- operator-facing services to internal APIs where appropriate

## Authorization Model

Authentication proves who a caller is.
Authorization decides what that caller is allowed to do.

Authorization decisions should be based on:
- verified caller identity
- request type
- policy
- target scope
- required approvals or integrity artifacts

## Service Expectations

### Vigil
May investigate, call tools, emit events, and produce runbook proposals.

Vigil must not directly bypass execution controls.

### Reviewer
May evaluate proposals and emit review decisions.

Reviewer does not perform execution.

### Executor
May perform a strictly bounded runbook action only when policy, approvals, and integrity requirements are satisfied.

### Observer
May produce referrals and findings.

Observer must not directly trigger execution.

## Secret Minimization

The preferred design is to minimize long-lived shared secrets for internal trust wherever possible.

Short-lived workload identity should be the default.

## Trust Boundary Principle

Any service that can influence operational decisions should have:
- explicit identity
- auditable requests
- clear authorization rules
- bounded responsibilities