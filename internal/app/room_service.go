package app

import (
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type RoomService interface {
	Save(o domain.Room) (domain.Room, error)
	FindForOrganization(oId uint64) ([]domain.Room, error)
	Find(id uint64) (interface{}, error)
	Update(m domain.Room) (domain.Room, error)
	Delete(id uint64) error
}

type roomService struct {
	roomRepo database.RoomRepository
}

func NewRoomService(ro database.RoomRepository) RoomService {
	return &roomService{
		roomRepo: ro,
	}
}

func (s roomService) Save(o domain.Room) (domain.Room, error) {
	m, err := s.roomRepo.Save(o)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return domain.Room{}, err
	}

	return m, nil
}

func (s roomService) FindForOrganization(oId uint64) ([]domain.Room, error) {
	roms, err := s.roomRepo.FindForOrganization(oId)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return nil, err
	}

	return roms, nil
}

func (s roomService) Find(id uint64) (interface{}, error) {
	rom, err := s.roomRepo.FindById(id)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return nil, err
	}

	return rom, nil
}

func (s roomService) Update(m domain.Room) (domain.Room, error) {
	rom, err := s.roomRepo.Update(m)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return domain.Room{}, err
	}

	return rom, nil
}

func (s roomService) Delete(id uint64) error {
	err := s.roomRepo.Delete(id)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return err
	}

	return nil
}
