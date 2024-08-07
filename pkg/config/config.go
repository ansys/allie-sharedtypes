package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

//////////////////////////////////////////
// Read Config variables from Config file
//////////////////////////////////////////

// InitGlobalConfigFromFile reads the configuration file and initializes the Config object.
//
// Parameters:
//   - fileName: The name of the configuration file.
//   - requiredProperties: The list of required properties.
//   - optionalDefaultValues: The map of optional properties and their default values.
func InitGlobalConfigFromFile(fileName string, requiredProperties []string, optionalDefaultValues map[string]interface{}) {
	var config Config
	configResult := readYaml(fileName, config)
	// Validate mandatory config properties
	validateConfig(configResult, requiredProperties)

	// Assign to global config
	GlobalConfig = &configResult

	// Set optional properties if missing
	defineOptionalProperties(GlobalConfig, optionalDefaultValues)
}

// readYaml reads the yaml specified in `fileName` parameter and saves it to `config_struct`
//
// Parameters:
//   - fileName: The name of the configuration file.
//   - configStruct: Struct with the parameters of the YAML to read.
func readYaml(fileName string, configStruct Config) Config {
	// Read the YAML file into a byte slice
	data, err := os.ReadFile(fileName)
	if err != nil {
		message := "config.yaml file is missing from directory or not accessible"
		pan := writeStringToFile(message)
		if pan != nil {
			panic(pan)
		}
		pan2 := writeInterfaceToFile(err)
		if pan2 != nil {
			panic(pan2)
		}
		panic(err)
	}

	// Unmarshal the YAML data into the config object
	err = yaml.Unmarshal(data, &configStruct)
	if err != nil {
		// Create a new Config struct with field names and types
		configStruct := reflect.TypeOf(configStruct)
		var fieldList []string
		for i := 0; i < configStruct.NumField(); i++ {
			field := configStruct.Field(i)
			fieldList = append(fieldList, fmt.Sprintf("%q: %s", field.Name, field.Type.String()))
		}

		// Define error message
		message := fileName + " contains incorrect content. The allowed values are as follows: {"
		message += strings.Join(fieldList, ",")
		message = message + "}"

		// Write message and error to error file
		pan := writeStringToFile(message)
		if pan != nil {
			panic(pan)
		}
		pan2 := writeInterfaceToFile(err)
		if pan2 != nil {
			panic(pan2)
		}
		panic(err)
	}
	return configStruct
}

// validateConfig checks for mandatory entries in the configuration and validates chosen models.
//
// Parameters:
//   - config: The configuration object to validate.
//   - requiredProperties: The list of required properties.
func validateConfig(config Config, requiredProperties []string) {
	// Check if all mandatory properties are present
	configValue := reflect.ValueOf(config)

	for _, property := range requiredProperties {
		field := configValue.FieldByName(property)

		if !field.IsValid() || field.IsZero() {
			message := fmt.Sprintf("config.yaml is missing mandatory property '%v': ", property)

			// Write message to error file
			pan := writeInterfaceToFile(message)
			if pan != nil {
				panic(pan)
			}
			panic(message)
		}
	}
}

// defineOptionalProperties sets optional properties for the configuration.
//
// Parameters:
//   - config: The configuration object to validate.
//   - optionalDefaultValues: The map of optional properties and their default values.
func defineOptionalProperties(config *Config, optionalDefaultValues map[string]interface{}) {
	defer func() {
		r := recover()
		if r != nil {
			// Write message to error file
			message := fmt.Sprintf("Error in defineOptionalProperties: %v", r)
			pan := writeStringToFile(message)
			if pan != nil {
				panic(pan)
			}
		}
	}()

	// Iterate over the optional default values
	for key, defaultValue := range optionalDefaultValues {
		// Use reflection to check if the field exists and is set to its zero value
		fieldValue := reflect.ValueOf(config).Elem().FieldByName(key)
		if fieldValue.IsValid() && isZeroValue(fieldValue) {
			// Set the default value using reflection
			if reflect.TypeOf(defaultValue) == fieldValue.Type() {
				fieldValue.Set(reflect.ValueOf(defaultValue))
			} else {
				// Handle type mismatch
				message := fmt.Sprintf("Type mismatch for key '%s': expected %v, got %v", key, fieldValue.Type(), reflect.TypeOf(defaultValue))
				pan := writeStringToFile(message)
				if pan != nil {
					panic(pan)
				}
			}
		}
	}
}

// isZeroValue checks if a reflect.Value is zero for its type.
//
// Parameters:
//   - v: The reflect.Value to check.
//
// Returns:
//   - bool: True if the value is zero, false otherwise.
func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

/////////////////////////////////////////
// Create or update Config file from CLI
/////////////////////////////////////////

// CreateUpdateConfigFileFromCLI reads and updates the configuration file based on command-line arguments.
//
// Parameters:
//   - fileName: The name of the configuration file.
func CreateUpdateConfigFileFromCLI(fileName string) {
	// Checking for any command-line arguments
	if len(os.Args) == 1 {
		fmt.Println("No command line options given; full config will be retrieved from existing config.yaml file")
		return
	}

	// Create a new config to store command-line options
	cliConfig := Config{}

	// Use reflection to create flags for each field in Config
	createFlags(reflect.ValueOf(&cliConfig).Elem(), "")

	// Parse the flags
	flag.Parse()

	// Checking if config.yaml file already exists
	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		// If it doesn't exist, create a new one
		fmt.Println("config.yaml file does not exist. Creating a new one...")

		file, _ := yaml.Marshal(cliConfig)
		_ = os.WriteFile(fileName, file, 0644)
	} else {
		// If it does exist, open and append to it
		fmt.Println("config.yaml file exists. Appending command line options...")

		file, _ := os.ReadFile(fileName)
		config := Config{}
		err := yaml.Unmarshal(file, &config)
		if err != nil {
			message := "error in yaml.Unmarshal:"
			pan := writeStringToFile(message)
			if pan != nil {
				panic(pan)
			}
			pan2 := writeInterfaceToFile(err)
			if pan2 != nil {
				panic(pan2)
			}
			panic(err)
		}

		// Use reflection to update fields that were set in the command line
		valCli := reflect.ValueOf(&cliConfig).Elem()
		valConfig := reflect.ValueOf(&config).Elem()
		t := reflect.TypeOf(cliConfig)

		for i := 0; i < t.NumField(); i++ {
			cliField := valCli.Field(i)
			configField := valConfig.Field(i)
			if !reflect.DeepEqual(cliField.Interface(), reflect.Zero(cliField.Type()).Interface()) {
				configField.Set(cliField)
			}
		}

		// Write back to the file
		file, _ = yaml.Marshal(config)
		_ = os.WriteFile(fileName, file, 0644)
	}
}

// CreateFlags initializes command-line flags for configuration.
//
// Parameters:
//   - val: The value to create flags for.
//   - prefix: The prefix to use for the flags.
func createFlags(val reflect.Value, prefix string) {
	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := strings.ToUpper(prefix + field.Name)
		switch field.Type.Kind() {
		case reflect.String:
			flag.StringVar(val.Field(i).Addr().Interface().(*string), name, "", "config option")
		case reflect.Int:
			flag.IntVar(val.Field(i).Addr().Interface().(*int), name, 0, "config option")
		case reflect.Bool:
			flag.BoolVar(val.Field(i).Addr().Interface().(*bool), name, false, "config option")
		case reflect.Slice:
			if field.Type.Elem().Kind() == reflect.String {
				flag.Var((*flagStringSlice)(val.Field(i).Addr().Interface().(*[]string)), name, "config option")
			}
		case reflect.Struct:
			createFlags(val.Field(i), name+"_")
		}
	}
}

///////////////////////
// Helper Functions
///////////////////////

// GetGlobalConfigAsJSON returns the global configuration as a JSON string.
//
// Returns:
//   - string: The global configuration as a JSON string.
func GetGlobalConfigAsJSON() string {
	jsonData, err := json.Marshal(GlobalConfig)
	if err != nil {
		return ""
	}
	return string(jsonData)
}

///////////////////////
// Error file creator
///////////////////////

// writeInterfaceToFile writes interface data to an error log file.
//
// Parameters:
//   - data: The data to write to the file.
//
// Returns:
//   - error: An error if there was an issue writing to the file.
func writeInterfaceToFile(data interface{}) error {
	var file *os.File
	var err error

	// Get file name
	filename := "error.log"

	// Create file
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// If the file does not exist, create a new file.
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
	} else {
		// If the file already exists, open it in append mode.
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	// Write to file
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Add time
	timestamp := timeToString(time.Now())

	// Write to file
	line := fmt.Sprintf("%s: %s\n", timestamp, string(jsonData))
	_, err = file.Write([]byte(line))
	if err != nil {
		return err
	}

	return nil
}

// writeStringToFile writes string data to an error log file.
//
// Parameters:
//   - data: The data to write to the file.
//
// Returns:
//   - error: An error if there was an issue writing to the file.
func writeStringToFile(data string) error {
	var file *os.File
	var err error

	// Get file name
	filename := "error.log"

	// Add time
	timestamp := timeToString(time.Now())

	// Change string
	data = timestamp + ": " + data

	if _, err = os.Stat(filename); os.IsNotExist(err) {
		// File does not exist, create a new file
		file, err = os.Create(filename)
		if err != nil {
			return err
		}
	} else {
		// File exists, open it for appending
		file, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
	}
	defer file.Close()

	// Append data to file with a new line
	_, err = fmt.Fprintln(file, data)
	return err
}

// timeToString converts a time.Time value to a formatted string.
//
// Parameters:
//   - t: The time.Time value to convert.
//
// Returns:
//   - string: The formatted string.
func timeToString(t time.Time) string {
	layout := "2006-01-02 15:04:05.000"
	return t.Format(layout)
}
