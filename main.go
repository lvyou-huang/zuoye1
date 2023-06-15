package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	var guess int
	for {
		fmt.Println("请输入猜测的数字")
		_, err := fmt.Scanf("%d\n", &guess)
		if err != nil {
			fmt.Println("非数字")
			continue
		}
		if num == guess {
			fmt.Println("猜对了")
			break
		} else if guess > num {
			fmt.Println("猜大了")
		} else if guess < num {
			fmt.Println("猜小了")
		}
	}
	fmt.Println("结束")
}
