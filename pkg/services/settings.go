package services

import (
	"time"
	"strconv"
	"github.com/demas/observer/pkg/datastore"
	"github.com/demas/observer/pkg/services/service_locator"
	"github.com/demas/observer/pkg/common"
)

var ds datastore.IDataStore

func init() {
	ds = service_locator.GetDataStore()
}

const lastStackSyncTimeKey = "lastStackSyncTime"
const stackSyncTimeOffset = 2000

func GetLastStackSyncTime() (int64) {

	defaultLastStackSyncTime := time.Now().Unix() - stackSyncTimeOffset

	lastSyncTimeStr := ds.GetSettings(lastStackSyncTimeKey)
	if lastSyncTimeStr == "" {
		return defaultLastStackSyncTime
	}

	result, err := strconv.ParseInt(lastSyncTimeStr, 10, 64)
	if err != nil {
		return defaultLastStackSyncTime
	}

	return result
}

func SetLastStackSyncTime(value int64) {
	ds.SetSettings(lastStackSyncTimeKey, common.Int64ToString(value))
}
