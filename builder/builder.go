package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	GetHouse() House
}

func GetHouseBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	} else if builderType == "igloo" {
		return newIglooBuilder()
	}

	return nil
}
