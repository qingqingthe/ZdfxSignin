package forum

import (
	"github.com/LovesAsuna/ForumSignin/util"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type huahuo struct {
	name,
	baseUrl,
	cookie string
}

func (huahuo *huahuo) FormHash() (string, bool) {
	return FormHash(huahuo)
}

func (huahuo *huahuo) Name() string {
	return huahuo.name
}

func (huahuo *huahuo) BasicUrl() string {
	return huahuo.baseUrl
}

func (huahuo *huahuo) Cookie() string {
	return huahuo.cookie
}

func NewHuaHuoClient() Sign {
	cookie := os.Getenv("HUAHUO_COOKIE")
	name := "花火"
	baseUrl := "https://www.sayhuahuo.com/"
	if len(cookie) == 0 {
		return NewNoCookieClient(name)
	}
	util.Debug(name, "cookie:", cookie)
	client := huahuo{
		name,
		baseUrl,
		cookie,
	}
	return &client
}

func (huahuo *huahuo) Sign() (<-chan string, bool) {
	signUrl := huahuo.baseUrl + "plugin.php?id=dsu_paulsign:sign&operation=qiandao&infloat=1&inajax=1"
	data := make(url.Values)
	hashChannel := make(chan string)
	go func() {
		util.Debug("尝试获取", huahuo.name, "的hash")
		hash, ok := huahuo.FormHash()
		util.Debug("hash结果:", hash, ok)
		if !ok {
			hashChannel <- ""
		} else {
			hashChannel <- hash
		}
	}()
	data["qdxq"] = []string{"kx"}
	data["qdmode"] = []string{"1"}
	data["todaysay"] = []string{"签到"}
	data["fastreply"] = []string{"0"}
	hash := <-hashChannel
	if len(hash) == 0 {
		return nil, false
	}
	data["formhash"] = []string{hash}
	util.Debug("建立", huahuo.name, "的签到请求")
	req, err := http.NewRequest("POST", signUrl, strings.NewReader(data.Encode()))
	util.Debug("请求结果", err == nil)
	if err != nil {
		return nil, false
	}
	req.Header.Set("Cookie", huahuo.cookie)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	c := make(chan string)
	go func() {
		util.Debug("发送", huahuo.name, "的签到请求")
		res, err := client.Do(req)
		if err != nil {
			c <- err.Error()
		}
		util.Debug("获取", huahuo.name, "的签到结果")
		c <- util.GetText(res, "div.c", "div.c")
		close(c)
	}()
	return c, true
}
