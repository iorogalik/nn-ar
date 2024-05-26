package database

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const DevicesTableName = "devices"

type DeviceCategory string

const (
	SENSOR   DeviceCategory = "SENSOR"
	ACTUATOR DeviceCategory = "ACTUATOR"
)

type device struct {
	Id               uint64         `db:"id,omitempty"`
	OrganizationId   uint64         `db:"organization_id"`
	RoomId           uint64         `db:"room_id"`
	GUID             string         `db:"guid"`
	InventoryNumber  string         `db:"inventory_number"`
	SerialNumber     string         `db:"serial_number"`
	Characteristics  string         `db:"characteristics"`
	Category         DeviceCategory `db:"device_category"`
	Units            string         `db:"units"`
	PowerConsumption float64        `db:"power_consumption"`
	CreatedDate      time.Time      `db:"created_date"`
	UpdatedDate      time.Time      `db:"updated_date"`
	DeletedDate      *time.Time     `db:"deleted_date"`
}

type DeviceRepository interface {
	Save(dv domain.Device) (domain.Device, error)
	Update(dv domain.Device) (domain.Device, error)
	FindForRoom(mId uint64) ([]domain.Device, error)
	FindById(id uint64) (domain.Device, error)
	Delete(id uint64) error
}

type deviceRepository struct {
	coll db.Collection
	sess db.Session
}

func NewDeviceRepository(dbSession db.Session) DeviceRepository {
	return deviceRepository{
		coll: dbSession.Collection(DevicesTableName),
		sess: dbSession,
	}
}

func (r deviceRepository) Save(dv domain.Device) (domain.Device, error) {
	dev := r.mapDomainToModel(dv)
	dv.CreatedDate, dv.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&dv)
	if err != nil {
		return domain.Device{}, err
	}
	dv = r.mapModelToDomain(dev)
	return dv, nil
}

func (r deviceRepository) Update(dv domain.Device) (domain.Device, error) {
	dev := r.mapDomainToModel(dv)
	dev.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": dev.Id, "deleted_date": nil}).Update(&dev)
	if err != nil {
		return domain.Device{}, err
	}
	dv = r.mapModelToDomain(dev)
	return dv, nil
}

func (r deviceRepository) FindForRoom(mId uint64) ([]domain.Device, error) {
	var devs []device
	err := r.coll.Find(db.Cond{"room_id": mId, "deleted_date": nil}).All(&devs)
	if err != nil {
		return nil, err
	}
	res := r.mapModelToDomainCollection(devs)
	return res, nil
}

func (r deviceRepository) FindById(id uint64) (domain.Device, error) {
	var dev device
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&dev)
	if err != nil {
		return domain.Device{}, err
	}
	dv := r.mapModelToDomain(dev)
	return dv, nil
}

func (r deviceRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r deviceRepository) mapDomainToModel(dv domain.Device) device {
	return device{
		Id:               dv.Id,
		OrganizationId:   dv.OrganizationId,
		RoomId:           dv.RoomId,
		GUID:             dv.GUID,
		InventoryNumber:  dv.InventoryNumber,
		SerialNumber:     dv.SerialNumber,
		Characteristics:  dv.Characteristics,
		Category:         DeviceCategory(dv.Category),
		Units:            dv.Units,
		PowerConsumption: dv.PowerConsumption,
		CreatedDate:      dv.CreatedDate,
		UpdatedDate:      dv.UpdatedDate,
		DeletedDate:      dv.DeletedDate,
	}
}

func (r deviceRepository) mapModelToDomain(dv device) domain.Device {
	return domain.Device{
		Id:               dv.Id,
		OrganizationId:   dv.OrganizationId,
		RoomId:           dv.RoomId,
		GUID:             dv.GUID,
		InventoryNumber:  dv.InventoryNumber,
		SerialNumber:     dv.SerialNumber,
		Characteristics:  dv.Characteristics,
		Category:         string(dv.Category),
		Units:            dv.Units,
		PowerConsumption: dv.PowerConsumption,
		CreatedDate:      dv.CreatedDate,
		UpdatedDate:      dv.UpdatedDate,
		DeletedDate:      dv.DeletedDate,
	}
}

func (r deviceRepository) mapModelToDomainCollection(devs []device) []domain.Device {
	var devices []domain.Device
	for _, dv := range devs {
		dev := r.mapModelToDomain(dv)
		devices = append(devices, dev)
	}
	return devices
}
