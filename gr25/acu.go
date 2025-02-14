package gr25

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

func RunACUTests(mqttClient *mqtt.Client, db *gorm.DB) {
	SendACUStatusOne(mqttClient, db)
	SendACUStatusTwo(mqttClient, db)
	SendACUStatusThree(mqttClient, db)
	SendACUPrecharge(mqttClient, db)
	SendACUConfigChargeParameters(mqttClient, db)
	SendACUConfigOperationalParameters(mqttClient, db)
	SendACUCellDataOne(mqttClient, db)
	SendACUCellDataTwo(mqttClient, db)
	SendACUCellDataThree(mqttClient, db)
	SendACUCellDataFour(mqttClient, db)
	SendACUCellDataFive(mqttClient, db)
	SendDCDCStatus(mqttClient, db)
}

func SendACUStatusOne(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x007,
		Name: "ACU Status One Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"acu_accumulator_voltage": 0,
			"acu_ts_voltage":          0,
			"acu_accumulator_current": 0,
			"acu_accumulator_soc":     0,
			"acu_glv_soc":             0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendACUStatusTwo(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x008,
		Name: "ACU Status Two Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"acu_20v_voltage":               0,
			"acu_12v_voltage":               0,
			"acu_sdc_voltage":               0,
			"acu_min_cell_voltage":          0,
			"acu_max_cell_temp":             0,
			"acu_over_temp_error":           0,
			"acu_over_voltage_error":        0,
			"acu_under_voltage_error":       0,
			"acu_over_current_error":        0,
			"acu_under_current_error":       0,
			"acu_under_voltage_20v_warning": 0,
			"acu_under_voltage_12v_warning": 0,
			"acu_under_voltage_sdc_warning": 0,
			"acu_precharge_error":           0,
			"acu_ir_minus_precharge_state":  0, //idk if these are correct
			"acu_ir_plus_state":             0, //idk if these are correct
			"acu_software_latch":            0, //idk if these are correct
		},
	}
	test1.Run(mqttClient, db)
}

func SendACUStatusThree(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x009,
		Name: "ACU Status Three Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"acu_hv_input_voltage":  0,
			"acu_hv_output_voltage": 0,
			"acu_hv_input_current":  0,
			"acu_hv_output_current": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendACUPrecharge(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x00A,
		Name: "ACU Precharge Test 1",
		Data: []byte{0x00},
		ExpectedValues: map[string]interface{}{
			"acu_set_ts_active": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendACUConfigChargeParameters(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x00B,
		Name: "ACU Config Charge Parameters Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"charge_voltage": 0,
			"charge_current": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendACUConfigOperationalParameters(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x00C,
		Name: "ACU Config Operational Parameters Test 1",
		Data: []byte{0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"minimum_cell_voltage": 0,
			"max_cell_temperature": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendACUCellDataOne(mqttClient *mqtt.Client, db *gorm.DB) {
	data := make([]byte, 64)
	test1 := MessageTest{
		ID:   0x00D,
		Name: "ACU Cell Data 1 Test 1",
		Data: data, //might break
		ExpectedValues: map[string]interface{}{
			"acu_cell_0_voltage":  0,
			"acu_cell_0_temp":     0,
			"acu_cell_1_voltage":  0,
			"acu_cell_1_temp":     0,
			"acu_cell_2_voltage":  0,
			"acu_cell_2_temp":     0,
			"acu_cell_3_voltage":  0,
			"acu_cell_3_temp":     0,
			"acu_cell_4_voltage":  0,
			"acu_cell_4_temp":     0,
			"acu_cell_5_voltage":  0,
			"acu_cell_5_temp":     0,
			"acu_cell_6_voltage":  0,
			"acu_cell_6_temp":     0,
			"acu_cell_7_voltage":  0,
			"acu_cell_7_temp":     0,
			"acu_cell_8_voltage":  0,
			"acu_cell_8_temp":     0,
			"acu_cell_9_voltage":  0,
			"acu_cell_9_temp":     0,
			"acu_cell_10_voltage": 0,
			"acu_cell_10_temp":    0,
			"acu_cell_11_voltage": 0,
			"acu_cell_11_temp":    0,
			"acu_cell_12_voltage": 0,
			"acu_cell_12_temp":    0,
			"acu_cell_13_voltage": 0,
			"acu_cell_13_temp":    0,
			"acu_cell_14_voltage": 0,
			"acu_cell_14_temp":    0,
			"acu_cell_15_voltage": 0,
			"acu_cell_15_temp":    0,
			"acu_cell_16_voltage": 0,
			"acu_cell_16_temp":    0,
			"acu_cell_17_voltage": 0,
			"acu_cell_17_temp":    0,
			"acu_cell_18_voltage": 0,
			"acu_cell_18_temp":    0,
			"acu_cell_19_voltage": 0,
			"acu_cell_19_temp":    0,
			"acu_cell_20_voltage": 0,
			"acu_cell_20_temp":    0,
			"acu_cell_21_voltage": 0,
			"acu_cell_21_temp":    0,
			"acu_cell_22_voltage": 0,
			"acu_cell_22_temp":    0,
			"acu_cell_23_voltage": 0,
			"acu_cell_23_temp":    0,
			"acu_cell_24_voltage": 0,
			"acu_cell_24_temp":    0,
			"acu_cell_25_voltage": 0,
			"acu_cell_25_temp":    0,
			"acu_cell_26_voltage": 0,
			"acu_cell_26_temp":    0,
			"acu_cell_27_voltage": 0,
			"acu_cell_27_temp":    0,
			"acu_cell_28_voltage": 0,
			"acu_cell_28_temp":    0,
			"acu_cell_29_voltage": 0,
			"acu_cell_29_temp":    0,
			"acu_cell_30_voltage": 0,
			"acu_cell_30_temp":    0,
			"acu_cell_31_voltage": 0,
			"acu_cell_31_temp":    0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendACUCellDataTwo(mqttClient *mqtt.Client, db *gorm.DB) {
	data := make([]byte, 64)
	test1 := MessageTest{
		ID:   0x00E,
		Name: "ACU Cell Data 2 Test 1",
		Data: data, //might break
		ExpectedValues: map[string]interface{}{
			"acu_cell_32_voltage": 0,
			"acu_cell_32_temp":    0,
			"acu_cell_33_voltage": 0,
			"acu_cell_33_temp":    0,
			"acu_cell_34_voltage": 0,
			"acu_cell_34_temp":    0,
			"acu_cell_35_voltage": 0,
			"acu_cell_35_temp":    0,
			"acu_cell_36_voltage": 0,
			"acu_cell_36_temp":    0,
			"acu_cell_37_voltage": 0,
			"acu_cell_37_temp":    0,
			"acu_cell_38_voltage": 0,
			"acu_cell_38_temp":    0,
			"acu_cell_39_voltage": 0,
			"acu_cell_39_temp":    0,
			"acu_cell_40_voltage": 0,
			"acu_cell_40_temp":    0,
			"acu_cell_41_voltage": 0,
			"acu_cell_41_temp":    0,
			"acu_cell_42_voltage": 0,
			"acu_cell_42_temp":    0,
			"acu_cell_43_voltage": 0,
			"acu_cell_43_temp":    0,
			"acu_cell_44_voltage": 0,
			"acu_cell_44_temp":    0,
			"acu_cell_45_voltage": 0,
			"acu_cell_45_temp":    0,
			"acu_cell_46_voltage": 0,
			"acu_cell_46_temp":    0,
			"acu_cell_47_voltage": 0,
			"acu_cell_47_temp":    0,
			"acu_cell_48_voltage": 0,
			"acu_cell_48_temp":    0,
			"acu_cell_49_voltage": 0,
			"acu_cell_49_temp":    0,
			"acu_cell_50_voltage": 0,
			"acu_cell_50_temp":    0,
			"acu_cell_51_voltage": 0,
			"acu_cell_51_temp":    0,
			"acu_cell_52_voltage": 0,
			"acu_cell_52_temp":    0,
			"acu_cell_53_voltage": 0,
			"acu_cell_53_temp":    0,
			"acu_cell_54_voltage": 0,
			"acu_cell_54_temp":    0,
			"acu_cell_55_voltage": 0,
			"acu_cell_55_temp":    0,
			"acu_cell_56_voltage": 0,
			"acu_cell_56_temp":    0,
			"acu_cell_57_voltage": 0,
			"acu_cell_57_temp":    0,
			"acu_cell_58_voltage": 0,
			"acu_cell_58_temp":    0,
			"acu_cell_59_voltage": 0,
			"acu_cell_59_temp":    0,
			"acu_cell_60_voltage": 0,
			"acu_cell_60_temp":    0,
			"acu_cell_61_voltage": 0,
			"acu_cell_61_temp":    0,
			"acu_cell_62_voltage": 0,
			"acu_cell_62_temp":    0,
			"acu_cell_63_voltage": 0,
			"acu_cell_63_temp":    0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendACUCellDataThree(mqttClient *mqtt.Client, db *gorm.DB) {
	data := make([]byte, 64)
	test1 := MessageTest{
		ID:   0x00F,
		Name: "ACU Cell Data 3 Test 1",
		Data: data, //might break
		ExpectedValues: map[string]interface{}{
			"acu_cell_64_voltage": 0,
			"acu_cell_64_temp":    0,
			"acu_cell_65_voltage": 0,
			"acu_cell_65_temp":    0,
			"acu_cell_66_voltage": 0,
			"acu_cell_66_temp":    0,
			"acu_cell_67_voltage": 0,
			"acu_cell_67_temp":    0,
			"acu_cell_68_voltage": 0,
			"acu_cell_68_temp":    0,
			"acu_cell_69_voltage": 0,
			"acu_cell_69_temp":    0,
			"acu_cell_70_voltage": 0,
			"acu_cell_70_temp":    0,
			"acu_cell_71_voltage": 0,
			"acu_cell_71_temp":    0,
			"acu_cell_72_voltage": 0,
			"acu_cell_72_temp":    0,
			"acu_cell_73_voltage": 0,
			"acu_cell_73_temp":    0,
			"acu_cell_74_voltage": 0,
			"acu_cell_74_temp":    0,
			"acu_cell_75_voltage": 0,
			"acu_cell_75_temp":    0,
			"acu_cell_76_voltage": 0,
			"acu_cell_76_temp":    0,
			"acu_cell_77_voltage": 0,
			"acu_cell_77_temp":    0,
			"acu_cell_78_voltage": 0,
			"acu_cell_78_temp":    0,
			"acu_cell_79_voltage": 0,
			"acu_cell_79_temp":    0,
			"acu_cell_80_voltage": 0,
			"acu_cell_80_temp":    0,
			"acu_cell_81_voltage": 0,
			"acu_cell_81_temp":    0,
			"acu_cell_82_voltage": 0,
			"acu_cell_82_temp":    0,
			"acu_cell_83_voltage": 0,
			"acu_cell_83_temp":    0,
			"acu_cell_84_voltage": 0,
			"acu_cell_84_temp":    0,
			"acu_cell_85_voltage": 0,
			"acu_cell_85_temp":    0,
			"acu_cell_86_voltage": 0,
			"acu_cell_86_temp":    0,
			"acu_cell_87_voltage": 0,
			"acu_cell_87_temp":    0,
			"acu_cell_88_voltage": 0,
			"acu_cell_88_temp":    0,
			"acu_cell_89_voltage": 0,
			"acu_cell_89_temp":    0,
			"acu_cell_90_voltage": 0,
			"acu_cell_90_temp":    0,
			"acu_cell_91_voltage": 0,
			"acu_cell_91_temp":    0,
			"acu_cell_92_voltage": 0,
			"acu_cell_92_temp":    0,
			"acu_cell_93_voltage": 0,
			"acu_cell_93_temp":    0,
			"acu_cell_94_voltage": 0,
			"acu_cell_94_temp":    0,
			"acu_cell_95_voltage": 0,
			"acu_cell_95_temp":    0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendACUCellDataFour(mqttClient *mqtt.Client, db *gorm.DB) {
	data := make([]byte, 64)
	test1 := MessageTest{
		ID:   0x010,
		Name: "ACU Cell Data 4 Test 1",
		Data: data, //might break
		ExpectedValues: map[string]interface{}{
			"acu_cell_96_voltage":  0,
			"acu_cell_96_temp":     0,
			"acu_cell_97_voltage":  0,
			"acu_cell_97_temp":     0,
			"acu_cell_98_voltage":  0,
			"acu_cell_98_temp":     0,
			"acu_cell_99_voltage":  0,
			"acu_cell_99_temp":     0,
			"acu_cell_100_voltage": 0,
			"acu_cell_100_temp":    0,
			"acu_cell_101_voltage": 0,
			"acu_cell_101_temp":    0,
			"acu_cell_102_voltage": 0,
			"acu_cell_102_temp":    0,
			"acu_cell_103_voltage": 0,
			"acu_cell_103_temp":    0,
			"acu_cell_104_voltage": 0,
			"acu_cell_104_temp":    0,
			"acu_cell_105_voltage": 0,
			"acu_cell_105_temp":    0,
			"acu_cell_106_voltage": 0,
			"acu_cell_106_temp":    0,
			"acu_cell_107_voltage": 0,
			"acu_cell_107_temp":    0,
			"acu_cell_108_voltage": 0,
			"acu_cell_108_temp":    0,
			"acu_cell_109_voltage": 0,
			"acu_cell_109_temp":    0,
			"acu_cell_110_voltage": 0,
			"acu_cell_110_temp":    0,
			"acu_cell_111_voltage": 0,
			"acu_cell_111_temp":    0,
			"acu_cell_112_voltage": 0,
			"acu_cell_112_temp":    0,
			"acu_cell_113_voltage": 0,
			"acu_cell_113_temp":    0,
			"acu_cell_114_voltage": 0,
			"acu_cell_114_temp":    0,
			"acu_cell_115_voltage": 0,
			"acu_cell_115_temp":    0,
			"acu_cell_116_voltage": 0,
			"acu_cell_116_temp":    0,
			"acu_cell_117_voltage": 0,
			"acu_cell_117_temp":    0,
			"acu_cell_118_voltage": 0,
			"acu_cell_118_temp":    0,
			"acu_cell_119_voltage": 0,
			"acu_cell_119_temp":    0,
			"acu_cell_120_voltage": 0,
			"acu_cell_120_temp":    0,
			"acu_cell_121_voltage": 0,
			"acu_cell_121_temp":    0,
			"acu_cell_122_voltage": 0,
			"acu_cell_122_temp":    0,
			"acu_cell_123_voltage": 0,
			"acu_cell_123_temp":    0,
			"acu_cell_124_voltage": 0,
			"acu_cell_124_temp":    0,
			"acu_cell_125_voltage": 0,
			"acu_cell_125_temp":    0,
			"acu_cell_126_voltage": 0,
			"acu_cell_126_temp":    0,
			"acu_cell_127_voltage": 0,
			"acu_cell_127_temp":    0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendACUCellDataFive(mqttClient *mqtt.Client, db *gorm.DB) {
	data := make([]byte, 64)
	test1 := MessageTest{
		ID:   0x011,
		Name: "ACU Cell Data 5 Test 1",
		Data: data, //might break
		ExpectedValues: map[string]interface{}{
			"acu_cell_128_voltage": 0,
			"acu_cell_128_temp":    0,
			"acu_cell_129_voltage": 0,
			"acu_cell_129_temp":    0,
			"acu_cell_130_voltage": 0,
			"acu_cell_130_temp":    0,
			"acu_cell_131_voltage": 0,
			"acu_cell_131_temp":    0,
			"acu_cell_132_voltage": 0,
			"acu_cell_132_temp":    0,
			"acu_cell_133_voltage": 0,
			"acu_cell_133_temp":    0,
			"acu_cell_134_voltage": 0,
			"acu_cell_134_temp":    0,
			"acu_cell_135_voltage": 0,
			"acu_cell_135_temp":    0,
			"acu_cell_136_voltage": 0,
			"acu_cell_136_temp":    0,
			"acu_cell_137_voltage": 0,
			"acu_cell_137_temp":    0,
			"acu_cell_138_voltage": 0,
			"acu_cell_138_temp":    0,
			"acu_cell_139_voltage": 0,
			"acu_cell_139_temp":    0,
			"acu_cell_140_voltage": 0,
			"acu_cell_140_temp":    0,
			"acu_cell_141_voltage": 0, //this is empty
			"acu_cell_141_temp":    0, //this is empty
			"acu_cell_142_voltage": 0, //this is empty
			"acu_cell_142_temp":    0, //this is empty
			"acu_cell_143_voltage": 0, //this is empty
			"acu_cell_143_temp":    0, //this is empty
			"acu_cell_144_voltage": 0, //this is empty
			"acu_cell_144_temp":    0, //this is empty
			"acu_cell_145_voltage": 0, //this is empty
			"acu_cell_145_temp":    0, //this is empty
			"acu_cell_146_voltage": 0, //this is empty
			"acu_cell_146_temp":    0, //this is empty
			"acu_cell_147_voltage": 0, //this is empty
			"acu_cell_147_temp":    0, //this is empty
			"acu_cell_148_voltage": 0, //this is empty
			"acu_cell_148_temp":    0, //this is empty
			"acu_cell_149_voltage": 0, //this is empty
			"acu_cell_149_temp":    0, //this is empty
			"acu_cell_150_voltage": 0, //this is empty
			"acu_cell_150_temp":    0, //this is empty
			"acu_cell_151_voltage": 0, //this is empty
			"acu_cell_151_temp":    0, //this is empty
			"acu_cell_152_voltage": 0, //this is empty
			"acu_cell_152_temp":    0, //this is empty
			"acu_cell_153_voltage": 0, //this is empty
			"acu_cell_153_temp":    0, //this is empty
			"acu_cell_154_voltage": 0, //this is empty
			"acu_cell_154_temp":    0, //this is empty
			"acu_cell_155_voltage": 0, //this is empty
			"acu_cell_155_temp":    0, //this is empty
			"acu_cell_156_voltage": 0, //this is empty
			"acu_cell_156_temp":    0, //this is empty
			"acu_cell_157_voltage": 0, //this is empty
			"acu_cell_157_temp":    0, //this is empty
			"acu_cell_158_voltage": 0, //this is empty
			"acu_cell_158_temp":    0, //this is empty
			"acu_cell_159_voltage": 0, //this is empty
			"acu_cell_159_temp":    0, //this is empty
		},
	}
	test1.Run(mqttClient, db)
}

func SendDCDCStatus(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x012,
		Name: "DC DC Status Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"acu_input_voltage":  0,
			"acu_output_voltage": 0,
			"acu_input_current":  0,
			"acu_output_current": 0,
			"acu_dc_dc_temp":     0,
		},
	}
	test1.Run(mqttClient, db)
}
