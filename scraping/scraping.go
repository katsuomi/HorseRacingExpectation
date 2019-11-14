package scraping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
)

func Scraping(mapNameCount *map[string]int, mapNumberCount *map[int]int, mapNameNumber *map[string]int, horseArray *[]string) {
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

	// 検索
	searchResult, _ := api.GetSearch(`"パドック since:2019-11-10_15:10:00_JST until:2019-11-10_15:40:00_JST"`, nil)
	for i, tweet := range searchResult.Statuses {
		fmt.Printf("%s\n", tweet.FullText)
		fmt.Printf("%d\n", i)
	}
	fmt.Println(horseArray)
	fmt.Println(mapNameCount)
	fmt.Println(mapNumberCount)
	fmt.Println(mapNameNumber)
}

// TwitterAccount はTwitterの認証用の情報
type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}
