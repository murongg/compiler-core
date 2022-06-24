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

func TestAdvanceBy(t *testing.T) {
	t.Log("test advanceBy")

	context := ParseContext{}
	t.Run("test 1", func(t *testing.T) {
		input := "test a b c d"
		expected := "st a b c "
		context.Source = input
		parse.advanceBy(&context, 2)
		if context.Source != expected {
			t.Errorf("failed: input %s, expected %s, actual %s", input, expected, context.Source)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		input := "Hello World!"
		expected := "lo World"
		context.Source = input
		parse.advanceBy(&context, 3)
		if context.Source != expected {
			t.Errorf("failed: input %s, expected %s, actual %s", input, expected, context.Source)
		}
	})

}

func TestParseTextData(t *testing.T) {
	t.Log("test advanceBy")

	context := ParseContext{}
	t.Run("test 1", func(t *testing.T) {
		input := "test a b c d"
		expected := "te"
		sourceExpected := "st a b c "
		context.Source = input
		result := parse.ParseTextData(&context, 2)
		if result != expected {
			t.Errorf("failed: input %s, expected %s, actual %s", input, expected, result)
		}

		if context.Source != sourceExpected {
			t.Errorf("failed: input %s, expected %s, actual %s", input, sourceExpected, context.Source)
		}
	})

	t.Run("test 2", func(t *testing.T) {
		input := "Hello World!"
		expected := "Hel"
		sourceExpected := "lo World"
		context.Source = input
		result := parse.ParseTextData(&context, 3)
		if result != expected {
			t.Errorf("failed: input %s, expected %s, actual %s", input, expected, result)
		}

		if context.Source != sourceExpected {
			t.Errorf("failed: input %s, expected %s, actual %s", input, sourceExpected, context.Source)
		}
	})

}
