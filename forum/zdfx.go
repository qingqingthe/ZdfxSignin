package forum

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/LovesAsuna/ForumSignin/util"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
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
	client := Zdfx{
		name,
		baseUrl,
		cookie,
	}
	return &client
}

func (zdfx *Zdfx) Do() (<-chan string, bool) {
	c := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		wg.Wait()
		close(c)
	}()

	log.Debug("获取"+zdfx.name, "的hash和token")
	params, err := params(zdfx)
	if err != nil {
		go func() {
			c <- err.Error()
			wg.Add(-2)
		}()
		return c, false
	}
	hash := params[0]
	token := params[1]
	log.Debug(zdfx.name, "hash: ", hash)
	log.Debug(zdfx.name, "token: ", token)
	go func() {
		log.Debug("模拟", zdfx.name, "的签到操作")
		zdfx.sign(c, hash, token)
		wg.Done()
	}()

	go func() {
		log.Debug("模拟", zdfx.name, "的摇奖操作")
		zdfx.lottery(c, token)
		wg.Done()
	}()
	return c, true
}

func setCookie(sign Sign) chromedp.Action {
	cookies := strings.Split(sign.Cookie(), ";")
	slice := make([]string, 0)
	for _, c := range cookies {
		if len(c) == 0 {
			continue
		}
		params := strings.Trim(c, " ")
		slice = append(slice, strings.Split(params, "=")...)
	}
	u, _ := url.Parse(sign.BasicUrl())
	host := u.Host
	return chromedp.ActionFunc(
		func(ctx context.Context) error {
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			for i := 0; i < len(slice); i += 2 {
				err := network.SetCookie(slice[i], slice[i+1]).
					WithExpires(&expr).
					WithDomain(host).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func (zdfx *Zdfx) sign(c chan<- string, hash, token string) {
	req, err := http.NewRequest("GET", zdfx.baseUrl+"plugin.php?id=k_misign:sign", nil)
	if err != nil {
		c <- err.Error()
		return
	}
	req.Header.Set("Cookie", zdfx.Cookie())
	req.Header.Set("User-Agent", util.UA)

	data := req.URL.Query()
	data.Add("formhash", hash)
	data.Add("token", token)
	data.Add("operation", "qiandao")
	data.Add("inajax", "1")
	data.Add("format", "empty")
	data.Add("ajaxtarget", "JD_sign")
	req.URL.RawQuery = data.Encode()
	if err != nil {
		c <- err.Error()
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		c <- err.Error()
		return
	}
	defer resp.Body.Close()
	log.Debug("获取", zdfx.name, "的签到结果")
	c <- util.Text(resp, "root")
}

func (zdfx *Zdfx) lottery(c chan<- string, token string) {
	data := make(url.Values)
	data["token"] = []string{token}
	req, err := http.NewRequest("POST", zdfx.baseUrl+"plugin.php?id=yinxingfei_zzza:yaoyao", strings.NewReader(data.Encode()))
	req.Header.Set("Cookie", zdfx.Cookie())
	req.Header.Set("User-Agent", util.UA)
	req.Header.Set("Content-Type", util.URLEncoded)
	if err != nil {
		c <- err.Error()
		return
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		c <- err.Error()
	}
	log.Debug("获取", zdfx.name, "的摇奖结果")
	type result struct {
		Success bool   `json:"success"`
		Token   bool   `json:"token"`
		Point   string `json:"jifen"`
	}
	e := new(result)
	json.NewDecoder(resp.Body).Decode(e)
	var res string
	if e.Success {
		res = fmt.Sprintf("摇奖成功，获得%s点币\n", e.Point)
	} else {
		if e.Token {
			res = "你已经摇过奖了"
		} else {
			res = "token校验失败"
		}
	}
	c <- fmt.Sprint(res)
}

func params(sign Sign) (params []string, err error) {
	done := make(chan string)
	var hash string
	var ok bool
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		if ev, ok := ev.(*page.EventJavascriptDialogOpening); ok {
			done <- ev.Message
		}
	})
	tasks := chromedp.Tasks{
		chromedp.AttributeValue("#scbar_form input:nth-child(2)", "value", &hash, &ok),
		chromedp.EvaluateAsDevTools(`grecaptcha.execute('6Lfl9bwZAAAAADZ5gAwWyb7U2UynEMHR52oS8d9V', {action: 'create_comment'}).then(token => alert(token))`, nil),
	}
	err = chromedp.Run(ctx,
		setCookie(sign),
		chromedp.ActionFunc(func(cxt context.Context) error {
			_, err := page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(navigator, 'webdriver', { get: () => false, });").Do(cxt)
			if err != nil {
				return err
			}
			_, err = page.AddScriptToEvaluateOnNewDocument("Object.defineProperty(window, 'yzfile', { get: () => 0, });").Do(cxt)
			if err != nil {
				return err
			}
			return nil
		}),
		chromedp.Navigate(sign.BasicUrl()+"k_misign-sign.html"),
		tasks,
	)
	if err != nil {
		return
	}
	params = make([]string, 2)
	if ok {
		params[0] = hash
	}
	params[1] = <-done
	return
}
