package scraping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"net/url"
	"unicode"
	"sort"

	"github.com/ChimeraCoder/anaconda"
)

func Scraping(mapNameCount *map[string]int, mapNumberCount *map[string]int, mapNameNumber *map[string]string) {
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

	numConv := unicode.SpecialCase{
		// 半角の 0 から 9 に対する変換ルール
		unicode.CaseRange{
				0x0030, // Lo: 半角の 0
				0x0039, // Hi: 半角の 9
				[unicode.MaxCase]rune{
						0xff10 - 0x0030, // UpperCase で全角に変換
						0,               // LowerCase では変換しない
						0xff10 - 0x0030, // TitleCase で全角に変換
				},
		},
		// 全角の ０ から ９ に対する変換ルール
		unicode.CaseRange{
				0xff10, // Lo: 全角の ０
				0xFF19, // Hi: 全角の ９
				[unicode.MaxCase]rune{
						0,               // UpperCase では変換しない
						0x0030 - 0xff10, // LowerCase で半角に変換
						0,               // TitleCase では変換しない
				},
		},
	}

	v := url.Values{}
	v.Set("count","10000")
	v.Set("exclude","retweets")
	// 検索
	searchResult, _ := api.GetSearch(`"パドック since:2019-11-24_15:00:00_JST until:2019-11-24_15:40:00_JST"`, v)
	for _, tweet := range searchResult.Statuses {
		// fmt.Printf("%s\n", tweet.FullText)
		// fmt.Printf("%d\n", i)
		for key,_ := range *mapNameCount {
			if(strings.Contains(tweet.FullText,key) == true ){
				(*mapNameCount)[key] += 1
			}
		}
		for key,_ := range *mapNumberCount {
			if(strings.Contains(tweet.FullText,key) == true || strings.Contains(tweet.FullText,strings.ToUpperSpecial(numConv, key)) == true ){
				(*mapNumberCount)[key] += 1
			}
		}
	}

	for key,_ := range *mapNumberCount {
		// fmt.Println(key)
		for key2,_ := range *mapNameNumber {
			if( key == (*mapNameNumber)[key2] ){
				for key3,_ := range *mapNameCount {
					if( key3 == key2 ){
						(*mapNameCount)[key3] += (*mapNumberCount)[key]
					}
				}
			}
			// fmt.Println((*mapNameNumber)[key2])
		 	// fmt.Printf("%T",key2)
		 }
		//  fmt.Println((*mapNameNumber)[key])
		//  fmt.Printf("%T",(*mapNameNumber)[key])
	}

	// fmt.Println(mapNameCount)
	list := List{}
	for k, v := range *mapNameCount {
		e := Entry{k, v}
		list = append(list, e)
	}
	sort.Sort(list)
	fmt.Println(list)

	// fmt.Println(mapNumberCount)
	// fmt.Println(mapNameNumber)
}

// TwitterAccount はTwitterの認証用の情報
type TwitterAccount struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}

type Entry struct {
	name  string
	value int
}
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].name > l[j].name)
	} else {
		return (l[i].value > l[j].value)
	}
}