package main

import (
	"testing"

	"github.com/k0kubun/pp"
	elastigo "github.com/mattbaird/elastigo/lib"

	_ "bitbucket.org/ikeikeikeike/antenna/conf/inits"
	_ "bitbucket.org/ikeikeikeike/antenna/routers"

	"bitbucket.org/ikeikeikeike/antenna/lib/es"
	"bitbucket.org/ikeikeikeike/antenna/models"
)

func Analyzer() string {
	return `
	analysis: {
		analyzer: {
				searchkick_autocomplete_index: {
					type: "custom",
					tokenizer: "searchkick_autocomplete_ngram",
					filter: ["lowercase", "asciifolding"]
				},
				searchkick_autocomplete_search: {
					type: "custom",
					tokenizer: "keyword",
					filter: ["lowercase", "asciifolding"]
				}
			},
			filter: {
              searchkick_edge_ngram: {
                type: "edgeNGram",
                min_gram: 1,
                max_gram: 50
              }
			},
			tokenizer: {
				searchkick_autocomplete_ngram: {
				type: "edgeNGram",
				min_gram: 1,
				max_gram: 50
			}
		}
	}`
}

func TestScore(t *testing.T) {
	e := models.Entry{Id: 12}
	e.Read()
	e.LoadRelated()

	if r, err := es.Index(e.SearchData()); err != nil {
		t.Error(err)
	} else {
		pp.Println(r)
	}

	// Search Using Raw json String
	searchJson := `{
	    "query" : {
	        "term" : { "title" : "コンプティーク" }
	    }
	}`
	c := elastigo.NewConn()
	// out, err := c.Search("antenna", "EntryIndex", 12, searchJson)
	out, err := c.Search("antenna", "EntryIndex", nil, searchJson)
	pp.Println(err, out.Hits.Hits)
}
