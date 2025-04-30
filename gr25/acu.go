package gr25

import "fmt"

func GenerateACUTests() []MessageTest {
	var tests = []MessageTest{
		ACUStatusOneTest1,
		ACUStatusOneTest2,
		ACUStatusTwoTest1,
		ACUStatusTwoTest2,
		ACUStatusThreeTest1,
		ACUStatusThreeTest2,
	}

	for i := range tests {
		tests[i].Node = "acu"
		tests[i].Name = fmt.Sprintf("ACU → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("acu_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}
