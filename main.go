package main

import (
	"fmt"
	"go-crud/config"
	"go-crud/routers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDB()
	config.Initialize()
	fmt.Println("after ConnectDB")
	muxRouter := mux.NewRouter()

	routers.RegisterUserRoutes(muxRouter)
	http.Handle("/", muxRouter)
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
