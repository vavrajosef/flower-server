package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	api "github.com/vavrajosef/flower-server/api"
	filter "github.com/vavrajosef/flower-server/internal/flower_filter"
)

func GetFlower(conn *sql.DB, id int) (api.FlowerDetail, error) {
	rows, err := conn.Query("select Id, Name, Period, LastWatered from flower where Id='%d'", id)
	if err != nil {
		return api.FlowerDetail{}, err
	}
	defer rows.Close()

	for rows.Next() {
		return rowToFlowerDetail(rows), nil
	}
	return api.FlowerDetail{}, errors.New("Flower not found")
}

func GetAllFlowers(conn *sql.DB) ([]api.Flower, error) {
	rows, err := conn.Query("select Id, Name from flowers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []api.Flower
	for rows.Next() {
		result = append(result, rowToFlower(rows))
	}
	return result, nil
}

func getAllFlowerDetails(conn *sql.DB) ([]api.FlowerDetail, error) {
	rows, err := conn.Query("select * from flowers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []api.FlowerDetail
	for rows.Next() {
		result = append(result, rowToFlowerDetail(rows))
	}
	return result, nil
}

func UpdateFlower(conn *sql.DB, flowerToUpdate api.FlowerDetail) error {
	updateStmt := fmt.Sprintf("update flower set Name='%s', Period='%d', LastWatered='%s' where Id='%d'", flowerToUpdate.Name, flowerToUpdate.Period, flowerToUpdate.LastWatered.String(), flowerToUpdate.Id)
	_, err := conn.Exec(updateStmt)
	return err
}

func GetUnwateredFlowers(conn *sql.DB) ([]api.FlowerDetail, error) {
	flowers, err := getAllFlowerDetails(conn)
	if err != nil {
		return nil, err
	}
	result := filter.FilterUnwatered(flowers)
	return result, nil
}

func RemoveFlower(conn *sql.DB, id int) error {
	deleteStmt := fmt.Sprintf("delete from flower where Id = %d", id)
	_, err := conn.Exec(deleteStmt)
	return err
}

func AddFlower(conn *sql.DB, flower api.FlowerDetail) error {
	insertStmt := fmt.Sprintf("insert into flower('Name', 'Period', 'LastWatered') values('%s', '%d', '%s');", flower.Name, flower.Period, flower.LastWatered.String())
	_, err := conn.Exec(insertStmt)
	return err
}

func rowToFlowerDetail(rows *sql.Rows) api.FlowerDetail {
	var Id int
	var Name string
	var Period int
	var LastWatered time.Time
	rows.Scan(&Id, &Name, &Period, &LastWatered)
	return api.FlowerDetail{
		Id:          Id,
		Name:        Name,
		Period:      Period,
		LastWatered: LastWatered,
	}
}

func rowToFlower(rows *sql.Rows) api.Flower {
	var Id int
	var Name string
	rows.Scan(&Id, &Name)
	return api.Flower{
		Id:   Id,
		Name: Name,
	}
}
