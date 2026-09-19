package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	// 版本1 使用WithTimeout
	/*ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	done := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		done <- "work done"
	}()

	select {
	case <-ctx.Done():
		fmt.Println("timeout:", ctx.Err())
	case res := <-done:
		fmt.Println(res)
	}*/

	// 版本2-使用WithCancel
	/*ctx, cancelFunc := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("主动取消：", ctx.Err())
		}
	}(ctx)

	time.Sleep(1 * time.Second)
	cancelFunc()
	time.Sleep(2 * time.Second)*/

	// 版本3-使用WithValue
	//val := TraceCode("test-12345")
	//ctx := context.WithValue(context.Background(), "trace-id", val)
	//
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func(ctx context.Context, wg *sync.WaitGroup) {
	//	value := ctx.Value("trace-id")
	//	fmt.Println("trace-id", value)
	//	wg.Done()
	//}(ctx, &wg)
	//wg.Wait()

	// 版本4-用context.WithTimeout 包一个http.Get
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second) // 3秒接口
		w.Write([]byte("Hello World"))
	}))
	defer server.Close()

	// 用WithTimeout 包住请求，2秒内不反悔就取消
	timeout, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()

	req, err := http.NewRequestWithContext(timeout, "GET", server.URL, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// 这里会走到：context deadline exceeded
		fmt.Println("请求失败:", err)
		fmt.Println("ctx.Err():", timeout.Err())
		return
	}
	defer resp.Body.Close()
	fmt.Println("请求成功:", resp.Status)
}

type TraceCode string
