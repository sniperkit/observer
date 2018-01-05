package datastore

import (
	"github.com/demas/observer/pkg/models"
)

type IDataStore interface {

	// Settings
	GetSettings(key string) string
	SetSettings(key string, value string)

	// StackOverflow
	GetStackTags() []models.StackTag
	GetSecondTagByClassification(classification string) interface{}
	GetStackQuestionsByClassification(classification string) []models.StackQuestion
	GetStackQuestionsByClassificationAndDetails(classification string, details string) []models.StackQuestion
	InsertStackOverflowQuestions(questions map[string][]models.SOQuestion)
	SetStackQuestionAsReaded(question_id int)
	SetStackQuestionsAsReadedByClassification(classification string)
	SetStackQuestionsAsReadedByClassificationFromTime(classification string, t int64)
}

type DataStore struct {}
