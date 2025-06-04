package gr25

import "fmt"

func GenerateLVDCDCTests() []MessageTest {
	var tests = []MessageTest{
		DCDCStatusTest1,
		DCDCStatusTest2,
	}

	for i := range tests {
		tests[i].Node = "LV_DC-DC"
		tests[i].Name = fmt.Sprintf("LV DC-DC → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("LV_DC-DC_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}
