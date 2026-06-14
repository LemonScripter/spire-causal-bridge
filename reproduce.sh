#!/bin/bash
set -e

# DCC SPIRE Bridge: Automated Reproduction Script
# Validates the causal identity attestation logic.

echo "--- [1/3] Building SPIRE Causal Attestor Proof ---"
mkdir -p bin
go build -o bin/spire-dcc-proof src/main.go

echo "--- [2/3] Verifying Causal Attestation (Demo Mode) ---"

echo "Test A: Verified Workload (PID 5555)"
RESULT_A=$(./bin/spire-dcc-proof --demo --pid 5555)
echo "$RESULT_A"
if [[ "$RESULT_A" == *"ATTESTED"* ]]; then
    echo "[PASS] Verified workload successfully attested."
else
    echo "[FAIL] Verified workload rejected."
    exit 1
fi

echo -e "\nTest B: Orphaned Workload (PID 8888)"
if ./bin/spire-dcc-proof --demo --pid 8888 > /dev/null 2>&1; then
    echo "[FAIL] Orphaned workload allowed identity (Security Breach)."
    exit 1
else
    echo "[PASS] Orphaned workload rejected (Fail-Closed)."
fi

echo -e "\n--- [3/3] SPIFFE Integration ---"
echo "Upon successful attestation, the plugin issues the selector:"
echo "dcc:causal_chain:verified"

echo -e "\nSUCCESS: SPIRE Causal Attestation logic verified."
