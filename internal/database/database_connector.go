package database

import (
	"database/sql"

	api "github.com/vavrajosef/flower-server/api"
	filter "github.com/vavrajosef/flower-server/internal/flower_filter"
)

func GetFlower(conn *sql.DB, id string) (api.FlowerDetail, error) {

	return api.FlowerDetail{}, nil
}

func GetAllFlowers(conn *sql.DB) ([]api.Flower, error) {
	return []api.Flower{}, nil
}

func getAllFlowerDetails(conn *sql.DB) ([]api.FlowerDetail, error) {
	return []api.FlowerDetail{}, nil
}

func UpdateFlower(conn *sql.DB, id string) error {
	return nil
}

func GetUnwateredFlowers(conn *sql.DB) ([]api.FlowerDetail, error) {
	flowers, err := getAllFlowerDetails(conn)
	if err != nil {
		return nil, err
	}
	result := filter.FilterUnwatered(flowers)
	return result, nil
}

func RemoveFlower(conn *sql.DB, id string) error {
	return nil
}

func AddFlower(conn *sql.DB, flower api.FlowerDetail) error {
	return nil
}
