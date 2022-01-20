package api

import (
	"encoding/json"
	"fmt"
	"inventory-app/ent"
	"inventory-app/ent/item"
	"inventory-app/ent/tag"
	"net/http"
	"strconv"
)

type itemm struct {
	Title string   `json:"title"`
	Desc  string   `json:"desc"`
	Tags  []string `json:"tags"`
	Count int      `json:"count"`
}

func (api *API) Item(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	title := r.URL.Query().Get("title")

	if len(title) == 0 {
		w.WriteHeader(200)
		msg := ReqError{Err: "The length of tag's title cannot be 0."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	item, err := api.Db.
		Item.
		Query().
		Where(item.Title(title)).
		First(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		SendErrMsg(w, err)
		return
	}

	marshalled, err := json.Marshal(item)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	w.Write(marshalled)
}

func (api *API) AddItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	newItem := new(itemm)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(newItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		SendErrMsg(w, err)
		return
	}

	exists, err := api.Db.
		Item.
		Query().
		Where(item.Title(newItem.Title)).
		Exist(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	if exists {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "Item with that title already exists."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	tagCreators := make([]*ent.TagCreate, 0)
	for _, tagg := range newItem.Tags {
		tagExists, err := api.Db.
			Tag.
			Query().
			Where(tag.Title(tagg)).
			Exist(ctx)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			SendErrMsg(w, err)
			return
		}
		if !tagExists {
			tagCreators = append(tagCreators, api.Db.Tag.Create().SetTitle(tagg))
		}
	}

	_, err = api.Db.Tag.CreateBulk(tagCreators...).Save(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	itemCreator := api.Db.
		Item.
		Create().
		SetDesc(newItem.Desc).
		SetCount(newItem.Count).
		SetTitle(newItem.Title).
		SetTags(newItem.Tags)

	for _, tagg := range newItem.Tags {
		entTag, err := api.Db.
			Tag.
			Query().
			Where(tag.Title(tagg)).
			First(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			SendErrMsg(w, err)
			return
		}
		itemCreator = itemCreator.AddTag(entTag)
		fmt.Printf("adding tagg: %v\n", tagg)
	}

	entItem, err := itemCreator.Save(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		SendErrMsg(w, err)
		return
	}

	res := struct {
		Id int `json:"id"`
	}{Id: entItem.ID}
	marshalled, _ := json.Marshal(res)
	w.Write(marshalled)
}

func (api *API) DeleteItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	itemIdd, err := strconv.
		ParseInt(r.URL.Query().
			Get("id"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "The item id could not be parsed."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	itemID := int(itemIdd)

	_, err = api.Db.Item.Delete().Where(item.ID(itemID)).Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}
}

func (api *API) UpdateItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idd, err := strconv.
		ParseInt(r.URL.Query().
			Get("id"), 10, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		SendErrMsg(w, err)
		return
	}

	id := int(idd)

	updatedItem := new(itemm)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(updatedItem)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		SendErrMsg(w, err)
		return
	}

	itemWithTitle, _ := api.Db.Item.
		Query().
		Where(item.Title(updatedItem.Title)).
		First(ctx)

	if itemWithTitle != nil && itemWithTitle.ID != id {
		w.WriteHeader(http.StatusBadRequest)
		msg := ReqError{Err: "Item with that title already exists."}
		marshalled, _ := json.Marshal(msg)
		w.Write(marshalled)
		return
	}

	tagCreators := make([]*ent.TagCreate, 0)
	for _, tagg := range updatedItem.Tags {
		tagExists, err := api.Db.
			Tag.
			Query().
			Where(tag.Title(tagg)).
			Exist(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			SendErrMsg(w, err)
			return
		}
		if !tagExists {
			tagCreators = append(tagCreators, api.Db.Tag.Create().SetTitle(tagg))
		}
	}

	_, err = api.Db.Tag.CreateBulk(tagCreators...).Save(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	_, err = api.Db.Item.
		Update().
		Where(item.ID(id)).
		SetTitle(updatedItem.Title).
		SetDesc(updatedItem.Desc).
		SetCount(updatedItem.Count).
		SetTags(updatedItem.Tags).
		Save(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	itemm, err := api.Db.
		Item.
		Query().
		Where(item.ID(id)).
		First(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		SendErrMsg(w, err)
		return
	}

	itemUpdater := api.Db.
		Item.
		Update().
		Where(item.ID(id))

	entTags, err := api.Db.
		Item.
		QueryTag(itemm).
		All(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	// remove old tags
	for _, tagg := range entTags {
		itemUpdater = itemUpdater.RemoveTag(tagg)
	}

	// add new tags
	for _, tagg := range updatedItem.Tags {
		entTag, err := api.Db.
			Tag.
			Query().
			Where(tag.Title(tagg)).
			First(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			SendErrMsg(w, err)
			return
		}
		itemUpdater = itemUpdater.AddTag(entTag)
	}

	if itemUpdater != nil {
		_, err = itemUpdater.Save(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			SendErrMsg(w, err)
			return
		}
	}

	for _, tagg := range entTags {
		itemms, err := api.Db.Tag.QueryTagItems(tagg).All(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			SendErrMsg(w, err)
			return
		}

		if len(itemms) == 0 {
			fmt.Printf("deleting tag: %v\n", tagg)
			api.Db.
				Tag.
				Delete().
				Where(tag.Title(tagg.Title)).
				Exec(ctx)
		}
	}

	marshalled, _ := json.Marshal(itemm)
	w.Write(marshalled)
}

func (api *API) ItemsWithTag(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tagTitle := r.URL.Query().Get("tag")

	if len(tagTitle) == 0 {
		w.WriteHeader(200)
		return
	}

	entTag, err := api.Db.
		Tag.
		Query().
		Where(tag.Title(tagTitle)).
		First(ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		SendErrMsg(w, err)
		return
	}

	entItems, err := api.Db.
		Tag.
		QueryTagItems(entTag).
		Select(item.FieldTitle).
		Select(item.FieldCount).
		Select(item.FieldUpdatedAt).
		All(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	fmt.Printf("entItems: %v\n", entItems)

	marshalled, _ := json.Marshal(entItems)
	w.Write(marshalled)
}

func (api *API) ListItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	entItems, err := api.Db.
		Item.Query().
		Select(item.FieldTitle, item.FieldUpdatedAt, item.FieldCount).
		All(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		SendErrMsg(w, err)
		return
	}

	marshalled, _ := json.Marshal(entItems)
	w.Write(marshalled)
}
