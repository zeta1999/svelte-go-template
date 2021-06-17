package auth

import (
	"fmt"
	config "homefill/backend/config"
	"homefill/backend/errset"
	"homefill/backend/logs"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

// GetUserInfo : get user info from google auth
func GetUserInfo(state string, code string) ([]byte, error) {
	if state != config.State {
		logs.LogIt(logs.LogWarn, "GetUserInfo", "invalid config state", fmt.Errorf("state doesn't match"))
		return nil, errset.ErrInternalServer
	}

	token, err := config.GOOGLEAuthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		logs.LogIt(logs.LogWarn, "GetUserInfo", "auth code to token exchange failed", err)
		return nil, errset.ErrInternalServer
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		logs.LogIt(logs.LogWarn, "GetUserInfo", "unable to get user info", err)
		return nil, errset.ErrInternalServer
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logs.LogIt(logs.LogWarn, "GetUserInfo", "unable to read request body", err)
		return nil, errset.ErrInternalServer
	}

	return contents, nil
}
