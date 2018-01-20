package datastore

import (
	"github.com/demas/observer/pkg/models"
	"strings"
	"fmt"
)

// возвращает непрочитанные теги первого уровня
func (ds *DataStore) GetStackTags() []models.StackTag {

	result := []models.StackTag{}
	db.Where("unreaded > 0 and details = ''").Order("hidden, classification").Find(&result)
	return result
}

func (ds *DataStore) GetSecondTagByClassification(classification string) interface{} {

	// TODO: от этой структуры можно избавиться и просто возвращать коллекцию тегов
	type Result struct {
		Details string `json:"details"`
		Count   int    `json:"count"`
	}

	result := []Result{}
	db.Model(StackTag{}).
		Select("details, unreaded as count").
			Where("classification = ? and details != '' and unreaded > 0", classification).
				Order("hidden, details").Scan(&result)

	return result
}

func (ds *DataStore) GetStackQuestionsByClassification(classification string, limit int) []models.StackQuestion {

	result := []models.StackQuestion{}
	db.Model(StackQuestion{}).
		Where("classification = ? and readed = 0", classification).Order("score desc").Limit(limit).Find(&result)

	return result
}

func (ds *DataStore) GetStackQuestionsForRating() []models.StackQuestion {

	result := []models.StackQuestion{}

	db.Table("stack_questions").
		Joins("join stack_tags on stack_questions.classification = stack_tags.classification and stack_questions.details = stack_tags.details").
			Where("stack_questions.readed = 0 and stack_tags.hidden = 0").Scan(&result)

	return result
}

func (ds *DataStore) GetStackQuestionsByClassificationAndDetails(classification string, details string) []models.StackQuestion {

	result := []models.StackQuestion{}
	db.Model(StackQuestion{}).Where("classification = ? and details = ? and readed = 0", classification, details).
		Order("score desc").Limit(15).Find(&result)
	return result
}

func (ds *DataStore) InsertStackOverflowQuestions(questionsMap map[string][]models.SOQuestion) {

	for site, questions  := range questionsMap {
		for _, question := range questions {

			tx := db.Begin()

			// сохраняем вопрос
			dbQuestion := StackQuestion{}
			dbQuestion.Title = question.Title
			dbQuestion.Link = question.Link
			dbQuestion.QuestionId = question.Question_id
			dbQuestion.Tags = strings.Join(question.Tags[:], ",")
			dbQuestion.Score = question.Score
			dbQuestion.AnswerCount = question.Answer_count
			dbQuestion.ViewCount = question.View_count
			dbQuestion.UserId = question.Owner.User_id
			dbQuestion.UserReputation = question.Owner.Reputation
			dbQuestion.UserDisplayName = question.Owner.Display_name
			dbQuestion.UserProfileImage = question.Owner.Profile_image
			dbQuestion.Classification = question.Classification
			dbQuestion.Details = question.Details
			dbQuestion.CreationDate = question.Creation_date
			dbQuestion.Readed = 0
			dbQuestion.Favorite = 0
			dbQuestion.Classified = 1
			dbQuestion.Site = site
			tx.Save(&dbQuestion)

			// обновляем тег первого уровня
			stackTag := StackTag{}
			tx.Where("classification = ? and details = ''", dbQuestion.Classification).First(&stackTag)
			if stackTag.ID == 0 {
				stackTag.Classification = dbQuestion.Classification
				stackTag.Hidden = 0
				stackTag.Unreaded = 1
			} else {
				stackTag.Unreaded += 1
			}
			tx.Save(&stackTag)

			// обновляем тег второго уровня
			stackTag = StackTag{}
			tx.Where("classification = ? and details = ?", dbQuestion.Classification, dbQuestion.Details).First(&stackTag)
			if stackTag.ID == 0 {
				stackTag.Classification = dbQuestion.Classification
				stackTag.Details = dbQuestion.Details
				stackTag.Hidden = 0
				stackTag.Unreaded = 1
			} else {
				stackTag.Unreaded += 1
			}
			tx.Save(&stackTag)

			tx.Commit()
		}
	}
}

func (ds *DataStore) UpdateStackQuestionRating(id uint32, score int) {

	var question StackQuestion
	tx := db.Begin()
	tx.Model(StackQuestion{}).Where("question_id = ?", id).First(&question)
	question.Score = score

	if question.QuestionId == 48166311 {
		fmt.Println("dsad")

	}

	if score < 0 {

		stackTag := StackTag{}
		tx.Where("classification = ? and details = ''", question.Classification).First(&stackTag)
		stackTag.Unreaded -= 1
		tx.Save(&stackTag)

		stackTag = StackTag{}
		tx.Where("classification = ? and details = ?", question.Classification, question.Details).First(&stackTag)
		stackTag.Unreaded -= 1
		tx.Save(&stackTag)

		question.Readed = 1
	}
	tx.Save(&question)
	tx.Commit()
}

func (ds *DataStore) SetStackQuestionAsReaded(question_id int) {

	tx := db.Begin()

	var question StackQuestion
	tx.Model(StackQuestion{}).Where("question_id = ?", question_id).First(&question)
	question.Readed = 1
	tx.Save(&question)

	stackTag := StackTag{}
	tx.Where("classification = ? and details = ''", question.Classification).First(&stackTag)
	stackTag.Unreaded -= 1
	tx.Save(&stackTag)

	stackTag = StackTag{}
	tx.Where("classification = ? and details = ?", question.Classification, question.Details).First(&stackTag)
	stackTag.Unreaded -= 1
	tx.Save(&stackTag)

	tx.Commit()
}

func (ds *DataStore) SetStackQuestionsAsReadedByClassification(classification string) {

	tx := db.Begin()
	tx.Model(StackQuestion{}).Where("classification = ?", classification).UpdateColumn("readed", 1)
	tx.Model(StackTag{}).Where("classification = ?", classification).UpdateColumn("unreaded", 0)
	tx.Commit()
}

func (ds *DataStore) SetStackQuestionsAsReadedByClassificationFromTime(classification string, t int64) {

	// эту операцию имеет смысл делать только при выбранных 2-х тегах
}