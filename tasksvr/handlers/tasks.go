package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"victoriajuan/info344-in-class/tasksvr/models/tasks"

	"gopkg.in/mgo.v2/bson"
)

//TasksHandler handles requests for the /v1/tasks resource
func (ctx *Context) TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		//TODO: parse the `completed` query string param as a boolean
		//and call the GetAll() method on the tasks.Store to get the tasks
		tasks, err := ctx.tasksStore.GetAll(false)
		if err != nil {
			http.Error(w, fmt.Sprintf("error getting tasks: %v", err), http.StatusInternalServerError)
			return
		}
		respond(w, tasks)
	case "POST":
		//TODO: decode the request body into a tasks.NewTask
		//and insert it using the tasks.Store
		nt := &tasks.NewTask{}
		if err := json.NewDecoder(r.Body).Decode(nt); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}

		task, err := ctx.tasksStore.Insert(nt)
		if err != nil {
			http.Error(w, fmt.Sprintf("error inserting task: &v", err), http.StatusInternalServerError)
		}
		respond(w, task)
	default:
		http.Error(w, "method must be GET or POST", http.StatusMethodNotAllowed)
		return
	}
}

//SpecificTaskHandler handles requests for the /v1/tasks/...task-id... resource
func (ctx *Context) SpecificTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	oid := bson.ObjectIdHex(id)
	switch r.Method {
	case "PATCH":
		//TODO: decode the request body into a
		//tasks.TaskUpdates struct and pass that
		//to the Update() method on the tasks.Store
		tu := &tasks.TaskUpdates{}
		if err := json.NewDecoder(r.Body).Decode(tu); err != nil {
			http.Error(w, fmt.Sprintf("error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}
		tasks, err := ctx.tasksStore.Update(oid, tu); err != nil {
			http.Error(w, fmt.Sprintf("error updatring task: %v", err), http.StatusInternalServerError)
			return
		}
		respons(w, task)
	default:
		http.Error(w, "method must be PATCH", http.StatusMethodNotAllowed)
		return
	}
}
