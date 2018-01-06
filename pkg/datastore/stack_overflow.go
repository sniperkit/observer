package datastore

import (
	"github.com/demas/observer/pkg/models"
	"strings"
)

func (ds *DataStore) GetStackTags() []models.StackTag {

	result := []models.StackTag{}
	db.Where("unreaded > 0 and hidden = 0").Find(&result)
	return result
}

func (ds *DataStore) GetSecondTagByClassification(classification string) interface{} {

	type Result struct {
		Details string `json:"details"`
		Count   int    `json:"count"`
	}

	result := []Result{}
	db.Table("stack_questions").Select("details, count(id) as count").
		Where("classification = ? and readed= 0", classification).Group("details").Scan(&result)
	return result
}

func (ds *DataStore) GetStackQuestionsByClassification(classification string) []models.StackQuestion {

	result := []models.StackQuestion{}
	db.Model(StackQuestion{}).
		Where("classification = ? and readed = 0", classification).Order("score desc").Find(&result)

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

			db.Save(&dbQuestion)

			stackTag := StackTag{}
			db.Where("classification = ?", dbQuestion.Classification).First(&stackTag)

			if stackTag.ID == 0 {
				stackTag.Classification = dbQuestion.Classification
				stackTag.Hidden = 0
				stackTag.Unreaded = 1
			} else {
				stackTag.Unreaded += 1
			}

			db.Save(&stackTag)
		}
	}
}

func (ds *DataStore) SetStackQuestionAsReaded(question_id int) {

	var question StackQuestion
	db.Model(StackQuestion{}).Where("question_id = ?", question_id).First(&question)
	question.Readed = 1
	db.Save(&question)

	stackTag := StackTag{}
	db.Where("classification = ?", question.Classification).First(&stackTag)
	stackTag.Unreaded -= 1
	db.Save(&stackTag)
}

func (ds *DataStore) SetStackQuestionsAsReadedByClassification(classification string) {

	db.Model(StackQuestion{}).Where("classification = ?", classification).UpdateColumn("readed", 1)
	stackTag := StackTag{}
	db.Where("classification = ?", classification).First(&stackTag)
	stackTag.Unreaded = 0
	db.Save(&stackTag)
}

func (ds *DataStore) SetStackQuestionsAsReadedByClassificationFromTime(classification string, t int64) {

	var count int

	db.Model(&StackQuestion{}).Where("classification = ? and creationdate < ?", classification, t).Count(&count)
	db.Model(StackQuestion{}).Where("classification = ? and creationdate < ?", classification, t).
		UpdateColumn("readed", 1)

	stackTag := StackTag{}
	db.Where("classification = ?", classification).First(&stackTag)
	stackTag.Unreaded -= count
	db.Save(&stackTag)
}