# Switchboard

**Trustworthy AI-assisted infrastructure operations platform for investigation, review, guarded execution, and later proactive observation.**

Switchboard is a machine collaborator for operating infrastructure. It is being built to investigate incidents, gather evidence from operational and security systems, support conversational follow-up, propose bounded runbooks, and later execute approved actions under strict review and audit controls.

Over time, Switchboard is also intended to grow beyond reactive incident response into proactive observation: continuously analyzing infrastructure signals, identifying weak signals and anomalies, and referring suspicious patterns upstream for investigation.

This is not a generic multi-agent demo platform. It is a practical system for real infrastructure operations with explicit trust boundaries, durable audit trails, and clear separation between reasoning, review, and execution.

## Goals

Switchboard is designed to:

- investigate alerts and operational issues using real infrastructure data
- gather evidence from observability, security, and cloud systems
- support conversational follow-up during incidents
- propose bounded, parameterized runbooks
- enforce provenance and auditability for important decisions
- separate proposal, review, integrity recording, and execution
- eventually support guarded remediation
- eventually support proactive weak-signal observation and anomaly referral

## Design principles

- **Provenance matters.** Important actions and decisions should be attributable, explainable, and auditable.
- **Reasoning and execution are separate.** Agents may investigate and recommend; execution remains narrow, deterministic, and policy-driven.
- **Separation of duties matters.** Investigation, review, execution, observation, and integrity recording are distinct responsibilities.
- **Referrals before autonomy.** Observer-style components may refer issues upstream, but should not directly execute changes.
- **Strong internal identity.** Internal service trust is based on SPIFFE/SPIRE workload identity with SVIDs.

## Current direction

The current major milestone is **V2**, which focuses on turning the existing alert-triage capability into a more structured and auditable investigation system.

V2 includes:

- alert ingestion
- a cleaned-up MCP-style tool boundary
- Wazuh as a first-class data source
- conversational follow-up
- one structured parameterized runbook proposal path
- structured audit event emission

V2 stops at proposal and does not execute infrastructure changes.

## Planned components

### Vigil

Investigation and triage component.

Responsibilities include:

- receiving alerts
- gathering evidence
- calling tools
- summarizing likely explanations
- supporting follow-up questions
- proposing candidate runbooks
- emitting audit events

### Operator

Conversational entry point for interacting with the system across incidents, evidence, and actions.

### Reviewer

Review layer for proposed actions.

Planned reviewer classes include:

- deterministic rule-based review
- contextual agent-based review
- optional human review later

### Executor

Deterministic execution subsystem for approved runbooks.

The executor is intentionally **not** an agent.

### Observer

Future component for continuous weak-signal and anomaly detection.

Observers may generate referrals and findings, but do not directly execute actions.

## Repository layout

    cmd/
      vigil/
      operator/
      reviewer/
      executor/
      observer/

    internal/
      contracts/
      audit/
      identity/
      incidents/
      tools/
      runbooks/
      platform/

    services/
      vigil/
        internal/
        docs/
      operator/
        internal/
        docs/
      reviewer/
        internal/
        docs/
      executor/
        internal/
        docs/
      observer/
        internal/
        docs/

    docs/
      architecture/
      adr/

    schemas/

## Documentation

### Architecture docs

System-level design docs live in `docs/architecture/`.

Start with:

1. `docs/architecture/system-overview.md`
2. `docs/architecture/runbook-lifecycle.md`
3. `docs/architecture/audit-model.md`
4. `docs/architecture/identity-and-auth.md`

### ADRs

Architecture Decision Records live in `docs/adr/`.

These capture important decisions such as:

- why the executor is not an agent
- why proposal and execution are separate objects
- why audit events use a shared envelope
- why internal auth uses SPIFFE/SVIDs

### Schemas

Formal contracts live in `schemas/`.

Current core schemas include:

- audit event
- tool call request
- tool result
- runbook proposal
- review decision
- execution request
- execution result

## Status

Early architecture and contract phase.

The system vision is larger than the currently implemented surface area. Some components are already live in limited form, while others are being defined and built in stages.

The immediate goal is to establish strong foundations:

- clear contracts
- clear trust boundaries
- clear service responsibilities
- a durable audit model
- one real end-to-end investigation and proposal path

## Non-goals

Switchboard is not intended to be:

- a generic agent framework
- a demo-oriented “AI ops” wrapper with weak controls
- a system where reasoning directly implies execution
- a platform optimized for hype over auditability and safety

## Why this exists

Operating infrastructure well requires more than reacting to threshold breaches. Strong operators notice patterns, gather context, weigh risk, and act carefully. Switchboard is an attempt to build a better system around those realities using the new capabilities AI makes possible, without sacrificing trust, clarity, or control.

## License

MIT. Copy it, steal it, modify it, learn from it, share your improvements with me. Or don't. It's code, do what you want with it.