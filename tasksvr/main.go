package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"victoriajuan/info344-in-class/tasksvr/handlers"
	"victoriajuan/info344-in-class/tasksvr/models/tasks"

	mgo "gopkg.in/mgo.v2"
)

const defaultAddr = ":80"

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = defaultAddr
	}

	//TODO: make connection to the DBMS
	//construct the appropriate tasks.Store
	//construct the handlers.Context
	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalf("error dialing mongo: %v", err)
	}
	mongoStore := tasks.NewMongoStore(mongoSession, "tasks", "tasks")
	handlerCtx := handlers.NewHandlerContext(mongoStore)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/tasks", handlerCtx.TasksHandler)
	mux.HandleFunc("/v1/tasks/", handlerCtx.SpecificTaskHandler)

	fmt.Printf("server is listening at http://%s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
