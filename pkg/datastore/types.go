package datastore

import (
	"github.com/jinzhu/gorm"
)

type Settings struct {
	gorm.Model

	Key  string `gorm:"size:40"`
	Value string `gorm:"size:40"`
}

type StackTag struct {
	gorm.Model

	Classification string `gorm:"size:40"`
	Unreaded       int
	Hidden         int
}

type StackQuestion struct {
	gorm.Model

	Title            string `gorm:"size:500"`
	Link             string `gorm:"size:500"`
	QuestionId       uint32 `gorm:"column:questionid"`
	Tags             string `gorm:"size:300"`
	Score            int
	AnswerCount      int
	ViewCount        int
	UserId           int
	UserReputation   int
	UserDisplayName  string `gorm:"size:200"`
	UserProfileImage string `gorm:"size:100"`
	Classification   string `gorm:"size:40"`
	Details          string `gorm:"size:40"`
	Readed 			 int
	CreationDate     uint32
	Favorite         int
	Classified       int
	Site             string `gorm:"size:100"`
}