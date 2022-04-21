package util

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	"log"
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
	Debug("GetText结果:", html)
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

var debug = func() bool {
	flag := os.Getenv("DEBUG")
	if len(flag) == 0 {
		return false
	}
	debug, err := strconv.ParseBool(flag)
	if err != nil {
		return false
	}
	return debug
}()

func Debug(v ...any) {
	if debug {
		log.Println(v...)
	}
}
