package datastore

import (
	"github.com/demas/observer/pkg/models"
)

func (ds *DataStore) GetTags() []models.Tag {

	result := []models.Tag{}
	db.Find(&result)
	return result
}

func (ds *DataStore) GetTaggedItemsByTagId(tagId int) []models.TaggedItem {

	result := []models.TaggedItem{}
	db.Model(models.TaggedItem{}).Where("tagid = ?", tagId)
	return result
}

func (ds *DataStore) InsertTaggedItem(questionId int, tagId int) {

	// TODO: доделать
}

func (ds *DataStore) DeleteTaggedItem(id int) {

	// TODO: доделать
}