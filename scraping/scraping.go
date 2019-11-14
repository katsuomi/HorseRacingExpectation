package scraping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/ChimeraCoder/anaconda"
)

func Scraping() {
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

	// 検索 [ライトコード]
	searchResult, _ := api.GetSearch(`ライトコード`, nil)
	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}
}

// TwitterAccount はTwitterの認証用の情報
type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}
