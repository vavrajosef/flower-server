package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	server "github.com/vavrajosef/flower-server/internal/server"
)

var DBNAME string
var DBPASS string
var DBUSER string
var DBHOST string
var DBPORT string

func main() {
	DBHOST = "postgres"
	DBPORT = "5432"
	DBNAME = os.Getenv("DBNAME")
	DBUSER = os.Getenv("DBUSER")
	DBPASS = os.Getenv("DBPASS")
	router := httprouter.New()

	router.GET("/api/flowers", server.GetAllFlowers)
	router.GET("/api/flowers/:id", server.GetFlower)
	router.GET("/api/unwatered-flowers", server.GetFlowersToWater)
	router.PUT("/api/flowers/:id", server.UpdateFlower)
	router.POST("/api/flowers", server.AddFlower)
	router.DELETE("/api/flowers/:id", server.RemoveFlower)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func InjectContext(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "connection-string", createConnString())
		h(w, r.WithContext(ctx), ps)
	}
}

func createConnString() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DBHOST, DBPORT, DBUSER, DBPASS, DBNAME)
	return psqlInfo
}
