package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		who := "外卖小哥我："
		food := "鸡腿"
		fmt.Println(who, "送餐中……2s")
		time.Sleep(time.Second * 2)
		fmt.Println(who, "已送餐到门口，餐挂门把手了（缓冲区）")
		ch <- food
		fmt.Println(who, "订单已送达，开始送其他单")
	}()

	go func() {
		who := "你："
		fmt.Println(who, "等待外卖……")
		time.Sleep(time.Second * 3)
		fmt.Println(who, "磨磨唧唧开门中……3s")
		food := <-ch
		fmt.Println(who, "拿到", food, "开吃！")
	}()

	time.Sleep(time.Second * 5)
}
