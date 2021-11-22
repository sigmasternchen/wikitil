package twitter

import (
	"fmt"
	"github.com/dghubble/oauth1"
	twauth "github.com/dghubble/oauth1/twitter"
	. "randomarticle/internal/config"
)

const outOfBand = "oob"

func Login(config Config) (AccessConfig, error) {
	oauthConfig := oauth1.Config{
		ConsumerKey:    config.ConsumerKey,
		ConsumerSecret: config.ConsumerSecret,
		CallbackURL:    outOfBand,
		Endpoint:       twauth.AuthorizeEndpoint,
	}

	requestToken, _, err := oauthConfig.RequestToken()
	if err != nil {
		return AccessConfig{}, fmt.Errorf("could not get request token: %w", err)
	}

	authorizationURL, err := oauthConfig.AuthorizationURL(requestToken)
	if err != nil {
		return AccessConfig{}, fmt.Errorf("could not create authorization url: %w", err)
	}
	fmt.Printf("Open this URL in your browser:\n%s\n", authorizationURL.String())

	fmt.Printf("Paste your PIN here: ")
	var verifier string
	_, err = fmt.Scanf("%s", &verifier)
	if err != nil {
		return AccessConfig{}, err
	}

	accessToken, accessSecret, err := oauthConfig.AccessToken(requestToken, "secret does not matter", verifier)
	if err != nil {
		return AccessConfig{}, err
	}

	return AccessConfig{
		AccessToken:  accessToken,
		AccessSecret: accessSecret,
	}, nil
}