package main

import (
	"HorseRacingExpectation/scraping"
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func main() {
	id := "p201908050411" //ここにネットケイバの各レース毎のIDをぶち込む
	url := "https://race.netkeiba.com/?pid=race&id=" + id + "&mode=top"

	// Getリクエスト
	res, _ := http.Get(url)
	defer res.Body.Close()

	// 読み取り
	buf, _ := ioutil.ReadAll(res.Body)

	// 文字コード判定
	det := chardet.NewTextDetector()
	detRslt, _ := det.DetectBest(buf)
	// => EUC-JP

	// 文字コード変換
	bReader := bytes.NewReader(buf)
	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

	// HTMLパース
	doc, _ := goquery.NewDocumentFromReader(reader)

	var mapNameNumber map[string]int
	mapNameNumber = map[string]int{}
	var mapNameCount map[string]int
	mapNameCount = map[string]int{}
	var mapNumberCount map[int]int
	mapNumberCount = map[int]int{}
	horseArray := []string{}

	doc.Find(".txt_l").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find("a").Attr("href")
		if strings.Contains(href, "horse") == true {
			band := s.Find("a").Text()
			band2 := s.Parent().Find("td").Next().Next().First().Text()
			if band != "" && band2 != "" {
				band2Int, _ := strconv.Atoi(band2)
				mapNameNumber[band] = band2Int
				mapNameCount[band] = 0
				mapNumberCount[band2Int] = 0
				horseArray = append(horseArray, band)
			}
		}
	})
	// fmt.Printf("%d\n", mapNameNumber["ラッキーライラック"])
	// fmt.Printf("%d\n", mapNameCount)
	scraping.Scraping(&mapNameCount, &mapNumberCount, &mapNameNumber, &horseArray)
}
