package service_locator

import (
	"github.com/demas/observer/pkg/datastore"
)

func GetDataStore() datastore.IDataStore {
	return &datastore.DataStore{}
}
