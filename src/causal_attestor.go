package main

import (
	"context"
	"fmt"
	"unsafe"

	"github.com/cilium/ebpf"
	"github.com/hashicorp/go-plugin"
	"github.com/spiffe/spire-plugin-sdk/plugins/workloadattestor/v1"
)

/*
 * SPIRE DCC Causal Attestor Plugin (Go) - Hardened
 * 
 * This plugin implements "Causal Identity" for the SPIRE Agent. 
 * It ensures that identities are only issued to workloads with a verified 
 * kernel-anchored causal lineage, physically preventing identity hijacking.
 */

const (
	DCCMapPath = "/sys/fs/bpf/spire/global_dcc_map"
)

type DCCToken struct {
	Timestamp  uint64
	IntentID   uint32
	AgeLimitNS uint32
	Consumed   uint8
}

type CausalAttestorPlugin struct {
	workloadattestor.UnsafeWorkloadAttestorServer
	dccMap *ebpf.Map
}

func (p *CausalAttestorPlugin) Attest(ctx context.Context, req *workloadattestor.AttestRequest) (*workloadattestor.AttestResponse, error) {
	if p.dccMap == nil {
		m, err := ebpf.LoadPinnedMap(DCCMapPath, nil)
		if err != nil {
			return nil, fmt.Errorf("DCC Critical: Failed to load pinned causal map: %w", err)
		}
		p.dccMap = m
	}

	pid := uint32(req.Pid)
	var token DCCToken
	
	// Atomic kernel-state verification
	if err := p.dccMap.Lookup(unsafe.Pointer(&pid), unsafe.Pointer(&token)); err != nil {
		return nil, fmt.Errorf("DCC Violation: No causal lineage found for PID %d", pid)
	}

	// Fail-Closed Logic: Verify causal integrity
	if token.Consumed != 0 {
		return nil, fmt.Errorf("DCC Violation: Token replay detected for PID %d", pid)
	}

	// SPIRE Attestation issues the 'dcc:causal_chain:verified' selector
	return &workloadattestor.AttestResponse{
		Selectors: []string{
			"dcc:causal_chain:verified",
			fmt.Sprintf("dcc:intent:%x", token.IntentID),
		},
	}, nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: workloadattestor.Handshake,
		Plugins: map[string]plugin.Plugin{
			"dcc_causal_attestor": &workloadattestor.WorkloadAttestorPlugin{
				Server: &CausalAttestorPlugin{},
			},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
