package utils
import (
	"google.golang.org/api/oauth2/v2"
	"net/http"
)

var httpClient = &http.Client{}

func VerifyGoogleIdToken(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.New(httpClient)
	if err != nil {
		return nil, err
	}
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()

	return tokenInfo, nil
}

