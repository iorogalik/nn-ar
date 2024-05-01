package resources

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type OrgDto struct {
	Id          uint64    `json:"id"`
	UserId      uint64    `json:"userId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	Lat         float64   `json:"lat"`
	Lon         float64   `json:"lon"`
	CreatedDate time.Time `json:"createdDate"`
	UpdatedDate time.Time `json:"updatedDate"`
}

func (d OrgDto) DomainToDto(o domain.Organization) OrgDto {
	return OrgDto{
		Id:          o.Id,
		UserId:      o.UserId,
		Name:        o.Name,
		Description: o.Description,
		City:        o.City,
		Address:     o.Address,
		Lat:         o.Lat,
		Lon:         o.Lon,
		CreatedDate: o.CreatedDate,
		UpdatedDate: o.UpdatedDate,
	}
}
