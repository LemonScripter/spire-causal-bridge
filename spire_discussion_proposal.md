# Proposal: Causal Authorization as a complement to SPIRE identity

Hello SPIFFE/SPIRE Community,

This is a corrected framing of an earlier proposal. SPIRE answers **who** a
workload is. The idea here is **not** to change that, nor to gate SVID issuance
on anything transient. It is a separate **authorization** concern, enforced **at
the point of use**.

### The gap

Identity issuance is — correctly — decoupled from the causal event behind a
process. A correctly signed binary, if launched or driven by an exploit or a
logic bug, still holds a valid SVID because its identity is real. The risk lives
in the **asynchronous post-exploitation tail**: delayed C2 beacons, reverse
shells, orphaned exfil, self-initiated logic-bug actions — all presenting a valid
SVID.

### The proposal

A kernel-anchored **causal authorization** check, applied **when the SVID is
presented / used** (e.g. at the egress syscall), that admits an action only if it
falls inside a **fresh causal window of a genuine external stimulus**. It
consumes the already-attested identity; it does not modify SPIRE.

Why at use and not at issuance: the causal window is sub-millisecond, while an
X509-SVID TTL is ~an hour and cached — a microsecond-fresh property cannot live
inside an hour-long credential without being stale-by-construction.

### What distinguishes it from identity-based authorization

Identity-presentation authorization (JWT present / X509 ownership) cannot separate
a legitimate action from a self-initiated one when both carry the **same valid
SVID**. A causal check can. We have validated this on a real kernel: same
identity, two egress attempts differing only in causal provenance — the caused
one admitted, the self-initiated one denied fail-closed, with zero false denials
and zero causal false-positives (raw logs available on request).

### Scope

Addresses the **asynchronous tail**, not in-band synchronous exploitation inside
a genuine request window. v1 targets blocking-server workloads. Assumes a trusted
kernel.

### Foundation

Based on *The Causal Operating System: Digital Causal Closure for Autonomous
Systems* (DOI: 10.5281/zenodo.20384700).

Best regards,
**MetaSpace BioOS Team** — [metaspace.bio](https://metaspace.bio) | admin@metaspace.bio
