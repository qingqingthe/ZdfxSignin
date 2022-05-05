package forum

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
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
	log.Debug(name, "cookie:", cookie)
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
		log.Debug("模拟", zdfx.name, "的签到操作")
		ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
		zdfx.signInternal(ctx, cancel, c)
		defer cancel()
		wg.Done()
	}()

	go func() {
		log.Debug("模拟", zdfx.name, "的摇奖操作")
		ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
		zdfx.lottery(ctx, c)
		defer cancel()
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

func (zdfx *Zdfx) signInternal(ctx context.Context, cancel context.CancelFunc, c chan<- string) {
	log.Debug(zdfx.name, "模拟签到操作启动浏览器")
	sel := `#wp #JD_sign`
	cn := 0
	by := chromedp.ByFunc(func(ctx context.Context, n *cdp.Node) ([]cdp.NodeID, error) {
		cn++
		if cn >= 500 {
			errString := "操作超时，签到成功"
			log.Debug(errString)
			cancel()
			return nil, fmt.Errorf(errString)
		}
		id, count, err := dom.PerformSearch(sel).Do(ctx)
		if err != nil {
			return nil, err
		}

		if count < 1 {
			return []cdp.NodeID{}, nil
		}

		nodes, err := dom.GetSearchResults(id, 0, count).Do(ctx)
		if err != nil {
			return nil, err
		}

		return nodes, nil
	})
	err := chromedp.Run(ctx,
		zdfx.cookieSlice(),
		chromedp.Navigate(zdfx.baseUrl+`k_misign-sign.html`),
		chromedp.Click(sel, by, chromedp.NodeReady),
		chromedp.Reload(),
	)
	log.Debug(zdfx.name, "模拟签到操作完成，获取结果")
	if err != nil {
		if err == context.Canceled {
			c <- "已签到"
		} else {
			c <- err.Error()
		}
	} else {
		c <- "已签到"
	}
}

func (zdfx *Zdfx) lottery(ctx context.Context, c chan<- string) {
	log.Debug(zdfx.name, "模拟摇奖操作启动浏览器")
	var res string
	err := chromedp.Run(ctx,
		zdfx.cookieSlice(),
		chromedp.Navigate(zdfx.baseUrl+`plugin.php?id=yinxingfei_zzza:yaoyao`),
		chromedp.Click(`.num_box > .btn`, chromedp.NodeVisible),
		chromedp.Sleep(5*time.Second),
		chromedp.InnerHTML(`div #res`, &res),
	)
	log.Debug(zdfx.name, "模拟摇将操作完成，获取结果")
	if err != nil {
		c <- err.Error()
	} else {
		c <- res
	}
}
