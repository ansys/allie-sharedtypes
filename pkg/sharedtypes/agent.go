package sharedtypes

import (
	"encoding/json"

	"github.com/ansys/allie-sharedtypes/pkg/logging"
)

// Message represents the JSON message you are expecting
type SessionContext struct {
	SessionType string            `json:"session_type"`          // Type of session: "workflow", "exec"
	ApiKey      string            `json:"api_key"`               // API key for authentication, only relevant if "session_type" is "exec"
	JwtToken    string            `json:"jwt_token"`             // JWT token for authentication (optional)
	ExecId      string            `json:"exec_id,omitempty"`     // Unique identifier of connecting Exec, only relevant if "session_type" is "exec"
	WorkflowId  string            `json:"workflow_id,omitempty"` // Workflow ID, only relevant if "workflow_endpoint" is "custom"
	Variables   map[string]string `json:"variables,omitempty"`   // Variables to be passed to the workflow
	// Snapshot logic
	SnapshotId     string `json:"snapshot_id,omitempty"`     // Snapshot ID, only relevant if "session_type" is "workflow"; if defined, the given snapshot will retrived from the database
	WorkflowRunId  string `json:"workflow_run_id,omitempty"` // Workflow run ID, only relevant if "session_type" is "workflow"; if defined, mandatory if "snapshot_id" is defined in order to retrieve the snapshot from the database
	UserId         string `json:"user_id,omitempty"`         // User ID, only relevant if "session_type" is "workflow"; if defined, mandatory if "snapshot_id" is defined in order to retrieve the snapshot from the database
	StoreSnapshots bool   `json:"store_snapshots,omitempty"` // Store snapshots, only relevant if "session_type" is "workflow"; if true, all taken snapshots will be stored in the database
}

// SetSessionContext sets the SessionContext struct from the JSON payload
//
// Parameters:
//   - msg: the JSON payload
//
// Returns:
//   - SessionContext: the SessionContext struct
func ExtractSessionContext(ctx *logging.ContextMap, msg []byte) (SessionContext, error) {
	var SessionContext SessionContext

	// Unmarshal the JSON payload into your struct
	if err := json.Unmarshal(msg, &SessionContext); err != nil {
		logging.Log.Error(ctx, "Error decoding JSON:", err)
		return SessionContext, err
	} else {
		return SessionContext, nil
	}
}
