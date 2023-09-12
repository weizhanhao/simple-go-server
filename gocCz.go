package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main7() {
	rand.Seed(time.Now().UnixNano())

	// 生成一个0到99之间的随机数
	randomNum := rand.Intn(100)

	// 模块1的执行概率为70%
	if randomNum < 70 {
		fmt.Println("进入模块1")
	} else {
		// 模块2的执行概率为30%
		fmt.Println("进入模块2")
	}

	// 模块3的执行概率为50%
	if randomNum < 50 {
		fmt.Println("进入模块3")
		fmt.Println("进入模块3")
		fmt.Println("进入模块3")
	} else {
		// 模块4的执行概率为50%
		fmt.Println("进入模块4")
		fmt.Println("进入模块4")
		fmt.Println("进入模块4")
		fmt.Println("进入模块4")
	}
	// 模块5的执行概率为80%
	if randomNum < 80 {
		fmt.Println("进入模块5")
		fmt.Println("进入模块5")
		fmt.Println("进入模块5")
		fmt.Println("进入模块5")
	} else {
		// 模块6的执行概率为20%
		fmt.Println("进入模块6")
		fmt.Println("进入模块6")
		fmt.Println("进入模块6")
		fmt.Println("进入模块6")
	}
}
