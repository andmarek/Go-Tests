package main

import (
  "log" //Simple logging package, defines Logger type
        //Fatal[f|ln], Panic, Print std helper functions
  //"fmt"
  "net/http" //HTTP client and server implementations
             // Get, head, post, and post form make HTTP(S) requests
  "encoding/json"
  "github.com/gorilla/mux"
)

type User struct {
  ID string `json:"id"`
  Username string `json:"Username"`
  Password string `json:"Password"`
}

var users []User //var users is array of type User

func getUserHandler(w http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  for _, user := range users {
    if user.ID == vars["id"] {
      json.NewEncoder(w).Encode(user)
      return
    }
    json.NewEncoder(w).Encode(&User{})
  }
}

func getUsersHandler(w http.ResponseWriter, req *http.Request) {
  json.NewEncoder(w).Encode(users)
}

func createUserHandler(w http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  var user User
  _ = json.NewDecoder(req.Body).Decode(&user)
  user.ID = vars["id"]
  users = append(users, user)
  json.NewEncoder(w).Encode(users)
}

func deleteUserHandler(w http.ResponseWriter, req *http.Request) {
  vars := mux.Vars(req)
  for index, item := range users {
    if item.ID == vars["id"] {
      users = append(users[:index], users[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(users)
}

func main() {
  r := mux.NewRouter()

  users = append(users, User{ID: "1", Username: "Andrew", Password: "abracadabra"})

  users = append(users, User{ID: "2", Username: "haptic", Password: "dogs"})

  r.HandleFunc("/Users", getUsersHandler).Methods("GET")
  r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
  r.HandleFunc("/user/{id}", createUserHandler).Methods("POST")
  r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":12345", r)) //Server
}
