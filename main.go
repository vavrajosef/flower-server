package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	server "github.com/vavrajosef/flower-server/internal/server"
)

func main() {
	router := httprouter.New()
	router.GET("/api/flowers", server.GetAllFlowers)
	router.GET("/api/flowers/:id", server.GetFlower)
	router.GET("/api/unwatered-flowers", server.GetFlowersToWater)
	router.PUT("/api/flowers/:id", server.UpdateFlower)
	router.POST("/api/flowers", server.AddFlower)
	router.DELETE("/api/flowers/:id", server.RemoveFlower)

	log.Fatal(http.ListenAndServe(":8080", router))
}
