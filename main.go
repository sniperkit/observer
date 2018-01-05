package main

import (
	"github.com/demas/observer/pkg/services/fetchers"
	"github.com/demas/observer/pkg/rest_api"
)

func main() {

	// запускаем сборку новых вопросов в отдельной горутине
	(&fetchers.StackFetcher{}).Fetch()

	// поднимаем веб-службы
	rest_api.Serve()
}
