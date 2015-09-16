package main

import (
	"testing"
	"time"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"

	"bitbucket.org/ikeikeikeike/antenna/models"
)

func TestSQLForDelete(t *testing.T) {
	var summs []*models.Summary

	ago := time.Now().AddDate(0, -1, 0) // 1 month ago

	sqs := models.Summaries()
	sqs.Filter("created__lte", ago).All(&summs)
}
