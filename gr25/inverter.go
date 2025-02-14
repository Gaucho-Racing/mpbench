package gr25

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

func RunInverterTests(mqttClient *mqtt.Client, db *gorm.DB) {
	SendInverterStatusOne(mqttClient, db)
	SendInverterStatusTwo(mqttClient, db)
	SendInverterStatusThree(mqttClient, db)
}

func SendInverterStatusOne(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x013,
		Name: "Inverter Status 1 Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"inverter_ac_current": 0,
			"inverter_dc_current": 0,
			"inverter_motor_rpm":  0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendInverterStatusTwo(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x014,
		Name: "Inverter Status 2 Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"u_mosfet_temperature": 0,
			"v_mosfet_temperature": 0,
			"w_mosfet_temperature": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendInverterStatusThree(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x015,
		Name: "Inverter Status 3 Test 1",
		Data: []byte{0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"inverter_motor_temperature":    0,
			"inverter_over_voltage_faults":  0,
			"inverter_under_voltage_faults": 0,
			"inverter_overtemp_fault":       0, //different from spreadsheet
			"inverter_motor_overtemp_fault": 0,
			"inverter_transitor_fault":      0,
			"inverter_encoder_fault":        0,
			"inverter_can_fault":            0,
			"inverter_future_use":           0, //wtf
		},
	}
	test1.Run(mqttClient, db)
}
