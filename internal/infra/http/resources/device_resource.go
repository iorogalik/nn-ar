package resources

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type DevsDto struct {
	Devices []DevDto `json:"devices"`
}

type DevDto struct {
	Id               uint64  `json:"id"`
	OrganizationId   uint64  `json:"organizationId"`
	RoomId           uint64  `json:"roomId"`
	GUID             string  `json:"guid"`
	InventoryNumber  string  `json:"inventoryNumber"`
	SerialNumber     string  `json:"serialNumber"`
	Characteristics  string  `json:"characteristics"`
	Category         string  `json:"category"`
	Units            string  `json:"units"`
	PowerConsumption float64 `json:"powerconsumption"`
}

func (d DevDto) DomainToDto(dv domain.Device) DevDto {
	return DevDto{
		Id:               dv.Id,
		OrganizationId:   dv.OrganizationId,
		RoomId:           dv.RoomId,
		GUID:             dv.GUID,
		InventoryNumber:  dv.InventoryNumber,
		SerialNumber:     dv.SerialNumber,
		Characteristics:  dv.Characteristics,
		Category:         dv.Category,
		Units:            dv.Units,
		PowerConsumption: dv.PowerConsumption,
	}
}

func (d DevsDto) DomainToDto(devs []domain.Device) DevsDto {
	var devices []DevDto
	for _, dv := range devs {
		var dvDto DevDto
		dev := dvDto.DomainToDto(dv)
		devices = append(devices, dev)
	}
	response := DevsDto{
		Devices: devices,
	}
	return response
}
