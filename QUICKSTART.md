# Quickstart: DCC Causal Attestor for SPIRE

This guide provides a self-contained environment to verify the **Causal Identity** logic for SPIFFE/SPIRE.

## Prerequisites
- **Go** (1.21+)
- **Linux** or **macOS**

## Step 1: Build the Attestor Proof
```bash
git clone https://github.com/LemonScripter/spire-causal-bridge.git
cd spire-causal-bridge
make build
```

## Step 2: Run the Automated Proof
Execute the reproduction script to see the "Fail-Closed Attestation" logic in action:
```bash
./reproduce.sh
```

## Step 3: Production Deployment
For live attestation, the plugin must be registered in the SPIRE Agent configuration. It will automatically link with the BioOS kernel state to verify workload provenance before issuing SVIDs.

## Verification Scenarios
- **Attested:** Workloads with a verified causal lineage receive the `dcc:causal_chain:verified` selector.
- **Rejected:** Orphaned processes are denied identity issuance, physically preventing session hijacking.

---
*Production-Grade Research Prototype by MetaSpace BioOS Team*
