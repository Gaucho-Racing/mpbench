package gr25

import "fmt"

func GenerateTCMTests() []MessageTest {
	var tests = []MessageTest{
		TCMResourceTest1,
	}

	for i := range tests {
		tests[i].Node = "tcm"
		tests[i].Name = fmt.Sprintf("TCM → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("tcm_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}
