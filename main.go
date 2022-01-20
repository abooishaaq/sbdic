package main

import (
	"inventory-app/api"
	"inventory-app/ent"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:db?cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	api := new(api.API)
	api.Db = client

	r := mux.NewRouter()

	r.HandleFunc("/api/get-item", api.Item)
	r.HandleFunc("/api/add-item", api.AddItem)
	r.HandleFunc("/api/list-items", api.ListItems)
	r.HandleFunc("/api/delete-item", api.DeleteItem)
	r.HandleFunc("/api/update-item", api.UpdateItem)
	r.HandleFunc("/api/items-with-tag", api.ItemsWithTag)

	r.HandleFunc("/api/get-group", api.GetGroup)
	r.HandleFunc("/api/list-groups", api.ListGroups)
	r.HandleFunc("/api/create-group", api.CreateGroup)
	r.HandleFunc("/api/delete-group", api.DeleteGroup)
	r.HandleFunc("/api/group-of-item", api.GroupOfItem)
	r.HandleFunc("/api/get-group-items", api.GetGroupItems)
	r.HandleFunc("/api/add-item-to-group", api.AddItemToGroup)
	r.HandleFunc("/api/remove-item-from-group", api.RemoveItemFromGroup)

	http.Handle("/api/", r)
	http.Handle("/", http.FileServer(http.Dir("./client/public")))

	http.ListenAndServe(":3000", nil)
	defer client.Close()
}
