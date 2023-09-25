package database

import (
	api "github.com/vavrajosef/flower-server/api"
	filter "github.com/vavrajosef/flower-server/internal/flower_filter"
)

func GetFlower(id string) (api.FlowerDetail, error) {
	return api.FlowerDetail{}, nil
}

func GetAllFlowers() ([]api.Flower, error) {
	return []api.Flower{}, nil
}

func getAllFlowerDetails() ([]api.FlowerDetail, error) {
	return []api.FlowerDetail{}, nil
}

func UpdateFlower(id string) error {
	return nil
}

func GetUnwateredFlowers() ([]api.FlowerDetail, error) {
	flowers, err := getAllFlowerDetails()
	if err != nil {
		return nil, err
	}
	result := filter.FilterUnwatered(flowers)
	return result, nil
}

func RemoveFlower(id string) error {
	return nil
}

func AddFlower(flower api.FlowerDetail) error {
	return nil
}
