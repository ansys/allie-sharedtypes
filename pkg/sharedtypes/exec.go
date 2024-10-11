package sharedtypes

import "time"

// ExecRequest represents the requests that can be sent to allie-exec
type ExecRequest struct {
	Action               string                `json:"action"` // "execute", "append", "cancel", "status"
	InstructionGuid      string                `json:"instructionGuid"`
	ExecutionInstruction *ExecutionInstruction `json:"executionInstruction"`
}

// ExecutionInstruction contain an array of strings that represent the code to be executed in allie-exec
type ExecutionInstruction struct {
	CodeType string   `json:"codeType"` // "python", "bash"
	Code     []string `json:"Code"`
}

// ExecResponse represents the response that allie-exec sends back
type ExecResponse struct {
	Type             string            `json:"type"` // "status", "file", "error"
	InstructionGuid  string            `json:"instructionGuid"`
	Error            *ErrorResponse    `json:"error,omitempty"`
	ExecutionDetails *ExecutionDetails `json:"executionDetails,omitempty"`
	FileDetails      *FileDetails      `json:"fileDetails,omitempty"`
}

// ExecutionDetails represents the details of the execution
type ExecutionDetails struct {
	InstructionGuid     string        `json:"instructionGuid"`
	ClientGuid          string        `json:"clientGuid"`
	StartTime           time.Time     `json:"startTime"`
	TimeoutAt           time.Time     `json:"timeoutAt"`
	Response            string        `json:"response"`
	Status              string        `json:"status"` // "started", "running", "completed", "failed"
	LastResponseDiff    string        `json:"lastResponseDiff"`
	InterruptionChannel chan bool     `json:"-"`
	StdinChannel        chan []string `json:"-"`
	Cancelled           bool          `json:"-"`
}

// FileDetails contain parts of a file that is being sent
type FileDetails struct {
	FileName        string `json:"fileName"`
	FileSize        int64  `json:"fileSize"`
	FileChunkNumber int    `json:"fileChunkNumber"`
	FileChunk       []byte `json:"fileChunk"`
	IsLastChunk     bool   `json:"isLastChunk"`
}
