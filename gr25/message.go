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
	ID:   0x007,
	Name: "ACU Status One Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"accumulator_voltage": 0,
		"ts_voltage":          0,
		"accumulator_current": -327.68,
		"accumulator_soc":     0,
		"glv_soc":             0,
	},
}

var ACUStatusOneTest2 = MessageTest{
	ID:   0x007,
	Name: "ACU Status One Test 1",
	Data: []byte{0x84, 0x07, 0xd7, 0x12, 0xbf, 0x63, 0x80, 0xf3},
	ExpectedValues: map[string]interface{}{
		"accumulator_voltage": 19.24,
		"ts_voltage":          48.23,
		"accumulator_current": -72.33,
		"accumulator_soc":     31.372,
		"glv_soc":             95.294,
	},
}

var ACUStatusTwoTest1 = MessageTest{
	ID:   0x008,
	Name: "ACU Status Two Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"20v_voltage":               0,
		"12v_voltage":               0,
		"sdc_voltage":               0,
		"min_cell_voltage":          2,
		"max_cell_temp":             0,
		"over_temp_error":           0,
		"over_voltage_error":        0,
		"under_voltage_error":       0,
		"over_current_error":        0,
		"under_current_error":       0,
		"under_voltage_20v_warning": 0,
		"under_voltage_12v_warning": 0,
		"under_voltage_sdc_warning": 0,
		"precharge_error":           0,
		"ir_minus_state":            0,
		"ir_plus_state":             0,
		"software_latch":            0,
	},
}

var ACUStatusTwoTest2 = MessageTest{
	ID:   0x008,
	Name: "ACU Status Two Test 2",
	Data: []byte{0xbd, 0x66, 0x10, 0x28, 0xcc, 0x96, 0x0c},
	ExpectedValues: map[string]interface{}{
		"20v_voltage":               18.9,
		"12v_voltage":               10.2,
		"sdc_voltage":               1.6,
		"min_cell_voltage":          6,
		"max_cell_temp":             51,
		"over_temp_error":           0,
		"over_voltage_error":        1,
		"under_voltage_error":       1,
		"over_current_error":        0,
		"under_current_error":       1,
		"under_voltage_20v_warning": 0,
		"under_voltage_12v_warning": 0,
		"under_voltage_sdc_warning": 1,
		"precharge_error":           0,
		"ir_minus_state":            0,
		"ir_plus_state":             1,
		"software_latch":            1,
	},
}

var ACUStatusThreeTest1 = MessageTest{
	ID:   0x009,
	Name: "ACU Status Three Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"hv_input_voltage":  0,
		"hv_output_voltage": 0,
		"hv_input_current":  0,
		"hv_output_current": 0,
	},
}

var ACUStatusThreeTest2 = MessageTest{
	ID:   0x009,
	Name: "ACU Status Three Test 2",
	Data: []byte{0xff, 0xff, 0x7c, 0xea, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"hv_input_voltage":  655.35,
		"hv_output_voltage": 600.28,
		"hv_input_current":  0,
		"hv_output_current": 0,
	},
}

var ACUPrechargeTest1 = MessageTest{
	ID:   0x00A,
	Name: "ACU Status Three Test 1",
	Data: []byte{0x00},
	ExpectedValues: map[string]interface{}{
		"set_ts_active": 0,
	},
}

var ACUPrechargeTest2 = MessageTest{
	ID:   0x00A,
	Name: "ACU Status Three Test 2",
	Data: []byte{0x01},
	ExpectedValues: map[string]interface{}{
		"set_ts_active": 1,
	},
}
var ACUConfigChargeParametersTest1 = MessageTest{
	ID:   0x00B,
	Name: "ACU Config Charge Parameters Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"charge_voltage": 0,
		"charge_current": 0,
	},
}

var ACUConfigChargeParametersTest2 = MessageTest{
	ID:   0x00B,
	Name: "ACU Config Charge Parameters Test 2",
	Data: []byte{0xdb, 0x98, 0xdd, 0x31},
	ExpectedValues: map[string]interface{}{
		"charge_voltage": 3913.1,
		"charge_current": 1276.5,
	},
}

var ACUConfigOperationalParametersTest1 = MessageTest{
	ID:   0x00C,
	Name: "ACU Config Operational Parameters Test 1",
	Data: []byte{0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"min_cell_voltage":     2,
		"max_cell_temperature": 0,
	},
}
var ACUConfigOperationalParametersTest2 = MessageTest{
	ID:   0x00C,
	Name: "ACU Config Operational Parameters Test 2",
	Data: []byte{0x2c, 0x16},
	ExpectedValues: map[string]interface{}{
		"min_cell_voltage":     2.44,
		"max_cell_temperature": 32.5,
	},
}

var ACUCellDataOneTest1 = MessageTest{
	ID:   0x00D,
	Name: "ACU Cell Data One Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"cell0_voltage":  2,
		"cell0_temp":     0,
		"cell1_voltage":  2,
		"cell1_temp":     0,
		"cell2_voltage":  2,
		"cell2_temp":     0,
		"cell3_voltage":  2,
		"cell3_temp":     0,
		"cell4_voltage":  2,
		"cell4_temp":     0,
		"cell5_voltage":  2,
		"cell5_temp":     0,
		"cell6_voltage":  2,
		"cell6_temp":     0,
		"cell7_voltage":  2,
		"cell7_temp":     0,
		"cell8_voltage":  2,
		"cell8_temp":     0,
		"cell9_voltage":  2,
		"cell9_temp":     0,
		"cell10_voltage": 2,
		"cell10_temp":    0,
		"cell11_voltage": 2,
		"cell11_temp":    0,
		"cell12_voltage": 2,
		"cell12_temp":    0,
		"cell13_voltage": 2,
		"cell13_temp":    0,
		"cell14_voltage": 2,
		"cell14_temp":    0,
		"cell15_voltage": 2,
		"cell15_temp":    0,
		"cell16_voltage": 2,
		"cell16_temp":    0,
		"cell17_voltage": 2,
		"cell17_temp":    0,
		"cell18_voltage": 2,
		"cell18_temp":    2,
		"cell19_voltage": 2,
		"cell19_temp":    0,
		"cell20_voltage": 2,
		"cell20_temp":    0,
		"cell21_voltage": 2,
		"cell21_temp":    0,
		"cell22_voltage": 2,
		"cell22_temp":    0,
		"cell23_voltage": 2,
		"cell23_temp":    0,
		"cell24_voltage": 2,
		"cell24_temp":    0,
		"cell25_voltage": 2,
		"cell25_temp":    0,
		"cell26_voltage": 2,
		"cell26_temp":    0,
		"cell27_voltage": 2,
		"cell27_temp":    0,
		"cell28_voltage": 2,
		"cell28_temp":    0,
		"cell29_voltage": 2,
		"cell29_temp":    0,
		"cell30_voltage": 2,
		"cell30_temp":    0,
		"cell31_voltage": 2,
		"cell31_temp":    0,
	},
}

var ACUCellDataOneTest2 = MessageTest{
	ID:   0x00D,
	Name: "ACU Cell Data One Test 2",
	Data: []byte{
		0x0a, 0x89, 0x65, 0x04, 0x03, 0xd2, 0x32, 0x55, 0xc9, 0xb7,
		0xd6, 0xab, 0x7e, 0xfc, 0xa2, 0xf4, 0x59, 0x0b, 0x86, 0x80,
		0x42, 0x62, 0x39, 0xf1, 0x52, 0xe9, 0xae, 0x3f, 0x40, 0x88,
		0x8f, 0x70, 0x82, 0x18, 0xb4, 0x67, 0xc8, 0xa5, 0xaa, 0xb1,
		0x56, 0xdb, 0xb2, 0xfe, 0x1a, 0x1b, 0x5c, 0x2b, 0xb7, 0x45,
		0xc7, 0x96, 0xee, 0x07, 0xb8, 0xf2, 0xad, 0x8b, 0x8c, 0x0c,
		0x12, 0xe7, 0x86, 0xab},
	ExpectedValues: map[string]interface{}{
		"cell0_voltage":  2.1,
		"cell0_temp":     34.25,
		"cell1_voltage":  3.01,
		"cell1_temp":     1,
		"cell2_voltage":  2.03,
		"cell2_temp":     52.5,
		"cell3_voltage":  2.5,
		"cell3_temp":     21.25,
		"cell4_voltage":  4.01,
		"cell4_temp":     45.75,
		"cell5_voltage":  4.14,
		"cell5_temp":     42.75,
		"cell6_voltage":  3.26,
		"cell6_temp":     63,
		"cell7_voltage":  3.62,
		"cell7_temp":     61,
		"cell8_voltage":  2.89,
		"cell8_temp":     2.75,
		"cell9_voltage":  3.34,
		"cell9_temp":     32,
		"cell10_voltage": 2.66,
		"cell10_temp":    24.5,
		"cell11_voltage": 2.57,
		"cell11_temp":    60.25,
		"cell12_voltage": 2.82,
		"cell12_temp":    58.25,
		"cell13_voltage": 3.74,
		"cell13_temp":    15.75,
		"cell14_voltage": 2.64,
		"cell14_temp":    32,
		"cell15_voltage": 3.43,
		"cell15_temp":    28,
		"cell16_voltage": 3.3,
		"cell16_temp":    6,
		"cell17_voltage": 3.8,
		"cell17_temp":    25.75,
		"cell18_voltage": 4,
		"cell18_temp":    41.25,
		"cell19_voltage": 3.7,
		"cell19_temp":    44.25,
		"cell20_voltage": 2.86,
		"cell20_temp":    54.75,
		"cell21_voltage": 3.78,
		"cell21_temp":    63.5,
		"cell22_voltage": 2.26,
		"cell22_temp":    6.75,
		"cell23_voltage": 2.92,
		"cell23_temp":    10.75,
		"cell24_voltage": 3.83,
		"cell24_temp":    17.25,
		"cell25_voltage": 3.99,
		"cell25_temp":    37.5,
		"cell26_voltage": 4.38,
		"cell26_temp":    1.75,
		"cell27_voltage": 3.84,
		"cell27_temp":    60.5,
		"cell28_voltage": 3.73,
		"cell28_temp":    34.75,
		"cell29_voltage": 3.4,
		"cell29_temp":    3,
		"cell30_voltage": 2.18,
		"cell30_temp":    57.75,
		"cell31_voltage": 3.34,
		"cell31_temp":    42.75,
	},
}

var ACUCellDataTwoTest1 = MessageTest{
	ID:   0x00E,
	Name: "ACU Cell Data Two Test 1",
	Data: []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"cell32_voltage": 2,
		"cell32_temp":    0,
		"cell33_voltage": 2,
		"cell33_temp":    0,
		"cell34_voltage": 2,
		"cell34_temp":    0,
		"cell35_voltage": 2,
		"cell35_temp":    0,
		"cell36_voltage": 2,
		"cell36_temp":    0,
		"cell37_voltage": 2,
		"cell37_temp":    0,
		"cell38_voltage": 2,
		"cell38_temp":    0,
		"cell39_voltage": 2,
		"cell39_temp":    0,
		"cell40_voltage": 2,
		"cell40_temp":    0,
		"cell41_voltage": 2,
		"cell41_temp":    0,
		"cell42_voltage": 2,
		"cell42_temp":    0,
		"cell43_voltage": 2,
		"cell43_temp":    0,
		"cell44_voltage": 2,
		"cell44_temp":    0,
		"cell45_voltage": 2,
		"cell45_temp":    0,
		"cell46_voltage": 2,
		"cell46_temp":    0,
		"cell47_voltage": 2,
		"cell47_temp":    0,
		"cell48_voltage": 2,
		"cell48_temp":    0,
		"cell49_voltage": 2,
		"cell49_temp":    0,
		"cell50_voltage": 2,
		"cell50_temp":    0,
		"cell51_voltage": 2,
		"cell51_temp":    0,
		"cell52_voltage": 2,
		"cell52_temp":    0,
		"cell53_voltage": 2,
		"cell53_temp":    0,
		"cell54_voltage": 2,
		"cell54_temp":    0,
		"cell55_voltage": 2,
		"cell55_temp":    0,
		"cell56_voltage": 2,
		"cell56_temp":    0,
		"cell57_voltage": 2,
		"cell57_temp":    0,
		"cell58_voltage": 2,
		"cell58_temp":    0,
		"cell59_voltage": 2,
		"cell59_temp":    0,
		"cell60_voltage": 2,
		"cell60_temp":    0,
		"cell61_voltage": 2,
		"cell61_temp":    0,
		"cell62_voltage": 2,
		"cell62_temp":    0,
		"cell63_voltage": 2,
		"cell63_temp":    0,
	},
}

var ACUCellDataTwoTest2 = MessageTest{
	ID:   0x00E,
	Name: "ACU Cell Data Two Test 2",
	Data: []byte{
		0xce, 0xcb, 0xea, 0x11, 0x98, 0xc4, 0x8a, 0xa9, 0x61, 0x5f,
		0x06, 0xac, 0xbd, 0xf8, 0x42, 0xb1, 0xec, 0x77, 0xfa, 0xa7,
		0xc9, 0xc5, 0x2d, 0x53, 0xd7, 0x77, 0xf6, 0xdc, 0xf3, 0x1c,
		0x97, 0x50, 0x71, 0xd7, 0x5f, 0xe3, 0x8c, 0xc7, 0x34, 0x30,
		0xa5, 0x38, 0xd0, 0x34, 0x50, 0x10, 0x7c, 0x8e, 0x71, 0xc0,
		0xb3, 0x84, 0xf1, 0x09, 0xbc, 0xf1, 0xcc, 0x95, 0xd6, 0xc2,
		0x89, 0x64, 0x60, 0x35},
	ExpectedValues: map[string]interface{}{
		"cell32_voltage": 4.06,
		"cell32_temp":    50.75,
		"cell33_voltage": 4.34,
		"cell33_temp":    4.25,
		"cell34_voltage": 3.52,
		"cell34_temp":    49,
		"cell35_voltage": 3.38,
		"cell35_temp":    42.25,
		"cell36_voltage": 2.97,
		"cell36_temp":    23.75,
		"cell37_voltage": 2.06,
		"cell37_temp":    43,
		"cell38_voltage": 3.89,
		"cell38_temp":    62,
		"cell39_voltage": 2.66,
		"cell39_temp":    44.25,
		"cell40_voltage": 4.36,
		"cell40_temp":    29.75,
		"cell41_voltage": 4.5,
		"cell41_temp":    41.75,
		"cell42_voltage": 4.01,
		"cell42_temp":    49.25,
		"cell43_voltage": 2.45,
		"cell43_temp":    20.75,
		"cell44_voltage": 4.15,
		"cell44_temp":    29.75,
		"cell45_voltage": 4.46,
		"cell45_temp":    55.0,
		"cell46_voltage": 4.43,
		"cell46_temp":    7.0,
		"cell47_voltage": 3.51,
		"cell47_temp":    20.0,
		"cell48_voltage": 3.13,
		"cell48_temp":    53.75,
		"cell49_voltage": 2.95,
		"cell49_temp":    56.75,
		"cell50_voltage": 3.4,
		"cell50_temp":    49.75,
		"cell51_voltage": 2.52,
		"cell51_temp":    12.0,
		"cell52_voltage": 3.65,
		"cell52_temp":    14.0,
		"cell53_voltage": 4.08,
		"cell53_temp":    13.0,
		"cell54_voltage": 2.8,
		"cell54_temp":    4.0,
		"cell55_voltage": 3.24,
		"cell55_temp":    35.5,
		"cell56_voltage": 3.13,
		"cell56_temp":    48.0,
		"cell57_voltage": 3.79,
		"cell57_temp":    33.0,
		"cell58_voltage": 4.41,
		"cell58_temp":    2.25,
		"cell59_voltage": 3.88,
		"cell59_temp":    60.25,
		"cell60_voltage": 4.04,
		"cell60_temp":    37.25,
		"cell61_voltage": 4.14,
		"cell61_temp":    48.5,
		"cell62_voltage": 3.37,
		"cell62_temp":    25.0,
		"cell63_voltage": 2.96,
		"cell63_temp":    13.25,
	},
}

var ACUCellDataThreeTest1 = MessageTest{
	ID:   0x00F,
	Name: "ACU Cell Data Three Test 1",
	Data: []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"cell64_voltage": 2,
		"cell64_temp":    0,
		"cell65_voltage": 2,
		"cell65_temp":    0,
		"cell66_voltage": 2,
		"cell66_temp":    0,
		"cell67_voltage": 2,
		"cell67_temp":    0,
		"cell68_voltage": 2,
		"cell68_temp":    0,
		"cell69_voltage": 2,
		"cell69_temp":    0,
		"cell70_voltage": 2,
		"cell70_temp":    0,
		"cell71_voltage": 2,
		"cell71_temp":    0,
		"cell72_voltage": 2,
		"cell72_temp":    0,
		"cell73_voltage": 2,
		"cell73_temp":    0,
		"cell74_voltage": 2,
		"cell74_temp":    0,
		"cell75_voltage": 2,
		"cell75_temp":    0,
		"cell76_voltage": 2,
		"cell76_temp":    0,
		"cell77_voltage": 2,
		"cell77_temp":    0,
		"cell78_voltage": 2,
		"cell78_temp":    0,
		"cell79_voltage": 2,
		"cell79_temp":    0,
		"cell80_voltage": 2,
		"cell80_temp":    0,
		"cell81_voltage": 2,
		"cell81_temp":    0,
		"cell82_voltage": 2,
		"cell82_temp":    0,
		"cell83_voltage": 2,
		"cell83_temp":    0,
		"cell84_voltage": 2,
		"cell84_temp":    0,
		"cell85_voltage": 2,
		"cell85_temp":    0,
		"cell86_voltage": 2,
		"cell86_temp":    0,
		"cell87_voltage": 2,
		"cell87_temp":    0,
		"cell88_voltage": 2,
		"cell88_temp":    0,
		"cell89_voltage": 2,
		"cell89_temp":    0,
		"cell90_voltage": 2,
		"cell90_temp":    0,
		"cell91_voltage": 2,
		"cell91_temp":    0,
		"cell92_voltage": 2,
		"cell92_temp":    0,
		"cell93_voltage": 2,
		"cell93_temp":    0,
		"cell94_voltage": 2,
		"cell94_temp":    0,
		"cell95_voltage": 2,
		"cell95_temp":    0,
	},
}

var ACUCellDataThreeTest2 = MessageTest{
	ID:   0x00F,
	Name: "ACU Cell Data Three Test 2",
	Data: []byte{
		0xdc, 0x13, 0xfa, 0xf8, 0x2b, 0xd1, 0xea, 0x2a, 0x4e, 0x16,
		0x17, 0xca, 0x49, 0x9b, 0x66, 0x07, 0xa0, 0xf0, 0xfe, 0x38,
		0x97, 0xf5, 0xbb, 0xa0, 0x07, 0x5e, 0xfe, 0x58, 0x7c, 0x40,
		0x6e, 0xaf, 0x0e, 0x77, 0x2f, 0x89, 0xc1, 0xa3, 0xcb, 0x24,
		0x1f, 0x31, 0xe6, 0x87, 0x6f, 0xc5, 0x97, 0x84, 0xab, 0xf2,
		0x58, 0xfe, 0xfa, 0xf6, 0x4c, 0x9d, 0x00, 0xd7, 0x29, 0x4c,
		0x50, 0x81, 0xc9, 0x20},
	ExpectedValues: map[string]interface{}{
		"cell64_voltage": 4.2,
		"cell64_temp":    4.75,
		"cell65_voltage": 4.5,
		"cell65_temp":    62.0,
		"cell66_voltage": 2.43,
		"cell66_temp":    52.25,
		"cell67_voltage": 4.34,
		"cell67_temp":    10.5,
		"cell68_voltage": 2.78,
		"cell68_temp":    5.5,
		"cell69_voltage": 2.23,
		"cell69_temp":    50.5,
		"cell70_voltage": 2.73,
		"cell70_temp":    38.75,
		"cell71_voltage": 3.02,
		"cell71_temp":    1.75,
		"cell72_voltage": 3.6,
		"cell72_temp":    60.0,
		"cell73_voltage": 4.54,
		"cell73_temp":    14.0,
		"cell74_voltage": 3.51,
		"cell74_temp":    61.25,
		"cell75_voltage": 3.87,
		"cell75_temp":    40.0,
		"cell76_voltage": 2.07,
		"cell76_temp":    23.5,
		"cell77_voltage": 4.54,
		"cell77_temp":    22.0,
		"cell78_voltage": 3.24,
		"cell78_temp":    16.0,
		"cell79_voltage": 3.1,
		"cell79_temp":    43.75,
		"cell80_voltage": 2.14,
		"cell80_temp":    29.75,
		"cell81_voltage": 2.47,
		"cell81_temp":    34.25,
		"cell82_voltage": 3.93,
		"cell82_temp":    40.75,
		"cell83_voltage": 4.03,
		"cell83_temp":    9.0,
		"cell84_voltage": 2.31,
		"cell84_temp":    12.25,
		"cell85_voltage": 4.3,
		"cell85_temp":    33.75,
		"cell86_voltage": 3.11,
		"cell86_temp":    49.25,
		"cell87_voltage": 3.51,
		"cell87_temp":    33.0,
		"cell88_voltage": 3.71,
		"cell88_temp":    60.5,
		"cell89_voltage": 2.88,
		"cell89_temp":    63.5,
		"cell90_voltage": 4.5,
		"cell90_temp":    61.5,
		"cell91_voltage": 2.76,
		"cell91_temp":    39.25,
		"cell92_voltage": 2.0,
		"cell92_temp":    53.75,
		"cell93_voltage": 2.41,
		"cell93_temp":    19.0,
		"cell94_voltage": 2.8,
		"cell94_temp":    32.25,
		"cell95_voltage": 4.01,
		"cell95_temp":    8.0,
	},
}

var ACUCellDataFourTest1 = MessageTest{
	ID:   0x010,
	Name: "ACU Cell Data Four Test 1",
	Data: []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"cell96_voltage":  2,
		"cell96_temp":     0,
		"cell97_voltage":  2,
		"cell97_temp":     0,
		"cell98_voltage":  2,
		"cell98_temp":     0,
		"cell99_voltage":  2,
		"cell99_temp":     0,
		"cell100_voltage": 2,
		"cell100_temp":    0,
		"cell101_voltage": 2,
		"cell101_temp":    0,
		"cell102_voltage": 2,
		"cell102_temp":    0,
		"cell103_voltage": 2,
		"cell103_temp":    0,
		"cell104_voltage": 2,
		"cell104_temp":    0,
		"cell105_voltage": 2,
		"cell105_temp":    0,
		"cell106_voltage": 2,
		"cell106_temp":    0,
		"cell107_voltage": 2,
		"cell107_temp":    0,
		"cell108_voltage": 2,
		"cell108_temp":    0,
		"cell109_voltage": 2,
		"cell109_temp":    0,
		"cell110_voltage": 2,
		"cell110_temp":    0,
		"cell111_voltage": 2,
		"cell111_temp":    0,
		"cell112_voltage": 2,
		"cell112_temp":    0,
		"cell113_voltage": 2,
		"cell113_temp":    0,
		"cell114_voltage": 2,
		"cell114_temp":    0,
		"cell115_voltage": 2,
		"cell115_temp":    0,
		"cell116_voltage": 2,
		"cell116_temp":    0,
		"cell117_voltage": 2,
		"cell117_temp":    0,
		"cell118_voltage": 2,
		"cell118_temp":    0,
		"cell119_voltage": 2,
		"cell119_temp":    0,
		"cell120_voltage": 2,
		"cell120_temp":    0,
		"cell121_voltage": 2,
		"cell121_temp":    0,
		"cell122_voltage": 2,
		"cell122_temp":    0,
		"cell123_voltage": 2,
		"cell123_temp":    0,
		"cell124_voltage": 2,
		"cell124_temp":    0,
		"cell125_voltage": 2,
		"cell125_temp":    0,
		"cell126_voltage": 2,
		"cell126_temp":    0,
		"cell127_voltage": 2,
		"cell127_temp":    0,
	},
}

var ACUCellDataFourTest2 = MessageTest{
	ID:   0x010,
	Name: "ACU Cell Data Four Test 2",
	Data: []byte{
		0x49, 0xe6, 0xb1, 0x1b, 0xe5, 0x98, 0x12, 0x68, 0xf9, 0x95,
		0x54, 0x2d, 0x0d, 0x3d, 0x75, 0xc7, 0xfd, 0x5b, 0xb5, 0x92,
		0x06, 0x4e, 0x1b, 0x86, 0x5b, 0x9a, 0xb1, 0x20, 0x95, 0xd1,
		0x82, 0xbf, 0x50, 0x7e, 0x18, 0x7b, 0x50, 0x57, 0x4c, 0xc3,
		0xce, 0xa9, 0x87, 0x23, 0x87, 0x34, 0x13, 0x8d, 0xcc, 0xcb,
		0xd2, 0x13, 0x3a, 0xec, 0xa1, 0xe4, 0x1d, 0x5f, 0x0c, 0xc9,
		0x4f, 0xe8, 0x09, 0x20},
	ExpectedValues: map[string]interface{}{
		"cell96_voltage":  2.73,
		"cell96_temp":     57.5,
		"cell97_voltage":  3.77,
		"cell97_temp":     6.75,
		"cell98_voltage":  4.29,
		"cell98_temp":     38.0,
		"cell99_voltage":  2.18,
		"cell99_temp":     26.0,
		"cell100_voltage": 4.49,
		"cell100_temp":    37.25,
		"cell101_voltage": 2.84,
		"cell101_temp":    11.25,
		"cell102_voltage": 2.13,
		"cell102_temp":    15.25,
		"cell103_voltage": 3.17,
		"cell103_temp":    49.75,
		"cell104_voltage": 4.53,
		"cell104_temp":    22.75,
		"cell105_voltage": 3.81,
		"cell105_temp":    36.5,
		"cell106_voltage": 2.06,
		"cell106_temp":    19.5,
		"cell107_voltage": 2.27,
		"cell107_temp":    33.5,
		"cell108_voltage": 2.91,
		"cell108_temp":    38.5,
		"cell109_voltage": 3.77,
		"cell109_temp":    8.0,
		"cell110_voltage": 3.49,
		"cell110_temp":    52.25,
		"cell111_voltage": 3.3,
		"cell111_temp":    47.75,
		"cell112_voltage": 2.8,
		"cell112_temp":    31.5,
		"cell113_voltage": 2.24,
		"cell113_temp":    30.75,
		"cell114_voltage": 2.8,
		"cell114_temp":    21.75,
		"cell115_voltage": 2.76,
		"cell115_temp":    48.75,
		"cell116_voltage": 4.06,
		"cell116_temp":    42.25,
		"cell117_voltage": 3.35,
		"cell117_temp":    8.75,
		"cell118_voltage": 3.35,
		"cell118_temp":    13.0,
		"cell119_voltage": 2.19,
		"cell119_temp":    35.25,
		"cell120_voltage": 4.04,
		"cell120_temp":    50.75,
		"cell121_voltage": 4.1,
		"cell121_temp":    4.75,
		"cell122_voltage": 2.58,
		"cell122_temp":    59.0,
		"cell123_voltage": 3.61,
		"cell123_temp":    57.0,
		"cell124_voltage": 2.29,
		"cell124_temp":    23.75,
		"cell125_voltage": 2.12,
		"cell125_temp":    50.25,
		"cell126_voltage": 2.79,
		"cell126_temp":    58.0,
		"cell127_voltage": 2.09,
		"cell127_temp":    8.0,
	},
}

var ACUCellDataFiveTest1 = MessageTest{
	ID:   0x011,
	Name: "ACU Cell Data Five Test 1",
	Data: []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"cell128_voltage": 2,
		"cell128_temp":    0,
		"cell129_voltage": 2,
		"cell129_temp":    0,
		"cell130_voltage": 2,
		"cell130_temp":    0,
		"cell131_voltage": 2,
		"cell131_temp":    0,
		"cell132_voltage": 2,
		"cell132_temp":    0,
		"cell133_voltage": 2,
		"cell133_temp":    0,
		"cell134_voltage": 2,
		"cell134_temp":    0,
		"cell135_voltage": 2,
		"cell135_temp":    0,
		"cell136_voltage": 2,
		"cell136_temp":    0,
		"cell137_voltage": 2,
		"cell137_temp":    0,
		"cell138_voltage": 2,
		"cell138_temp":    0,
		"cell139_voltage": 2,
		"cell139_temp":    0,
		"cell140_voltage": 2,
		"cell140_temp":    0,
		"cell141_voltage": 2,
		"cell141_temp":    0,
		"cell142_voltage": 2,
		"cell142_temp":    0,
		"cell143_voltage": 2,
		"cell143_temp":    0,
		"cell144_voltage": 2,
		"cell144_temp":    0,
		"cell145_voltage": 2,
		"cell145_temp":    0,
		"cell146_voltage": 2,
		"cell146_temp":    0,
		"cell147_voltage": 2,
		"cell147_temp":    0,
		"cell148_voltage": 2,
		"cell148_temp":    0,
		"cell149_voltage": 2,
		"cell149_temp":    0,
		"cell150_voltage": 2,
		"cell150_temp":    0,
		"cell151_voltage": 2,
		"cell151_temp":    0,
		"cell152_voltage": 2,
		"cell152_temp":    0,
		"cell153_voltage": 2,
		"cell153_temp":    0,
		"cell154_voltage": 2,
		"cell154_temp":    0,
		"cell155_voltage": 2,
		"cell155_temp":    0,
		"cell156_voltage": 2,
		"cell156_temp":    0,
		"cell157_voltage": 2,
		"cell157_temp":    0,
		"cell158_voltage": 2,
		"cell158_temp":    0,
		"cell159_voltage": 2,
		"cell159_temp":    0,
	},
}

var ACUCellDataFiveTest2 = MessageTest{
	ID:   0x011,
	Name: "ACU Cell Data Five Test 2",
	Data: []byte{
		0x5b, 0x8b, 0x4f, 0x78, 0x56, 0xb7, 0x03, 0x27, 0xf4, 0x6d,
		0x88, 0xf4, 0x70, 0xd9, 0xa0, 0x81, 0x88, 0xe0, 0x0b, 0xa7,
		0x06, 0xc9, 0x39, 0x2d, 0x4f, 0xca, 0xda, 0x84, 0xd9, 0x68,
		0xc9, 0xd0, 0x9c, 0x75, 0xe9, 0x58, 0xd3, 0x05, 0x5c, 0x95,
		0x5a, 0x1c, 0xb0, 0x81, 0xe6, 0xf0, 0xd3, 0xb7, 0x25, 0x13,
		0xc6, 0xd2, 0xf5, 0xa1, 0xbe, 0x5c, 0x36, 0x1a, 0x3c, 0xfd,
		0x69, 0x27, 0xf1, 0x32},
	ExpectedValues: map[string]interface{}{
		"cell128_voltage": 2.91,
		"cell128_temp":    34.75,
		"cell129_voltage": 2.79,
		"cell129_temp":    30.0,
		"cell130_voltage": 2.86,
		"cell130_temp":    45.75,
		"cell131_voltage": 2.03,
		"cell131_temp":    9.75,
		"cell132_voltage": 4.44,
		"cell132_temp":    27.25,
		"cell133_voltage": 3.36,
		"cell133_temp":    61.0,
		"cell134_voltage": 3.12,
		"cell134_temp":    54.25,
		"cell135_voltage": 3.6,
		"cell135_temp":    32.25,
		"cell136_voltage": 3.36,
		"cell136_temp":    56.0,
		"cell137_voltage": 2.11,
		"cell137_temp":    41.75,
		"cell138_voltage": 2.06,
		"cell138_temp":    50.25,
		"cell139_voltage": 2.57,
		"cell139_temp":    11.25,
		"cell140_voltage": 2.79,
		"cell140_temp":    50.5,
		"cell141_voltage": 4.18,
		"cell141_temp":    33.0,
		"cell142_voltage": 4.17,
		"cell142_temp":    26.0,
		"cell143_voltage": 4.01,
		"cell143_temp":    52.0,
		"cell144_voltage": 3.56,
		"cell144_temp":    29.25,
		"cell145_voltage": 4.33,
		"cell145_temp":    22.0,
		"cell146_voltage": 4.11,
		"cell146_temp":    1.25,
		"cell147_voltage": 2.92,
		"cell147_temp":    37.25,
		"cell148_voltage": 2.9,
		"cell148_temp":    7.0,
		"cell149_voltage": 3.76,
		"cell149_temp":    32.25,
		"cell150_voltage": 4.3,
		"cell150_temp":    60.0,
		"cell151_voltage": 4.11,
		"cell151_temp":    45.75,
		"cell152_voltage": 2.37,
		"cell152_temp":    4.75,
		"cell153_voltage": 3.98,
		"cell153_temp":    52.5,
		"cell154_voltage": 4.45,
		"cell154_temp":    40.25,
		"cell155_voltage": 3.9,
		"cell155_temp":    23.0,
		"cell156_voltage": 2.54,
		"cell156_temp":    6.5,
		"cell157_voltage": 2.6,
		"cell157_temp":    63.25,
		"cell158_voltage": 3.05,
		"cell158_temp":    9.75,
		"cell159_voltage": 4.41,
		"cell159_temp":    12.5,
	},
}

var DCDCStatusTest1 = MessageTest{
	ID:   0x012,
	Name: "DC-DC Status Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"input_voltage":  0,
		"output_voltage": 0,
		"input_current":  0,
		"output_current": 0,
		"dc_dc_temp":     0,
	},
}

var DCDCStatusTest2 = MessageTest{
	ID:   0x012,
	Name: "DC-DC Status Test 2",
	Data: []byte{0xff, 0x71, 0x52, 0x2b, 0x63, 0xbb, 0xd6},
	ExpectedValues: map[string]interface{}{
		"input_voltage":  29.183,
		"output_voltage": 11.09,
		"input_current":  9.9,
		"output_current": 18.7,
		"dc_dc_temp":     214,
	},
}

var InverterStatusOneTest1 = MessageTest{
	ID:   0x013,
	Name: "Inverter Status One Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"AC_current": -327.68,
		"DC_current": 0,
		"motor_RPM":  -32768,
	},
}

var InverterStatusOneTest2 = MessageTest{
	ID:   0x013,
	Name: "Inverter Status One Test 2",
	Data: []byte{0x19, 0x33, 0xec, 0x87, 0x0d, 0xb1}, // add
	ExpectedValues: map[string]interface{}{
		"AC_current": 7.93,
		"DC_current": 347.96,
		"motor_RPM":  12557,
	},
}

// ? can sheet length/unit is wrong
var InverterStatusTwoTest1 = MessageTest{
	ID:   0x014,
	Name: "Inverter Status Two Test 1",
	Data: []byte{0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"U_MOSFET_temperature": -40,
		"V_MOSFET_temperature": -40,
		"W_MOSFET_temperature": -40,
	},
}

var InverterStatusTwoTest2 = MessageTest{
	ID:   0x014,
	Name: "Inverter Status Two Test 2",
	Data: []byte{0x73, 0x75, 0x04},
	ExpectedValues: map[string]interface{}{
		"U_MOSFET_temperature": 33,
		"V_MOSFET_temperature": 35,
		"W_MOSFET_temperature": -36,
	},
}

var InverterStatusThreeTest1 = MessageTest{
	ID:   0x015,
	Name: "Inverter Status Three Test 1",
	Data: []byte{0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"motor_temperature":    -40,
		"over_voltage_faults":  0,
		"under_voltage_fault":  0,
		"inv._overtemp_fault":  0,
		"motor_overtemp_fault": 0,
		"transistor_fault":     0,
		"encoder_fault":        0,
		"CAN_fault":            0,
		"future_use":           0,
	},
}

var InverterStatusThreeTest2 = MessageTest{
	ID:   0x015,
	Name: "Inverter Status Three Test 2",
	Data: []byte{0xf2, 0x2f},
	ExpectedValues: map[string]interface{}{
		"motor_temperature":    202,
		"over_voltage_faults":  0,
		"under_voltage_fault":  0,
		"inv._overtemp_fault":  1,
		"motor_overtemp_fault": 0,
		"transistor_fault":     1,
		"encoder_fault":        1,
		"CAN_fault":            1,
		"future_use":           1,
	},
}

var InverterConfigTest1 = MessageTest{
	ID:   0x016,
	Name: "Inverter Config Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"max_AC_current":         -327.68,
		"max_DC_current":         -327.68,
		"absolute_max_RPM_limit": -32768,
		"motor_direction":        0,
	},
}

var InverterConfigTest2 = MessageTest{
	ID:   0x016,
	Name: "Inverter Config Test 2",
	Data: []byte{0xfa, 0xa8, 0x54, 0xd3, 0x9f, 0xed, 0x01},
	ExpectedValues: map[string]interface{}{
		"max_AC_current":         104.9,
		"max_DC_current":         213.32,
		"absolute_max_RPM_limit": 38063,
		"motor_direction":        1,
	},
}

var InverterCommandTest1 = MessageTest{
	ID:   0x017,
	Name: "Inverter Command Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"set_AC_current":  -327.68,
		"set_DC_current":  -327.68,
		"RPM_limit":       -32768,
		"field_weakening": 0,
		"drive_enable":    0,
	},
}

var InverterCommandTest2 = MessageTest{
	ID:   0x017,
	Name: "Inverter Command Test 2",
	Data: []byte{0x66, 0x70, 0x0b, 0x3c, 0x0f, 0x4a, 0xed, 0x01},
	ExpectedValues: map[string]interface{}{
		"set_AC_current":  -39.94,
		"set_DC_current":  -173.97,
		"RPM_limit":       -13809,
		"field_weakening": 23.7,
		"drive_enable":    1,
	},
}

var FanStatusTest1 = MessageTest{
	ID:   0x018,
	Name: "Fan Status Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"fan_speed":     0,
		"input_voltage": 0,
		"input_current": 0,
	},
}

var FanStatusTest2 = MessageTest{
	ID:   0x018,
	Name: "Fan Status Test 2",
	Data: []byte{0xdf, 0x41, 0x9d, 0x1d},
	ExpectedValues: map[string]interface{}{
		"fan_speed":     16863,
		"input_voltage": 15.7,
		"input_current": 2.9,
	},
}
var FanCommandTest1 = MessageTest{
	ID:   0x019,
	Name: "Fan Command Test 1",
	Data: []byte{0x00},
	ExpectedValues: map[string]interface{}{
		"fan_command": 0,
	},
}

var FanCommandTest2 = MessageTest{
	ID:   0x019,
	Name: "Fan Command Test 2",
	Data: []byte{0xd3},
	ExpectedValues: map[string]interface{}{
		"fan_command": 211,
	},
}

var DashStatusTest1 = MessageTest{
	ID:   0x01A,
	Name: "Dash Status Test 1",
	Data: []byte{0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"BMS_LED":         0,
		"IMD_LED":         0,
		"BSPD_LED":        0,
		"TS_button_data":  0,
		"RTD_button_data": 0,
	},
}

var DashStatusTest2 = MessageTest{
	ID:   0x01A,
	Name: "Dash Status Test 2",
	Data: []byte{0x05, 0x1e, 0xF4},
	ExpectedValues: map[string]interface{}{
		"BMS_LED":         1,
		"IMD_LED":         0,
		"BSPD_LED":        1,
		"TS_button_data":  3,
		"RTD_button_data": 1.2, // check, 2s complement -12
	},
}

var DashConfigTest1 = MessageTest{
	ID:   0x01B,
	Name: "Dash Config Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"BMS_LED":        0,
		"IMD_LED":        0,
		"BSPD_LED":       0,
		"button_LED_1_R": 0,
		"button_LED_1_G": 0,
		"button_LED_1_B": 0,
		"button_LED_2_R": 0,
		"button_LED_2_G": 0,
		"button_LED_2_B": 0,
	},
}

var DashConfigTest2 = MessageTest{
	ID:   0x01B,
	Name: "Dash Config Test 2",
	Data: []byte{0x07, 0xd4, 0x8d, 0xde, 0x22, 0x20, 0xc5},
	ExpectedValues: map[string]interface{}{
		"BMS_LED":        1,
		"IMD_LED":        1,
		"BSPD_LED":       1,
		"button_LED_1_R": 212,
		"button_LED_1_G": 141,
		"button_LED_1_B": 222,
		"button_LED_2_R": 34,
		"button_LED_2_G": 32,
		"button_LED_2_B": 197,
	},
}

// // check
// var SteeringStatusTest1 = MessageTest{
// 	ID:   0x01C,
// 	Name: "Steering Status Test 1",
// 	Data: []byte{0x00, 0x00},
// 	ExpectedValues: map[string]interface{}{
// 		"current_encoder":    0, // ??? Spreadsheet scaling is inconsistent with min/max
// 		"torque_map_encoder": 0, // ??? Spreadsheet scaling is inconsistent with min/max
// 		"regen":              0, // ??? Spreadsheet scaling is inconsistent with min/max
// 		"button_1":           0,
// 		"button_2":           0,
// 		"button_3":           0,
// 		"button_4":           0,
// 	},
// }

// // check and fill
// var SteeringStatusTest2 = MessageTest{
// 	ID:   0x01C,
// 	Name: "Steering Status Test 2",
// 	Data: []byte{0x00, 0x00}, // need to fill
// 	ExpectedValues: map[string]interface{}{
// 		"current_encoder":    0, // ??? Spreadsheet scaling is inconsistent with min/max
// 		"torque_map_encoder": 0, // ??? Spreadsheet scaling is inconsistent with min/max
// 		"regen":              0, // ??? Spreadsheet scaling is inconsistent with min/max
// 		"button_1":           0,
// 		"button_2":           0,
// 		"button_3":           0,
// 		"button_4":           0,
// 	},
// }

// // reserved
// var SteeringConfigTest1 = MessageTest{
// 	ID:   0x01D,
// 	Name: "Steering Config Test 1",
// 	Data: []byte{0x00},
// 	ExpectedValues: map[string]interface{}{
// 		"reserved": 0, // ??? what
// 	},
// }

// var SteeringConfigTest2 = MessageTest{
// 	ID:   0x01D,
// 	Name: "Steering Config Test 2",
// 	Data: []byte{0x00},
// 	ExpectedValues: map[string]interface{}{
// 		"reserved": 0, // ??? what
// 	},
// }

var SAMBrakeIRTest1 = MessageTest{
	ID:   0x01E,
	Name: "SAM Break IR Test 1",
	Data: []byte{0x00},
	ExpectedValues: map[string]interface{}{
		"temp": 0,
	},
}

var SAMBrakeIRTest2 = MessageTest{
	ID:   0x01E,
	Name: "SAM Break IR Test 2",
	Data: []byte{0x89},
	ExpectedValues: map[string]interface{}{
		"temp": 137,
	},
}

var SAMTireTempTest1 = MessageTest{
	ID:   0x01F,
	Name: "SAM Tire Temp Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"outside_temp":        0,
		"outside_middle_temp": 0,
		"inside_middle_temp":  0,
		"inside_temp":         0,
	},
}

var SAMTireTempTest2 = MessageTest{
	ID:   0x01F,
	Name: "SAM Tire Temp Test 2",
	Data: []byte{0xfa, 0xb1, 0xda, 0x0a},
	ExpectedValues: map[string]interface{}{
		"outside_temp":        250,
		"outside_middle_temp": 177,
		"inside_middle_temp":  218,
		"inside_temp":         10,
	},
}

var SAMIMUTest1 = MessageTest{
	ID:   0x020,
	Name: "SAM IMU Test 1",
	Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"acceleration_in_X-axis":     -327.68,
		"acceleration_in_Y-axis":     -327.68,
		"acceleration_in_Z-axis":     -327.68,
		"angular_velocity_in_X-axis": -32.768,
		"angular_velocity_in_Y-axis": -32.768,
		"angular_velocity_in_Z-axis": -32.768,
	},
}
var SAMIMUTest2 = MessageTest{
	ID:   0x020,
	Name: "SAM IMU Test 2",
	Data: []byte{0x7d, 0xc2, 0xb7, 0x61, 0xaf, 0x20, 0x2e, 0x72, 0x5d, 0xac, 0xaf, 0xa1},
	ExpectedValues: map[string]interface{}{
		"acceleration_in_X-axis":     -170.21,
		"acceleration_in_Y-axis":     -77.53,
		"acceleration_in_Z-axis":     -244.01,
		"angular_velocity_in_X-axis": -3.538,
		"angular_velocity_in_Y-axis": 11.357,
		"angular_velocity_in_Z-axis": 8.623,
	},
}

// var SAMGPSOneTest1 = MessageTest{
// 	ID:             0x021,
// 	Name:           "SAM GPS 1 Test 1",
// 	Data:           []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
// 	ExpectedValues: map[string]interface{}{
// 		// what
// 	},
// }

// var SAMGPSOneTest2 = MessageTest{
// 	ID:             0x021,
// 	Name:           "SAM GPS 1 Test 2",
// 	Data:           []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // change
// 	ExpectedValues: map[string]interface{}{
// 		// what
// 	},
// }

// var SAMGPSTwoTest1 = MessageTest{
// 	ID:             0x022,
// 	Name:           "SAM GPS 2 Test 1",
// 	Data:           []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
// 	ExpectedValues: map[string]interface{}{
// 		// what
// 	},
// }

// var SAMGPSTwoTest2 = MessageTest{
// 	ID:             0x022,
// 	Name:           "SAM GPS 2 Test 2",
// 	Data:           []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, // change
// 	ExpectedValues: map[string]interface{}{
// 		// what
// 	},
// }
// var SAMGPSTimeTest1 = MessageTest{
// 	ID:             0x023,
// 	Name:           "SAM GPS Time Test 1",
// 	Data:           []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
// 	ExpectedValues: map[string]interface{}{
// 		// what
// 	},
// }

// var SAMGPSTimeTest2 = MessageTest{
// 	ID:             0x023,
// 	Name:           "SAM GPS Time Test 2",
// 	Data:           []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
// 	ExpectedValues: map[string]interface{}{
// 		// what
// 	},
// }

// var SAMGPSHeadingTest1 = MessageTest{
// 	ID:             0x024,
// 	Name:           "SAM GPS Heading Test 1",
// 	Data:           []byte{0x00, 0x00, 0x00, 0x00},
// 	ExpectedValues: map[string]interface{}{
// 		// what
// 	},
// }

// var SAMGPSHeadingTest2 = MessageTest{
// 	ID:             0x024,
// 	Name:           "SAM GPS Heading Test 2",
// 	Data:           []byte{0x00, 0x00, 0x00, 0x00},
// 	ExpectedValues: map[string]interface{}{
// 		// what
// 	},
// }

var SAMSusPotsTest1 = MessageTest{
	ID:   0x025,
	Name: "SAM Sus Pots Test 1",
	Data: []byte{0x00},
	ExpectedValues: map[string]interface{}{
		"suspension_angle": 0,
	},
}

var SAMSusPotsTest2 = MessageTest{
	ID:   0x025,
	Name: "SAM Sus Pots Test 2",
	Data: []byte{0x0f},
	ExpectedValues: map[string]interface{}{
		"suspension_angle": 15,
	},
}

var SAMTOFTest1 = MessageTest{
	ID:   0x026,
	Name: "SAM TOF Test 1",
	Data: []byte{0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"height": 0,
	},
}

var SAMTOFTest2 = MessageTest{
	ID:   0x026,
	Name: "SAM TOF Test 2",
	Data: []byte{0x00, 0x82},
	ExpectedValues: map[string]interface{}{
		"height": 130,
	},
}

var SAMRearWheelspeedTest1 = MessageTest{
	ID:   0x027,
	Name: "SAM Rear Wheelspeed Test 1",
	Data: []byte{0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"speed": -3276.8,
	},
}

var SAMRearWheelspeedTest2 = MessageTest{
	ID:   0x027,
	Name: "SAM Rear Wheelspeed Test 2",
	Data: []byte{0x0f, 0x44},
	ExpectedValues: map[string]interface{}{
		"speed": -1534.5,
	},
}

var SAMPushrodForceTest1 = MessageTest{
	ID:   0x028,
	Name: "SAM Pushrod Force Test 1",
	Data: []byte{0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"load_force": -3276.8,
	},
}
var SAMPushrodForceTest2 = MessageTest{
	ID:   0x028,
	Name: "SAM Pushrod Force Test 2",
	Data: []byte{0x27, 0xe9},
	ExpectedValues: map[string]interface{}{
		"load_force": 2691.9,
	},
}

var SpecificBrakeIRTest1 = MessageTest{
	ID:   0x02C,
	Name: "Specific Brake IR Test 1",
	Data: []byte{0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"wheel_identifier": 0,
		"temp":             0,
	},
}

var SpecificBrakeIRTest2 = MessageTest{
	ID:   0x02C,
	Name: "Specific Brake IR Test 2",
	Data: []byte{0x01, 0x12},
	ExpectedValues: map[string]interface{}{
		"wheel_identifier": 1,
		"temp":             18,
	},
}

// bottom ones skipped
// tcm status
// dash warning flags
// ecu ping information

// need to add ecu pedals data, spreadsheet data names are weird

var TCMResourceTest1 = MessageTest{
	ID:   0x02A,
	Name: "TCM Resource Utilization Test 1",
	Data: []byte{0x35, 0x1b, 0x58, 0x03, 0xB8, 0x4b, 0x59},
	ExpectedValues: map[string]interface{}{
		"CPU_util":    53,
		"GPU_util":    27,
		"RAM_util":    88,
		"disk_util":   3,
		"power_usage": 18.4,
		"CPU_temp":    75,
		"GPU_temp":    89,
	},
}

var TCMResourceTest2 = MessageTest{
	ID:   0x02A,
	Name: "TCM Resource Utilization Test 2",
	Data: []byte{0x00, 0x00, 0x01, 0x02, 0x03, 0x00, 0x00},
	ExpectedValues: map[string]interface{}{
		"CPU_util":    0,
		"GPU_util":    0,
		"RAM_util":    1,
		"disk_util":   2,
		"power_usage": 0.3,
		"CPU_temp":    0,
		"GPU_temp":    0,
	},
}

var TCMResourceTest3 = MessageTest{
	ID:   0x02A,
	Name: "TCM Resource Utilization Test 3",
	Data: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
	ExpectedValues: map[string]interface{}{
		"CPU_util":    255,
		"GPU_util":    255,
		"RAM_util":    255,
		"disk_util":   255,
		"power_usage": 25.5,
		"CPU_temp":    255,
		"GPU_temp":    255,
	},
}
