package main

import (
	"./KeyCount"
	"testing"
)

func TestCreateMap(t *testing.T) {
	var keyCount KeyCount.KeyCount
	keyCount = KeyCount.NewKeyCountMap()
	kType := keyCount.PrintType()
	if kType != "map" {
		t.Error("not a map, is:" + kType)
	}
}
func TestCreateArray(t *testing.T) {
	var keyCount KeyCount.KeyCount
	keyCount = KeyCount.NewKeyCountArray()
	kType := keyCount.PrintType()
	if kType != "array" {
		t.Error("not a array, is:" + kType)
	}
}
