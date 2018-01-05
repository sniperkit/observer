package datastore

import "github.com/demas/observer/pkg/models"

type IDataStore interface {

	// Settings
	GetSettings(key string) string
	SetSettings(key string, value string)

	// StackOverflow
	InsertStackOverflowQuestions(questions map[string][]models.SOQuestion)
}

type DataStore struct {}
