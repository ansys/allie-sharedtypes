package config

import (
	"fmt"
	"strings"
)

// Config contains all the configuration settings for the Allie service.
type Config struct {

	// Logging
	///////////
	LOG_LEVEL string `yaml:"LOG_LEVEL" json:"LOGLEVEL"`
	// Local Logs
	LOCAL_LOGS          bool   `yaml:"LOCAL_LOGS" json:"LOCALLOGS"`
	LOCAL_LOGS_LOCATION string `yaml:"LOCAL_LOGS_LOCATION" json:"LOCALLOGSLOCATION"`
	// Datadog Logs
	DATADOG_LOGS        bool   `yaml:"DATADOG_LOGS" json:"DATADOGLOGS"`
	STAGE               string `yaml:"STAGE" json:"STAGE"`
	VERSION             string `yaml:"VERSION" json:"VERSION"`
	SERVICE_NAME        string `yaml:"SERVICE_NAME" json:"SERVICENAME"`
	ERROR_FILE_LOCATION string `yaml:"ERROR_FILE_LOCATION" json:"ERRORFILELOCATION"`
	LOGGING_URL         string `yaml:"LOGGING_URL" json:"LOGGINGURL"`
	LOGGING_API_KEY     string `yaml:"LOGGING_API_KEY" json:"LOGGINGAPIKEY"`
	DATADOG_SOURCE      string `yaml:"DATADOG_SOURCE" json:"DATADOGSOURCE"`
	// Datadog Metrics
	DATADOG_METRICS bool   `yaml:"DATADOG_METRICS" json:"DATADOGMETRICS"`
	METRICS_URL     string `yaml:"METRICS_URL" json:"METRICSURL"`

	// SSL Settings
	/////////////////
	USE_SSL                   bool   `yaml:"USE_SSL" json:"USESSL"`
	SSL_CERT_PUBLIC_KEY_FILE  string `yaml:"SSL_CERT_PUBLIC_KEY_FILE" json:"SSLCERTPUBLICKEYFILE"`
	SSL_CERT_PRIVATE_KEY_FILE string `yaml:"SSL_CERT_PRIVATE_KEY_FILE" json:"SSLCERTPRIVATEKEYFILE"`

	// Azure Key Vault Settings
	////////////////////////////
	EXTRACT_CONFIG_FROM_AZURE_KEY_VAULT bool   `yaml:"EXTRACT_CONFIG_FROM_AZURE_KEY_VAULT" json:"EXTRACTCONFIGFROMAZUREKEYVAULT"`
	AZURE_KEY_VAULT_NAME                string `yaml:"AZURE_KEY_VAULT_NAME" json:"AZUREKEYVAULTNAME"`
	AZURE_MANAGED_IDENTITY_ID           string `yaml:"AZURE_MANAGED_IDENTITY_ID" json:"AZUREMANAGEDIDENTITYID"`

	// Allie Agent
	///////////////
	PRODUCTION_MODE            bool   `yaml:"PRODUCTION_MODE" json:"PRODUCTIONMODE"` // If true, the agent error messages will be generic
	AGENT_PORT                 string `yaml:"AGENT_PORT" json:"AGENTPORT"`
	WORKFLOW_API_KEY           string `yaml:"WORKFLOW_API_KEY" json:"WORKFLOWAPIKEY"`
	WORKFLOW_STORE_PATH        string `yaml:"WORKFLOW_STORE_PATH" json:"WORKFLOWSTOREPATH"`
	BINARY_STORE_PATH          string `yaml:"BINARY_STORE_PATH" json:"BINARYSTOREPATH"`
	NUMBER_OF_WORKFLOW_WORKERS int    `yaml:"NUMBER_OF_WORKFLOW_WORKERS" json:"NUMBEROFWORKFLOWWORKERS"`
	// External Function Endpoints
	EXTERNALFUNCTIONS_ENDPOINT string `yaml:"EXTERNALFUNCTIONS_ENDPOINT" json:"EXTERNALFUNCTIONSENDPOINT"`
	FLOWKIT_PYTHON_ENDPOINT    string `yaml:"FLOWKIT_PYTHON_ENDPOINT" json:"FLOWKITPYTHONENDPOINT"`
	FLOWKIT_PYTHON_API_KEY     string `yaml:"FLOWKIT_PYTHON_API_KEY" json:"FLOWKITPYTHONAPIKEY"`
	// Authentication & Authorization
	AUTH_TYPE                   string `yaml:"AUTH_TYPE" json:"AUTHTYPE"` // "ansys_gpt"
	AZURE_AD_AUTHENTICATION_URL string `yaml:"AZURE_AD_AUTHENTICATION_URL" json:"AZUREADAUTHENTICATIONURL"`
	ANSYS_AUTHORIZATION_URL     string `yaml:"ANSYS_AUTHORIZATION_URL" json:"ANSYSAUTHORIZATIONURL"`
	// Private Workflows
	LOAD_PRIVATE_WORKFLOWS   bool   `yaml:"LOAD_PRIVATE_WORKFLOWS" json:"LOADPRIVATEWORKFLOWS"`
	GITHUB_USER              string `yaml:"GITHUB_USER" json:"GITHUBUSER"`
	GITHUB_TOKEN             string `yaml:"GITHUB_TOKEN" json:"GITHUBTOKEN"`
	PRIVATE_WORKFLOWS_FOLDER string `yaml:"PRIVATE_WORKFLOWS_FOLDER" json:"PRIVATEWORKFLOWSFOLDER"`
	// Exec Settings
	EXEC_ENDPOINT                        string `yaml:"EXEC_ENDPOINT" json:"EXECENDPOINT"`
	EXEC_AGENT_API_KEY                   string `yaml:"EXEC_AGENT_API_KEY" json:"EXECAGENTAPIKEY"`
	MONGO_DB_FOR_MULTI_AGENT             bool   `yaml:"MONGO_DB_FOR_MULTI_AGENT" json:"MONGODBFORMULTIAGENT"`
	MONGO_DB_ENDPOINT                    string `yaml:"MONGO_DB_ENDPOINT" json:"MONGODBENDPOINT"`
	MILLISECONDS_MONGODB_UPDATE_INTERVAL int    `yaml:"MILLISECONDS_MONGODB_UPDATE_INTERVAL" json:"MILLISECONDSMONGODBUPDATEINTERVAL"`
	EXEC_FILE_STORE_PATH                 string `yaml:"EXEC_FILE_STORE_PATH" json:"EXECFILESTOREPATH"`

	// Allie Flowkit
	/////////////////
	EXTERNALFUNCTIONS_GRPC_PORT string `yaml:"EXTERNALFUNCTIONS_GRPC_PORT" json:"EXTERNALFUNCTIONSGRPCPORT"`
	FLOWKIT_API_KEY             string `yaml:"FLOWKIT_API_KEY" json:"FLOWKITAPIKEY"`
	// Allie Modules
	LLM_HANDLER_ENDPOINT  string `yaml:"LLM_HANDLER_ENDPOINT" json:"LLMHANDLERENDPOINT"`
	KNOWLEDGE_DB_ENDPOINT string `yaml:"KNOWLEDGE_DB_ENDPOINT" json:"KNOWLEDGEDBENDPOINT"`

	// Allie LLM
	/////////////
	WEBSERVER_PORT         string `yaml:"WEBSERVER_PORT" json:"WEBSERVERPORT"`
	MODELS_CONFIG_LOCATION string `yaml:"MODELS_CONFIG_LOCATION" json:"MODELSCONFIGLOCATION"`
	LLM_API_KEY            string `yaml:"LLM_API_KEY" json:"LLMAPIKEY"`

	// Allie DB
	///////////
	WEBSERVER_PORT_DB                  string   `yaml:"WEBSERVER_PORT_DB" json:"WEBSERVERPORTDB"`
	EMBEDDINGS_DIMENSIONS              int      `yaml:"EMBEDDINGS_DIMENSIONS" json:"EMBEDDINGSDIMENSIONS"`
	MILVUS_INDEX_TYPE                  string   `yaml:"MILVUS_INDEX_TYPE" json:"MILVUSINDEXTYPE"`
	MILVUS_METRIC_TYPE                 string   `yaml:"MILVUS_METRIC_TYPE" json:"MILVUSMETRICTYPE"` // cosine, l2 or ip
	MILVUS_HOST                        string   `yaml:"MILVUS_HOST" json:"MILVUSHOST"`
	MILVUS_PORT                        string   `yaml:"MILVUS_PORT" json:"MILVUSPORT"`
	NEO4J_DB                           bool     `yaml:"NEO4J_DB" json:"NEO4JDB"`
	NEO4J_URI                          string   `yaml:"NEO4J_URI" json:"NEO4JURI"`
	NEO4J_USERNAME                     string   `yaml:"NEO4J_USERNAME" json:"NEO4JUSERNAME"`
	NEO4J_PASSWORD                     string   `yaml:"NEO4J_PASSWORD" json:"NEO4JPASSWORD"`
	TEMP_COLLECTION_NAME               string   `yaml:"TEMP_COLLECTION_NAME" json:"TEMPCOLLECTIONNAME"`
	ELASTICSEARCH_HOST                 string   `yaml:"ELASTICSEARCH_HOST" json:"ELASTICSEARCHHOST"`
	ELASTICSEARCH_PORT                 string   `yaml:"ELASTICSEARCH_PORT" json:"ELASTICSEARCHPORT"`
	ELASTICSEARCH_USERNAME             string   `yaml:"ELASTICSEARCH_USERNAME" json:"ELASTICSEARCHUSERNAME"`
	ELASTICSEARCH_PASSWORD             string   `yaml:"ELASTICSEARCH_PASSWORD" json:"ELASTICSEARCHPASSWORD"`
	ELASTICSEARCH_INDEX_TYPE           string   `yaml:"ELASTICSEARCH_INDEX_TYPE" json:"ELASTICSEARCHINDEXTYPE"`                     // cosineSimilarity or dotProduct
	ELASTICSEARCH_TRUSTED_CERTIFICATES []string `yaml:"ELASTICSEARCH_TRUSTED_CERTIFICATES" json:"ELASTICSEARCHTRUSTEDCERTIFICATES"` // list of paths to trusted certificates
	ELASTICSEARCH_INSECURE_CONNECTION  bool     `yaml:"ELASTICSEARCH_INSECURE_CONNECTION" json:"ELASTICSEARCHINSECURECONNECTION"`
	DATABASE_TYPE                      string   `yaml:"DATABASE_TYPE" json:"DATABASETYPE"` // milvus or elasticsearch

	// Allie Exec
	//////////////
	EXEC_ID             string `yaml:"EXEC_ID" json:"EXECID"`
	WEBSERVER_PORT_EXEC string `yaml:"WEBSERVER_PORT_EXEC" json:"WEBSERVERPORTEXEC"`
	EXEC_API_KEY        string `yaml:"EXEC_API_KEY" json:"EXECAPIKEY"`
	// Python executable name
	PYTHON_EXECUTABLE string `yaml:"PYTHON_EXECUTABLE" json:"PYTHONEXECUTABLE"`
	BASH_EXECUTABLE   string `yaml:"BASH_EXECUTABLE" json:"BASHEXECUTABLE"`
	// File transfer
	WATCH_FOLDER_PATH              string `yaml:"WATCH_FOLDER_PATH" json:"WATCHFOLDERPATH"`
	MILLISECONDS_SINCE_LAST_CHANGE int    `yaml:"MILLISECONDS_SINCE_LAST_CHANGE" json:"MILLISECONDSSINCELASTCHANGE"`
	// Agent connection
	AGENT_ENDPOINT string `yaml:"AGENT_ENDPOINT" json:"AGENTENDPOINT"`

	// Workflow Store
	WORKFLOW_CONFIG_VARIABLES map[string]string `yaml:"WORKFLOW_CONFIG_VARIABLES" json:"WORKFLOWCONFIGVARIABLES"`
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
