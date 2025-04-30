// unit test
package address_test

import (
	. "introduction-tests/address"
	"testing"
)

type scenaryTest struct {
	addressToTest string
	expected      string
}

func TestTypeOfAddress(t *testing.T) {
	scenarysFromTest := []scenaryTest{
		{"Road 123", "Road"},
		{"Street 456", "Street"},
		{"Avenue 789", "Avenue"},
		{"Boulevard 101", "Boulevard"},
		{"Drive 112", "Drive"},
		{"Lane 131", "Lane"},
		{"Court 415", "Court"},
		{"Place 161", "Place"},
		{"Way 718", "Way"},
		{"Circle 192", "Circle"},
		{"Unknown 202", "unknown"},
	}
	for _, scenary := range scenarysFromTest {
		addressReceived := TypeOfAddress(scenary.addressToTest)
		if addressReceived != scenary.expected {
			t.Errorf("Expected %s, but got %s", scenary.expected, addressReceived)
		}
	}
}
