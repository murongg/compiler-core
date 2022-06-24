package main

import (
	"testing"
)

var parse = new(Parse)

func TestStartsWithEndTagOpen(t *testing.T) {
	t.Log("test ToLowerCase")
	input := "<div></div>"
	t.Run("ToLowerCase div", func(t *testing.T) {
		result := parse.StartsWithEndTagOpen(input, "div")
		if result {
			t.Errorf("failed: input %s, expected %t, actual %t", input, true, result)
		}
	})

	t.Run("ToLowerCase span", func(t *testing.T) {
		result := parse.StartsWithEndTagOpen(input, "span")
		if result != false {
			t.Errorf("failed:input %s, expected %t, actual %t", input, false, result)
		}
	})
}

func TestIsEnd(t *testing.T) {
	t.Log("test IsEnd")
	context := ParseContext{
		Source: "",
	}

	ancestors := []*Node{{
		Tag: "span",
	}}
	t.Run("IsEnd have close tag", func(t *testing.T) {
		input := "<div><span></span></div>"
		context.Source = input
		result := parse.IsEnd(&context, ancestors)
		if result {
			t.Errorf("failed: input %s, expected %t, actual %t", input, true, result)
		}
	})

	t.Run("IsEnd no close tag", func(t *testing.T) {
		input := "<div><span></div>"
		context.Source = input
		result := parse.IsEnd(&context, ancestors)
		if result != false {
			t.Errorf("failed: input %s, expected %t, actual %t", input, false, result)
		}
	})
}
