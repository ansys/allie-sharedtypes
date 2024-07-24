package typeconverters

import (
	"strings"
)

// JSONToGo converts a JSON data type to a Go data type
func JSONToGo(jsonType string) string {
	if strings.HasPrefix(jsonType, "array<") && strings.HasSuffix(jsonType, ">") {
		elementType := jsonType[6 : len(jsonType)-1]
		return "[]" + JSONToGo(elementType)
	}

	switch {
	case jsonType == "string":
		return "string"
	case jsonType == "string(binary)":
		return "[]byte"
	case jsonType == "number":
		return "float64" // Default to float64 for general numeric values
	case jsonType == "integer":
		return "int"
	case jsonType == "boolean":
		return "bool"
	case jsonType == "object":
		return "map[string]interface{}"
	case jsonType == "null":
		return "interface{}"
	default:
		return "interface{}"
	}
}

// GoToJSON converts a Go data type to a JSON data type
func GoToJSON(goType string) string {
	if strings.HasPrefix(goType, "[]") && goType != "[]byte" {
		elementType := goType[2:]
		return "array<" + GoToJSON(elementType) + ">"
	}

	switch goType {
	case "string":
		return "string"
	case "float32", "float64":
		return "number"
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64":
		return "integer"
	case "bool":
		return "boolean"
	case "map[string]interface{}":
		return "object"
	case "interface{}":
		return "null"
	case "[]byte":
		return "string(binary)"
	default:
		return "string"
	}
}
