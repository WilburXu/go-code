package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/go-redis/redis/v7"
)

const Rfc3399Mills = "2006-01-02T15:04:05.000-07:00"

var redisClient *redis.Client

type Task struct {
	ID          string `redis_idx:"primary"`
	Type        string `redis_idx:"i_type_phase_updated_time:0:type" json:"type"`
	UserID      string
	FileSize    int64
	UpdatedTime string `redis_idx:"i_type_phase_updated_time:2:updated_time" json:"updated_time"`
	Phase       int    `redis_idx:"i_type_phase_updated_time:1:phase" json:"phase"`
	Params      map[string]string
}

var taskReflectType = reflect.TypeOf(Task{})

// KEYS index1 index2...
// ARGV score1 member1 score2 member2...
func createIndexScript() string {
	return `
		local firstKey = KEYS[1]
		table.remove(KEYS, 1)
		local firstField = ARGV[1]
		table.remove(ARGV, 1)
		local firstVal = ARGV[1]
		table.remove(ARGV, 1)
		redis.call('HSet', firstKey, firstField, firstVal)

		for i, key in ipairs(KEYS) do
			local member = ARGV[i]
			redis.call('ZADD', key, 0, member)
		end

		return 'ok'
	`
}

// KEYS index1 index2...
// ARGV member1 member2...
func deleteIndexScript() string {
	return `
		local firstKey = KEYS[1]
		table.remove(KEYS, 1)
		local firstField = ARGV[1]
		table.remove(ARGV, 1)
		redis.call('HDel', firstKey, firstField)

		for i, key in ipairs(KEYS) do
			local member = ARGV[i]
			redis.call('ZREM', key, member)
		end

		return 'ok'
	`
}

func updateIndexScript() string {
	return `
			local firstKey = KEYS[1]
			table.remove(KEYS, 1)
			local firstField = ARGV[1]
			table.remove(ARGV, 1)
			local firstVal = ARGV[1]
			table.remove(ARGV, 1)

			local ZSetToDelete = KEYS[1]
			local newZSetToAdd = KEYS[2]

			-- 删除指定的 ZSET 中的成员
			for i = 1, #ARGV/2 do
				redis.call("ZREM", ZSetToDelete, ARGV[i])
			end

			-- 重新设置 hash
			redis.call('HSet', firstKey, firstField, firstVal)

			-- 新增另一个 ZSET
			for i = #ARGV/2+1, #ARGV do
				local member = ARGV[i]
				redis.call("ZADD", newZSetToAdd, 0, member)
			end

			return 'ok'
	`
}

// KEYS index1 index2...
// ARGV values
func searchScoreScript() string {
	return `
		local indexes = KEYS
		local values = ARGV
		local res = {}

		for i, key in ipairs(indexes) do
			local min = values[i*2-1]
			local max = values[i*2]

			if not max then
				max = min
			end

			if min and max then
				local members = redis.call('ZRANGEBYSCORE', key, min, max)
				if members then
					table.insert(res, members)
				end
			end
		end
		return res
	`
}

// KEYS index1 index2...
// ARGV values
func searchLexScript() string {
	return `
		local indexes = KEYS
		local values = ARGV
		local res = {}

		for i, key in ipairs(indexes) do
			local min = values[i*2-1]
			local max = values[i*2]

			if not max then
				max = min
			end

			if min and max then
				local members = redis.call('ZRANGEBYLEX', key, min, max)
				if members then
					table.insert(res, members)
				end
			end
		end
		return res
	`
}

func generateIndexKeys(prefix string) (keys []string) {
	for i := 0; i < taskReflectType.NumField(); i++ {
		f := taskReflectType.Field(i)
		if name := f.Tag.Get("redis_idx"); name != "" {
			keys = append(keys, fmt.Sprintf("%s:%s", prefix, name))
		}
	}
	return keys
}

type indexField struct {
	field    int
	priority string
	name     string
}

type SortIndexFields []indexField

func (a SortIndexFields) Len() int           { return len(a) }
func (a SortIndexFields) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortIndexFields) Less(i, j int) bool { return a[i].priority < a[j].priority }

type transform func(v interface{}) (interface{}, bool)

func transformUpdatedTime(v interface{}) (interface{}, bool) {
	assertV, ok := v.(string)
	if !ok {
		return "", false
	}

	t, err := ParseTime(assertV)
	if err != nil {
		return "", false
	}

	return t.UnixMilli(), true
}

// GetTaskBufferIndex 获取索引
func GetTaskBufferIndex(indexName string) (keys []string, err error) {
	indexKeys := make(map[string][]indexField)
	for i := 0; i < taskReflectType.NumField(); i++ {
		f := taskReflectType.Field(i)
		redisIdx := f.Tag.Get("redis_idx")
		if redisIdx == "" {
			continue
		}

		indexes := strings.Split(redisIdx, ",")
		for _, idx := range indexes {
			if idx == "primary" {
				continue
			}

			args := strings.Split(idx, ":")
			if len(args) != 3 {
				return nil, fmt.Errorf("invalid index: %s", idx)
			}

			if args[0] != indexName {
				continue
			}

			name := args[0]
			field := indexField{i, "0", args[2]}
			if len(args) > 1 {
				field = indexField{i, args[1], args[2]}
			}
			if _, ok := indexKeys[name]; !ok {
				indexKeys[name] = []indexField{field}
			} else {
				indexKeys[name] = append(indexKeys[name], field)
			}
		}
	}

	fields, ok := indexKeys[indexName]
	if !ok {
		return keys, nil
	}

	sort.Sort(SortIndexFields(fields))
	for _, f := range fields {
		keys = append(keys, f.name)
	}

	return keys, nil
}

func generateIndexValues(prefix string, task *Task, transFunc map[string]transform) (keys []string, values []interface{}, err error) {
	if task == nil {
		return nil, nil, nil
	}

	var (
		primaryValue interface{}
		indexKeys    = make(map[string][]indexField)
	)

	values = make([]interface{}, 0)
	v := reflect.ValueOf(*task)
	for i := 0; i < taskReflectType.NumField(); i++ {
		f := taskReflectType.Field(i)
		redisIdx := f.Tag.Get("redis_idx")
		if redisIdx == "" {
			continue
		}

		indexes := strings.Split(redisIdx, ",")
		for _, idx := range indexes {
			if idx == "primary" {
				primaryValue, err = toAdaptRedisType(v.Field(i).Interface())
				if err != nil {
					return nil, nil, err
				}
				continue
			}

			args := strings.Split(idx, ":")
			if len(args) != 3 {
				return nil, nil, fmt.Errorf("invalid index: %s", idx)
			}

			name := args[0]
			field := indexField{i, "0", args[2]}
			if len(args) > 1 {
				field = indexField{i, args[1], args[2]}
			}
			if _, ok := indexKeys[name]; !ok {
				indexKeys[name] = []indexField{field}
			} else {
				indexKeys[name] = append(indexKeys[name], field)
			}
		}
	}
	b := strings.Builder{}
	for name, fields := range indexKeys {
		keys = append(keys, fmt.Sprintf("%s:%s", prefix, name))
		b.Reset()
		sort.Sort(SortIndexFields(fields))
		//indexKeys[name] = fields
		//score := "0"
		for i, field := range fields {
			vv := v.Field(field.field).Interface()
			if i == 0 && isNumber(vv) && len(fields) == 1 {
				//score = fmt.Sprintf("%v", vv)
				continue
			}

			if trans, exist := transFunc[field.name]; exist {
				transVal, ok := trans(vv)
				if !ok {
					return nil, nil, fmt.Errorf("invalid index val: %s", v)
				}

				b.WriteString(fmt.Sprintf("%v:", transVal))
			} else {
				b.WriteString(fmt.Sprintf("%v:", vv))
			}
		}
		b.WriteString(fmt.Sprintf("%v", primaryValue))
		values = append(values, b.String())
		//values = append(values, score, b.String())
	}

	return keys, values, nil
}

var ErrUnsupportedDataType = fmt.Errorf("unsupported data type")

// 将interface转换成string或number
func toAdaptRedisType(v interface{}) (interface{}, error) {
	switch vv := v.(type) {
	case []byte:
		return string(vv), nil
	case string:
		return vv, nil
	case uint, int, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return vv, nil
	default:
		return 0, ErrUnsupportedDataType
	}
}

func isNumber(v interface{}) bool {
	switch v.(type) {
	case uint, int, uint8, uint16, uint32, uint64, int8, int16, int32, int64, float32, float64:
		return true
	default:
		return false
	}
}

func main() {
	redisClient = NewRedisClient(RedisConfig{
		Addr: "localhost:6379",
	})
	scriptHandler := map[string]*redis.Script{
		"create":      redis.NewScript(createIndexScript()),
		"delete":      redis.NewScript(deleteIndexScript()),
		"update":      redis.NewScript(updateIndexScript()),
		"searchScore": redis.NewScript(searchScoreScript()),
		"searchLex":   redis.NewScript(searchLexScript()),
	}

	tasks := []*Task{
		{
			ID:          "task1",
			Type:        "offline",
			UserID:      "1000021430",
			FileSize:    100,
			Phase:       4,
			UpdatedTime: "2020-10-22T19:17:40.645+08:00",
			Params:      map[string]string{"a": "1", "b": "2"},
		},
		{
			ID:          "task2",
			Type:        "offline",
			UserID:      "1000021430",
			FileSize:    200,
			Phase:       3,
			UpdatedTime: "2020-10-22T19:17:41.645+08:00",
			Params:      map[string]string{"a": "1", "b": "2"},
		},
		{
			ID:          "task3",
			Type:        "offline",
			UserID:      "1000021430",
			FileSize:    300,
			Phase:       2,
			UpdatedTime: "2020-10-22T19:17:42.645+08:00",
			Params:      map[string]string{"a": "1", "b": "2"},
		},
		{
			ID:          "task4",
			Type:        "offline",
			UserID:      "1000021430",
			FileSize:    400,
			Phase:       1,
			UpdatedTime: "2020-10-22T19:17:43.645+08:00",
			Params:      map[string]string{"a": "1", "b": "2"},
		},
	}

	for _, t := range tasks {
		keys, values, err := generateIndexValues(fmt.Sprintf("task_index:%s", t.UserID), t, map[string]transform{
			"updated_time": transformUpdatedTime,
		})
		if err != nil {
			log.Fatalf("generateIndexValues error: %v", err)
		}

		keys = append([]string{fmt.Sprintf("task_%s", t.UserID)}, keys...)
		values = append([]interface{}{t.ID, string(toJSON(t))}, values...)

		res, err := scriptHandler["create"].Run(redisClient, keys, values...).Result()
		log.Println(res, err)
	}

	// 查询
	//keys, _ := redisClient.Keys("*").Result()
	//for _, key := range keys {
	//	values, _ := redisClient.ZRangeWithScores(key, 0, -1).Result()
	//	log.Println(key)
	//	for _, v := range values {
	//		log.Println("\tscore:", v.Score, "member:", v.Member)
	//	}
	//}

	//resLex, err := scriptHandler["searchScore"].Run(redisClient, keys, 0, "+inf", 0, "+inf", 200, 300).Result()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//for _, v := range resLex.([]interface{}) {
	//	log.Println(v)
	//}

	// 搜索
	resLex, err := scriptHandler["searchLex"].Run(redisClient, []string{"task_index:1000021430:i_type_phase_updated_time"}, "[offline:1:", "[offline:1:\xff").Result()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range resLex.([]interface{}) {
		log.Println(v)
	}

	// 更新
	for _, t := range tasks {
		// 1. 获取旧索引
		deleteKeys, deleteValues, err := generateIndexValues(fmt.Sprintf("task_index:%s", t.UserID), t, map[string]transform{
			"updated_time": transformUpdatedTime,
		})
		if err != nil {
			log.Fatalf("generateIndexValues error: %v", err)
		}

		fmt.Println(deleteKeys, deleteValues)

		t.Phase = 9
		addKeys, addValues, err := generateIndexValues(fmt.Sprintf("task_index:%s", t.UserID), t, map[string]transform{
			"updated_time": transformUpdatedTime,
		})
		if err != nil {
			log.Fatalf("generateIndexValues error: %v", err)
		}

		keys := append(deleteKeys[:], addKeys[:]...)
		args := append(deleteValues[:], addValues[:]...)

		keys = append([]string{fmt.Sprintf("task_%s", t.UserID)}, keys...)
		args = append([]interface{}{t.ID, string(toJSON(t))}, args...)

		res, err := scriptHandler["update"].Run(redisClient, keys, args...).Result()
		log.Println(res, err)
		fmt.Println(addKeys, addValues)

	}

	// 删除
	for _, t := range tasks {
		keys, values, err := generateIndexValues(fmt.Sprintf("task_index:%s", t.UserID), t, map[string]transform{
			"updated_time": transformUpdatedTime,
		})
		if err != nil {
			log.Fatalf("generateIndexValues error: %v", err)
		}

		keys = append([]string{fmt.Sprintf("task_%s", t.UserID)}, keys...)
		values = append([]interface{}{t.ID}, values...)

		res, err := scriptHandler["delete"].Run(redisClient, keys, values...).Result()
		log.Println(res, err)
	}
}

// ParseTime 解析Rfc3399Mills格式时间
func ParseTime(s string) (time.Time, error) {
	t, err := time.Parse(Rfc3399Mills, s)
	return t, err
}

func toJSON(v interface{}) []byte {
	bs, err := json.Marshal(v)
	if err != nil {
		//log.Errorf("json marshal type:%T, value:%+v,  err: %v", v, v, err)
		return nil
	}
	return bs
}
