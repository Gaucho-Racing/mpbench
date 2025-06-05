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

func GenerateFanController2Tests() []MessageTest {
	var tests = []MessageTest{
		FanStatusTest1,
		FanStatusTest2,
	}

	for i := range tests {
		tests[i].Node = "fan_controller_2"
		tests[i].Name = fmt.Sprintf("Fan Controller 2 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("fan_controller_2_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateFanController3Tests() []MessageTest {
	var tests = []MessageTest{
		FanStatusTest1,
		FanStatusTest2,
	}

	for i := range tests {
		tests[i].Node = "fan_controller_3"
		tests[i].Name = fmt.Sprintf("Fan Controller 3 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("fan_controller_3_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateFanController4Tests() []MessageTest {
	var tests = []MessageTest{
		FanStatusTest1,
		FanStatusTest2,
	}

	for i := range tests {
		tests[i].Node = "fan_controller_4"
		tests[i].Name = fmt.Sprintf("Fan Controller 4 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("fan_controller_4_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateFanController5Tests() []MessageTest {
	var tests = []MessageTest{
		FanStatusTest1,
		FanStatusTest2,
	}

	for i := range tests {
		tests[i].Node = "fan_controller_5"
		tests[i].Name = fmt.Sprintf("Fan Controller 5 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("fan_controller_5_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateFanController6Tests() []MessageTest {
	var tests = []MessageTest{
		FanStatusTest1,
		FanStatusTest2,
	}

	for i := range tests {
		tests[i].Node = "fan_controller_6"
		tests[i].Name = fmt.Sprintf("Fan Controller 6 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("fan_controller_6_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateFanController7Tests() []MessageTest {
	var tests = []MessageTest{
		FanStatusTest1,
		FanStatusTest2,
	}

	for i := range tests {
		tests[i].Node = "fan_controller_7"
		tests[i].Name = fmt.Sprintf("Fan Controller 7 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("fan_controller_7_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateFanController8Tests() []MessageTest {
	var tests = []MessageTest{
		FanStatusTest1,
		FanStatusTest2,
	}

	for i := range tests {
		tests[i].Node = "fan_controller_8"
		tests[i].Name = fmt.Sprintf("Fan Controller 8 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("fan_controller_8_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}
