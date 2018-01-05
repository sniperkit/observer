package main

import (
	"github.com/demas/observer/pkg/services/fetchers"
)

func main() {
	(&fetchers.StackFetcher{}).Fetch()
}
