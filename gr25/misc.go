package gr25

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

// contains only the functions (they do not yet have a home)

func SendDashStatus(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x01A,
		Name: "Dash Status Test 1",
		Data: []byte{0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"bms_led_status":  0,
			"imd_led_status":  0,
			"bspd_led_status": 0,
			"ts_button_data":  0,
			"rtd_button_data": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendDashConfig(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x01B,
		Name: "Dash Config Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"bms_led_config":   0,
			"imd_led_config":   0,
			"bspd_led_config":  0,
			"button_led_one_r": 0,
			"button_led_one_g": 0,
			"button_led_one_b": 0,
			"button_led_two_r": 0,
			"button_led_two_g": 0,
			"button_led_two_b": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendSteeringStatus(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x01C,
		Name: "Steering Status Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"current_encoder":    0, // why is it written like this
			"torque_map_encoder": 0,
			"regen":              0,
			"button_one":         0,
			"button_two":         0,
			"button_three":       0,
			"button_four":        0,
		},
	}
	test1.Run(mqttClient, db)
}

// reserved
func SendSteeringConfig(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:             0x01D,
		Name:           "Steering Config Test 1",
		Data:           []byte{},
		ExpectedValues: map[string]interface{}{},
	}
	test1.Run(mqttClient, db)
}
