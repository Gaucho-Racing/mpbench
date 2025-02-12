package gr25

import "mpbench/utils"

func SendECUStatusOne(mqttPort int, dbPort int) {
	test1 := MessageTest{
		ID:   0x003,
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"ecu_state":               0,
			"ecu_status_acu":          0,
			"ecu_status_inv_one":      0,
			"ecu_status_inv_two":      0,
			"ecu_status_inv_three":    0,
			"ecu_status_inv_four":     0,
			"ecu_status_fan_one":      0,
			"ecu_status_fan_two":      0,
			"ecu_status_fan_three":    0,
			"ecu_status_fan_four":     0,
			"ecu_status_fan_five":     0,
			"ecu_status_fan_six":      0,
			"ecu_status_fan_seven":    0,
			"ecu_status_fan_eight":    0,
			"ecu_status_dash":         0,
			"ecu_status_steering":     0,
			"ecu_power_level":         0,
			"ecu_torque_map":          0,
			"ecu_max_cell_temp":       0,
			"ecu_acu_state_of_charge": 0,
			"ecu_glv_state_of_charge": 0,
		},
	}
	status := test1.Run(mqttPort, dbPort)
	if !status {
		utils.SugarLogger.Infof("❌ TEST FAILED: %d ECU Status One", test1.ID)
		return
	}
	utils.SugarLogger.Infof("✅ TEST PASSED: %d ECU Status One", test1.ID)
}
