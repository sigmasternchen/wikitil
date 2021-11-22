package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	appConfig "wikitil/internal/config"
)
import "github.com/dghubble/oauth1"

var client *twitter.Client

func Init(appConfig appConfig.Config, access appConfig.AccessConfig)  {
	config := oauth1.NewConfig(appConfig.ConsumerKey, appConfig.ConsumerSecret)
	token := oauth1.NewToken(access.AccessToken, access.AccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client = twitter.NewClient(httpClient)
}