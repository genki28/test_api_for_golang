package main

import(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Item representation
type Item struct {
	Title 		string `json:"title"`
	Description string `json:"description"`
}

// Global, static list of items
var itemList = []Item {
	Item{Title: "Item A", Description: "The first item"},
	Item{Title: "Item B", Description: "The second item"},
}

// Controller for the / route(home)
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("テスト成功")
	// fmt.Fprintf(w, "This is the home page, Welcom!")
}

// Controller for the /items route
func returnAllItems(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, itemList)
}

// Controller for the /items/{id} route
func returnSingleItem(w http.ResponseWriter, r *http.Request) {
	// Get query parameters using Mux
	vars := mux.Vars(r)

	// Convert {id} parameter from string to int
	key, err := strconv.Atoi(vars["id"])

	// If {id} parameter is not valid int
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// If Item with ID of {id} does not exist
	if key >= len(itemList) {
		respondWithError(w, http.StatusNotFound, "Item does not exist")
		return
	}

	respondWithJson(w, http.StatusOK, itemList[key])
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/items", returnAllItems)
	myRouter.HandleFunc("/items/{id}", returnSingleItem)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}