# Vigil Conversation Model

## Purpose

This document defines how conversational follow-up works in Vigil.

The goal is to support real incident investigation over multiple turns without treating the entire transcript as the primary context source. Each turn should be built from a curated investigation packet assembled from durable operational state, conversational state, and the most relevant recent artifacts.

This model is designed to:
- preserve continuity across follow-up questions
- keep context windows efficient and focused
- avoid replaying full transcripts by default
- keep operational truth separate from conversational phrasing
- support later evolution into proposal, review, and proactive modes

## Core Principle

Incident state is the canonical operational record.

Conversation state is the dialogue layer attached to that incident.

Each LLM call should receive a layered investigation packet, not the full incident transcript.

## State Model

## Incident State

Incident state is the durable operational truth for an investigation.

It contains structured records such as:
- alert or referral metadata
- audit events
- tool call requests
- tool results
- findings
- hypotheses
- proposal objects
- later review and execution objects
- evidence references
- structured investigation summaries
- degradation and failure notes

Incident state should remain useful even if the conversation transcript is removed.

## Conversation State

Conversation state is the interactive dialogue layer associated with an incident.

It contains:
- operator messages
- assistant responses
- turn summaries
- current conversational focus
- active goals
- open questions
- retrieval notes for prior conversational context

Conversation state is not the canonical operational truth.
It exists to maintain useful interaction continuity.

## Summary Layers

The conversation model uses multiple summary layers rather than one undifferentiated transcript summary.

### Incident Summary

The incident summary is the rolling structured summary of the operational state.

It should capture:
- what is happening
- what evidence has been gathered
- leading hypotheses
- what has been ruled out
- important tool findings
- relevant degradation or missing dependencies
- proposal status if any
- next likely operational questions

### Conversation Summary

The conversation summary is the rolling summary of the dialogue flow.

It should capture:
- what the operator has asked
- what the assistant has already explained
- what the current conversational focus is
- what level of detail the operator is asking for
- what thread of the investigation is active right now

### Open Questions

Open questions represent unresolved issues that matter to the investigation or the conversation.

Examples:
- Is the stale ring member still actively checking in?
- Is the issue limited to Mimir, or also affecting Loki?
- Do we have enough evidence to propose a runbook safely?

### Active Goals

Active goals represent the immediate purpose of the current investigation thread.

Examples:
- determine likely cause of alert
- confirm whether issue is isolated or systemic
- gather enough evidence to support a bounded runbook proposal
- verify whether a proposed action is still relevant

### Current Conversational Focus

Current conversational focus is a short summary of what this turn is really about.

This should help the model stay aligned with the operator’s intent without re-reading large prior transcripts.

Examples:
- determine whether stale ring entries are genuinely stale
- compare Mimir symptoms to recent Loki behavior
- decide whether to move from investigation to proposal

## Investigation Packet

Each LLM turn should be built from an investigation packet.

The investigation packet is the curated context sent to the model for a given turn.

It should contain only the information needed for the current question or action.

## Investigation Packet Layers

### Tier 0: Current Turn Inputs

Always include in full:
- current operator message
- current active alert or referral object, if one exists
- newly returned tool results for this turn, if any
- any directly relevant proposal or runbook object being actively discussed

This is the highest-priority context.

### Tier 1: Recent Critical Artifacts

Usually include a limited number of recent important artifacts in fuller structured form, such as:
- recent tool results
- latest proposal object
- latest review-related object later
- recent important findings

These should be represented as normalized structured objects, not raw unbounded transcripts.

### Tier 2: Rolling Summaries

Always include:
- current incident summary
- current conversation summary
- open questions
- active goals
- current conversational focus

This is the main continuity layer across turns.

### Tier 3: Archived History and Retrieval

Older transcript content, older tool results, and raw artifacts should be stored but not included by default.

They should only be pulled into the packet when:
- the operator asks about them
- the current question depends on exact prior details
- retrieval logic identifies them as directly relevant

## Tool Result Representation

Prior tool results should not be represented only as pasted raw output.

They should be represented in three layers:

### Structured Result Object
The primary representation of a prior tool result should be a structured object containing:
- tool name
- tool call id
- status
- summary
- normalized structured data
- evidence references
- raw reference if needed

### Summary Integration
Important findings from prior tool results should be incorporated into the rolling incident summary.

### Raw Artifact Reference
If exact details are needed, the raw result should be retrieved by reference rather than included by default in every turn.

## Transcript Strategy

Full transcripts should not be sent to the model by default on every turn.

Instead:
- recent turns may be summarized
- important facts should be promoted into incident summary
- important conversational intent should be promoted into conversation summary
- full prior transcript content should be retrieval-only unless directly needed

This keeps the model focused on the current operational problem.

## Context Window Strategy

The default strategy is summary-first, artifact-selective, retrieval-on-demand.

This means:
- current input is always full
- current alert or referral is always full
- current summaries are always present
- recent critical structured artifacts are usually present
- old transcript content and raw artifacts are retrieved only when needed

## Token Budget Strategy

Prompt construction should use context buckets rather than one undifferentiated blob.

A typical priority order is:

1. system and mode instructions
2. current operator message
3. current alert or referral
4. incident summary
5. conversation summary
6. open questions
7. active goals
8. current conversational focus
9. recent critical artifacts
10. relevant runbook metadata
11. older retrieved artifacts
12. tool definitions

If context pressure is high, the system should reduce less important context first.

Suggested degradation order:
1. drop older retrieved artifacts
2. compress conversation summary
3. compress narrative portions of incident summary
4. retain structured facts and current turn inputs
5. never drop critical current-turn artifacts required to answer safely

## Modes

The same conversation system may support different operational modes.

Initial modes may include:
- investigate
- clarify
- propose

Later modes may include:
- review support
- post-action follow-up
- proactive referral handling
- routine maintenance planning

Mode affects what is prioritized in the investigation packet.

### Investigate Mode
Emphasize:
- alert/referral details
- recent evidence
- hypotheses
- open questions

### Clarify Mode
Emphasize:
- current operator question
- current conversational focus
- concise state summary
- the exact prior artifacts relevant to the clarification

### Propose Mode
Emphasize:
- latest relevant evidence
- runbook metadata
- hard constraints
- verification expectations
- unresolved blocking questions

## Summarization Strategy

Summarization should be hybrid, not purely free-form.

### Structured State Updates
After each turn, the system should update structured fields such as:
- leading hypotheses
- established facts
- ruled-out causes
- open questions
- active goals
- current conversational focus
- proposal status
- degradation notes

### Narrative Summaries
The system may also generate short narrative summaries to improve readability and LLM continuity.

Structured state is the primary operational memory.
Narrative summaries are supportive, not canonical.

## Relationship Between Conversation and Incident

Conversation state and incident state should be linked but not collapsed into one object.

A single incident may have:
- no conversation yet
- one ongoing conversation
- later multiple conversations or interaction threads

The incident remains the canonical operational record.
The conversation remains the interaction layer.

## Design Intent

This model is intended to preserve context quality over long investigations, keep token usage under control, and prevent the system from drifting into transcript replay as a substitute for structured operational memory.

The key rule is:

Each turn should be built from a layered investigation packet assembled from durable incident state, conversational summaries, recent critical artifacts, and selective retrieval when needed.