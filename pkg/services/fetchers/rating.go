package fetchers

import (
	"github.com/demas/observer/pkg/common"
	"strings"
	"fmt"
	"io/ioutil"
	"github.com/demas/observer/pkg/models"
	"encoding/json"
	"net/http"
)

func processQuestions(ids []uint32, site string) {

	// преобразуем id-ки в строку
	ids_as_str := []string{}
	for _, id := range ids {
		ids_as_str = append(ids_as_str, common.Uint32ToString(id))
	}

	ids_joined := strings.Join(ids_as_str, ";")
	url := fmt.Sprintf(soRatingUrl, ids_joined, site, key())
	log.Debugf("fetching score: %s", url)

	// fetch data
	res, err := http.Get(url)
	if err != nil {
		log.Error(err.Error())
	}

	jsn, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Error(err.Error())
	}

	// decode
	var p models.SOResponse

	err = json.Unmarshal(jsn, &p)
	if err != nil {
		log.Error(err.Error())
	} else {

		// получили вопросы и можем обновить их рейтинг
		for _, q := range p.Items {
			ds.UpdateStackQuestionRating(q.Question_id, q.Score)
		}
	}
}

func fetchRating() {

	// собираем нужные нам id-ки вопросов в разрезе сайта
	siteToQuestions := map[string][]uint32{}

	for _, question := range ds.GetStackQuestionsForRating() {

		ids, exists := siteToQuestions[question.Site]
		if exists {
			ids = append(ids, question.QuestionId)
			siteToQuestions[question.Site] = ids
		} else {
			siteToQuestions[question.Site] = []uint32 { question.QuestionId }
		}
	}

	// разбиваем запросы на партии и передаем на обработку
	for site, ids := range siteToQuestions {

		chunkSize := 49
		var chunk []uint32
		for len(ids) >= chunkSize {
			chunk, ids = ids[:chunkSize], ids[chunkSize:]
			processQuestions(chunk, site)
		}

		if len(ids) > 0 {
			processQuestions(ids[:len(ids)], site)
		}
	}
}
