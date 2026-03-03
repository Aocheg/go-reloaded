package main

import (
	"reflect"
	"testing"
)

//////////////////////////////////////////////////
// TEST CAPITALIZE
//////////////////////////////////////////////////

func TestCapitalize(t *testing.T) {
	tests := map[string]string{
		"hello": "Hello",
		"HELLO": "Hello",
		"hELLo": "Hello",
		"":      "",
	}

	for input, expected := range tests {
		result := Capitalize(input)
		if result != expected {
			t.Errorf("Capitalize(%q) = %q; want %q", input, result, expected)
		}
	}
}

//////////////////////////////////////////////////
// TEST FIX A → AN
//////////////////////////////////////////////////

func TestFixAAn(t *testing.T) {
	input := []string{"There", "is", "a", "apple"}
	expected := []string{"There", "is", "an", "apple"}

	result := FixAAn(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FixAAn failed.\nGot: %v\nWant: %v", result, expected)
	}
}

//////////////////////////////////////////////////
// TEST APPLY MODIFIERS (HEX & BIN)
//////////////////////////////////////////////////

func TestApplyModifiersHexBin(t *testing.T) {
	input := []string{"1E", "(", "hex", ")", "10", "(", "bin", ")"}
	expected := []string{"30", "2"}

	result := ApplyModifiers(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ApplyModifiers hex/bin failed.\nGot: %v\nWant: %v", result, expected)
	}
}

//////////////////////////////////////////////////
// TEST APPLY MODIFIERS (UP, LOW, CAP)
//////////////////////////////////////////////////

func TestApplyModifiersCase(t *testing.T) {
	input := []string{"hello", "(", "up", ")"}
	expected := []string{"HELLO"}

	result := ApplyModifiers(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ApplyModifiers up failed.\nGot: %v\nWant: %v", result, expected)
	}
}

//////////////////////////////////////////////////
// TEST APPLY MODIFIERS WITH NUMBER
//////////////////////////////////////////////////

func TestApplyModifiersWithNumber(t *testing.T) {
	input := []string{"this", "is", "great", "(", "up", ",", "2", ")"}
	expected := []string{"this", "IS", "GREAT"}

	result := ApplyModifiers(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ApplyModifiers (up, 2) failed.\nGot: %v\nWant: %v", result, expected)
	}
}

//////////////////////////////////////////////////
// TEST TOKENIZE
//////////////////////////////////////////////////

func TestTokenize(t *testing.T) {
	input := "Hello, world!"
	expected := []string{"Hello", ",", "world", "!"}

	result := Tokenize(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Tokenize failed.\nGot: %v\nWant: %v", result, expected)
	}
}

//////////////////////////////////////////////////
// TEST REBUILD (PUNCTUATION)
//////////////////////////////////////////////////

func TestRebuildPunctuation(t *testing.T) {
	input := []string{"Hello", ",", "world", "!"}
	expected := "Hello, world!"

	result := Rebuild(input)

	if result != expected {
		t.Errorf("Rebuild failed.\nGot: %q\nWant: %q", result, expected)
	}
}

//////////////////////////////////////////////////
// TEST APOSTROPHE
//////////////////////////////////////////////////

func TestRebuildApostrophe(t *testing.T) {
	input := []string{"'", "awesome", "'"}
	expected := "'awesome'"

	result := Rebuild(input)

	if result != expected {
		t.Errorf("Rebuild apostrophe failed.\nGot: %q\nWant: %q", result, expected)
	}
}
