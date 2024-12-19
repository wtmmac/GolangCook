package query

const subjectQueryDSL = `{
	"size": 0,
	"query": {
		"bool": {
			"filter": [
				{
					"range": {
						"@timestamp": {
						"format": "strict_date_optional_time",
						"gte": "%s",
						"lte": "%s"
						}
					}
				},
				{
					"match_phrase": {
						"uri": "/subject"
					}
				},
				{
					"match_phrase": {
						"status": "404"
					}
				}
			]
		}
	},
	"aggs": {
		"uid": {
			"cardinality": {
				"field": "uid.keyword"
			}
		},
		"total": {
			"value_count": {
				"field": "_index"
			}
		},
		"agg_uri": {
			"terms": {
				"field": "uri.keyword",
				"order": {
					"_count": "desc"
				},
				"size": 10
			},
			"aggs": {
				"1": {
					"cardinality": {
						"field": "uid.keyword"
					}
				},
				"2": {
					"cardinality": {
						"field": "clientIp.keyword"
					}
				}
			}
		}
	}
}`

func init() {
	const indexPrefix = "logstash-api.my.tv.sohu.com-"

	QueryManager.BuildTimeRangeQuery("subjectQueryDSL", indexPrefix, subjectQueryDSL)
}
