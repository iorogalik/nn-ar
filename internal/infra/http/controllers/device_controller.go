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

type DeviceController struct {
	deviceService app.DeviceService
}

func NewDeviceController(os app.DeviceService) DeviceController {
	return DeviceController{
		deviceService: os,
	}
}

func (c DeviceController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dev, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		dev, err = c.deviceService.Save(dev)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			if err.Error() == "access denied" {
				Forbidden(w, err)
			} else {
				InternalServerError(w, err)
			}
			return
		}

		var devDto resources.DevDto
		Created(w, devDto.DomainToDto(dev))
	}
}

func (c DeviceController) FindForRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dev := r.Context().Value(DeviceKey).(domain.Device)
		devs, err := c.deviceService.FindForRoom(dev.Id)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			InternalServerError(w, err)
			return
		}

		var devsDto resources.DevsDto
		response := devsDto.DomainToDto(devs)
		Success(w, response)
	}
}

func (c DeviceController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dev := r.Context().Value(DeviceKey).(domain.Device)
		rom := r.Context().Value(RoomKey).(domain.Room)

		if dev.RoomId != rom.Id {
			err := fmt.Errorf("access denied")
			Forbidden(w, err)
			return
		}

		var devDto resources.DevDto
		Success(w, devDto.DomainToDto(dev))
	}
}

func (c DeviceController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rom := r.Context().Value(RoomKey).(domain.Room)
		dev, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		device := r.Context().Value(DeviceKey).(domain.Device)
		if dev.RoomId != rom.Id {
			err := fmt.Errorf("access denied")
			Forbidden(w, err)
			return
		}

		device.GUID = dev.GUID
		device.InventoryNumber = dev.InventoryNumber
		device.SerialNumber = dev.SerialNumber
		device.Characteristics = dev.Characteristics
		device.Category = dev.Category
		device.Units = dev.Units
		device.PowerConsumption = dev.PowerConsumption
		device, err = c.deviceService.Update(device)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			InternalServerError(w, err)
			return
		}

		var devDto resources.DevDto
		Success(w, devDto.DomainToDto(device))
	}
}

func (c DeviceController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dev := r.Context().Value(DeviceKey).(domain.Device)
		rom := r.Context().Value(RoomKey).(domain.Room)

		if dev.RoomId != rom.Id {
			err := fmt.Errorf("access denied")
			Forbidden(w, err)
			return
		}

		err := c.deviceService.Delete(dev.Id)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
