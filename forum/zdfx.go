package forum

import (
	"context"
	"github.com/LovesAsuna/ForumSignin/util"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type Zdfx struct {
	name,
	baseUrl,
	cookie string
}

func (zdfx *Zdfx) FormHash() (string, bool) {
	return FormHash(zdfx)
}

func (zdfx *Zdfx) Name() string {
	return zdfx.name
}

func (zdfx *Zdfx) BasicUrl() string {
	return zdfx.baseUrl
}

func (zdfx *Zdfx) Cookie() string {
	return zdfx.cookie
}

func NewZdfxClient() Sign {
	cookie := os.Getenv("ZDFX_COOKIE")
	name := "终点"
	baseUrl := "https://bbs.zdfx.net/"
	if len(cookie) == 0 {
		return NewNoCookieClient(name)
	}
	util.Debug(name, "cookie:", cookie)
	client := Zdfx{
		name,
		baseUrl,
		cookie,
	}
	return &client
}

func (zdfx *Zdfx) Sign() (<-chan string, bool) {
	c := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		wg.Wait()
		close(c)
	}()

	go func() {
		util.Debug("模拟", zdfx.name, "的签到操作")
		ctx, _ := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		zdfx.signInternal(ctx, c)
		cancel()
		wg.Done()
	}()

	go func() {
		util.Debug("模拟", zdfx.name, "的摇奖操作")
		ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
		zdfx.lottery(ctx, c)
		cancel()
		wg.Done()
	}()
	return c, true
}

func (zdfx *Zdfx) cookieSlice() chromedp.Action {
	cookies := strings.Split(zdfx.cookie, ";")
	slice := make([]string, 0)
	for _, c := range cookies {
		if len(c) == 0 {
			continue
		}
		params := strings.Trim(c, " ")
		slice = append(slice, strings.Split(params, "=")...)
	}
	return chromedp.ActionFunc(
		func(ctx context.Context) error {
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			for i := 0; i < len(slice); i += 2 {
				err := network.SetCookie(slice[i], slice[i+1]).
					WithExpires(&expr).
					WithDomain("bbs.zdfx.net").
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (zdfx *Zdfx) signInternal(ctx context.Context, c chan<- string) {
	util.Debug(zdfx.name, "模拟签到操作启动浏览器")
	err := chromedp.Run(ctx,
		zdfx.cookieSlice(),
		chromedp.Navigate(zdfx.baseUrl+`k_misign-sign.html`),
		chromedp.Click(`#JD_sign`),
	)
	util.Debug(zdfx.name, "模拟签到操作完成，获取结果")
	if err != nil {
		if err == context.DeadlineExceeded {
			c <- "已签到"
		} else {
			c <- err.Error()
		}
	}
}

func (zdfx *Zdfx) lottery(ctx context.Context, c chan<- string) {
	util.Debug(zdfx.name, "模拟摇奖操作启动浏览器")
	var res string
	err := chromedp.Run(ctx,
		zdfx.cookieSlice(),
		chromedp.Navigate(zdfx.baseUrl+`plugin.php?id=yinxingfei_zzza:yaoyao`),
		chromedp.Click(`.num_box > .btn`, chromedp.NodeVisible),
		chromedp.Sleep(5*time.Second),
		chromedp.InnerHTML(`div #res`, &res),
	)
	util.Debug(zdfx.name, "模拟摇将操作完成，获取结果")
	if err != nil {
		c <- err.Error()
	} else {
		c <- res
	}
}
