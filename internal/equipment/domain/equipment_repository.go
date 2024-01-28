package domain

type EquipmentRepository interface {
	Save(equipment *Equipment) error
	FindOneById(id string) (*Equipment, error)
}
