package main

import (
  "context"
  "errors"
  "fmt"
  "golang.org/x/sync/errgroup"
  "net/http"
  "os"
  "os/signal"
  "time"
)


// 基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

func main() {
  g, ctx := errgroup.WithContext(context.Background())
  
  s := http.NewServer()
  
  // http server
  g.Go(func() error {
		return server.ListenAndServe()
	})

  g.Go(func() error {
    select {
		case <-ctx.Done():
      fmt.Println("exit")
      s.Shutdown(context.TODO())
		}
  })
  
  // signal
  g.Go(func() error {
    c := make(chan os.Signal)
    signal.Notify(c)
    
    select {
      case <-ctx.Done():
        return ctx.Err()
      case <-c:
        return nil
      }
    }
  })

  fmt.Println(g.Wait())
}
