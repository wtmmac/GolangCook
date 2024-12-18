package query

import "fmt"

func init() {
	queryDsl := `{
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
						  "uri": "/comment"
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
	fmt.Printf(queryDsl, "2019-07-01", "2019-07-02")
}
