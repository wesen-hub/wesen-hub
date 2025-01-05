package dispatcher

import (
	"awesomeProject/src/counter"
	"awesomeProject/src/data/redis"
	"fmt"
	"time"
)

func Run() {
	for {
		if counter.Get() < 20 {
			dispatch()
		} else {
			time.Sleep(10 * time.Second)
		}
	}
}

var key = "files"

func dispatch() {
	//分批获取待处理文件列表
	// 列表的键
	listKey := "my_list"

	// 批处理大小
	var batchSize int64 = 100

	// 获取列表长度
	listLength := redis.LLen(listKey)
	if listLength < 0 {
		//todo
		return
	}
	//获取处理中列表
	listDoing := getHandling()
	var start int64
	// 分批读取列表
	for start = 0; start < listLength; start += batchSize {
		end := start + batchSize - 1
		if end >= int64(listLength) {
			end = int64(listLength) - 1
		}

		// 使用LRANGE命令读取当前批次的元素
		batch, err := redis.LRange(listKey, start, end)
		if err != nil {
			panic(err)
		}

		// 处理当前批次的元素
		for _, element := range batch {
			fmt.Println(element)
			if existInList(listDoing, element) {
				break
			}
			//获取锁
			//获取文章内容
			//计数器加1
			//提交任务给携程池
		}

	}
}

func getHandling() []string {
	return nil
}

func existInList(list []string, val string) bool {
	return false
}
