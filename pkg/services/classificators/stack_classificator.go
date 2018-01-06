package classificators

import (
	"github.com/demas/observer/pkg/models"
	"github.com/demas/observer/pkg/services"
	"os"
	"io/ioutil"
	"encoding/json"
	"strings"
)

type FirstLevelRule struct {
	Site string `json:"site"`
	Include string `json:"include"`
	Result string `json:"result"`
}

type SecondLevelRule struct {
	Site string  `json:"site"`
	First string  `json:"first"`
	Include string  `json:"include"`
	Result string `json:"result"`
}

type StackClassificatorRules struct {
	StopTags []string `json:"stop_tags"`
	FirstLevelRules []FirstLevelRule `json:"first_level_rules"`
	SecondLevelRules []SecondLevelRule `json:"second_level_rules"`
}

const rulesFileName = "./classificator.json"

var rules StackClassificatorRules
var log = services.GetLogger("classificator")

func init() {

	jsonFile, err := os.Open(rulesFileName)
	if err != nil {
		log.Error(err.Error())

		jsonFile, err = os.Open("/go/bin/classificator.json")
		if err != nil {
			log.Error(err.Error())
			panic(err)
		}
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &rules)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}
}

// Check tag existence in the question
func containTag(q models.SOQuestion, tag string) (bool) {

	for _, question_tag := range q.Tags {
		if question_tag == tag {
			return true
		}
	}
	return false
}

// check tag existence in the questions by slice of tags
func containAnyOfTags(q models.SOQuestion, tags []string) (bool) {

	for _, tag := range tags {
		if containTag(q, tag) {
			return true
		}
	}
	return false
}

// Check stop tag existance in the question
func containStopTag(q models.SOQuestion) (bool, string) {

	for _, stop_tag := range rules.StopTags {
		if containTag(q, stop_tag) {
			return true, stop_tag
		}
	}
	return false, ""
}

func firstLevelClassification(q models.SOQuestion, site string) (string) {

	for _, flr := range rules.FirstLevelRules {
		if site == flr.Site {
			if flr.Include == "*" {
				return flr.Result
			} else {
				ruleTags := strings.Split(flr.Include, ",")
				if containAnyOfTags(q, ruleTags) {
					return flr.Result
				}
			}
		}
	}
	return "general"
}

func secondLevelClassification(q models.SOQuestion, site string, first string) (string) {

	for _, slr := range rules.SecondLevelRules {

		if (slr.Site == site)  && (slr.First == first) && containTag(q, slr.Include) {
			return slr.Result
		}
	}
	return "general"
}


func ClassifyStackQuestions(questions []models.SOQuestion, site string) {

	for i, stackQuestion := range questions {

		contain_stop_tag, stop_tag := containStopTag(stackQuestion)
		if contain_stop_tag {
			log.Debugf("Classificator: stop tag [%s] ", stop_tag)
			questions[i].Classification = "remove"
			questions[i].Details = "remove"
			continue;
		}

		flr := firstLevelClassification(stackQuestion, site)
		questions[i].Classification = flr
		log.Debugf("Classificator I: %s", flr)

		slr := secondLevelClassification(stackQuestion, site, flr)
		questions[i].Details = slr
		log.Debugf("Classificator II: %s", slr)
	}
}
