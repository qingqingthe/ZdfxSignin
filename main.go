package main

import (
	"github.com/LovesAsuna/ForumSignin/forum"
	"github.com/LovesAsuna/ForumSignin/util"
	"log"
	"sync"
)

type client []forum.Sign

func main() {
	clients := client{
		forum.NewHuaHuoClient(),
		forum.NewZdfxClient(),
	}
	wg := sync.WaitGroup{}
	for _, client := range clients {
		wg.Add(1)
		go func(client forum.Sign) {
			util.Debug("开始执行", client.Name(), "的签到操作")
			c, _ := client.Sign()
			util.Debug(client.Name(), "签到完成，打印结果")
			for m := range c {
				log.Println(m)
			}
			wg.Done()
		}(client)
	}
	wg.Wait()
}
