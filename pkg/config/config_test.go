package config

import (
	"reflect"
	"testing"
)

// TestDefineOptionalProperties tests the defineOptionalProperties function
func TestDefineOptionalProperties(t *testing.T) {
	// Setup test cases
	tests := []struct {
		name                  string
		initialConfig         Config
		optionalDefaultValues map[string]interface{}
		expectedConfig        Config
	}{
		{
			name: "Set some optional defaults",
			initialConfig: Config{
				LOG_LEVEL:                  "",
				LOCAL_LOGS:                 false,
				NUMBER_OF_WORKFLOW_WORKERS: 0,
				WEBSERVER_PORT:             "",
				SERVICE_NAME:               "",
			},
			optionalDefaultValues: map[string]interface{}{
				"LOG_LEVEL":                  "info",
				"LOCAL_LOGS":                 true,
				"NUMBER_OF_WORKFLOW_WORKERS": 5,
				"WEBSERVER_PORT":             "8080",
				"SERVICE_NAME":               "AaliService",
			},
			expectedConfig: Config{
				LOG_LEVEL:                  "info",
				LOCAL_LOGS:                 true,
				NUMBER_OF_WORKFLOW_WORKERS: 5,
				WEBSERVER_PORT:             "8080",
				SERVICE_NAME:               "AaliService",
			},
		},
		{
			name: "Partial defaults applied",
			initialConfig: Config{
				LOG_LEVEL:                  "debug",
				LOCAL_LOGS:                 false,
				NUMBER_OF_WORKFLOW_WORKERS: 0,
				WEBSERVER_PORT:             "",
				SERVICE_NAME:               "ExistingService",
			},
			optionalDefaultValues: map[string]interface{}{
				"LOG_LEVEL":                  "info",
				"LOCAL_LOGS":                 true,
				"NUMBER_OF_WORKFLOW_WORKERS": 10,
				"WEBSERVER_PORT":             "9090",
				"SERVICE_NAME":               "AaliService",
			},
			expectedConfig: Config{
				LOG_LEVEL:                  "debug",
				LOCAL_LOGS:                 true,
				NUMBER_OF_WORKFLOW_WORKERS: 10,
				WEBSERVER_PORT:             "9090",
				SERVICE_NAME:               "ExistingService",
			},
		},
		{
			name: "No changes needed",
			initialConfig: Config{
				LOG_LEVEL:                  "warn",
				LOCAL_LOGS:                 true,
				NUMBER_OF_WORKFLOW_WORKERS: 8,
				WEBSERVER_PORT:             "8081",
				SERVICE_NAME:               "CustomService",
			},
			optionalDefaultValues: map[string]interface{}{
				"LOG_LEVEL":                  "info",
				"LOCAL_LOGS":                 false,
				"NUMBER_OF_WORKFLOW_WORKERS": 5,
				"WEBSERVER_PORT":             "8080",
				"SERVICE_NAME":               "AaliService",
			},
			expectedConfig: Config{
				LOG_LEVEL:                  "warn",
				LOCAL_LOGS:                 true,
				NUMBER_OF_WORKFLOW_WORKERS: 8,
				WEBSERVER_PORT:             "8081",
				SERVICE_NAME:               "CustomService",
			},
		},
	}

	// Run each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the initial config for each test case
			config := tt.initialConfig

			// Call the function under test
			defineOptionalProperties(&config, tt.optionalDefaultValues)

			// Compare the result with the expected config
			if !reflect.DeepEqual(config, tt.expectedConfig) {
				t.Errorf("Expected config %+v, got %+v", tt.expectedConfig, config)
			}
		})
	}
}
