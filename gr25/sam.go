package gr25

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

func RunSAMTests(mqttClient *mqtt.Client, db *gorm.DB) {
	SendSAMBrakeIR(mqttClient, db)
	SendSAMTireTemp(mqttClient, db)
	SendSAMIMU(mqttClient, db)
	SendSAMGPSOne(mqttClient, db)
	SendSAMGPSTwo(mqttClient, db)
	SendSAMGPSTime(mqttClient, db)
	SendSAMGPSHeading(mqttClient, db)
	SendSAMSusPots(mqttClient, db)
	SendSAMTOF(mqttClient, db)
	SendSAMRearWheelspeed(mqttClient, db)
	SendSAMPushrodForce(mqttClient, db)
}

func SendSAMBrakeIR(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x01E,
		Name: "Sam Brake IR Test 1",
		Data: []byte{0x00},
		ExpectedValues: map[string]interface{}{
			"sam_brake_temp": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendSAMTireTemp(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x01F,
		Name: "SAM Tire Temp Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"sam_tire_outside_temp":        0,
			"sam_tire_outside_middle_temp": 0,
			"sam_tire_inside_middle_temp":  0,
			"sam_tire_inside_temp":         0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendSAMIMU(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x020,
		Name: "SAM IMU Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"sam_accel_x": 0,
			"sam_accel_y": 0,
			"sam_accel_z": 0,
			"sam_gyro_x":  0,
			"sam_gyro_y":  0,
			"sam_gyro_z":  0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendSAMGPSOne(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x021,
		Name: "SAM GPS 1 Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"sam_gps_latitude":  0,
			"sam_gps_longitude": 0,
		},
	}
	test1.Run(mqttClient, db)
}
func SendSAMGPSTwo(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x022,
		Name: "SAM GPS 2 Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"sam_gps_accuracy": 0,
			"sam_gps_attitude": 0,
		},
	}
	test1.Run(mqttClient, db)
}
func SendSAMGPSTime(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x023,
		Name: "SAM GPS Time Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"sam_gps_time":         0,
			"sam_gps_time_of_week": 0,
		},
	}
	test1.Run(mqttClient, db)
}
func SendSAMGPSHeading(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x024,
		Name: "SAM GPS Heading Test 1",
		Data: []byte{0x00, 0x00, 0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"sam_gps_heading": 0,
		},
	}
	test1.Run(mqttClient, db)
}
func SendSAMSusPots(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x025,
		Name: "SAM Sus Pots Test 1",
		Data: []byte{0x00},
		ExpectedValues: map[string]interface{}{
			"sam_suspension_angle": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendSAMTOF(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x026,
		Name: "SAM TOF Test 1",
		Data: []byte{0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"sam_height": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendSAMRearWheelspeed(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x027,
		Name: "SAM Rear Wheelspeed Test 1",
		Data: []byte{0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"sam_speed": 0,
		},
	}
	test1.Run(mqttClient, db)
}

func SendSAMPushrodForce(mqttClient *mqtt.Client, db *gorm.DB) {
	test1 := MessageTest{
		ID:   0x028, //differs from sheet (26 seems to be a typo)
		Name: "SAM Pushrod Force Test 1",
		Data: []byte{0x00, 0x00},
		ExpectedValues: map[string]interface{}{
			"sam_load_force": 0,
		},
	}
	test1.Run(mqttClient, db)
}
