# ADR 0009: Tool boundary is standardized around MCP-style concepts

- Status: accepted
- Date: 2026-03-27

## Context

This system depends heavily on tool use.

Components such as Vigil, future reviewer components, and future observer components need to interact with:
- observability systems
- cloud APIs
- security systems
- runbook metadata
- future operational services

Without a consistent tool boundary, tool integrations tend to become ad hoc:
- inconsistent argument shapes
- inconsistent response formats
- inconsistent error handling
- inconsistent evidence tracking
- inconsistent auditability

One option is to keep tool integrations entirely custom and service-specific.

Another option is to standardize the tool boundary around concepts similar to the Model Context Protocol (MCP), such as:
- named tools
- structured arguments
- structured results
- explicit error handling
- discoverable or describable tool interfaces
- a consistent contract between caller and tool

At this stage, the system benefits from standardizing around those concepts even if the exact external wire protocol is not yet final.

## Decision

The tool boundary will be standardized around MCP-style concepts.

This means tools should follow a consistent model that includes:
- stable tool names
- structured request arguments
- structured response data
- normalized summaries
- structured errors
- evidence references where appropriate
- clear caller and correlation context

The current system schemas and contracts for tool requests and tool results reflect this direction.

This ADR does **not** decide that Switchboard must use MCP as its external transport or protocol in every case.

That protocol-level decision is deferred to a future ADR.

## Rationale

### Standardization scales better
As a solo operator building a growing operational system, standards reduce cognitive load and make systems easier to evolve.

### Better interoperability
A standardized tool model makes it easier for multiple services and components to consume the same tool capabilities over time.

### Better auditability
Consistent request and result envelopes make it much easier to audit tool usage and correlate it with investigations and decisions.

### Better ergonomics
Consistent tool shapes reduce repeated glue code and make behavior more predictable.

### Better future optionality
Adopting MCP-style concepts now preserves the option to align more directly with MCP later without forcing that transport decision immediately.

## Consequences

### Positive
- more predictable tool integrations
- easier future reuse across services
- easier reasoning about tool behavior
- cleaner audit event generation
- smoother path toward future protocol standardization if desired

### Negative
- introduces abstraction overhead compared with fully custom tool code
- may feel heavier for very simple integrations
- may not perfectly fit every future tool pattern, especially long-running or streaming cases
- could encourage premature protocol thinking if applied too rigidly

## Alternatives considered

### Fully custom per-service tool integrations
Rejected because it would create inconsistency and long-term maintenance friction.

### Adopt MCP protocol everywhere immediately
Rejected for now because the current need is contract standardization, not immediate protocol lock-in.

### Standardize later after more tools exist
Rejected because inconsistent early integrations create cleanup work and make later standardization harder.

## Notes

The current decision is about **conceptual and contract standardization**, not mandatory transport.

Future ADRs may decide:
- whether to use MCP directly as an external protocol
- where MCP fits well
- where in-process or service-local interfaces remain more practical

The current goal is to make tool use consistent, predictable, and auditable.