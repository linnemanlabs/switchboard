# Tool Contracts

## Purpose

This document defines the contract and behavior expectations for tool calls used by Vigil and later related components.

The goal is to make every tool invocation look structurally similar even when the underlying tools are very different.

This improves:
- consistency
- auditability
- error handling
- evidence tracking
- future reviewer and observer integration

## Contract Structure

A tool interaction has two primary objects:

- `tool-call-request`
- `tool-result`

The request describes what was asked.
The result describes what happened.

## Request Contract

Tool requests use `schemas/tool-call-request.schema.json`.

A request must include:
- unique tool call identity
- correlation identity
- tool name
- request timestamp
- caller identity
- structured arguments

A request may also include:
- incident context
- timeout
- priority
- read-only flag
- human-readable argument summary
- idempotency key
- additional metadata

## Response Contract

Tool responses use `schemas/tool-result.schema.json`.

A response must include:
- the same tool call identity
- correlation identity
- tool name
- status
- timing
- normalized request information
- normalized result information

A response may also include:
- incident context
- tool version or kind
- structured error data
- evidence references
- tags
- metadata

## Behavioral Rules

## Every tool must return a normalized summary
Every successful tool response must include a short summary describing what the tool found or did.

This summary should be easy for:
- operators
- agents
- logs
- audit events

to consume without parsing the full raw result.

## Structured data should be preserved
Whenever possible, tools should return structured result data rather than only unstructured text.

This helps with:
- later reasoning
- review
- verification
- UI presentation
- follow-up investigation

## Errors must be structured
Tool failures should return structured errors with:
- code
- message
- retryable flag
- optional details

Free-form string errors alone are not sufficient.

## Evidence should be referenced explicitly
When a tool returns evidence relevant to an investigation, it should provide `evidence_refs` whenever practical.

Examples:
- log query result references
- metric query references
- Wazuh finding references
- AWS object or instance references
- stored raw result references

## Large raw payloads should be referenced, not embedded blindly
If a result is large, verbose, or expensive to keep inline, the response may:
- set `data_truncated` to true
- store the raw output elsewhere
- include `raw_ref`

## Read-only vs mutating tools
Investigative tools should generally be marked `read_only: true`.

Tools that can change state must be clearly marked `read_only: false`.

This distinction matters for:
- policy
- audit
- future reviewer logic
- execution safety boundaries

## Idempotency
Mutating tool calls should support idempotency where practical through `idempotency_key`.

## Correlation and incident usage

### correlation_id
Used to tie related activity together across a single investigation path.

### incident_id
Used when the tool call belongs to a specific incident context.

Not every tool call requires an incident, but most Vigil investigations should carry one.

## Recommended Flow

1. Build a `tool-call-request`.
2. Emit a `tool.called` audit event.
3. Execute the tool.
4. Build a `tool-result`.
5. Emit either:
   - `tool.completed`
   - `tool.failed`

The `tool_result.tool_call_id` should match the request `tool_call_id`.

## Design Guidance

The tool contract should stay:
- strict at the top level
- flexible in tool-specific arguments and data
- easy to validate
- easy to extend

The contract is intended to normalize the envelope, not eliminate tool-specific richness.

## Initial Expectations for Vigil

For V2, every first-class tool should aim to provide:
- stable naming
- structured arguments
- structured result data
- short result summary
- structured errors
- evidence references where possible

This is enough to make the MCP-style boundary real without overengineering the individual tools too early.