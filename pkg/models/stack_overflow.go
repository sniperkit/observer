package models

type SOUser struct {
	Reputation    int
	User_id       int
	User_type     string
	Accept_rate   int
	Profile_image string
	Display_name  string
	Link          string
}

type SOQuestion struct {

	Tags               []string
	Owner              SOUser
	Is_answered        bool
	View_count         int
	Answer_count       int
	Score              int
	Last_activity_date uint32
	Creation_date      uint32
	Question_id        uint32
	Link               string
	Title              string
	Classification	   string
	Details            string
}

type SOResponse struct {
	Items           []SOQuestion
	Has_more        bool
	Quota_max       int
	Quota_remaining int
}

type StackTag struct {

	Id 			   int `json:"Id"`
	Classification string `json:"Classification"`
	Unreaded       int `json:"Unreaded"`
	Hidden         int `json:"Hidden"`
}


type StackQuestion struct {

	Id               int    `json:"id"`
	Title            string `json:"title"`
	Link             string `json:"link"`
	QuestionId       uint32    `json:"questionid"`
	Tags             string `json:"tags"`
	Score            int    `json:"score"`
	AnswerCount      int    `json:"answercount"`
	ViewCount        int    `json:"viewcount"`
	UserId           int	`json:"userid"`
	UserReputation   int 	`json:"userreputation"`
	UserDisplayName  string `json:"userdisplayname"`
	UserProfileImage string `json:"userprofileimage"`
	Classification   string `json:"classification"`
	Details          string `json:"details"`
	Readed 			 int    `json:"readed"`
	CreationDate     uint32  `json:"creationdate"`
	Favorite         int    `json:"favorite"`
	Classified       int    `json:"classified"`
	Site             string `json:"site"`
}
