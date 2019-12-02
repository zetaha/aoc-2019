package main

import "testing"

func TestFuelUpper(t *testing.T) {
	got := fuelCounterUpper(14)
	expected := 2
	if got != expected {
		t.Errorf("fuel count wrong got %v, expected %v", got, expected)
	}

	got = fuelCounterUpper(1969)
	expected = 654
	if got != expected {
		t.Errorf("fuel count wrong got %v, expected %v", got, expected)
	}

	got = fuelCounterUpper(100756)
	expected = 33583

	if got != expected {
		t.Errorf("fuel count wrong got %v, expected %v", got, expected)
	}
}

func TestTotalFuel(t *testing.T) {
	got := totalFuel(1969)
	expected := 966
	if got != expected {
		t.Errorf("fuel count wrong got %v, expected %v", got, expected)
	}
}
