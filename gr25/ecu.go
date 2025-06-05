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
		ACUPrechargeTest1,
		ACUPrechargeTest2,
		ACUConfigChargeParametersTest1,
		ACUConfigChargeParametersTest2,
		ACUConfigOperationalParametersTest1,
		ACUConfigOperationalParametersTest2,
		InverterConfigTest1,
		InverterConfigTest2,
		InverterCommandTest1,
		InverterCommandTest2,
		FanCommandTest1,
		FanCommandTest2,
		DashConfigTest1,
		DashConfigTest2,
		DashWarningFlagsTest1,
		DashWarningFlagsTest2,
		ECUPedalsDataTest1,
		ECUPedalsDataTest2,
	}

	for i := range tests {
		tests[i].Node = "ecu"
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
