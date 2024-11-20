package structs

// FunctionDefinition is a struct that contains the id, name, description, package, inputs and outputs of a function
type FunctionDefinition struct {
	Name        string           `json:"name" yaml:"name"`
	DisplayName string           `json:"display_name" yaml:"display_name"`
	Description string           `json:"description" yaml:"description"`
	Category    string           `json:"category" yaml:"category"` // "data_extraction", "generic", "knowledge_db", "llm_handler", "ansys_gpt"
	Type        string           `json:"type" yaml:"type"`         // "go", "python"
	Path        string           `json:"path" yaml:"path"`         // only for python functions
	Inputs      []FunctionInput  `json:"inputs" yaml:"inputs"`
	Outputs     []FunctionOutput `json:"outputs" yaml:"outputs"`
}

// FlowKitPythonFunction is a struct that contains the name, path, description, inputs, outputs and definitions of a FlowKit-Python function
type FlowKitPythonFunction struct {
	Name        string           `json:"name"`
	Path        string           `json:"path"`
	Description string           `json:"description"`
	Category    string           `json:"category"`
	DisplayName string           `json:"display_name"`
	Inputs      []FunctionInput  `json:"inputs"`
	Outputs     []FunctionOutput `json:"outputs"`
	Definitions interface{}      `json:"definitions"`
}

// FunctionDefinitionShort is a struct that contains the id, name and description of a function
type FunctionDefinitionShort struct {
	Id          string `json:"id" yaml:"id"` // Unique identifier
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
}

// FunctionInput is a struct that contains the name, type, go type and options of a function input
type FunctionInput struct {
	Name    string   `json:"name" yaml:"name"`
	Type    string   `json:"type" yaml:"type"` // string, number, boolean, json
	GoType  string   `json:"go_type" yaml:"go_type"`
	Options []string `json:"options" yaml:"options"` // only applicable if not empty
}

// FunctionOutput is a struct that contains the name, type and go type of a function output
type FunctionOutput struct {
	Name   string `json:"name" yaml:"name"`
	Type   string `json:"type" yaml:"type"` // string, number, boolean, json
	GoType string `json:"go_type" yaml:"go_type"`
}

// FilledInputOutput is a struct that contains the name, go type and value of a filled input/output
type FilledInputOutput struct {
	Name   string      `json:"name" yaml:"name"`
	GoType string      `json:"go_type" yaml:"go_type"`
	Value  interface{} `json:"value" yaml:"value"`
}
