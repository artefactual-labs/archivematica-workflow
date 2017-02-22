package pb

import "testing"

var workflow = []byte(`{
	"chains": {
		"e7d78b8d-49d9-40fb-bd5e-8621167fc5fd": {
			"description": {
				"en": "Transfer in progress",
				"es": "Comenzando transferencia"
			},
			"link_id": "8b4415a7-eb12-40cd-b63f-544a0a342cc6"
		}
	},
	"links": {
		"8b4415a7-eb12-40cd-b63f-544a0a342cc6": {
			"config": {
				"@manager": "linkTaskManagerChoice",
				"@model": "MicroServiceChainChoice",
				"chain_choices": [
					"5a823b8b-e10d-4613-9fb9-8f6a795f03f3",
					"15139de2-c338-4055-a89e-5ff74a7ab630",
					"505544b0-a754-4722-a938-eac79a7cdce1"
				]
			},
			"description": {
				"en": "Approve standard transfer",
				"es": "Aprobar transferencia estándar"
			},
			"exit_codes": {
				"0": {
					"job_status": "Completed successfully",
					"link_id": "e7d7ee72-0e30-4f71-9980-725fc118f6cc"
				},
				"1": {
					"job_status": "Failed",
					"link_id": "f260d63d-d21d-42ae-adfd-65e711a59699"
				},
				"255": {
					"job_status": "Failed",
					"link_id": "bcabd5e2-c93e-4aaa-af6a-9a74d54e8bf0"
				}
			},
			"fallback_job_status": "Failed",
			"fallback_link_id": "83a9f35b-01ee-4d3a-86a6-62fffbede302",
			"group": {
				"en": "Approve transfer",
				"es": "Aprobar transferencia"
			}
		}
	},
	"watched_directories": [
		{
			"chain_id": "e7d78b8d-49d9-40fb-bd5e-8621167fc5fd",
			"only_dirs": true,
			"path": "/foobar",
			"unit_type": "Transfer"
		}
	]
}`)

func TestWorkflowData_UnmarshalJSON(t *testing.T) {
	w := &WorkflowData{
		Chains:             map[string]*Chain{},
		Links:              map[string]*Link{},
		WatchedDirectories: []*WatchedDirectory{},
	}
	if err := w.UnmarshalJSON(workflow); err != nil {
		t.Errorf("WorkflowData.UnmarshalJSON() error = %v", err)
	}

	t.Run("JSON to proto (chains)", func(t *testing.T) {
		if count := len(w.Chains); count != 1 {
			t.Errorf("w.Chains contains %d item(s), there should be only one", count)
		}
		const chainID = "e7d78b8d-49d9-40fb-bd5e-8621167fc5fd"
		chain, ok := w.Chains[chainID]
		if !ok {
			t.Errorf("Chain %s could not be found", chainID)
		}
		if chain.Id != chainID {
			t.Errorf("Chain ID is %v, should be %s", chain.Id, chainID)
		}
	})

	t.Run("JSON to proto (links)", func(t *testing.T) {
		if count := len(w.Chains); count != 1 {
			t.Errorf("w.Links contains %d item(s), there should be only one", count)
		}
		const linkID = "8b4415a7-eb12-40cd-b63f-544a0a342cc6"
		link, ok := w.Links[linkID]
		if !ok {
			t.Errorf("Link %s could not be found", linkID)
		}
		if link.Id != linkID {
			t.Errorf("Link ID is %v, should be %s", link.Id, linkID)
		}
		if link.FallbackJobStatus != Job_FAILED {
			t.Errorf("Expected job status was JOB_FAILED, got %s", Job_Status_name[int32(link.FallbackJobStatus)])
		}
	})

	t.Run("JSON to proto (watched directories)", func(t *testing.T) {
		if count := len(w.WatchedDirectories); count != 1 {
			t.Errorf("w.WatchedDirectories contains %d item(s), there should be only one", count)
		}
	})
}
