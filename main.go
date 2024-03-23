package main

import (
	"strings"
)

func DeleteFields(data interface{}, fields []string) interface{} {
	switch t := data.(type) {
	case map[string]interface{}:
		for k, v := range t {
			if ContainsAny(k, fields) { // Check if the field name is in the slice
				delete(t, k)
			} else {
				newValue := DeleteFields(v, fields)
				t[k] = newValue
			}
		}
	case []interface{}:
		for i, v := range t {
			newValue := DeleteFields(v, fields)
			t[i] = newValue
		}
	}
	return data
}

// Helper function to check if a string is present in any part of a slice element
func ContainsAny(s string, slice []string) bool {
	for _, item := range slice {
		if strings.Contains(s, item) { // Use strings.Contains for efficient matching
			return true
		}
	}
	return false
}

func main() {
}
