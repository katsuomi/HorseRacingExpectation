package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func main() {
		id := "p201908050411" //ここにネットケイバの各レース毎のIDをぶち込む
    url := "https://race.netkeiba.com/?pid=race&id="+id+"&mode=top"

    // Getリクエスト
    res, _ := http.Get(url)
    defer res.Body.Close()

    // 読み取り
    buf, _ := ioutil.ReadAll(res.Body)

    // 文字コード判定
    det := chardet.NewTextDetector()
    detRslt, _ := det.DetectBest(buf)
    fmt.Println(detRslt.Charset)
    // => EUC-JP

    // 文字コード変換
    bReader := bytes.NewReader(buf)
    reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

    // HTMLパース
    doc, _ := goquery.NewDocumentFromReader(reader)

    // // titleを抜き出し
    // rslt := doc.Find("body").("div").Text()
    // fmt.Println(rslt)

		doc.Find(".txt_l").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			href, _ := s.Find("a").Attr("href")
			if (strings.Contains(href, "horse") == true) {
				band := s.Find("a").Text()
				fmt.Printf("Review %d: %s\n", i, band)
			}
		})

}