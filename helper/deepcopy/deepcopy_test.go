package deepcopy

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	original := map[string]interface{}{
		"key1": "value1",
		"key2": 42,
		"key3": []int{1, 2, 3},
		"key4": map[string]interface{}{
			"nestedKey1": "nestedValue1",
		},
	}

	copy := Map(original)

	// Check if the copy is equal to the original
	if !reflect.DeepEqual(original, copy) {
		t.Errorf("Expected copy to be %v, but got %v", original, copy)
	}

	// Modify the copy and check if the original is unchanged
	copy["key1"] = "newValue1"
	copy["key3"].([]int)[0] = 99
	copy["key4"].(map[string]interface{})["nestedKey1"] = "newNestedValue1"

	if original["key1"] == "newValue1" {
		t.Errorf("Expected original key1 to be 'value1', but got %v", original["key1"])
	}

	if original["key3"].([]int)[0] == 99 {
		t.Errorf("Expected original key3[0] to be 1, but got %v", original["key3"].([]int)[0])
	}

	if original["key4"].(map[string]interface{})["nestedKey1"] == "newNestedValue1" {
		t.Errorf("Expected original nestedKey1 to be 'nestedValue1', but got %v", original["key4"].(map[string]interface{})["nestedKey1"])
	}
}
