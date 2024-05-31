package requests

import (
	"errors"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/google/uuid"
)

type DeviceRequest struct {
	RoomId           *uint64   `json:"roomId,omitempty"`
	GUID             uuid.UUID `json:"guid"`
	InventoryNumber  string    `json:"inventoryNumber"`
	SerialNumber     string    `json:"serialNumber"`
	Characteristics  string    `json:"characteristics"`
	Category         string    `json:"category"`
	Units            *string   `json:"units,omitempty"`
	PowerConsumption *float64  `json:"powerConsumption,omitempty"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
	if r.Category != "SENSOR" && r.Category != "ACTUATOR" {
		return domain.Device{}, errors.New("invalid category")
	}
	if r.Category == "SENSOR" && r.Units == nil {
		return domain.Device{}, errors.New("units is required for SENSOR")
	}

	if r.Category == "ACTUATOR" && r.PowerConsumption == nil {
		return domain.Device{}, errors.New("power consumption is required for ACTUATOR")
	}

	return domain.Device{
		RoomId:           r.RoomId,
		GUID:             r.GUID,
		InventoryNumber:  r.InventoryNumber,
		SerialNumber:     r.SerialNumber,
		Characteristics:  r.Characteristics,
		Category:         r.Category,
		Units:            r.Units,
		PowerConsumption: r.PowerConsumption,
	}, nil
}

type SetRoomRequest struct {
	RoomId uint64 `json:"room_id"`
}

func (r SetRoomRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		RoomId: &r.RoomId,
	}, nil
}
