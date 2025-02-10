package gr25

import "github.com/google/uuid"

func SendECUStatusOne(port int) {
	data := []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
	id := uuid.New().String()
}
