package code

import (
	"fmt"
	"sync"
)

func main() {
	// 使用WaitGroup等待所有goroutine完成
	var wg sync.WaitGroup
	wg.Add(3)
	// 创建三个channel，分别用于控制A、B、C的打印
	printACh := make(chan struct{})
	printBCh := make(chan struct{})
	printCCh := make(chan struct{})
	defer close(printACh)
	defer close(printBCh)
	defer close(printCCh)
	// 打印'A'的goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-printACh
			fmt.Print("A")
			printBCh <- struct{}{}
		}
	}()

	// 打印'B'的goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-printBCh
			fmt.Print("B")
			printCCh <- struct{}{}
		}
	}()

	// 打印'C'的goroutine
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-printCCh
			fmt.Print("C")
			printACh <- struct{}{}
		}
	}()
	//启动，开始打印
	printACh <- struct{}{}
	wg.Wait()
}
