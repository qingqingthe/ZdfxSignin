package forum

import (
	"github.com/LovesAsuna/ForumSignin/util"
	log "github.com/sirupsen/logrus"
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
	client := huahuo{
		name,
		baseUrl,
		cookie,
	}
	return &client
}

func (huahuo *huahuo) Do() (<-chan string, bool) {
	signUrl := huahuo.baseUrl + "plugin.php?id=dsu_paulsign:sign&operation=qiandao&infloat=1&inajax=1"
	data := url.Values{}
	hashChannel := make(chan string)
	go func() {
		log.Debug("获取"+huahuo.name, "的hash")
		hash, ok := huahuo.FormHash()
		log.Debug(huahuo.name, "hash: ", hash)
		if !ok {
			hashChannel <- ""
		} else {
			hashChannel <- hash
		}
	}()
	data.Add("qdxq", "kx")
	data.Add("qdmode", "1")
	data.Add("todaysay", "签到")
	data.Add("fastreply", "0")
	hash := <-hashChannel
	if len(hash) == 0 {
		return nil, false
	}
	data.Add("formhash", hash)
	req, err := http.NewRequest("POST", signUrl, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, false
	}
	req.Header.Set("Cookie", huahuo.cookie)
	req.Header.Set("Content-Type", util.URLEncoded)
	req.Header.Set("User-Agent", util.UA)
	c := make(chan string)
	go func() {
		log.Debug("发送", huahuo.name, "的签到请求")
		resp, err := client.Do(req)
		if err != nil {
			c <- err.Error()
		}
		log.Debug("获取", huahuo.name, "的签到结果")
		c <- util.ParseText(resp, "div.c", "div.c")
		close(c)
	}()
	return c, true
}
