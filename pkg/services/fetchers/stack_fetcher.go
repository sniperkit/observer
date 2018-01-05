package fetchers

import (
	"fmt"
	"strconv"
	"time"

	"os"
	"github.com/demas/observer/pkg/common"
	"github.com/demas/observer/pkg/services"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/demas/observer/pkg/services/classificators"
	"github.com/demas/observer/pkg/models"
	"github.com/demas/observer/pkg/datastore"
	"github.com/demas/observer/pkg/services/service_locator"
)

type IStackFetcher interface {
	Fetch()
}

type StackFetcher struct {}

var ds datastore.IDataStore

func init() {
	ds = service_locator.GetDataStore()
}

// TODO: где то в настройках должно быть
var soSites = [8]string{"stackoverflow", "security", "codereview", "softwareengineering", "ru.stackoverflow", "superuser",
	"unix", "serverfault"}

const maxSOPages = 50
const soBaseUrl = "https://api.stackexchange.com/2.2/questions?page=%d&pagesize=100&fromdate=%d&order=asc&sort=creation&site=%s%s"
const soKeyEnvVariable = "SOKEY"

var log = services.GetLogger("main")




func heartBeat() {
	fmt.Println("working...")
}

func key() string {

	if os.Getenv(soKeyEnvVariable) != "" {
		return fmt.Sprintf("&key=%s", os.Getenv(soKeyEnvVariable))
	}

	return ""
}

func getNewMassages(fromTime int64, site string) []models.SOQuestion {

	var result []models.SOQuestion
	page := 1
	has_more := true

	for has_more && page <= maxSOPages {

		url := fmt.Sprintf(soBaseUrl, page, fromTime, site, key())
		log.Debugf("fetching site: %s", url)

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
			result = append(result, p.Items...)
		}

		has_more = p.Has_more
		page = page + 1
	}

	return result
}

func fetchQuestions() {

	lastSyncTime := services.GetLastStackSyncTime()
	currentTime := time.Now().Unix()

	allQuestions := map[string][]models.SOQuestion{}

	for _, soSite := range soSites {
		soQuestions := getNewMassages(lastSyncTime, soSite)
		classificators.ClassifyStackQuestions(soQuestions, soSite)
		allQuestions[soSite] = soQuestions
	}

	ds.InsertStackOverflowQuestions(allQuestions)
	services.SetLastStackSyncTime(currentTime)
}

func (f *StackFetcher) Fetch() {
	fmt.Println("Fetch stack questions ...")

	// TODO: в случае недоступности Postgre ждать пока он поднимется
	// wait 1 minute to start postgresql
	//timer := time.NewTimer(time.Second * 1)
	//<-timer.C

	syncInterval, err := strconv.Atoi(os.Getenv("SYNCINTERVAL"))
	if err != nil {
		syncInterval = 10
	}

	go common.DoEvery(time.Minute * time.Duration(syncInterval), fetchQuestions)
	common.DoEvery(time.Second * 10, heartBeat)
}