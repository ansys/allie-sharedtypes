package typeconverters

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/ansys/allie-sharedtypes/pkg/sharedtypes"
)

// JSONToGo converts a JSON data type to a Go data type.
//
// Parameters:
//
//	jsonType: The JSON data type to convert.
//
// Returns:
//
//	string: The Go data type.
//	error: An error if the JSON data type is not supported.
func JSONToGo(jsonType string) (string, error) {
	if strings.HasPrefix(jsonType, "array<") && strings.HasSuffix(jsonType, ">") {
		elementType := jsonType[6 : len(jsonType)-1]
		arrayType, err := JSONToGo(elementType)
		if err != nil {
			return "", err
		}

		return "[]" + arrayType, nil
	}

	switch {
	case jsonType == "string":
		return "string", nil
	case jsonType == "string(binary)":
		return "[]byte", nil
	case jsonType == "number":
		return "float64", nil // Default to float64 for general numeric values
	case jsonType == "integer":
		return "int", nil
	case jsonType == "boolean":
		return "bool", nil
	default:
		return "", fmt.Errorf("not supported JSON type: %s", jsonType)
	}
}

// GoToJSON converts a Go data type to a JSON data type.
//
// Parameters:
//
//	goType: The Go data type to convert.
//
// Returns:
//
//	string: The JSON data type.
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
	case "[]byte":
		return "string(binary)"
	default:
		return "object"
	}
}

// ConvertStringToGivenType converts a string to a given Go type.
//
// Parameters:
// - value: a string containing the value to convert
// - goType: a string containing the Go type to convert to
//
// Returns:
// - output: an interface containing the converted value
// - err: an error containing the error message
func ConvertStringToGivenType(value string, goType string) (output interface{}, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("panic occured in convertStringToGivenType: %v", r)
		}
	}()

	switch goType {
	case "string":
		return value, nil
	case "float32":
		if value == "" {
			value = "0"
		}
		return strconv.ParseFloat(value, 32)
	case "float64":
		if value == "" {
			value = "0"
		}
		return strconv.ParseFloat(value, 64)
	case "int":
		if value == "" {
			value = "0"
		}
		return strconv.Atoi(value)
	case "uint32":
		if value == "" {
			value = "0"
		}
		valueUint64, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return nil, err
		}
		return uint32(valueUint64), nil
	case "bool":
		if value == "" {
			value = "false"
		}
		return strconv.ParseBool(value)
	case "[]string":
		if value == "" {
			value = "[]"
		}
		output := []string{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "[]float32":
		if value == "" {
			value = "[]"
		}
		output := []float32{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "[]float64":
		if value == "" {
			value = "[]"
		}
		output := []float64{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "[]int":
		if value == "" {
			value = "[]"
		}
		output := []int{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "[]bool":
		if value == "" {
			value = "[]"
		}
		output := []bool{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "[][]float32":
		if value == "" {
			value = "[]"
		}
		output := [][]float32{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "map[string]string":
		if value == "" {
			value = "{}"
		}
		output := map[string]string{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "map[string]float64":
		if value == "" {
			value = "{}"
		}
		output := map[string]float64{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "map[string]int":
		if value == "" {
			value = "{}"
		}
		output := map[string]int{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "map[string]bool":
		if value == "" {
			value = "{}"
		}
		output := map[string]bool{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "map[string][]string":
		if value == "" {
			value = "{}"
		}
		output := map[string][]string{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "DbArrayFilter":
		if value == "" {
			value = "{}"
		}
		output := sharedtypes.DbArrayFilter{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "DbFilters":
		if value == "" {
			value = "{}"
		}
		output := sharedtypes.DbFilters{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "[]DbJsonFilter":
		if value == "" {
			value = "[]"
		}
		output := []sharedtypes.DbJsonFilter{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "[]DbResponse":
		if value == "" {
			value = "[]"
		}
		output := []sharedtypes.DbResponse{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "[]HistoricMessage":
		if value == "" {
			value = "[]"
		}
		output := []sharedtypes.HistoricMessage{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	case "*chan string":
		var output *chan string
		output = nil
		return output, nil

	case "[]AnsysGPTDefaultFields":
		if value == "" {
			value = "[]"
		}
		output := []sharedtypes.AnsysGPTDefaultFields{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil

	case "[]ACSSearchResponse":
		if value == "" {
			value = "[]"
		}
		output := []sharedtypes.ACSSearchResponse{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil

	case "[]AnsysGPTCitation":
		if value == "" {
			value = "[]"
		}
		output := []sharedtypes.AnsysGPTCitation{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil

	case "[]DataExtractionDocumentData":
		if value == "" {
			value = "[]"
		}
		output := []sharedtypes.DataExtractionDocumentData{}
		err := json.Unmarshal([]byte(value), &output)
		if err != nil {
			return nil, err
		}
		return output, nil
	}

	return nil, fmt.Errorf("unsupported GoType: '%s'", goType)
}

// ConvertGivenTypeToString converts a given Go type to a string.
//
// Parameters:
// - value: an interface containing the value to convert
// - goType: a string containing the Go type to convert from
//
// Returns:
// - string: a string containing the converted value
// - err: an error containing the error message
func ConvertGivenTypeToString(value interface{}, goType string) (output string, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("panic occured in ConvertGivenTypeToString: %v", r)
		}
	}()

	switch goType {
	case "string":
		return value.(string), nil
	case "float32":
		return strconv.FormatFloat(float64(value.(float32)), 'f', -1, 32), nil
	case "float64":
		return strconv.FormatFloat(value.(float64), 'f', -1, 64), nil
	case "int":
		return strconv.Itoa(value.(int)), nil
	case "uint32":
		return strconv.FormatUint(uint64(value.(uint32)), 10), nil
	case "bool":
		return strconv.FormatBool(value.(bool)), nil
	case "[]string":
		output, err := json.Marshal(value.([]string))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[]float32":
		output, err := json.Marshal(value.([]float32))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[]float64":
		output, err := json.Marshal(value.([]float64))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[]int":
		output, err := json.Marshal(value.([]int))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[]bool":
		output, err := json.Marshal(value.([]bool))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[][]float32":
		output, err := json.Marshal(value.([][]float32))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "map[string]string":
		output, err := json.Marshal(value.(map[string]string))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "map[string]float64":
		output, err := json.Marshal(value.(map[string]float64))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "map[string]int":
		output, err := json.Marshal(value.(map[string]int))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "map[string]bool":
		output, err := json.Marshal(value.(map[string]bool))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "map[string][]string":
		output, err := json.Marshal(value.(map[string][]string))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "DbArrayFilter":
		output, err := json.Marshal(value.(sharedtypes.DbArrayFilter))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "DbFilters":
		output, err := json.Marshal(value.(sharedtypes.DbFilters))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[]DbJsonFilter":
		output, err := json.Marshal(value.([]sharedtypes.DbJsonFilter))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[]DbResponse":
		output, err := json.Marshal(value.([]sharedtypes.DbResponse))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[]HistoricMessage":
		output, err := json.Marshal(value.([]sharedtypes.HistoricMessage))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "*chan string":
		return "", nil
	case "[]AnsysGPTDefaultFields":
		output, err := json.Marshal(value.([]sharedtypes.AnsysGPTDefaultFields))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[]ACSSearchResponse":
		output, err := json.Marshal(value.([]sharedtypes.ACSSearchResponse))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[]AnsysGPTCitation":
		output, err := json.Marshal(value.([]sharedtypes.AnsysGPTCitation))
		if err != nil {
			return "", err
		}
		return string(output), nil
	case "[]DataExtractionDocumentData":
		output, err := json.Marshal(value.([]sharedtypes.DataExtractionDocumentData))
		if err != nil {
			return "", err
		}
		return string(output), nil
	}

	return "", fmt.Errorf("unsupported GoType: '%s'", goType)
}

// DeepCopy deep copies the source interface to the destination interface.
//
// Parameters:
// - src: an interface containing the source
// - dst: an interface containing the destination
//
// Returns:
// - err: an error containing the error message
func DeepCopy(src, dst interface{}) (err error) {
	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("panic occured in DeepCopy: %v", r)
		}
	}()

	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, dst)
}
