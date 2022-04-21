package forum

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type Discuz interface {
	Sign
	FormHash() (string, bool)
}

func FormHash(discuz Discuz) (string, bool) {
	req, _ := http.NewRequest("GET", discuz.BasicUrl(), nil)
	req.Header.Set("Cookie", discuz.Cookie())
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Charset", "UTF-8")
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err.Error(), false
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	Debug(doc.Html())
	if err != nil {
		return err.Error(), false
	}
	return doc.Find("#scbar_form").Find("input:nth-child(2)").Attr("value")
}
