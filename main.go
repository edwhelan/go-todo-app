package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Todo struct {
	ID		int 		`json:"id"`
	UserName	string		`json:"user_name"`
	Title	string		`json:"title"`
	TextField string	`json:"text_field"`

}
//slice of todos
var todos []Todo

// get all todos
//example call: localhost:8080/api/todos
func getTodos(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
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

	log.Fatal(http.ListenAndServe(":8080", r))
}

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//
//	"github.com/gorilla/mux"
//)
//
//type event struct {
//	ID		string `json:"ID"`
//	Title 	string `json:"Title"`
//	Description string `json:"Description"`
//}
//
//type allEvents []event
//
//var events = allEvents{
//	{
//		ID: 		"1",
//		Title:		"Introduction to Golang",
//		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
//	},
//}
//
//func createEvent(w http.ResponseWriter, r *http.Request){
//	var newEvent event
//	reqBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
//	}
//
//	json.Unmarshal(reqBody, &newEvent)
//	events = append(events, newEvent)
//	w.WriteHeader(http.StatusCreated)
//
//	json.NewEncoder(w).Encode(newEvent)
//}
//
//func getOneEvent(w http.ResponseWriter, r *http.Request){
//	eventID := mux.Vars(r)["id"]
//
//	for _,singleEvent := range events {
//		if singleEvent.ID == eventID {
//			json.NewEncoder(w).Encode(singleEvent)
//		}
//	}
//}
//
//func getAllEvents(w http.ResponseWriter, r *http.Request){
//	json.NewEncoder(w).Encode(events)
//}
//
//func homeLink(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Welcome home!")
//}
//
//func main() {
//	router := mux.NewRouter().StrictSlash(true)
//	router.HandleFunc("/", homeLink)
//	router.HandleFunc("/event", createEvent)
//	router.HandleFunc("/events/{id}", getOneEvent)
//	log.Fatal(http.ListenAndServe(":8080", router))
//	//fmt.Println("Hello World!")
//}