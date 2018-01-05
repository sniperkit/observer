package datastore

func (ds *DataStore) GetSettings(key string) string {

	settings := Settings{}
	db.Where("key = ?", key).First(&settings)
	return settings.Value
}

func (ds *DataStore) SetSettings(key string, value string) {

	settings := Settings{}
	db.Where("key = ?", key).First(&settings)
	settings.Key = key
	settings.Value = value
	db.Save(&settings)
}