package gr25

import "fmt"

func GenerateECUTests() []MessageTest {
	var tests = []MessageTest{
		ECUStatusOneTest1,
		ECUStatusOneTest2,
		ECUStatusTwoTest1,
		ECUStatusTwoTest2,
		ECUStatusThreeTest1,
		ECUStatusThreeTest2,
	}

	for i := range tests {
		tests[i].Name = fmt.Sprintf("ECU → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("ecu_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}
