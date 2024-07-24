package typeconverters

import (
	"testing"
)

func TestJSONToGo(t *testing.T) {
	tests := []struct {
		jsonType  string
		goType    string
		expectErr bool
	}{
		{"string", "string", false},
		{"string(binary)", "[]byte", false},
		{"number", "float64", false},
		{"integer", "int", false},
		{"boolean", "bool", false},
		{"array<string>", "[]string", false},
		{"array<number>", "[]float64", false},
		{"array<integer>", "[]int", false},
		{"array<boolean>", "[]bool", false},
		{"array<object>", "", true},
		{"unsupportedType", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.jsonType, func(t *testing.T) {
			got, err := JSONToGo(tt.jsonType)
			if (err != nil) != tt.expectErr {
				t.Errorf("JSONToGo(%s) error = %v, expectErr %v", tt.jsonType, err, tt.expectErr)
				return
			}
			if got != tt.goType {
				t.Errorf("JSONToGo(%s) = %v, want %v", tt.jsonType, got, tt.goType)
			}
		})
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
		{"interface{}", "object"},
		{"[]string", "array<string>"},
		{"[]float64", "array<number>"},
		{"[]int", "array<integer>"},
		{"unknown", "object"},
	}

	for _, test := range tests {
		result := GoToJSON(test.goType)
		if result != test.jsonType {
			t.Errorf("GoToJSON(%q) = %q; want %q", test.goType, result, test.jsonType)
		}
	}
}
