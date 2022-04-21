package forum

import (
	"crypto/tls"
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
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
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
