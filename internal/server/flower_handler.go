package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	api "github.com/vavrajosef/flower-server/api"
	db "github.com/vavrajosef/flower-server/internal/database"
)

func GetAllFlowers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sqlConnection, err := openConnection(r.Context())
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
		return
	}
	defer sqlConnection.Close()
	flowers, err := db.GetAllFlowers(sqlConnection)
	if err == nil {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(flowers)
	} else {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
	}
}

func GetFlower(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	flowerId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
		return
	}
	sqlConnection, err := openConnection(r.Context())
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
		return
	}
	defer sqlConnection.Close()
	flower, err := db.GetFlower(sqlConnection, flowerId)
	if err == nil {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(flower)
	} else {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
	}
}

func GetFlowersToWater(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sqlConnection, err := openConnection(r.Context())
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
		return
	}
	defer sqlConnection.Close()
	flowers, err := db.GetUnwateredFlowers(sqlConnection)
	if err == nil {
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(flowers)
	} else {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
	}
}

func UpdateFlower(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var flowerDetail api.FlowerDetail
	decoder.Decode(&flowerDetail)
	sqlConnection, err := openConnection(r.Context())
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
		return
	}
	defer sqlConnection.Close()
	err = db.UpdateFlower(sqlConnection, flowerDetail)
	if err == nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
	}
}

func AddFlower(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var flowerDetail api.FlowerDetail
	decoder.Decode(&flowerDetail)
	sqlConnection, err := openConnection(r.Context())
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
		return
	}
	defer sqlConnection.Close()
	err = db.AddFlower(sqlConnection, flowerDetail)
	if err == nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
	}
}

func RemoveFlower(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	flowerId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
		return
	}
	sqlConnection, err := openConnection(r.Context())
	if err != nil {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
		return
	}
	defer sqlConnection.Close()
	err = db.RemoveFlower(sqlConnection, flowerId)
	if err == nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
	}
}

func openConnection(ctx context.Context) (*sql.DB, error) {
	connString := ctx.Value("connection-string").(string)
	return sql.Open("postgres", connString)
}
