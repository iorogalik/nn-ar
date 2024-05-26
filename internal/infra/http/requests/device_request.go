package requests

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type DeviceRequest struct {
	RoomId           uint64  `json:"roomId"`
	GUID             string  `json:"guid"`
	InventoryNumber  string  `json:"inventoryNumber"`
	SerialNumber     string  `json:"serialNumber"`
	Characteristics  string  `json:"characteristics"`
	Category         string  `json:"category"`
	Units            string  `json:"units"`
	PowerConsumption float64 `json:"powerConsumption"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
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
