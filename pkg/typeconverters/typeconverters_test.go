package typeconverters

import (
	"testing"
)

func TestJSONToGo(t *testing.T) {
	tests := []struct {
		jsonType string
		goType   string
	}{
		{"string", "string"},
		{"string(binary)", "[]byte"},
		{"number", "float64"},
		{"integer", "int"},
		{"boolean", "bool"},
		{"object", "map[string]interface{}"},
		{"null", "interface{}"},
		{"array<string>", "[]string"},
		{"array<number>", "[]float64"},
		{"array<object>", "[]map[string]interface{}"},
		{"unknown", "interface{}"},
	}

	for _, test := range tests {
		result := JSONToGo(test.jsonType)
		if result != test.goType {
			t.Errorf("JSONToGo(%q) = %q; want %q", test.jsonType, result, test.goType)
		}
	}
}

func TestGoToJSON(t *testing.T) {
	tests := []struct {
		goType   string
		jsonType string
	}{
		{"string", "string"},
		{"[]byte", "string(binary)"},
		{"float32", "number"},
		{"float64", "number"},
		{"int", "integer"},
		{"int8", "integer"},
		{"int16", "integer"},
		{"int32", "integer"},
		{"int64", "integer"},
		{"uint", "integer"},
		{"uint8", "integer"},
		{"uint16", "integer"},
		{"uint32", "integer"},
		{"uint64", "integer"},
		{"bool", "boolean"},
		{"map[string]interface{}", "object"},
		{"interface{}", "null"},
		{"[]string", "array<string>"},
		{"[]float64", "array<number>"},
		{"[]int", "array<integer>"},
		{"unknown", "string"},
	}

	for _, test := range tests {
		result := GoToJSON(test.goType)
		if result != test.jsonType {
			t.Errorf("GoToJSON(%q) = %q; want %q", test.goType, result, test.jsonType)
		}
	}
}
