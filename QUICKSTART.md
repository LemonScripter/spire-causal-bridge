# Quickstart: DCC Causal Authorization for SPIRE

This guide builds the **reference integration sketch** and runs the in-repo
**logic unit tests** for the causal-authorization decision (token presence /
freshness). It is *not* the empirical kernel validation — see `VERIFICATION.md`
for the separation between logic tests and the real-kernel measurement.

## Prerequisites
- **Go** (1.21+)
- **Linux** or **macOS**

## Step 1: Build the reference sketch
```bash
git clone https://github.com/LemonScripter/spire-causal-bridge.git
cd spire-causal-bridge
make build
```

## Step 2: Run the logic unit tests
```bash
./reproduce.sh
```
This exercises the decision logic (fail-closed on absent/stale causal token) on
an in-memory map with synthetic PIDs.

## Step 3: Where enforcement actually happens
The enforced architecture is **authorization at the point of use** — the causal
check is applied when the SVID is presented / used (e.g. at the egress syscall),
consuming the already-attested identity. It does **not** gate SVID issuance and
does **not** modify SPIRE. The kernel producer/guard is under controlled
disclosure (patent pending); see `README.md`.

## Decision scenarios (logic)
- **Admit:** an action inside a fresh causal window of a genuine external event.
- **Deny (fail-closed):** a self-initiated / orphaned action with no fresh causal
  token — even if it carries a valid SVID.

---
*Research Prototype — MetaSpace BioOS Team*
