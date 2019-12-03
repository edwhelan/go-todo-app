package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID			int 		`json:"id"`
	UserName	string		`json:"user_name"`
	Title		string		`json:"title"`
	TextField 	string		`json:"text_field"`
}

//slice of todos
var todos []Todo

// get all todos
//example call: localhost:8080/api/todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(todos)
}

//make a newtodo
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTodo Todo
	_ = json.NewDecoder(r.Body).Decode(&newTodo)

	newTodo.ID = todos[len(todos)-1].ID + 1

	todos = append(todos, newTodo)
	_ = json.NewEncoder(w).Encode(newTodo)
}
//update existingtodo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//get params and convert id to int type
	params := mux.Vars(r)
	todoId, _ := strconv.Atoi(params["id"])

	for index, item := range todos {
		if item.ID == todoId {
			retainedId := item.ID     //retain this ID
			todos = append(todos[:index], todos[index+1:]...) //delete the newtodo?

		var todo Todo 									//instantiate new user
			_ = json.NewDecoder(r.Body).Decode(&todo)  	//fill out info from body
			todo.ID = retainedId						//use retained id
			todos = append(todos, todo) 				//re-add the newtodo
			_ = json.NewEncoder(w).Encode(todo)  		//return the newtodo
			return
		}
	}
	_ = json.NewEncoder(w).Encode(&Todo{})
}

// delete a newtodo
func deletetodo(w http.ResponseWriter, r*http.Request){
	w.Header().Set("Content-Type", "application/json")

	//get params and convert id to int type
	params := mux.Vars(r)
	todoId, _ := strconv.Atoi(params["id"])

	for index, item := range todos {
		if item.ID == todoId{
			todos = append(todos[:index], todos[index+1:]...)
			break
		}
	}
	_ = json.NewEncoder(w).Encode(todos)

}

func main(){
	//initialize router
	r := mux.NewRouter()

	//fuctional todos
	// @todo implement db
	todos = append(todos, Todo{ID: 1, UserName:"Ed", Title: "Feed Cat", TextField: "feed cats 1 scoop of food per .2 hours"})
	todos = append(todos, Todo{ID: 2, UserName:"Ed", Title: "Feed dogs", TextField: "feed dogs 1 scoop of food per 5 hours"})

	//route handling
	r.HandleFunc("/api/todos", getTodos).Methods("GET")
	r.HandleFunc("/api/newtodo", createTodo).Methods("POST")
	r.HandleFunc("/api/updatetodo/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/api/deletetodo/{id}", deletetodo).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8080", r))
}
