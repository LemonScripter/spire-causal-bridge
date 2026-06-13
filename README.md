# DCC Causal Attestor for SPIFFE/SPIRE

## Overview
The **DCC Causal Attestor** is a professional plugin for the SPIRE Agent that implements **Causal Identity**. It ensures that SPIFFE identities (SVIDs) are only issued to workloads that can prove a hardware-anchored causal lineage via the **BioOS Digital Causal Closure (DCC)** kernel module.

## The Problem: Identity Hijacking
Standard workload attestation relies on static attributes (binary hashes, Kubernetes labels). If a legitimate workload is compromised after startup, or if a rogue process mimics these attributes, SPIRE will still issue an identity. This is because identity is currently disconnected from the *intent* that launched the process.

## The Solution: Causal Identity
By integrating DCC into the attestation flow, we bind identity to causality:
1. **Out-of-Band Verification:** When a workload calls the SPIRE Agent, the DCC plugin queries the kernel's `global_dcc_map` based on the workload's PID.
2. **Selector Issuance:** The `dcc:causal_chain:verified` selector is only granted if a valid, non-expired Causal Token exists for that process.
3. **Hardened Trust:** Identity becomes a proof of legitimate execution origin, not just a proof of binary integrity.

## Scientific Background
This integration is based on the following formal research:
- [The Causal Operating System: Digital Causal Closure for Autonomous Systems](https://doi.org/10.5281/zenodo.20384700)
- [BioOS Causal Constitution (PDF)](https://bioos.metaspace.bio/bioos_causal_constitution_en.pdf)

## Components
- **`causal_attestor.go`**: SPIRE Workload Attestor plugin (gRPC-based).
- **`verify_spire.py`**: Logic verification suite ensuring identity issuance only for verified causal chains.

## Upstreaming Proposal
We propose the inclusion of Causal Attestation as a standardized security layer for SPIRE deployments in autonomous and high-stakes environments.

---
*Created by MetaSpace BioOS | [metaspace.bio](https://metaspace.bio) | [admin@metaspace.bio](mailto:admin@metaspace.bio)*
