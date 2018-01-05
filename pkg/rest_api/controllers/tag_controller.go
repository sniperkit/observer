package controllers

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func GetTags(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return ds.GetTags(), nil
}

func GetTaggedItemsByTagId(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	tagId, _ :=  strconv.Atoi(mux.Vars(r)["tagId"])
	return ds.GetTaggedItemsByTagId(tagId), nil
}

func InsertTaggedItem (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	type SetTaggedItemStruct struct {
		ItemId int `json:"itemId"`
		TagId int `json:"tagId"`
		Source int `json:"source"`
	}

	bodyData := new(SetTaggedItemStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		if bodyData.Source == 1 {
			ds.InsertTaggedItem(bodyData.ItemId, bodyData.TagId)
		} //else if bodyData.Source == 2 {
			//(&postgres.TagsService{}).InsertTaggedItemFromRss(bodyData.ItemId, bodyData.TagId)
	//	}
	}

	return bodyData, nil
}

func DeleteTaggedItem(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	id, _ :=  strconv.Atoi(mux.Vars(r)["id"])
	ds.DeleteTaggedItem(id)
	return ResultOk{"ok"}, nil
}

