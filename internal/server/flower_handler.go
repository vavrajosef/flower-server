package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	api "github.com/vavrajosef/flower-server/api"
	db "github.com/vavrajosef/flower-server/internal/database"
)

func GetAllFlowers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	flowers, err := db.GetAllFlowers()
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
	flowerId := ps.ByName("id")
	flower, err := db.GetFlower(flowerId)
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
	flowers, err := db.GetUnwateredFlowers()
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
	flowerId := ps.ByName("id")
	err := db.UpdateFlower(flowerId)
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
	err := db.AddFlower(flowerDetail)
	if err == nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
	}
}

func RemoveFlower(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	flowerId := ps.ByName("id")
	err := db.RemoveFlower(flowerId)
	if err == nil {
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
		fmt.Println("Error occured: ", err.Error())
		json.NewEncoder(w).Encode("Internal error")
	}
}
