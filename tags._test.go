package main

import "testing"

func TestInvalidTagsNoEnd(t *testing.T) {
	invalidString := "{tag1"
	tags, _ := Tags(invalidString)
	if len(tags) != 0 {
		t.Errorf("Expected empty set of tags")
	}
}

func TestInvalidTagsNoStart(t *testing.T) {
	invalidString := "tag1}"
	tags, _ := Tags(invalidString)
	if len(tags) != 0 {
		t.Errorf("Expected empty set of tags")
	}
}

func TestSingleTag(t *testing.T) {
	invalidString := "{tag1}"
	tags, _ := Tags(invalidString)
	if len(tags) != 1 {
		t.Errorf("Expected empty set of tags")
	}
	if tags[0] != "tag1" {
		t.Errorf("Expected tags[0] to be tag1. Found %s", tags[0])
	}
}

func TestTwoTagsWhitespace(t *testing.T) {
	invalidString := "{tag1, tag2}"
	tags, _ := Tags(invalidString)
	if len(tags) != 2 {
		t.Errorf("Expected empty set of tags")
	}
	if tags[0] != "tag1" {
		t.Errorf("Expected tags[0] to be tag1. Found '%s'", tags[0])
	}
	if tags[1] != "tag2" {
		t.Errorf("Expected tags[1] to be tag2. Found '%s'", tags[1])
	}
}

func TestTwoTagsNoWhitespace(t *testing.T) {
	invalidString := "{tag1,tag2}"
	tags, _ := Tags(invalidString)
	if len(tags) != 2 {
		t.Errorf("Expected empty set of tags")
	}
	if tags[0] != "tag1" {
		t.Errorf("Expected tags[0] to be tag1. Found '%s'", tags[0])
	}
	if tags[1] != "tag2" {
		t.Errorf("Expected tags[1] to be tag2. Found '%s'", tags[1])
	}
}
