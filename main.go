package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/ChimeraCoder/anaconda"
)

func main() {
	// Json読み込み
	raw, error := ioutil.ReadFile("./twitterAccount.json")
	if error != nil {
		fmt.Println(error.Error())
		return
	}

	var twitterAccount TwitterAccount
	// 構造体にセット
	json.Unmarshal(raw, &twitterAccount)

	// 認証
	api := anaconda.NewTwitterApiWithCredentials(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret, twitterAccount.ConsumerKey, twitterAccount.ConsumerSecret)
	v := url.Values{}
	v.Set("count","10000")
	// v.Set("exclude","retweets")
	// 検索
	searchResult, _ := api.GetSearch(`"プログラミング初学者"`, v)
	for _, tweet := range searchResult.Statuses {
		fmt.Printf("%d\n", tweet.User.Id)
		api.FollowUserId(tweet.User.Id,v)
		fmt.Println("--------------------------------------------------------------")
	}
}

// TwitterAccount はTwitterの認証用の情報
type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}



// package main

// import (
// 	"HorseRacingExpectation/scraping"
// 	"bytes"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"

// 	"github.com/PuerkitoBio/goquery"
// 	"github.com/saintfish/chardet"
// 	"golang.org/x/net/html/charset"
// )

// func main() {
// 	id := "c201905050811" //ここにネットケイバの各レース毎のIDをぶち込む
// 	url := "https://race.netkeiba.com/?pid=race&id=" + id + "&mode=top"

// 	// Getリクエスト
// 	res, _ := http.Get(url)
// 	defer res.Body.Close()

// 	// 読み取り
// 	buf, _ := ioutil.ReadAll(res.Body)

// 	// 文字コード判定
// 	det := chardet.NewTextDetector()
// 	detRslt, _ := det.DetectBest(buf)
// 	// => EUC-JP

// 	// 文字コード変換
// 	bReader := bytes.NewReader(buf)
// 	reader, _ := charset.NewReaderLabel(detRslt.Charset, bReader)

// 	// HTMLパース
// 	doc, _ := goquery.NewDocumentFromReader(reader)

// 	var mapNameNumber map[string]string
// 	mapNameNumber = map[string]string{}
// 	var mapNameCount map[string]int
// 	mapNameCount = map[string]int{}
// 	var mapNumberCount map[string]int
// 	mapNumberCount = map[string]int{}

// 	doc.Find(".txt_l").Each(func(i int, s *goquery.Selection) {
// 		href, _ := s.Find("a").Attr("href")
// 		if strings.Contains(href, "horse") == true {
// 			band := s.Find("a").Text()
// 			band2 := s.Parent().Find("td").Next().Next().First().Text()
// 			if band != "" && band2 != "" {
// 				mapNameNumber[band] = band2
// 				mapNameCount[band] = 0
// 				mapNumberCount[band2] = 0
// 			}
// 		}
// 	})
// 	// fmt.Printf("%d\n", mapNameNumber["ラッキーライラック"])
// 	// fmt.Printf("%d\n", mapNameCount)
// 	scraping.Scraping(&mapNameCount, &mapNumberCount, &mapNameNumber)
// }
