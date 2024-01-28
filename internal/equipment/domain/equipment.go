package domain

type Equipment struct {
	Id        int64
	Version   int64
	Name      string
	Quantity  int64
	Aggregate Aggregate
}

type Aggregate struct {
	Event []map[string]interface{}
}

func CreateNewEquipment(name string, quantity int64) (Equipment, error) {
	equipment := &Equipment{
		Name:     name,
		Quantity: quantity,
	}

	equipment.equipmentPushEvent("CREATE_EQUIPMENT")
	return *equipment, nil
}

func (equipment *Equipment) equipmentPushEvent(eventName string) {
	event := map[string]interface{}{
		"event": eventName,
		"data":  equipment,
	}

	equipment.Aggregate.Event = append(equipment.Aggregate.Event, event)
}
