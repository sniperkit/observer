package classificators

import (
	"github.com/demas/observer/pkg/services"
	"github.com/demas/observer/pkg/models"
)


var log = services.GetLogger("stack_classificator")

// Check tag existence in the question
func containTag(q models.SOQuestion, tag string) (bool) {

	for _, question_tag := range q.Tags {
		if question_tag == tag {
			return true
		}
	}

	return false
}


// Check stop tag existance in the question
func containStopTag(q models.SOQuestion) (bool, string) {

	for _, stop_tag := range stop_tags {
		if containTag(q, stop_tag) {
			return true, stop_tag
		}
	}

	return false, ""
}

func firstLevelClassification(q models.SOQuestion, site string) (string) {

	for _, flr := range firstLevelRules {

		if site == flr.Site {

			if flr.Include == "*" {
				return flr.Result
			} else if containTag(q, flr.Include) {
				return flr.Result
			}
		}
	}

	return ""
}

func secondLevelClassification(q models.SOQuestion, site string, first string) (string) {

	for _, slr := range secondLevelRules {

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
