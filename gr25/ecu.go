package gr25

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

func RunECUTests(mqttClient *mqtt.Client, db *gorm.DB) {
	SendECUStatusOne(mqttClient, db)
}

func SendECUStatusOne(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
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
	test1.Run(mqttClient, db)
}
