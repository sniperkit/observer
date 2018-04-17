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
	GetSecondTagByClassification(classification string) []models.StackTag
	GetStackQuestionsForRating() []models.StackQuestion
	GetStackQuestionsByClassification(classification string, limit int) []models.StackQuestion
	GetStackQuestionsByClassificationAndDetails(classification string, details string) []models.StackQuestion
	GetStackQuestionsByClassificationDetailsAndSorting(classification string, details string, sorting string) []models.StackQuestion
	InsertStackOverflowQuestions(questions map[string][]models.SOQuestion)
	UpdateStackQuestionRating(id uint32, score int)
	SetStackQuestionAsReaded(question_id int)
	SetStackQuestionsAsReadedByClassification(classification string)
	SetStackQuestionsAsReadedByClassificationFromTime(classification string, t int64)
	ChangeStackTagVisibility(tagId int)

	// Tags
	GetTags() []models.Tag
	GetTaggedItemsByTagId(tagId int) []models.TaggedItem
	InsertTaggedItem(questionId int, tagId int)
	DeleteTaggedItem(id int)
}

type DataStore struct {}
