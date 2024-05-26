package app

import (
	"errors"
	"log"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
)

type RoomService interface {
	Save(m domain.Room, uId uint64) (domain.Room, error)
	FindForOrganization(oId uint64) ([]domain.Room, error)
	Find(id uint64) (interface{}, error)
	Update(m domain.Room) (domain.Room, error)
	Delete(id uint64) error
}

type roomService struct {
	roomRepo               database.RoomRepository
	organizationRepository database.OrganizationRepository
}

func NewRoomService(ro database.RoomRepository, or database.OrganizationRepository) RoomService {
	return &roomService{
		roomRepo:               ro,
		organizationRepository: or,
	}
}

func (s roomService) Save(m domain.Room, uId uint64) (domain.Room, error) {
	org, err := s.organizationRepository.FindById(m.OrganizationId)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return domain.Room{}, err
	}

	if org.UserId != uId {
		err = errors.New("access denied")
		log.Panicf("RoomService: %s", err)
		return domain.Room{}, err
	}

	m, err = s.roomRepo.Save(m)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return domain.Room{}, err
	}

	return m, nil
}

func (s roomService) FindForOrganization(oId uint64) ([]domain.Room, error) {
	rooms, err := s.roomRepo.FindForOrganization(oId)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return nil, err
	}

	return rooms, nil
}

func (s roomService) Find(id uint64) (interface{}, error) {
	room, err := s.roomRepo.FindById(id)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return nil, err
	}

	return room, nil
}

func (s roomService) Update(m domain.Room) (domain.Room, error) {
	room, err := s.roomRepo.Update(m)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return domain.Room{}, err
	}

	return room, nil
}

func (s roomService) Delete(id uint64) error {
	err := s.roomRepo.Delete(id)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return err
	}

	return nil
}
