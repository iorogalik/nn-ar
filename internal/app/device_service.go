package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type DeviceService interface {
	Save(dv domain.Device) (domain.Device, error)
	FindForRoom(mId uint64) ([]domain.Device, error)
	Find(id uint64) (interface{}, error)
	Update(d domain.Device) (domain.Device, error)
	Delete(id uint64) error
}

type deviceService struct {
	deviceRepo database.DeviceRepository
}

func NewDeviceService(de database.DeviceRepository) DeviceService {
	return &deviceService{
		deviceRepo: de,
	}
}

func (s deviceService) Save(dv domain.Device) (domain.Device, error) {
	dv, err := s.deviceRepo.Save(dv)
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

func (s deviceService) Delete(id uint64) error {
	err := s.deviceRepo.Delete(id)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return err
	}

	return nil
}
