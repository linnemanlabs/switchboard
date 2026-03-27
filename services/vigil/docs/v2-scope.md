# Vigil V2 Scope

## Purpose

Vigil V2 turns the current live triage system into a more structured and auditable investigation component.

The focus of V2 is not broad autonomy.
The focus is improving investigation quality, tool consistency, auditability, and runbook proposal structure.

## In Scope

V2 includes the following capabilities.

### Alert ingestion
Vigil receives an alert and creates or continues an incident context.

### MCP-style tool boundary
Tools are exposed through a cleaner and more uniform interface so Vigil interacts with them consistently.

### Wazuh as a first-class data source
Wazuh is available alongside observability and AWS context as part of normal investigation flow.

### Conversational follow-up
After an initial triage, the operator can continue asking questions and requesting deeper investigation.

### One parameterized runbook proposal path
Vigil can propose one bounded runbook with parameters and supporting evidence.

### Audit event emission
Important actions and decisions emit structured audit events.

## Out of Scope

The following are explicitly out of scope for V2.

### Autonomous execution
Vigil must not directly execute infrastructure changes in V2.

### Continuous weak-signal observer
A continuous observer for full log-flow anomaly detection is not part of V2.

### Offline local-model full-stream analysis
Full-time local-model analysis across all logs is a future phase.

### Large runbook catalog
V2 only needs one strong runbook path, not a large general library.

### Generalized multi-agent orchestration
V2 is not a generic planner/executor platform.

## Responsibilities

Vigil is responsible for:
- receiving alerts
- managing incident context
- calling tools
- gathering evidence
- summarizing likely explanations
- supporting follow-up questions
- proposing bounded next actions
- emitting audit events

## Non-Responsibilities

Vigil is not responsible for:
- directly making changes
- bypassing approval paths
- acting without evidence
- replacing deterministic execution controls
- operating as a catch-all orchestration layer for the full future system

## Expected Inputs

Vigil consumes:
- alert payloads
- operator follow-up questions
- tool results
- incident history
- runbook metadata
- future reviewer feedback

## Expected Outputs

Vigil produces:
- investigation summaries
- tool call records
- evidence references
- runbook proposals
- audit events
- conversational responses

## First Runbook Target

The initial runbook candidate is removal of stale distributors from Loki or Mimir rings.

This is a good first runbook because it is:
- bounded
- low blast radius
- operationally useful
- observable before and after
- suitable for learning how proposals, constraints, and review should work

## Runbook Guidance Model

V2 should distinguish between:
- hints that make a runbook relevant
- hard constraints that should block execution or review
- verification checks that confirm success afterward

This distinction should remain visible in both code and schema design.

## Success Criteria

V2 is successful when:
- Vigil can investigate an alert using a normalized tool interface
- Wazuh is a normal part of that investigation flow
- the operator can continue the incident conversationally
- Vigil can produce one structured runbook proposal
- the important steps are emitted as structured audit events

## Notes

V2 should be shipped as a usable milestone, documented publicly, and treated as a stable foundation for later reviewer, transparency, and execution components.