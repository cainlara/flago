package flago

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// GetArgsMap parses the command-line arguments passed to the program
// and returns a map with each flag (without leading dashes) mapped
// to its corresponding value.
//
// It expects arguments to follow the format: -key value.
// For example, running the program as:
//
//	myapp -name John -age 30
//
// would result in a map: {"name": "John", "age": "30"}.
//
// If no arguments are passed or if the input is invalid,
// it returns an error.
func GetArgsMap() (map[string]any, error) {
	args := os.Args

	if len(args) == 0 {
		return map[string]any{}, errEmptyInput
	}

	return getPurgedArgsMap(args)
}

// GetArgsStruct populates a struct with command-line arguments.
//
// This function parses the command-line arguments provided via os.Args,
// maps them into key-value pairs, and then assigns those values to the
// corresponding fields in a user-defined struct. The keys from the arguments
// are automatically converted to CamelCase to match Go struct field names.
//
// The target struct must be passed as a pointer. Each argument key is matched
// against struct field names (case-insensitive, via CamelCase conversion).
//
// Parameters:
//   - outputPointer: A pointer to a struct that will be populated with the parsed arguments.
//   - ignoreUnknow: If true, unknown fields (i.e., fields in the arguments that
//     do not match any struct fields) are ignored. If false, an error is returned
//     when an unknown field is encountered.
//
// Returns:
//   - error: An error is returned if:
//   - os.Args is empty
//   - the outputPointer is not a pointer to a struct
//   - argument values cannot be assigned to struct fields due to type mismatch
//   - a required field is not found in the struct and ignoreUnknow is false
//
// Example:
//
//	type Config struct {
//	    Port int
//	    Debug bool
//	}
//
//	var cfg Config
//	err := GetArgsStruct(&cfg, true)
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// Notes:
//   - Argument parsing is delegated to internal helpers, including getPurgedArgsMap
//     which transforms the raw arguments into a usable map[string]any format.
//   - Keys in the map are converted to CamelCase using the ToCamelCase function
//     before being matched to struct fields.
func GetArgsStruct(outputPointer any, ignoreUnknow bool) error {
	args := os.Args

	if len(args) == 0 {
		return errEmptyInput
	}

	mapedArgs, err := getPurgedArgsMap(args)
	if err != nil {
		return err
	}

	return getArgsStructFromArgs(mapedArgs, outputPointer, ignoreUnknow)
}

func getPurgedArgsMap(args []string) (map[string]any, error) {
	if len(args) == 0 {
		return map[string]any{}, errEmptyInput
	}

	purgedArgs := args[1:]

	if len(purgedArgs) == 0 {
		return map[string]any{}, errEmptyInput
	}

	if len(purgedArgs)%2 != 0 {
		purgedArgs = append(purgedArgs, "")
	}

	mappedArgs := make(map[string]any, len(purgedArgs)/2)

	for index := 0; index < len(purgedArgs)-1; index += 2 {
		argName := strings.TrimPrefix(purgedArgs[index], dash)

		mappedArgs[argName] = purgedArgs[index+1]
	}

	return mappedArgs, nil
}

func getArgsStructFromArgs(input map[string]any, outputPointer any, ignoreUnknow bool) error {
	val := reflect.ValueOf(outputPointer)

	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return errors.New("outputPointer argument must be a pointer to a struct")
	}

	val = val.Elem()

	for key, value := range input {
		fieldName := ToCamelCase(key)
		field := val.FieldByName(fieldName)

		if !field.IsValid() {
			if !ignoreUnknow {
				return fmt.Errorf("no such field [%s] in provided output struct", fieldName)
			}
			continue
		}
		if !field.CanSet() {
			return fmt.Errorf("cannot set field %s", fieldName)
		}

		fieldType := field.Type()
		valType := reflect.TypeOf(value)

		if valType == nil || !valType.AssignableTo(fieldType) {
			return fmt.Errorf("cannot assign value of type %v to field %s (type %v)", valType, fieldName, fieldType)
		}

		field.Set(reflect.ValueOf(value))
	}

	return nil
}
