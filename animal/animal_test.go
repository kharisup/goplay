package main

import (
	"testing"
)

// // define input-result struct type
// type TestDataItem struct {
// 	inputs   []Dog  // inputs to `Add` function
// 	result   string // result of `Add` function
// 	hasError bool   // does `Add` function returns error
// }

func TestAnimal(t *testing.T) {

	// // input-result data items
	// dataItems := []TestDataItem{
	// 	{[]Dog{4}, "Bark", false},
	// 	{[]Dog{2}, "Bark", false},
	// }

	var obj Animal
	obj = Dog{legs: 4}
	result := obj.Talk()
	expected := "Bark"

	if result != expected {
		t.Errorf("Animal Dog Talk() failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("Animal Dog Talk() failed, expected %v, got %v", expected, result)
	}

	result = obj.Walking()
	expected = "Walking using 4 legs"
	if result != expected {
		t.Errorf("Animal Dog Talk() failed, expected %v, got %v", expected, result)
	} else {
		t.Logf("Animal Dog Talk() failed, expected %v, got %v", expected, result)
	}

	obj = Chicken{legs: 2}
	result = obj.Talk()
	if result != "Petok" {
		t.Errorf("Animal Chicken Talk() failed, expected %v, got %v", "Petok", result)
	} else {
		t.Logf("Animal Chicken Talk() failed, expected %v, got %v", "Petok", result)
	}
}
