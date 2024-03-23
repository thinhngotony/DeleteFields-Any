package main

import (
	"reflect"
	"testing"
)

// deleteFieldsTestCase defines a test case for the DeleteFields function.
type deleteFieldsTestCase struct {
	name     string
	data     interface{}
	fields   []string
	expected interface{}
}

// deleteFieldsTests contains all test cases for the DeleteFields function.
var deleteFieldsTests = []deleteFieldsTestCase{
	{
		name: "Delete single field from map",
		data: map[string]interface{}{
			"name": "John Doe",
			"age":  30,
		},
		fields: []string{"age"},
		expected: map[string]interface{}{
			"name": "John Doe",
		},
	},
	{
		name: "Delete nested field from map",
		data: map[string]interface{}{
			"name": "John Doe",
			"age":  30,
			"address": map[string]interface{}{
				"city": "New York",
				"zip":  "10001",
			},
		},
		fields: []string{"zip"},
		expected: map[string]interface{}{
			"name": "John Doe",
			"age":  30,
			"address": map[string]interface{}{
				"city": "New York",
			},
		},
	},
	{
		name: "Delete field from nested map",
		data: map[string]interface{}{
			"name": "John Doe",
			"age":  30,
			"address": map[string]interface{}{
				"city": "New York",
				"zip":  "10001",
			},
		},
		fields: []string{"city"},
		expected: map[string]interface{}{
			"name": "John Doe",
			"age":  30,
			"address": map[string]interface{}{
				"zip": "10001",
			},
		},
	},
	// Add more test cases as needed.
}

func TestDeleteFields(t *testing.T) {
	for _, tc := range deleteFieldsTests {
		t.Run(tc.name, func(t *testing.T) {
			result := DeleteFields(tc.data, tc.fields)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}
