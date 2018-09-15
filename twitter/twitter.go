package twitter

import (
	"github.com/ChimeraCoder/anaconda"
	"github.com/kazuhito-m/earthquake-twitter-bot/config"
	"log"
)

type Twitter struct {
	twitterApi *anaconda.TwitterApi
}

func CreateTwitter(settings config.TwitterSettings) Twitter {
	twitterApi := createTwitterApi(settings)
	return Twitter{twitterApi}
}

func createTwitterApi(s config.TwitterSettings) *anaconda.TwitterApi {
	anaconda.SetConsumerKey(s.ConsumerKey)
	anaconda.SetConsumerSecret(s.ConsumerSecret)
	return anaconda.NewTwitterApi(s.AccessToken, s.AccessTokenSecret)
}

func (t Twitter) Tweet(content string) {
	//api := t.twitterApi
	//v := url.Values{}
	//_, err := api.PostTweet(content, v)
	//if err == nil {
	//	return
	//}

	// TODO このダミー実装の削除
	log.Println("Tweet()メソッドが呼ばれました。")
	log.Println(content)
}
