package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
)

type RoomController struct {
	roomService app.RoomService
}

func NewRoomController(os app.RoomService) RoomController {
	return RoomController{
		roomService: os,
	}
}

func (c RoomController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		rom, err := requests.Bind(r, requests.RoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController: %s", err)
			BadRequest(w, err)
			return
		}

		rom, err = c.roomService.Save(rom, user.Id)
		if err != nil {
			log.Printf("RoomController: %s", err)
			if err.Error() == "access denied" {
				Forbidden(w, err)
			} else {
				InternalServerError(w, err)
			}
			return
		}

		var romDto resources.RomDto
		Created(w, romDto.DomainToDto(rom))
	}
}

func (c RoomController) FindForOrganization() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		org := r.Context().Value(OrgKey).(domain.Organization)
		roms, err := c.roomService.FindForOrganization(org.Id)
		if err != nil {
			log.Printf("RoomController: %s", err)
			InternalServerError(w, err)
			return
		}

		var romsDto resources.RomsDto
		response := romsDto.DomainToDto(roms)
		Success(w, response)
	}
}

func (c RoomController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		org := r.Context().Value(OrgKey).(domain.Organization)
		rom := r.Context().Value(RoomKey).(domain.Room)

		if rom.OrganizationId != org.Id {
			err := fmt.Errorf("access denied")
			Forbidden(w, err)
			return
		}

		var romDto resources.RomDto
		Success(w, romDto.DomainToDto(rom))
	}
}

func (c RoomController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		org := r.Context().Value(OrgKey).(domain.Organization)
		rom, err := requests.Bind(r, requests.RoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController: %s", err)
			BadRequest(w, err)
			return
		}

		room := r.Context().Value(RoomKey).(domain.Room)
		if room.OrganizationId != org.Id {
			err := fmt.Errorf("access denied")
			Forbidden(w, err)
			return
		}

		room.OrganizationId = rom.OrganizationId
		room.Name = rom.Name
		room.Description = rom.Description
		room, err = c.roomService.Update(room)
		if err != nil {
			log.Printf("RoomController: %s", err)
			InternalServerError(w, err)
			return
		}

		var romDto resources.RomDto
		Success(w, romDto.DomainToDto(room))
	}
}

func (c RoomController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		org := r.Context().Value(OrgKey).(domain.Organization)
		rom := r.Context().Value(RoomKey).(domain.Room)

		if rom.OrganizationId != org.Id {
			err := fmt.Errorf("access denied")
			Forbidden(w, err)
			return
		}

		err := c.roomService.Delete(rom.Id)
		if err != nil {
			log.Printf("RoomController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
