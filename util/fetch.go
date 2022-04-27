package util

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func GetText(res *http.Response, success string, fail string) string {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(res.Body)
	if err != nil {
		return err.Error()
	}
	html := buf.String()
	log.Debug("GetText结果:", html)
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))
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

func init() {
	if ok, _ := strconv.ParseBool(os.Getenv("DEBUG")); ok {
		log.SetLevel(log.DebugLevel)
	}
}
