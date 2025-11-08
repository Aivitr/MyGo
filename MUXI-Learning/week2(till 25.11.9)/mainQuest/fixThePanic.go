package main

import (
	"fmt"
	"sync"
	"time"
)

type message struct {
	Topic     string
	Partition int32
	Offset    int64
}

type FeedEventDM struct {
	Type    string
	UserID  int
	Title   string
	Content string
}

type MSG struct {
	ms        message
	feedEvent FeedEventDM
}

const ConsumeNum = 5

func main() {
	var consumeMSG []MSG
	var lastConsumeTime time.Time // 记录上次消费的时间
	msgs := make(chan MSG)
	var mu sync.Mutex // 保护 consumeMSG 的并发读写

	// 源源不断生产信息
	go func() {
		for i := 0; ; i++ {
			msgs <- MSG{
				ms: message{
					Topic:     "消费主题",
					Partition: 0,
					Offset:    0,
				},
				feedEvent: FeedEventDM{
					Type:    "grade",
					UserID:  i,
					Title:   "成绩提醒",
					Content: "您的成绩是xxx",
				},
			}
			time.Sleep(100 * time.Millisecond) // 模拟真实生产速率
		}
	}()

	// 不断接收消息进行消费
	for msg := range msgs {
		// 加锁保护append操作
		mu.Lock()
		consumeMSG = append(consumeMSG, msg)
		mu.Unlock()

		//  达到额定值：批量消费
		if len(consumeMSG) >= ConsumeNum {
			mu.Lock()
			// 拷贝需要消费的数据（避免引用原切片）
			batch := make([]MSG, ConsumeNum)
			copy(batch, consumeMSG[:ConsumeNum]) // 复制前 5 条
			// 截取原切片：移除已消费的数据
			consumeMSG = consumeMSG[ConsumeNum:]
			mu.Unlock()

			// 异步消费：传拷贝后的数据，避免闭包绑定问题
			go func(data []MSG) {
				fn(data)
			}(batch)

			lastConsumeTime = time.Now() // 更新消费时间
		} else if !lastConsumeTime.IsZero() && time.Since(lastConsumeTime) > 5*time.Minute {
			mu.Lock()
			// 按实际长度拷贝（避免越界）
			batchLen := len(consumeMSG)
			if batchLen > 0 {
				batch := make([]MSG, batchLen)
				copy(batch, consumeMSG)
				consumeMSG = consumeMSG[:0] // 清空原切片
				mu.Unlock()

				// 异步消费：传拷贝后的数据
				go func(data []MSG) {
					fn(data)
				}(batch)

				lastConsumeTime = time.Now()
			} else {
				mu.Unlock() // 无数据时也要解锁，避免死锁
			}
		}
	}
}

func fn(m []MSG) {
    fmt.Printf("本次消费了%d条消息\n", len(m))
}