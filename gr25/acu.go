package gr25

import "fmt"

func GenerateACUTests() []MessageTest {
	var tests = []MessageTest{
		ACUStatusOneTest1,
		ACUStatusOneTest2,
		ACUStatusTwoTest1,
		ACUStatusTwoTest2,
		ACUStatusThreeTest1,
		ACUStatusThreeTest2,
		ACUConfigChargeParametersTest1,
		ACUConfigChargeParametersTest2,
		ACUConfigOperationalParametersTest1,
		ACUConfigOperationalParametersTest2,
		ACUCellDataOneTest1,
		ACUCellDataOneTest2,
		ACUCellDataTwoTest1,
		ACUCellDataTwoTest2,
		ACUCellDataThreeTest1,
		ACUCellDataThreeTest2,
		ACUCellDataFourTest1,
		ACUCellDataFourTest2,
		ACUCellDataFiveTest1,
		ACUCellDataFiveTest2,
	}

	for i := range tests {
		tests[i].Node = "acu"
		tests[i].Name = fmt.Sprintf("ACU → %s", tests[i].Name)
		newExpectedValues := make(map[string]interface{})
		for k, v := range tests[i].ExpectedValues {
			signalName := fmt.Sprintf("acu_%s", k)
			newExpectedValues[signalName] = v
		}
		tests[i].ExpectedValues = newExpectedValues
	}
	return tests
}
