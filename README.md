# DCC Causal Authorization for SPIFFE/SPIRE

[![Status](https://img.shields.io/badge/Status-Research--Prototype-blue)](VERIFICATION.md)
[![Layer](https://img.shields.io/badge/Layer-Authorization%20(not%20identity)-orange)](README.md)
[![DOI](https://img.shields.io/badge/DOI-10.5281%2Fzenodo.20384700-purple)](https://doi.org/10.5281/zenodo.20384700)

A kernel-anchored **causal authorization** layer that *complements* SPIRE.

SPIRE answers **who** a workload is — a durable, structural identity (the SVID).
This layer answers a different question, **at the moment an action is taken**:
*was this specific action caused by a genuine external stimulus, or is it
self-initiated?* It denies security-relevant actions issued **outside a fresh
causal window of a real external event** — **even when the workload holds a
valid SVID**.

> **This is authorization, not authentication — enforced at the point of use,
> not at issuance.** It does **not** modify SPIRE or its identity semantics. It
> *consumes* the already-attested identity and adds one property the SVID cannot
> carry: the **freshness of the action's external cause**.

## The gap it closes

A correctly signed, correctly attested workload still gets a valid SVID if it is
hijacked by an exploit or a logic bug. Identity is real; intent is not. The
attacker's value is mostly in the **asynchronous post-exploitation tail**:

- reverse shells and **delayed C2 beacons**,
- orphaned / background **exfiltration**,
- self-initiated actions from a logic bug.

All of these present a **valid SVID**. An identity check admits them. A *causal*
check does not: there is no fresh external cause behind a self-initiated action.

**Out of scope (stated plainly):** in-band synchronous exploitation that runs
*inside* a genuine request window. This layer addresses the asynchronous tail,
not synchronous in-band abuse.

## Why at the point of use, not at issuance

The legitimate causal window is on the order of the stimulus→syscall latency
(**sub-millisecond**; measured ~100 µs in our validation). An X509-SVID TTL is on
the order of **an hour** and is cached. A property that is fresh for microseconds
cannot be baked into a credential that lives for an hour without being
**stale-by-construction**. So the freshness check must happen where the SVID is
**presented / used**, not where it is issued.

## Status (honest ledger)

- **Formal model — `[PROVEN]`.** A machine-checked (Z3) coverage theorem shows
  the layer is a **strict, non-regressive tightening** of SPIRE: it never opens a
  hole SPIRE closed (no regression) and provably denies actions SPIRE alone would
  admit (strictly stronger). The whole guarantee rests on one load-bearing
  assumption — *producer fidelity* — which is validated empirically.
- **Empirical validation — `[MEASURED]`.** The discriminator was validated on a
  **real Linux kernel (6.x, BPF LSM)**. One workload, one identity, two egress
  attempts to the same destination differing **only** in causal provenance: the
  **data-caused** egress is **admitted**, the **self-initiated** egress is
  **denied fail-closed**, with **zero false denials** and **zero causal
  false-positives** across the run (caused-egress window ~100 µs). An
  identity-presentation check would admit both.

The kernel producer/guard implementation and the raw measurement logs are under
**controlled disclosure (patent pending)** and are available to SPIRE
maintainers and researchers on request: **admin@metaspace.bio**.

## Scientific foundation

- **Research paper:** *The Causal Operating System: Digital Causal Closure for
  Autonomous Systems* — [DOI: 10.5281/zenodo.20384700](https://doi.org/10.5281/zenodo.20384700)

`src/` contains a **reference integration sketch** only; the enforced
architecture is authorization-at-point-of-use as described above (see the note in
`src/`).

---
*MetaSpace.Bio Logic Project | [metaspace.bio](https://metaspace.bio) | admin@metaspace.bio*
