package gr25

import "fmt"

func GenerateSteeringWheelTests() []MessageTest {
	var tests = []MessageTest{
		//SteeringStatusTest1,
		//SteeringStatusTest2
	}

	for i := range tests {
		tests[i].Node = "steering_wheel"
		tests[i].Name = fmt.Sprintf("Steering Wheel → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("steering_wheel_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}
