package gr25

var ECUStatusOneTest1 = MessageTest{
	ID:   0x003,
	Name: "ECU Status One Test 1",
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

var ECUStatusOneTest2 = MessageTest{
	ID:   0x003,
	Name: "ECU Status One Test 2",
	Data: []byte{0x12, 0x42, 0xFF, 0x00, 0x31, 0x82, 0x58, 0x72},
	ExpectedValues: map[string]interface{}{
		"ecu_state":               18,
		"ecu_status_acu":          0,
		"ecu_status_inv_one":      1,
		"ecu_status_inv_two":      0,
		"ecu_status_inv_three":    0,
		"ecu_status_inv_four":     0,
		"ecu_status_fan_one":      0,
		"ecu_status_fan_two":      0,
		"ecu_status_fan_three":    0,
		"ecu_status_fan_four":     1,
		"ecu_status_fan_five":     1,
		"ecu_status_fan_six":      1,
		"ecu_status_fan_seven":    1,
		"ecu_status_fan_eight":    1,
		"ecu_status_dash":         0,
		"ecu_status_steering":     1,
		"ecu_power_level":         3,
		"ecu_torque_map":          1,
		"ecu_max_cell_temp":       32.5,
		"ecu_acu_state_of_charge": 34.509804,
		"ecu_glv_state_of_charge": 44.705882,
	},
}

var ECUStatusTwoTest1 = MessageTest{
	ID:   0x004,
	Name: "ECU Status Two Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"ecu_tractive_system_voltage": 0,
		"ecu_vehicle_speed":           0,
		"ecu_fr_wheel_rpm":            0,
		"ecu_fl_wheel_rpm":            0,
	},
}

var ECUStatusThreeTest1 = MessageTest{
	ID:   0x005,
	Name: "ECU Status Three Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"ecu_rr_wheel_rpm": 0,
		"ecu_rl_wheel_rpm": 0,
	},
}
