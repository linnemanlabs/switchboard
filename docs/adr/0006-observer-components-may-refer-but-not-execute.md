# ADR 0006: Observer components may refer but not execute

- Status: accepted
- Date: 2026-03-27

## Context

A future version of this system will include observer-style components that continuously watch logs, telemetry, and other operational signals for weak signals, drift, anomalies, and unusual correlations.

These observer components are intended to identify potentially important situations that may not yet have triggered conventional alerts.

One possible design is to let observer components both detect and directly trigger remediation or execution paths.

Another possible design is to restrict observer components to generating referrals that are then investigated and handled by other components.

Because observer components will operate on large volumes of noisy, incomplete, and probabilistic data, and because their role is discovery rather than action, they should not directly initiate infrastructure changes.

## Decision

Observer components may generate referrals, findings, and investigation prompts, but they may not directly execute runbooks or bypass the normal proposal, review, and execution path.

Observer outputs may:
- create or enrich incident context
- send a referral to Vigil
- attach evidence references
- assign urgency or confidence
- suggest investigation directions

Observer outputs may not:
- directly trigger infrastructure execution
- bypass review or approval
- broaden their own authority into execution control

## Rationale

### Discovery is not execution
Observers exist to notice weak signals and unusual patterns.
That role is fundamentally different from deciding and performing infrastructure changes.

### Weak-signal analysis is probabilistic
Observer findings may be valuable while still being uncertain, noisy, or incomplete.
That is acceptable for referrals, but not sufficient for direct action.

### Better safety boundaries
Separating observer duties from execution reduces the risk that a high-volume pattern detector becomes an unsafe actor.

### Cleaner architecture
A referral model keeps observer logic focused on detection and pattern formation rather than remediation logic.

### Better explainability
When observers refer issues upstream, the investigation and action path remains easier to audit and reason about.

## Consequences

### Positive
- safer anomaly-detection architecture
- cleaner separation of duties
- easier to experiment with observer heuristics and models
- easier to tolerate observer false positives
- better long-term maintainability

### Negative
- adds one more handoff in the path from detection to remediation
- observer findings may require additional investigation before action can be proposed
- can feel slower than direct autonomous action

## Alternatives considered

### Observer may directly execute low-risk actions
Rejected for now because it weakens the architectural separation between detection and action.

### Observer may directly create execution requests
Rejected because it would let probabilistic pattern detection bypass the normal investigation and review path.

### Observer and Vigil merged into one component
Rejected because the responsibilities are materially different and benefit from distinct design and tuning.

## Notes

Observers may still be highly valuable and high-urgency.
They simply express that value through referrals, not direct action.

This decision reinforces the general system principle that action should remain narrower and more deterministic than detection.