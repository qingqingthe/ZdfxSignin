package forum

import (
	"context"
	"crypto/tls"
	"github.com/LovesAsuna/ForumSignin/util"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"net/http"
	"net/url"
	"strings"
	"time"
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

func setCookie(sign Sign) chromedp.Action {
	cookies := strings.Split(sign.Cookie(), ";")
	slice := make([]string, 0)
	for _, c := range cookies {
		if len(c) == 0 {
			continue
		}
		params := strings.Trim(c, " ")
		array := strings.Split(params, "=")
		for _, s := range array {
			if s != "" {
				slice = append(slice, s)
			}
		}
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
