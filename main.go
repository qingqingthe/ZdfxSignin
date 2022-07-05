package main

import (
	"github.com/LovesAsuna/ForumSignin/forum"
	_ "github.com/LovesAsuna/ForumSignin/util"
	log "github.com/sirupsen/logrus"
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
			log.Debug("开始执行", client.Name(), "的签到操作")
			c, res := client.Do()
			if res {
				for m := range c {
					log.Println(m)
				}
			}
			wg.Done()
		}(client)
	}
	wg.Wait()
}
