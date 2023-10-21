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

	router.GET("/api/flowers", CORS(InjectContext(server.GetAllFlowers)))
	router.GET("/api/flowers/:id", CORS(InjectContext(server.GetFlower)))
	router.GET("/api/unwatered-flowers", CORS(InjectContext(server.GetFlowersToWater)))
	router.PUT("/api/flowers/:id", CORS(InjectContext(server.UpdateFlower)))
	router.POST("/api/flowers", CORS(InjectContext(server.AddFlower)))
	router.DELETE("/api/flowers/:id", CORS(InjectContext(server.RemoveFlower)))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func InjectContext(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "connection-string", createConnString())
		h(w, r.WithContext(ctx), ps)
	}
}

func CORS(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

		if r.Method == "OPTIONS" {
			http.Error(w, "No Content", http.StatusNoContent)
			return
		}
		next(w, r, ps)
	}
}

func createConnString() string {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DBHOST, DBPORT, DBUSER, DBPASS, DBNAME)
	return psqlInfo
}
