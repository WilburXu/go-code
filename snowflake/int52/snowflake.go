package snowflake

import (
	"sync"
	"time"
)

var instance *Worker
var once sync.Once

// GetInstance 单例模式
func GetInstance() *Worker {
	once.Do(func() {
		instance = NewWorker(1)
	})
	return instance
}

/*
 * 算法解释
 * SnowFlake的结构如下(每部分用-分开):<br>
 * 0 - 0000000000 0000000000 0000000000 0000000000 0 - 00000 - 00000 - 000000000000 <br>
 * 1位标识，由于long基本类型在Java中是带符号的，最高位是符号位，正数是0，负数是1，所以id一般是正数，最高位是0<br>
 * 41位时间截(毫秒级)，注意，41位时间截不是存储当前时间的时间截，而是存储时间截的差值（当前时间截 - 开始时间截)
 * 得到的值），这里的的开始时间截，一般是我们的id生成器开始使用的时间，由我们程序来指定的（如下的epoch属性）。
 * 41位的时间截，可以使用69年，年T = (1L << 41) / (1000L * 60 * 60 * 24 * 365) = 69<br>
 * 10位的数据机器位，可以部署在1024个节点，包括5位datacenterId和5位workerId<br>
 * 12位序列，毫秒内的计数，12位的计数顺序号支持每个节点每毫秒(同一机器，同一时间截)产生4096个ID序号<br>
 * 加起来刚好64位，为一个Long型。<br>
 * SnowFlake的优点是，整体上按照时间自增排序，并且整个分布式系统内不会产生ID碰撞(由数据中心ID和机器ID作区分)，并且效率较高，经测试，SnowFlake每秒能够产生26万ID左右。
 */
const (
	workerBits  uint8 = 4
	numberBits  uint8 = 16
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	epoch       int64 = 1152629100 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID	2142-08-18 05:13:15 才会不够用
)

// Worker worker
type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerID  int64
	number    int64
}

// NewWorker new
func NewWorker(workerID int64) *Worker {
	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		workerID:  workerID,
		number:    0,
	}
}

// NextID 获取id
func (w *Worker) NextID() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().Unix()
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().Unix()
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}

	ID := (now-epoch)<<timeShift | (w.workerID << workerShift) | (w.number)
	return ID
}
