// Copyright (C) 2025 ANSYS, Inc. and/or its affiliates.
// SPDX-License-Identifier: MIT
//
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
