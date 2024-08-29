package config

import (
	"fmt"
	"strings"
)

// Config contains all the configuration settings for the Allie service.
type Config struct {

	// Logging
	///////////
	LOG_LEVEL string `yaml:"LOG_LEVEL"`
	// Local Logs
	LOCAL_LOGS          bool   `yaml:"LOCAL_LOGS"`
	LOCAL_LOGS_LOCATION string `yaml:"LOCAL_LOGS_LOCATION"`
	// Datadog Logs
	DATADOG_LOGS        bool   `yaml:"DATADOG_LOGS"`
	STAGE               string `yaml:"STAGE"`
	VERSION             string `yaml:"VERSION"`
	SERVICE_NAME        string `yaml:"SERVICE_NAME"`
	ERROR_FILE_LOCATION string `yaml:"ERROR_FILE_LOCATION"`
	LOGGING_URL         string `yaml:"LOGGING_URL"`
	LOGGING_API_KEY     string `yaml:"LOGGING_API_KEY"`
	DATADOG_SOURCE      string `yaml:"DATADOG_SOURCE"`
	// Datadog Metrics
	DATADOG_METRICS bool   `yaml:"DATADOG_METRICS"`
	METRICS_URL     string `yaml:"METRICS_URL"`

	// SSL Settings
	/////////////////
	USE_SSL                   bool   `yaml:"USE_SSL"`
	SSL_CERT_PUBLIC_KEY_FILE  string `yaml:"SSL_CERT_PUBLIC_KEY_FILE"`
	SSL_CERT_PRIVATE_KEY_FILE string `yaml:"SSL_CERT_PRIVATE_KEY_FILE"`

	// Allie Agent
	///////////////
	AGENT_PORT                 string `yaml:"AGENT_PORT,omitempty"`
	WORKFLOW_API_KEY           string `yaml:"WORKFLOW_API_KEY,omitempty"`
	WORKFLOW_STORE_PATH        string `yaml:"WORKFLOW_STORE_PATH,omitempty"`
	NUMBER_OF_WORKFLOW_WORKERS int    `yaml:"NUMBER_OF_WORKFLOW_WORKERS,omitempty"`
	BINARY_STORE_PATH          string `yaml:"BINARY_STORE_PATH,omitempty"`
	EXTERNALFUNCTIONS_ENDPOINT string `yaml:"EXTERNALFUNCTIONS_ENDPOINT"`
	FLOWKIT_PYTHON_ENDPOINT    string `yaml:"FLOWKIT_PYTHON_ENDPOINT"`
	FLOWKIT_PYTHON_API_KEY     string `yaml:"FLOWKIT_PYTHON_API_KEY"`

	// Allie Flowkit
	/////////////////
	EXTERNALFUNCTIONS_GRPC_PORT string `yaml:"EXTERNALFUNCTIONS_GRPC_PORT,omitempty"`
	FLOWKIT_API_KEY             string `yaml:"FLOWKIT_API_KEY,omitempty"`
	// Allie Modules
	LLM_HANDLER_ENDPOINT  string `yaml:"LLM_HANDLER_ENDPOINT,omitempty"`
	KNOWLEDGE_DB_ENDPOINT string `yaml:"KNOWLEDGE_DB_ENDPOINT,omitempty"`
	EXEC_ENDPOINT         string `yaml:"EXEC_ENDPOINT,omitempty"`
	// Azure Cognitive Services
	ACS_ENDPOINT    string `yaml:"ACS_ENDPOINT,omitempty"`
	ACS_API_KEY     string `yaml:"ACS_API_KEY,omitempty"`
	ACS_API_VERSION string `yaml:"ACS_API_VERSION,omitempty"`

	// Allie LLM
	/////////////
	WEBSERVER_PORT         string `yaml:"WEBSERVER_PORT"`
	MODELS_CONFIG_LOCATION string `yaml:"MODELS_CONFIG_LOCATION"`

	// Allie DB
	///////////
	WEBSERVER_PORT_DB                  string   `yaml:"WEBSERVER_PORT_DB"`
	EMBEDDINGS_DIMENSIONS              int      `yaml:"EMBEDDINGS_DIMENSIONS"`
	MILVUS_INDEX_TYPE                  string   `yaml:"MILVUS_INDEX_TYPE"`
	MILVUS_METRIC_TYPE                 string   `yaml:"MILVUS_METRIC_TYPE"` // cosine, l2 or ip
	MILVUS_HOST                        string   `yaml:"MILVUS_HOST"`
	MILVUS_PORT                        string   `yaml:"MILVUS_PORT"`
	NEO4J_DB                           bool     `yaml:"NEO4J_DB"`
	NEO4J_URI                          string   `yaml:"NEO4J_URI"`
	NEO4J_USERNAME                     string   `yaml:"NEO4J_USERNAME"`
	NEO4J_PASSWORD                     string   `yaml:"NEO4J_PASSWORD"`
	TEMP_COLLECTION_NAME               string   `yaml:"TEMP_COLLECTION_NAME"`
	ELASTICSEARCH_HOST                 string   `yaml:"ELASTICSEARCH_HOST"`
	ELASTICSEARCH_PORT                 string   `yaml:"ELASTICSEARCH_PORT"`
	ELASTICSEARCH_USERNAME             string   `yaml:"ELASTICSEARCH_USERNAME"`
	ELASTICSEARCH_PASSWORD             string   `yaml:"ELASTICSEARCH_PASSWORD"`
	ELASTICSEARCH_INDEX_TYPE           string   `yaml:"ELASTICSEARCH_INDEX_TYPE"`           // cosineSimilarity or dotProduct
	ELASTICSEARCH_TRUSTED_CERTIFICATES []string `yaml:"ELASTICSEARCH_TRUSTED_CERTIFICATES"` // list of paths to trusted certificates
	ELASTICSEARCH_INSECURE_CONNECTION  bool     `yaml:"ELASTICSEARCH_INSECURE_CONNECTION"`
	DATABASE_TYPE                      string   `yaml:"DATABASE_TYPE"` // milvus or elasticsearch

	// Allie Exec
	//////////////
	WEBSERVER_PORT_EXEC            string `yaml:"WEBSERVER_PORT_EXEC"`
	EXECUTABLE                     string `yaml:"EXECUTABLE"`
	WATCH_FOLDER_PATH              string `yaml:"WATCH_FOLDER_PATH"`
	MILLISECONDS_SINCE_LAST_CHANGE int    `yaml:"MILLISECONDS_SINCE_LAST_CHANGE"`
}

// Initialize conifg dict
var GlobalConfig *Config

// flagStringSlice is a custom flag type for string slices.
type flagStringSlice []string

// String returns a string representation of the flagStringSlice.
//
// Returns:
//   - string: The string representation of the flagStringSlice.
func (fss *flagStringSlice) String() string {
	return fmt.Sprintf("%v", *fss)
}

// Set sets the value of the flagStringSlice.
//
// Parameters:
//   - value: The value to set.
//
// Returns:
//   - error: An error if there was an issue setting the value.
func (fss *flagStringSlice) Set(value string) error {
	*fss = strings.Split(value, ",")
	return nil
}
