package snowflake

import (
	"log"
	"strconv"
	"testing"
)

/*
 * 将十进制数字转化为二进制字符串
 */
func convertToBin(num int64) string {
	s := ""
	if num == 0 {
		return "0"
	}
	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// 将数字强制性转化为字符串
		s = strconv.FormatInt(lsb, 10) + s
	}
	return s
}

func TestSnowFlake(t *testing.T) {
	ch := make(chan int64, 65535)
	count := 65535
	// 并发 goroutine ID生成
	for i := 0; i < count; i++ {
		go func() {
			id := GetInstance().NextID()
			ch <- id
		}()
	}
	defer close(ch)
	m := make(map[int64]int)
	v := 0
	for i := 0; i < count; i++ {
		id := <-ch
		// map中存在为id的key,说明生成的 ID有重复
				_, ok := m[id]
		log.Println(id)
		if ok {
			v++
			t.Error("ID is not unique!")
		}
		// id作为key存入map
		m[id] = i
		//fmt.Println(id)
		//fmt.Println(convertToBin(id))
	}
	//t.Error(v)

}
