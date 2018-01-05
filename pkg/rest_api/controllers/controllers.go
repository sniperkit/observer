package controllers

import (
	"github.com/demas/observer/pkg/datastore"
	"github.com/demas/observer/pkg/services/service_locator"
	"net/http"
	"log"
	"encoding/json"
)

var ds datastore.IDataStore

func init() {
	ds = service_locator.GetDataStore()
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type UniversalPostStruct struct {
	Id int `json:"id"`
	FromTime int64 `json:"fromTime"`
	Tag string `json:"tag"`
}

type HandlerFunc func(w http.ResponseWriter, req *http.Request) (interface{}, error)

func WrapHandler(handler HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		data, err := handler(w, req)

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		} else {
			w.Header().Add("Content-Type", "application/json")
			resp, _ := json.Marshal(data)
			w.Write(resp)
		}
	}
}

func PostWrapHandler(handler HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {

		bodyData, err := handler(w, req)

		if err != nil {
			log.Println(err)
			w.WriteHeader(500)
		} else {
			w.Header().Add("Content-Type", "application/json")
			resp, _ := json.Marshal(bodyData)
			w.Write(resp)
		}
	}
}

