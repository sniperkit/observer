package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func GetStackTags(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return ds.GetStackTags(), nil
}

func GetSecondTagsByClassification(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return ds.GetSecondTagByClassification(mux.Vars(r)["classification"]), nil
}

func GetStackQuestionsByClassification(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return ds.GetStackQuestionsByClassification(mux.Vars(r)["classification"], 15), nil
}

func GetStackQuestionsByClassificationAndDetails(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	return ds.GetStackQuestionsByClassificationAndDetails(mux.Vars(r)["classification"],
		mux.Vars(r)["details"]), nil
}

func SetStackQuestionAsReaded (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	type SetStackQuestionAsReadedStruct struct {
		QuestionId int `json:"questionid"`
	}

	bodyData := new(SetStackQuestionAsReadedStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		ds.SetStackQuestionAsReaded(bodyData.QuestionId)
	}

	return bodyData, err
}

func SetStackQuestionsAsReaded (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(UniversalPostStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		ds.SetStackQuestionsAsReadedByClassification(bodyData.Tag)
	}

	return bodyData, err
}

func SetStackQuestionsAsReadedFromTime (w http.ResponseWriter, r *http.Request) (interface{}, error) {

	bodyData := new(UniversalPostStruct)
	err := json.NewDecoder(r.Body).Decode(&bodyData)
	if err == nil {
		ds.SetStackQuestionsAsReadedByClassificationFromTime(bodyData.Tag, bodyData.FromTime)
	}

	return bodyData, err
}

