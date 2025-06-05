package gr25

import "fmt"

// ------------------------------------------------------------
// GR INVERTERS (inverter_1, inverter_2, inverter_3, inverter_4)
// ------------------------------------------------------------

func GenerateGRInverter1Tests() []MessageTest {
	var tests = []MessageTest{
		InverterStatusOneTest1,
		InverterStatusOneTest2,
		InverterStatusTwoTest1,
		InverterStatusTwoTest2,
		InverterStatusThreeTest1,
		InverterStatusThreeTest2,
	}

	for i := range tests {
		tests[i].Node = "inverter_1"
		tests[i].Name = fmt.Sprintf("Inverter 1 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("inverter_1_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateGRInverter2Tests() []MessageTest {
	var tests = []MessageTest{
		InverterStatusOneTest1,
		InverterStatusOneTest2,
		InverterStatusTwoTest1,
		InverterStatusTwoTest2,
		InverterStatusThreeTest1,
		InverterStatusThreeTest2,
	}

	for i := range tests {
		tests[i].Node = "inverter_2"
		tests[i].Name = fmt.Sprintf("Inverter 2 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("inverter_2_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateGRInverter3Tests() []MessageTest {
	var tests = []MessageTest{
		InverterStatusOneTest1,
		InverterStatusOneTest2,
		InverterStatusTwoTest1,
		InverterStatusTwoTest2,
		InverterStatusThreeTest1,
		InverterStatusThreeTest2,
	}

	for i := range tests {
		tests[i].Node = "inverter_3"
		tests[i].Name = fmt.Sprintf("Inverter 3 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("inverter_3_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateGRInverter4Tests() []MessageTest {
	var tests = []MessageTest{
		InverterStatusOneTest1,
		InverterStatusOneTest2,
		InverterStatusTwoTest1,
		InverterStatusTwoTest2,
		InverterStatusThreeTest1,
		InverterStatusThreeTest2,
	}

	for i := range tests {
		tests[i].Node = "inverter_4"
		tests[i].Name = fmt.Sprintf("Inverter 4 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("inverter_4_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

// ------------------------------------------------------------
// DTI INVERTERS (inverter_dti)
// ------------------------------------------------------------