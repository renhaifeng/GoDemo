package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 0)

	go func() {
		who := "送完外卖的我："
		time.Sleep(time.Second)
		fmt.Println(who, "writing……1 hour")
		ch <- "文章"
		fmt.Println(who, "分享成功")
	}()

	go func() {
		who := "你："
		fmt.Println(who, "摸鱼……")
		read := <-ch
		fmt.Println(who, "获得阅读", read, ", 并点赞")
	}()

	time.Sleep(time.Second * 5)
}
