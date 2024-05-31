package app

import (
	"errors"
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type DeviceService interface {
	Save(dv domain.Device, uId uint64) (domain.Device, error)
	FindForRoom(mId uint64) ([]domain.Device, error)
	Find(id uint64) (interface{}, error)
	Update(dv domain.Device) (domain.Device, error)
	SetDeviceToRoom(deviceId, roomId uint64) error
	RemoveDeviceFromRoom(deviceId uint64) error
	Delete(id uint64) error
}

type deviceService struct {
	deviceRepo database.DeviceRepository
	roomRepo   database.RoomRepository
	orgRepo    database.OrganizationRepository
}

func NewDeviceService(de database.DeviceRepository, ro database.RoomRepository, or database.OrganizationRepository) DeviceService {
	return &deviceService{
		deviceRepo: de,
		roomRepo:   ro,
		orgRepo:    or,
	}
}

func (s deviceService) Save(dv domain.Device, uId uint64) (domain.Device, error) {
	var (
		rom domain.Room
		org domain.Organization
		err error
	)
	if dv.RoomId != nil {
		rom, err = s.roomRepo.FindById(*dv.RoomId)
		if err != nil {
			log.Printf("DeviceService: %s", err)
			return domain.Device{}, err
		}
	}

	org, err = s.orgRepo.FindById(rom.OrganizationId)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return domain.Device{}, err
	}

	if org.UserId != uId {
		err = errors.New("access denied")
		log.Panicf("DeviceService: %s", err)
		return domain.Device{}, err
	}

	dv, err = s.deviceRepo.Save(dv)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return domain.Device{}, err
	}

	return dv, nil
}

func (s deviceService) FindForRoom(mId uint64) ([]domain.Device, error) {
	devices, err := s.deviceRepo.FindForRoom(mId)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return nil, err
	}

	return devices, nil
}

func (s deviceService) Find(id uint64) (interface{}, error) {
	device, err := s.deviceRepo.FindById(id)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return nil, err
	}

	return device, nil
}

func (s deviceService) Update(dv domain.Device) (domain.Device, error) {
	device, err := s.deviceRepo.Update(dv)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return domain.Device{}, err
	}

	return device, nil
}

func (s deviceService) SetDeviceToRoom(deviceId, roomId uint64) error {
	err := s.deviceRepo.SetDeviceToRoom(deviceId, roomId)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return err
	}

	return nil
}

func (s deviceService) RemoveDeviceFromRoom(deviceId uint64) error {
	err := s.deviceRepo.RemoveDeviceFromRoom(deviceId)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return err
	}

	return nil
}

func (s deviceService) Delete(id uint64) error {
	err := s.deviceRepo.Delete(id)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return err
	}

	return nil
}
