# System Overview

## Purpose

This system is being built to act as a trustworthy machine collaborator for operating infrastructure.

It is not intended to be a generic multi-agent demo platform. It is a practical operational system for investigating incidents, gathering evidence, proposing bounded actions, enforcing auditability, and later executing approved runbooks under strict controls.

The system is designed to improve the safety, speed, and quality of infrastructure operations by combining:
- real observability data
- security telemetry
- operational memory
- explicit trust boundaries
- auditable decision and execution flows

## Current State

Today, Vigil is already live in a limited form and performs alert triage.

V2 formalizes and extends that into a more structured system with:
- a cleaned-up MCP-style tool boundary
- Wazuh as a first-class data source
- conversational follow-up on investigations
- one parameterized runbook proposal path
- a consistent audit event model

## Core Principles

### Provenance for every important action
Every meaningful step should be attributable and auditable:
- who or what initiated it
- what evidence supported it
- what decision was made
- what policy or guardrails applied
- what happened next

### Separation of reasoning from execution
Reasoning may be broad and probabilistic.
Execution must be narrow, deterministic, constrained, and easy to audit.

### Safe degradation and fail-closed execution
When dependencies fail, the system should degrade investigative capability carefully where safe, but block approval and execution when required controls are missing.

### Separation of duties
Different responsibilities should be owned by different components:
- investigate
- review
- execute
- observe
- record

### Referrals before autonomy
Early-stage agents should refer issues upstream rather than directly take action.
Only the execution subsystem should perform changes.

### Strong service identity
Internal service-to-service trust is based on SPIFFE/SPIRE identity using SVIDs rather than shared static secrets.

## Failure Handling

Switchboard is designed to prefer reduced capability over unsafe capability.

As a general rule:
- investigation and observation may degrade in bounded ways when non-critical dependencies fail
- approval and execution must fail closed when required controls, reviews, or integrity artifacts are unavailable

This means the system may continue operating with partial capability during investigative workflows, but it must not proceed with controlled action when required safety or trust requirements are not satisfied.

The full failure-handling policy is defined in `docs/adr/0010-failure-modes-default-to-safe-degradation-and-blocked-execution.md`.

## Planned Components

## Vigil
Vigil is the investigation and triage component.

Responsibilities:
- receive alerts
- gather evidence from tools
- summarize likely causes
- support conversational follow-up
- propose candidate runbooks
- emit audit events

Vigil does not directly execute changes.

## Tool Boundary
Tools provide access to operational data and actions through a consistent interface.

Initial tool/data sources:
- observability stack
- AWS context
- Wazuh

The tool boundary should normalize:
- requests
- responses
- errors
- timing
- evidence references

## Reviewer
A reviewer component will evaluate proposed actions before execution.

Planned reviewer types:
- agent-based reviewer for contextual blast radius analysis
- deterministic rule-based reviewer for hard safety checks

Reviewer outputs should be auditable and separable from both proposal and execution.

## Executor
The executor performs approved runbooks.

The executor is intentionally not an agent.
It should be deterministic, typed, heavily constrained, and policy-driven.

The executor should refuse to run without required approval and integrity evidence.

## Transparency / Integrity Layer
Before execution, requests should be recorded in a transparency log and timestamped.

The executor should require integrity artifacts such as:
- transparency log reference
- timestamp authority proof
- signatures or equivalent integrity references

## Observer
A future observer component will continuously analyze log and telemetry flows for weak signals and behavioral anomalies.

The observer does not execute actions.
It produces referrals to Vigil.

## Memory Model

The system has at least two distinct forms of memory.

### Operational memory
Operational memory records concrete system activity and decisions, such as:
- alerts
- triages
- tool calls
- findings
- runbook proposals
- reviews
- execution requests
- execution results
- follow-up notes

This is expected to be primarily relational and event-oriented.

### Statistical / behavioral memory
Behavioral memory captures patterns, baselines, rare events, drift, unusual correlations, and weak signals over time.

This may require separate storage, processing, and retrieval mechanisms from operational memory.

These two memory classes are intentionally separated because they serve different purposes.

## V2 Scope Summary

V2 includes:
- alert ingestion
- MCP-style tool boundary
- Wazuh as first-class data source
- conversational follow-up
- one parameterized runbook proposal path
- audit event emission

V2 does not include:
- autonomous execution
- continuous observer agent
- offline full-log local-model analysis
- generalized multi-agent orchestration platform
- broad runbook library

## High-Level Flow

1. An alert is received.
2. Vigil investigates using available tools.
3. Tool interactions and findings are recorded as audit events.
4. The operator can continue the investigation conversationally.
5. Vigil may propose a parameterized runbook.
6. The proposal is recorded as an auditable object.
7. Later versions add review, transparency logging, and guarded execution.

## Design Intent

The system may grow into a large and complex operational platform over time.
That is acceptable and expected.

The goal is not minimalism for its own sake.
The goal is a coherent system with clear trust boundaries, explicit responsibilities, and durable operational value.