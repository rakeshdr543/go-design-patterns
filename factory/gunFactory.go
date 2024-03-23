package factory

import "errors"

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return NewAK47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, errors.New("wrong gun type passed")
}
