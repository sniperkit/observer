package routers

import (
	"github.com/gorilla/mux"
	"github.com/demas/observer/pkg/rest_api/controllers"
)

func InitRoutes() *mux.Router {

	router := mux.NewRouter().PathPrefix("/api").Subrouter().StrictSlash(true)

	// stack
	router.HandleFunc("/stack/tags", controllers.WrapHandler(controllers.GetStackTags))
	router.HandleFunc("/stack/secondtags/{classification}", controllers.WrapHandler(controllers.GetSecondTagsByClassification))
	router.HandleFunc("/stack/questions/{classification}", controllers.WrapHandler(controllers.GetStackQuestionsByClassification))
	router.HandleFunc("/stack/questions/{classification}/{details}", controllers.WrapHandler(controllers.GetStackQuestionsByClassificationAndDetails))
	router.HandleFunc("/stack/question-as-read", controllers.PostWrapHandler(controllers.SetStackQuestionAsReaded)).Methods("POST")
	router.HandleFunc("/stack/tags/as-read", controllers.PostWrapHandler(controllers.SetStackQuestionsAsReaded)).Methods("POST")
	router.HandleFunc("/stack/tags/from-time/as-read", controllers.PostWrapHandler(controllers.SetStackQuestionsAsReadedFromTime)).Methods("POST")

	// tags
	router.HandleFunc("/tags", controllers.WrapHandler(controllers.GetTags))
	router.HandleFunc("/tags/items/{tagId}", controllers.WrapHandler(controllers.GetTaggedItemsByTagId)).Methods("GET")
	router.HandleFunc("/tags/add-item", controllers.PostWrapHandler(controllers.InsertTaggedItem)).Methods("POST")
	router.HandleFunc("/tags/items/{id}", controllers.PostWrapHandler(controllers.DeleteTaggedItem)).Methods("DELETE")

	return router
}
