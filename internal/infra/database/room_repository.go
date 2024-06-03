package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const RoomsTableName = "rooms"

type room struct {
	Id             uint64     `db:"id,omitempty"`
	Name           string     `db:"name"`
	OrganizationId uint64     `db:"organization_id"`
	Description    string     `db:"description"`
	CreatedDate    time.Time  `db:"created_date"`
	UpdatedDate    time.Time  `db:"updated_date"`
	DeletedDate    *time.Time `db:"deleted_date"`
}

type RoomRepository interface {
	Save(r domain.Room) (domain.Room, error)
	FindForOrganization(oId uint64) ([]domain.Room, error)
	FindById(id uint64) (domain.Room, error)
	Update(m domain.Room) (domain.Room, error)
	Delete(id uint64) error
}

type roomRepository struct {
	coll db.Collection
	sess db.Session
}

func NewRoomRepository(dbSession db.Session) RoomRepository {
	return &roomRepository{
		coll: dbSession.Collection(RoomsTableName),
		sess: dbSession,
	}
}

func (r roomRepository) Save(m domain.Room) (domain.Room, error) {
	rom := r.mapDomainToModel(m)
	rom.CreatedDate, rom.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&rom)
	if err != nil {
		return domain.Room{}, err
	}
	m = r.mapModelToDomain(rom)
	return m, nil
}

func (r roomRepository) FindForOrganization(oId uint64) ([]domain.Room, error) {
	var roms []room
	err := r.coll.Find(db.Cond{"organization_id": oId, "deleted_date": nil}).All(&roms)
	if err != nil {
		return nil, err
	}
	res := r.mapModelToDomainCollection(roms)
	return res, nil
}

func (r roomRepository) FindById(id uint64) (domain.Room, error) {
	var rom room
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&rom)
	if err != nil {
		return domain.Room{}, err
	}
	m := r.mapModelToDomain(rom)
	return m, nil
}

func (r roomRepository) Update(m domain.Room) (domain.Room, error) {
	rom := r.mapDomainToModel(m)
	rom.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": rom.Id, "deleted_date": nil}).Update(&rom)
	if err != nil {
		return domain.Room{}, err
	}
	m = r.mapModelToDomain(rom)
	return m, nil
}

func (r roomRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r roomRepository) mapDomainToModel(d domain.Room) room {
	return room{
		Id:             d.Id,
		Name:           d.Name,
		OrganizationId: d.OrganizationId,
		Description:    d.Description,
		CreatedDate:    d.CreatedDate,
		UpdatedDate:    d.UpdatedDate,
		DeletedDate:    d.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomain(d room) domain.Room {
	return domain.Room{
		Id:             d.Id,
		Name:           d.Name,
		OrganizationId: d.OrganizationId,
		Description:    d.Description,
		CreatedDate:    d.CreatedDate,
		UpdatedDate:    d.UpdatedDate,
		DeletedDate:    d.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomainCollection(roms []room) []domain.Room {
	var rooms []domain.Room
	for _, m := range roms {
		rom := r.mapModelToDomain(m)
		rooms = append(rooms, rom)
	}
	return rooms
}
