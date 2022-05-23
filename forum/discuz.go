package forum

import (
	"crypto/tls"
	"github.com/LovesAsuna/ForumSignin/util"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Discuz interface {
	Sign
	FormHash() (string, bool)
}

var client = &http.Client{
	Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			MaxVersion: tls.VersionTLS12,
		},
	},
}

func FormHash(discuz Discuz) (string, bool) {
	req, _ := http.NewRequest("GET", discuz.BasicUrl(), nil)
	req.Header.Set("Cookie", discuz.Cookie())
	req.Header.Set("User-Agent", util.UA)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err.Error(), false
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err.Error(), false
	}
	return doc.Find("#scbar_form").Find("input:nth-child(2)").Attr("value")
}
