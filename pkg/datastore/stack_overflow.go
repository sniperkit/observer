package datastore

import (
	"github.com/demas/observer/pkg/models"
	"strings"
)

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
