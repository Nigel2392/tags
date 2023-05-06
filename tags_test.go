package tags_test

import (
	"testing"

	"github.com/Nigel2392/tags"
)

func TestTagmap(t *testing.T) {
	var tags = tags.ParseTags("tag1=value1,value2,value3;tag2=value1,value2;")
	if len(tags) != 2 {
		t.Error("Expected 2 tags, got", len(tags))
	}
	if len(tags["tag1"]) != 3 {
		t.Error("Expected 3 values for tag1")
	}
	if len(tags["tag2"]) != 2 {
		t.Error("Expected 2 values for tag2")
	}
	if tags["tag1"][0] != "value1" {
		t.Error("Expected value1 for tag1")
	}
	if tags["tag1"][1] != "value2" {
		t.Error("Expected value2 for tag1")
	}
	if tags["tag1"][2] != "value3" {
		t.Error("Expected value3 for tag1")
	}
	if tags["tag2"][0] != "value1" {
		t.Error("Expected value1 for tag2")
	}
	if tags["tag2"][1] != "value2" {
		t.Error("Expected value2 for tag2")
	}
}

type TestTagmapStruct struct {
	Field1 string `tag:"tag1=value1,value2,value3;tag2=value1,value2;"`
	Field2 string `tag:"tag1=value1,value2,value3;tag2=value1,value2;tag3=value1;"`
}

func TestFromStruct(t *testing.T) {
	var tags = tags.FromStruct(TestTagmapStruct{}, "tag", ";", "=", ",")
	if len(tags) != 2 {
		t.Error("Expected 2 tags, got", len(tags))
	}
	if len(tags["Field1"]) != 2 {
		t.Error("Expected 2 tags for Field1")
	}
	if len(tags["Field2"]) != 3 {
		t.Error("Expected 3 tags for Field2")
	}
	if tags["Field1"]["tag1"][0] != "value1" {
		t.Error("Expected value1 for Field1.tag1")
	}
	if tags["Field1"]["tag1"][1] != "value2" {
		t.Error("Expected value2 for Field1.tag1")
	}
	if tags["Field1"]["tag1"][2] != "value3" {
		t.Error("Expected value3 for Field1.tag1")
	}
	if tags["Field1"]["tag2"][0] != "value1" {
		t.Error("Expected value1 for Field1.tag2")
	}
	if tags["Field1"]["tag2"][1] != "value2" {
		t.Error("Expected value2 for Field1.tag2")
	}
	if tags["Field2"]["tag1"][0] != "value1" {
		t.Error("Expected value1 for Field2.tag1")
	}
	if tags["Field2"]["tag1"][1] != "value2" {
		t.Error("Expected value2 for Field2.tag1")
	}
	if tags["Field2"]["tag1"][2] != "value3" {
		t.Error("Expected value3 for Field2.tag1")
	}
	if tags["Field2"]["tag2"][0] != "value1" {
		t.Error("Expected value1 for Field2.tag2")
	}
	if tags["Field2"]["tag2"][1] != "value2" {
		t.Error("Expected value2 for Field2.tag2")
	}
	if tags["Field2"]["tag3"][0] != "value1" {
		t.Error("Expected value1 for Field2.tag3")
	}
}
