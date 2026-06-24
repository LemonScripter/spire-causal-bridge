# Verification

This project keeps a strict separation between **logic unit tests** (in this
repo) and the **empirical kernel validation** (under controlled disclosure).
Nothing here is presented as more than it is.

## 1. In-repo logic unit tests — `[unit test, not a measurement]`

`tests/verify_spire.py` and `reproduce.sh` are **unit tests of the integration
logic** (token presence / freshness / expiry on an in-memory map). They run in
microseconds and use synthetic PIDs. They demonstrate the *decision logic*; they
are **not** a measurement of the real kernel producer and must not be read as
empirical evidence of producer fidelity.

## 2. Formal model — `[PROVEN]`

A machine-checked (Z3) coverage theorem establishes that the layer is a strict,
non-regressive tightening of SPIRE (no regression ∧ strictly stronger), modulo
one load-bearing assumption: **producer fidelity** — that a causal token is
stamped *only* on a genuine external event. The model proves the state machine;
it does not prove the kernel code. That gap is closed by §3.

## 3. Empirical kernel validation — `[MEASURED]` (raw logs on request)

The discriminator was validated on a **real Linux kernel (6.x, BPF LSM)**:

- One workload, one identity, two egress attempts to the **same** destination,
  differing **only** in causal provenance.
- **Data-caused egress:** admitted, with a measured causal-token age of ~100 µs.
- **Self-initiated egress (same identity):** denied fail-closed (`EPERM`).
- Over the run: **0 false denials**, **0 causal false-positives**, and a clean
  separation between the caused and self-initiated token-age distributions
  (orders of magnitude).

These results validate the *producer-fidelity* assumption of §2 operationally.
The kernel implementation and the raw, timestamped logs are under controlled
disclosure (patent pending); available to maintainers/researchers on request:
**admin@metaspace.bio**.

---
*MetaSpace.Bio Logic Project | [metaspace.bio](https://metaspace.bio) | admin@metaspace.bio*
