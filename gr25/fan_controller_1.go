package gr25

import "fmt"

func GenerateFanController1Tests() []MessageTest {
	var tests = []MessageTest{
		FanStatusTest1,
		FanStatusTest2,
	}

	for i := range tests {
		tests[i].Node = "fan_controller_1"
		tests[i].Name = fmt.Sprintf("Fan Controller 1 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("fan_controller_1_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}
