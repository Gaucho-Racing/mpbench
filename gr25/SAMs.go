package gr25

import "fmt"

func GenerateSAM1Tests() []MessageTest {
	var tests = []MessageTest{
		SAMBrakeIRTest1,
		SAMBrakeIRTest2,
	}

	for i := range tests {
		tests[i].Node = "SAM1"
		tests[i].Name = fmt.Sprintf("SAM1 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("SAM1_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateSAM2Tests() []MessageTest {
	var tests = []MessageTest{
		SAMBrakeIRTest1,
		SAMBrakeIRTest2,
	}

	for i := range tests {
		tests[i].Node = "SAM2"
		tests[i].Name = fmt.Sprintf("SAM2 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("SAM2_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateSAM3Tests() []MessageTest {
	var tests = []MessageTest{
		SAMBrakeIRTest1,
		SAMBrakeIRTest2,
	}

	for i := range tests {
		tests[i].Node = "SAM3"
		tests[i].Name = fmt.Sprintf("SAM3 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("SAM3_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateSAM4Tests() []MessageTest {
	var tests = []MessageTest{
		SAMBrakeIRTest1,
		SAMBrakeIRTest2,
	}

	for i := range tests {
		tests[i].Node = "SAM4"
		tests[i].Name = fmt.Sprintf("SAM4 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("SAM4_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}

func GenerateSAM5Tests() []MessageTest {
	var tests = []MessageTest{
		SAMIMUTest1,
		SAMIMUTest2,
		SAMTOFTest1,
		SAMTOFTest2,
	}

	for i := range tests {
		tests[i].Node = "SAM5"
		tests[i].Name = fmt.Sprintf("SAM5 → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("SAM5_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}
