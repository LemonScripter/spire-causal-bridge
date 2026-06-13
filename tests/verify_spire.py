import unittest
import time

# SPIRE DCC Causal Attestor Logic Verification
# Verifies that identities (SVIDs) are only issued for verified causal lineages.

class TestSPIREDCCAttestor(unittest.TestCase):
    def setUp(self):
        self.dcc_map = {}
        self.CAUSALITY_WINDOW_NS = 500 * 1000 * 1000 # 500ms

    def issue_token(self, pid):
        self.dcc_map[pid] = {
            "ts": time.time_ns(),
            "consumed": False
        }

    def spire_attest_check(self, pid):
        # Simulation of the Attest() call in the SPIRE DCC Plugin
        now = time.time_ns()
        
        if pid not in self.dcc_map:
            return False, "BLOCK: NO_TOKEN"
            
        token = self.dcc_map[pid]
        if now - token["ts"] > self.CAUSALITY_WINDOW_NS:
            return False, "BLOCK: EXPIRED"
            
        # For attestation, we don't 'consume' the token, as it might be used 
        # by multiple attestors or sensors during the same window.
        
        return True, "dcc:causal_chain:verified"

    def test_reject_orphaned_workload(self):
        # Workload with no token must fail attestation
        pid = 5555
        success, result = self.spire_attest_check(pid)
        self.assertFalse(success)
        self.assertEqual(result, "BLOCK: NO_TOKEN")

    def test_approve_verified_workload(self):
        # Workload with fresh token must receive the selector
        pid = 5555
        self.issue_token(pid)
        success, result = self.spire_attest_check(pid)
        self.assertTrue(success)
        self.assertEqual(result, "dcc:causal_chain:verified")

    def test_reject_expired_lineage(self):
        # Stale causal link must fail attestation
        pid = 5555
        self.issue_token(pid)
        self.dcc_map[pid]["ts"] -= (self.CAUSALITY_WINDOW_NS + 1000)
        success, result = self.spire_attest_check(pid)
        self.assertFalse(success)
        self.assertEqual(result, "BLOCK: EXPIRED")

if __name__ == "__main__":
    print("--- Running SPIRE DCC Causal Attestor Tests ---")
    unittest.main()
