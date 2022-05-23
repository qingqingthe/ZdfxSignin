package util

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func ParseText(resp *http.Response, success string, fail string) string {
	doc, _ := doc(resp)
	if element := doc.Find(success); element.Size() == 0 {
		if element = doc.Find(fail); element.Size() != 0 {
			return element.Text()
		} else {
			return "无法解析HTML"
		}
	} else {
		return element.Text()
	}
}

func Text(resp *http.Response, selectors ...string) (s string) {
	doc, _ := doc(resp)
	selection := doc.Selection
	if selectors != nil {
		for _, selector := range selectors {
			selection = doc.Find(selector)
		}
	}
	if selection.Text() == "" {
		s, _ = selection.Html()
	} else {
		s = selection.Text()
	}
	return
}

func doc(res *http.Response) (doc *goquery.Document, err error) {
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return
	}
	html := buf.String()
	log.Debug("Response Body: ", html)
	doc, err = goquery.NewDocumentFromReader(strings.NewReader(html))
	return
}
