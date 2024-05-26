package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type RomsDto struct {
	Rooms []RomDto `json:"rooms"`
}

type RomDto struct {
	Id             uint64    `json:"id"`
	OrganizationId uint64    `json:"organizationId"`
	Name           string    `json:"name"`
	Description    string    `json:"description,somitempty"`
	CreatedDate    time.Time `json:"createdDate"`
	UpdatedDate    time.Time `json:"updatedDate"`
}

func (d RomDto) DomainToDto(m domain.Room) RomDto {
	return RomDto{
		Id:             m.Id,
		OrganizationId: m.OrganizationId,
		Name:           m.Name,
		Description:    m.Description,
		CreatedDate:    m.CreatedDate,
		UpdatedDate:    m.UpdatedDate,
	}
}

func (d RomsDto) DomainToDto(roms []domain.Room) RomsDto {
	var rooms []RomDto
	for _, m := range roms {
		var mDto RomDto
		rom := mDto.DomainToDto(m)
		rooms = append(rooms, rom)
	}
	response := RomsDto{
		Rooms: rooms,
	}
	return response
}
