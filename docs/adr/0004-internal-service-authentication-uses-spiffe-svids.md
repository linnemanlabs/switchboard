# ADR 0004: Internal service authentication uses SPIFFE/SVIDs

- Status: accepted
- Date: 2026-03-27

## Context

This system will contain multiple cooperating internal services, including components such as:
- Vigil
- reviewer
- executor
- transparency/integrity services
- future observer components

These services need strong mutual authentication and clear workload identity.

One option is to use shared static secrets, API keys, or long-lived certificates for service-to-service trust.

Another option is to use workload identity via SPIFFE/SPIRE with short-lived SVIDs.

Because this system is explicitly focused on trust, auditability, and constrained operational authority, internal authentication should use short-lived, machine-verifiable workload identity rather than shared static credentials.

## Decision

Internal service-to-service authentication will use SPIFFE identities issued through SPIRE, with SVID-based mutual authentication.

Internal services should identify themselves with explicit SPIFFE IDs.
Where practical, those identities should also be reflected in audit records, approval records, and execution records.

Authorization decisions remain separate from authentication and must be enforced using policy in addition to verified identity.

## Rationale

### Strong workload identity
SPIFFE/SVIDs provide explicit, machine-verifiable identity for workloads rather than relying on shared secrets.

### Short-lived credentials
Short-lived SVIDs reduce the operational and security risk associated with long-lived static credentials.

### Better trust boundaries
Using explicit workload identity makes it easier to define which service is allowed to call which other service and under what conditions.

### Better auditability
Internal actions can be tied directly to authenticated service identities rather than generic tokens or credentials.

### Good fit for system goals
This system is already designed around provenance, trust, and bounded responsibilities.
SPIFFE/SVIDs align naturally with those goals.

## Consequences

### Positive
- stronger service identity
- reduced secret management burden for internal trust
- better alignment between authn and audit trails
- easier future policy enforcement based on caller identity
- good foundation for least-privilege internal RPC/API design

### Negative
- requires operating SPIRE and related identity plumbing
- increases implementation complexity compared with simple shared tokens
- may require more upfront design around SPIFFE ID naming and authorization rules

## Alternatives considered

### Shared API keys or bearer tokens
Rejected because they weaken identity granularity, increase secret handling burden, and are a poor fit for a trust-heavy internal system.

### Long-lived internal certificates without SPIFFE identity model
Rejected because they provide weaker workload identity semantics and less consistent service naming/authorization patterns.

### Mixed model with ad hoc auth per service
Rejected because it would create uneven trust guarantees and increase operational complexity over time.

## Notes

This decision covers internal service authentication.

It does not eliminate the need for explicit authorization policy.
Authentication answers who the caller is.
Authorization answers what the caller is allowed to do.

SPIFFE ID format may evolve, but the principle is stable:
internal trust should be based on strong, short-lived workload identity.