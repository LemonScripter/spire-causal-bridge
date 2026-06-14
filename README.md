# DCC Causal Attestor for SPIFFE/SPIRE

[![Status](https://img.shields.io/badge/Status-Hardened--Prototype-blue)](ROADMAP.md)
[![Project](https://img.shields.io/badge/BioOS-Causal--Security-green)](https://metaspace.bio)

## Hardened Architecture: Causal Identity Attestation

This plugin implements **Causal Workload Attestation** for SPIRE. It transforms the concept of "Identity" from binary integrity to **Hardware-Anchored Intent Verification**.

### Hardened Implementation

- **Direct eBPF Map Lookup:** The plugin interacts directly with the `global_dcc_map` pinned at `/sys/fs/bpf/spire/` to verify process lineage before issuing an SVID.
- **Fail-Closed Attestation:** If a process lacks a valid, hardware-anchored Causal Token, the attestation fails with a `DCC Violation`.
- **Selector Enrichment:** Successfully attested workloads receive the `dcc:causal_chain:verified` selector, enabling fine-grained, intent-aware SPIRE registration entries.

### Security Guarantees

1. **Anti-Hijacking:** Identities are only issued to processes that can prove they were launched through an authorized causal event.
2. **Replay Protection:** Atomic token consumption ensures that a single intent cannot be used to spawn multiple unauthorized identities.
3. **Fail-Closed Integrity:** Identity issuance is physically dependent on kernel-level causal closure.

### Scientific Foundation

This implementation is based on the [BioOS Causal Constitution (DOI: 10.5281/zenodo.20384700)](https://doi.org/10.5281/zenodo.20384700).

---
*Verified by MetaSpace BioOS Team | [metaspace.bio](https://metaspace.bio)*
