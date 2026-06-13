# Proposal: Zero-Trust Causal Attestation for SPIRE Agents

Hello SPIFFE/SPIRE Community,

Current workload attestation in SPIRE is highly effective at verifying *environment* and *integrity* (via hashes, labels, and metadata). However, there is an emerging need to verify **Causality**: the proof of intent that originated the workload's execution.

### The Problem: Contextless Identity
Identity issuance is currently disconnected from the causal event that triggered the process. A correctly signed binary, if launched autonomously by an exploit or a logic bug, still qualifies for an SVID because its environmental selectors remain valid.

### Proposal: Causal Identity Attestor
I propose a new **Workload Attestor Plugin** for the SPIRE Agent that leverages **Digital Causal Closure (DCC)**. This plugin verifies if a process is part of a hardware-anchored causal chain before granting identity-defining selectors.

### Proof of Concept
I have developed a PoC plugin that interfaces with a kernel-level DCC module to attest workloads based on their causal lineage:
[https://github.com/LemonScripter/spire-causal-bridge](https://github.com/LemonScripter/spire-causal-bridge)

### Key Benefits:
1. **Dynamic Trust:** Identity is only issued if the process can prove it was caused by a verified event (e.g., a signed scheduler command).
2. **Orphaned Workload Prevention:** Orphaned processes (launched without a valid token) are denied identity, preventing them from accessing sensitive mesh resources.
3. **Formal Foundation:** Based on peer-reviewed research on Causal Operating Systems (DOI: 10.5281/zenodo.20384700).

We believe Causal Attestation is the next logical step in the evolution of Zero-Trust identity for autonomous and cloud-native workloads.

Best regards,

**MetaSpace BioOS Team**
[metaspace.bio](https://metaspace.bio) | [admin@metaspace.bio](mailto:admin@metaspace.bio)
