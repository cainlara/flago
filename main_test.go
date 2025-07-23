package flago

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

type TestMapScenario struct {
	Description string
	Input       []string
	Expected    map[string]any
	Err         error
}

var testCases_map = []TestMapScenario{
	{"Simple", []string{"go_file_name", "alice", "1", "bob", "2", "charlie", "3"}, map[string]any{"alice": "1", "bob": "2", "charlie": "3"}, nil},
	{"Incomplete", []string{"go_file_name", "alice", "1", "bob", "2", "charlie"}, map[string]any{"alice": "1", "bob": "2", "charlie": ""}, nil},
	{"Remove Dash", []string{"go_file_name", "-alice", "1", "bob", "2", "-charlie"}, map[string]any{"alice": "1", "bob": "2", "charlie": ""}, nil},
	{"Empty Error", []string{}, map[string]any{}, errEmptyInput},
}

func TestGetArgsMap(t *testing.T) {
	for _, testCase := range testCases_map {
		t.Run(testCase.Description, func(t *testing.T) {
			actual, err := getPurgedArgsMap(testCase.Input)

			if err != nil && !isExpectedError(err) {
				t.Errorf("Arguments parsing failed: %v\n", err)
			}

			equals := reflect.DeepEqual(testCase.Expected, actual)

			if !equals {
				t.Error("Output comparison failed")
			}
		})
	}
}

func TestGetArgsStructOk(t *testing.T) {
	t.Run("Struct OK", func(t *testing.T) {
		type config struct {
			Source string
			Size   int
			Skip   bool
		}

		expected := config{"path/to/file", 100, true}
		var actual config

		input := map[string]any{"source": "path/to/file", "size": 100, "skip": true}
		err := getArgsStructFromArgs(input, &actual, false)
		if err != nil && !isExpectedError(err) {
			t.Errorf("Unexpected error: %v", err)
		}

		equals := reflect.DeepEqual(expected, actual)
		if !equals {
			t.Error("Output comparison failed")
		}
	})
}

func TestGetArgsOkIgnoreFields(t *testing.T) {
	t.Run("Ignore Unknown Fields", func(t *testing.T) {
		type config struct {
			Source string
			Size   int
			Limit  int
		}

		var actual config

		input := map[string]any{"source": "path/to/file", "size": 100, "skip": true}
		err := getArgsStructFromArgs(input, &actual, true)
		if err != nil && !isExpectedError(err) {
			t.Errorf("Unexpected error: %v", err)
		}

		if actual.Source == "" || actual.Size == 0 {
			t.Error("Expected valid value fields contain zero-value")
		}

		if actual.Limit != 0 {
			t.Error("Expected zero-value field contain valid value")
		}
	})
}

func TestGetArgsStructFails(t *testing.T) {
	t.Run("Fails With Unknown Field", func(t *testing.T) {
		type config struct {
			Source string
			Size   int
			Limit  int
		}

		var actual config

		input := map[string]any{"source": "path/to/file", "size": 100, "skip": true}
		err := getArgsStructFromArgs(input, &actual, false)
		if err != nil && !isExpectedError(err) {
			t.Errorf("Unexpected error: %v", err)
		}

		actualError := errors.New(strings.ReplaceAll(err.Error(), "Skip", ""))

		if actualError.Error() != errUnknownField.Error() {
			t.Errorf("Unexpected error: %v", err)
		}
	})
}

func isExpectedError(err error) bool {
	return err == errEmptyInput || strings.HasPrefix(err.Error(), "no such field")
}
