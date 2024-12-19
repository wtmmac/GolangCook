package query

import (
	"fmt"
	"sync"
	"time"
)

// 全局查询管理器实例
var QueryManager = NewQueryManager()

// queryManager 管理查询和索引的生成
type queryManager struct {
	mu      sync.RWMutex
	queries map[string]string
	indexes map[string]string // 存储查询名称到索引的映射
}

// NewQueryManager 创建一个新的查询管理器实例
func NewQueryManager() *queryManager {
	return &queryManager{
		queries: make(map[string]string),
		indexes: make(map[string]string),
	}
}

// Register 注册新的查询DSL和对应的索引
func (qm *queryManager) Register(name, query, index string) {
	qm.mu.Lock()
	defer qm.mu.Unlock()
	qm.queries[name] = query
	qm.indexes[name] = index
}

// Get 获取指定名称的查询DSL和索引
func (qm *queryManager) Get(name string) (string, string, bool) {
	qm.mu.RLock()
	defer qm.mu.RUnlock()
	query, okQ := qm.queries[name]
	index, okI := qm.indexes[name]
	return query, index, okQ && okI
}

// BuildTimeRangeQuery 构建时间范围查询并注册到管理器
func (qm *queryManager) BuildTimeRangeQuery(name, indexPrefix, query string) {
	startTime, endTime, err := rangeTime()
	if err != nil {
		panic(err)
	}
	queryDsl := fmt.Sprintf(query, startTime, endTime)
	index := GenerateURL(indexPrefix)
	qm.Register(name, queryDsl, index)
}

// ExecuteAllQueries 遍历所有查询并执行它们
func (qm *queryManager) ExecuteAllQueries(execFunc func(name, query, index string) error) error {
	qm.mu.RLock()
	defer qm.mu.RUnlock()

	for name, query := range qm.queries {
		index, ok := qm.indexes[name]
		if !ok {
			return fmt.Errorf("index not found for query %s", name)
		}
		if err := execFunc(name, query, index); err != nil {
			return fmt.Errorf("failed to execute query %s on index %s: %v", name, index, err)
		}
	}
	return nil
}

func GenerateURL(indexPrefix string) string {
	const baseURL = "http://10.18.19.11:9200/"
	const searchPath = "/_search"

	// 获取当前时间和倒推1小时的时间
	now := time.Now()
	hourAgo := now.Add(-24 * time.Hour)

	// 格式化日期
	currentDate := now.Format("2006.01.02")
	hourAgoDate := hourAgo.Format("2006.01.02")

	// 生成索引部分
	indexPart := indexPrefix + currentDate
	if currentDate != hourAgoDate {
		indexPart += "," + indexPrefix + hourAgoDate
	}

	// 生成完整的 URL
	return baseURL + indexPart + searchPath
}

func rangeTime() (startTime, endTime string, err error) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "", "", err
	}

	now := time.Now()
	beforeOneHour := now.Add(time.Duration(-24 * time.Hour))

	nowInShanghai := now.In(location)
	endTime = nowInShanghai.Format("2006-01-02T15:04:05.000Z07:00")

	beforeOneHourInShanghai := beforeOneHour.In(location)
	startTime = beforeOneHourInShanghai.Format("2006-01-02T15:04:05.000Z07:00")

	return startTime, endTime, nil
}
