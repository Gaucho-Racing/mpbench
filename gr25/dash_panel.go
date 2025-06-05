package gr25

import "fmt"

func GenerateDashPanelTests() []MessageTest {
	var tests = []MessageTest{
		DashStatusTest1,
		DashStatusTest2,
	}

	for i := range tests {
		tests[i].Node = "dash_panel"
		tests[i].Name = fmt.Sprintf("Dash Panel → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("dash_panel_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}
