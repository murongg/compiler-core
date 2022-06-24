package main

import (
	"testing"
)

func TestToLowerCase(t *testing.T) {
	t.Log("test ToLowerCase")
	input := "aSdb"
	result := ToLowerCase(input)
	expected := "asdb"
	t.Run("ToLowerCase", func(t *testing.T) {
		if result != expected {
			t.Errorf("failed: input %s, expected %s, actual %s", input, expected, result)
		}
	})
}

func TestStringSlice(t *testing.T) {
	t.Log("test StringSlice")
	input := "seewrwe"

	t.Run("Slice 1", func(t *testing.T) {
		result := StringSlice(input, 2, 3)
		expected := "e"
		if result != expected {
			t.Errorf("failed: input %s, expected %s, actual %s", input, expected, result)
		}
	})
	t.Run("Slice 2", func(t *testing.T) {
		result := StringSlice(input, 2, 4)
		expected := "ew"
		if result != expected {
			t.Errorf("failed: input %s, expected %s, actual %s", input, expected, result)
		}
	})
}
