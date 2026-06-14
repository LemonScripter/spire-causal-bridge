# DCC Causal Attestor for SPIFFE/SPIRE

[![Verify](https://github.com/LemonScripter/spire-causal-bridge/actions/workflows/verify.yml/badge.svg)](https://github.com/LemonScripter/spire-causal-bridge/actions/workflows/verify.yml)
[![Status](https://img.shields.io/badge/Status-Hardened--Prototype-blue)](https://github.com/LemonScripter/spire-causal-bridge/blob/master/ROADMAP.md)
[![Project](https://img.shields.io/badge/BioOS-Causal--Security-green)](https://metaspace.bio)
[![DOI](https://img.shields.io/badge/DOI-10.5281%2Fzenodo.20384700-purple)](https://doi.org/10.5281/zenodo.20384700)

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

### Scientific & Technical Foundation

This implementation is based on the following formal specifications and research:

- **Research Paper:** [The Causal Operating System: Digital Causal Closure for Autonomous Systems](https://doi.org/10.5281/zenodo.20384700) (DOI: 10.5281/zenodo.20384700)
- **Formal Specification:** [BioOS Causal Constitution (PDF)](https://bioos.metaspace.bio/bioos_causal_constitution_en.pdf)

---
*MetaSpace.Bio Logic Project | [metaspace.bio](https://metaspace.bio) | [admin@metaspace.bio](mailto:admin@metaspace.bio)*
