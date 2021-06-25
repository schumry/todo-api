package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/schumry/todo-api/todo"
)

var (
	todos todo.Todos
)

func get(w http.ResponseWriter, req *http.Request) {
	b, _ := json.Marshal(todos.Get())
	w.Write(b)
}

func delete(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodDelete {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Println("Got error")
		}
		todos.Delete(int64(id))
	}
}

func put(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPut {
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Println("Got error")
		}
		todos.Complete(int64(id))
	}
}

func add(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		var todo todo.Todo
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Println("Got error")
		}
		json.Unmarshal(b, &todo)
		todo = todos.Add(todo)
		c, _ := json.Marshal(todo)
		w.Write(c)
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
		handlers.AllowedMethods([]string{"GET", "OPTIONS", "POST", "DELETE", "PUT"}),
	)

	var r = mux.NewRouter()
	r.Use(cors)
	todos = todo.NewTodos()

	r.HandleFunc("/get", get)
	r.HandleFunc("/delete/{id:[0-9]+}", delete)
	r.HandleFunc("/add", add)
	r.HandleFunc("/headers", headers)
	r.HandleFunc("/complete/{id:[0-9]+}", put)

	http.ListenAndServe(":8090", handlers.LoggingHandler(os.Stdout, r))
}
