package api

import (
	"encoding/json"
	"fmt"
	"inventory-app/ent/group"
	"inventory-app/ent/item"
	"net/http"
	"strconv"
)

type groupp struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func (api *API) GetGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	groupTitle := r.URL.Query().Get("title")

	if len(groupTitle) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The title of the group cannot be blank."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	entGroup, err := api.Db.Group.Query().Where(group.Title(groupTitle)).First(ctx)
	if err != nil {
		msg := ReqError{Err: "The group does not exist."}
		w.WriteHeader(http.StatusNotFound)
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	marshalled, _ := json.Marshal(entGroup)
	w.Write(marshalled)
}

func (api *API) CreateGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	newGroup := new(groupp)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(newGroup)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		SendErrMsg(w, err)
		return
	}

	if len(newGroup.Title) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The length of group's title cannot be 0."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	err = api.Db.Group.Create().SetTitle(newGroup.Title).SetDesc(newGroup.Desc).Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}
}

func (api *API) DeleteGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idd, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The group id cannot be parsed."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	id := int(idd)

	_, err = api.Db.Group.Delete().Where(group.ID(id)).Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}
}

func (api *API) GetGroupItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idd, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The group id cannot be parsed."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	id := int(idd)

	entGroup, err := api.Db.Group.Query().Where(group.ID(id)).First(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The group does not exist."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	entItems, err := api.Db.Group.QueryGroupItems(entGroup).All(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	marshalled, err := json.Marshal(entItems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	w.Write(marshalled)
}

func (api *API) ListGroups(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var entGroups []struct {
		Title string `json:"title"`
		ID    int    `json:"id"`
	}
	err := api.Db.Group.Query().Select(group.FieldTitle, group.FieldID).Scan(ctx, &entGroups)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		SendErrMsg(w, err)
		return
	}

	marshalled, _ := json.Marshal(entGroups)

	w.Write(marshalled)
}

func (api *API) GroupOfItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	itemIdd, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The item id cannot be parsed."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	itemId := int(itemIdd)

	entItem, err := api.Db.Item.Query().Where(item.ID(itemId)).First(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The item does not exist."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	fmt.Printf("entItem: %v\n", entItem)

	entGroup, err := api.Db.Item.QueryGroup(entItem).Select(group.FieldTitle, group.FieldID).First(ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		SendErrMsg(w, err)
		return
	}

	marshalled, _ := json.Marshal(entGroup)
	w.Write(marshalled)
}

func (api *API) AddItemToGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	groupIdd, err := strconv.ParseInt(r.URL.Query().Get("group"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The group id cannot be parsed."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	groupId := int(groupIdd)

	itemIdd, err := strconv.ParseInt(r.URL.Query().Get("item"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The item id cannot be parsed."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	itemId := int(itemIdd)

	entGroup, err := api.Db.Group.Query().Where(group.ID(groupId)).First(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The group does not exist."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	err = api.Db.Item.Update().Where(item.ID(itemId)).ClearGroup().Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	_, err = api.Db.Item.Update().Where(item.ID(itemId)).AddGroup(entGroup).Save(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The item does not exist."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}
}

func (api *API) RemoveItemFromGroup(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	itemIdd, err := strconv.ParseInt(r.URL.Query().Get("item"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The item id cannot be parsed."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	itemId := int(itemIdd)

	fmt.Println(itemId)

	err = api.Db.Item.Update().Where(item.ID(itemId)).ClearGroup().Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}
}
