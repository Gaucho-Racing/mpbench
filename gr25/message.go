package gr25

var ECUStatusOneTest1 = MessageTest{
	ID:   0x003,
	Name: "ECU Status One Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"state":               0,
		"acu_status":          0,
		"inv_one_status":      0,
		"inv_two_status":      0,
		"inv_three_status":    0,
		"inv_four_status":     0,
		"fan_one_status":      0,
		"fan_two_status":      0,
		"fan_three_status":    0,
		"fan_four_status":     0,
		"fan_five_status":     0,
		"fan_six_status":      0,
		"fan_seven_status":    0,
		"fan_eight_status":    0,
		"dash_status":         0,
		"steering_status":     0,
		"power_level":         0,
		"torque_map":          0,
		"max_cell_temp":       0,
		"acu_state_of_charge": 0,
		"glv_state_of_charge": 0,
	},
}

var ECUStatusOneTest2 = MessageTest{
	ID:   0x003,
	Name: "ECU Status One Test 2",
	Data: []byte{0x12, 0x42, 0xFF, 0x00, 0x31, 0x82, 0x58, 0x72},
	ExpectedValues: map[string]interface{}{
		"state":               18,
		"acu_status":          0,
		"inv_one_status":      1,
		"inv_two_status":      0,
		"inv_three_status":    0,
		"inv_four_status":     0,
		"fan_one_status":      0,
		"fan_two_status":      1,
		"fan_three_status":    0,
		"fan_four_status":     1,
		"fan_five_status":     1,
		"fan_six_status":      1,
		"fan_seven_status":    1,
		"fan_eight_status":    1,
		"dash_status":         1,
		"steering_status":     1,
		"power_level":         3,
		"torque_map":          1,
		"max_cell_temp":       32.5,
		"acu_state_of_charge": 34.509804,
		"glv_state_of_charge": 44.705882,
	},
}

var ECUStatusTwoTest1 = MessageTest{
	ID:   0x004,
	Name: "ECU Status Two Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"tractive_system_voltage": 0,
		"vehicle_speed":           0,
		"fr_wheel_rpm":            -3276.8,
		"fl_wheel_rpm":            -3276.8,
	},
}

var ECUStatusTwoTest2 = MessageTest{
	ID:   0x004,
	Name: "ECU Status Two Test 2",
	Data: []byte{0xa0, 0xa5, 0x9c, 0x95, 0xd1, 0xfe, 0x77, 0x3b},
	ExpectedValues: map[string]interface{}{
		"tractive_system_voltage": 424,
		"vehicle_speed":           383,
		"fr_wheel_rpm":            3246.5,
		"fl_wheel_rpm":            -1754.5,
	},
}

var ECUStatusThreeTest1 = MessageTest{
	ID:   0x005,
	Name: "ECU Status Three Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"rr_wheel_rpm": -3276.8,
		"rl_wheel_rpm": -3276.8,
	},
}

var ECUStatusThreeTest2 = MessageTest{
	ID:   0x005,
	Name: "ECU Status Three Test 2",
	Data: []byte{0xfe, 0xf4, 0x9e, 0xb0},
	ExpectedValues: map[string]interface{}{
		"rr_wheel_rpm": 2995,
		"rl_wheel_rpm": 1244.6,
	},
}

var ACUStatusOneTest1 = MessageTest{
	ID: 0x007,
	Name: "ACU Status One Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"accumulator_voltage": 0,
		"ts_voltage": 0,
		"accumulator_current": -327.68,
		"accumulator_soc": 0,
		"glv_soc": 0,
	},
}

var ACUStatusOneTest2 = MessageTest{
	ID: 0x007,
	Name: "ACU Status One Test 1",
	Data: []byte{0x84, 0x07, 0xd7, 0x12, 0xbf, 0x63, 0x80, 0xf3},
	ExpectedValues: map[string]interface{}{
		"accumulator_voltage": 19.24,
		"ts_voltage": 48.23,
		"accumulator_current": -72.33,
		"accumulator_soc": 31.372,
		"glv_soc": 95.294,
	},
}

var ACUStatusTwoTest1 = MessageTest{
	ID: 0x008,
	Name: "ACU Status Two Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"20v_voltage": 0,
		"12v_voltage": 0,
		"sdc_voltage": 0,
		"min_cell_voltage": 2,
		"max_cell_temp": 0,
		"over_temp_error": 0,
		"over_voltage_error": 0,
		"under_voltage_error": 0,
		"over_current_error": 0,
		"under_current_error": 0,
		"under_voltage_20v_warning": 0,
		"under_voltage_12v_warning": 0,
		"under_voltage_sdc_warning": 0,
		"precharge_error": 0,
		"ir_minus_state": 0,
		"ir_plus_state": 0,
		"software_latch": 0,
	},
}

var ACUStatusTwoTest2 = MessageTest{
	ID: 0x008,
	Name: "ACU Status Two Test 2",
	Data: []byte{0xbd, 0x66, 0x10, 0x28, 0xcc, 0x96, 0x0c},
	ExpectedValues: map[string]interface{}{
		"20v_voltage": 18.9,
		"12v_voltage": 10.2,
		"sdc_voltage": 1.6,
		"min_cell_voltage": 6,
		"max_cell_temp": 51,
		"over_temp_error": 0,
		"over_voltage_error": 1,
		"under_voltage_error": 1,
		"over_current_error": 0,
		"under_current_error": 1,
		"under_voltage_20v_warning": 0,
		"under_voltage_12v_warning": 0,
		"under_voltage_sdc_warning": 1,
		"precharge_error": 0,
		"ir_minus_state": 0,
		"ir_plus_state": 1,
		"software_latch": 1,
	},
}