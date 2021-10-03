package facebook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type App struct {
	AppId                string
	AppSecret            string
	RedirectUri          string
	EnableAppsecretProof bool
}

func New(appID, appSecret, redirectUri string) *App {
	return &App{
		AppId:       appID,
		AppSecret:   appSecret,
		RedirectUri: redirectUri,
	}
}

func (app *App) AppAccessToken() string {
	return app.AppId + "|" + app.AppSecret
}

func (app *App) ParseSignedRequest(signedRequest string) (res Result, err error) {
	strs := strings.SplitN(signedRequest, ".", 2)

	if len(strs) != 2 {
		err = fmt.Errorf("facebook: invalid signed request format")
		return
	}

	sig, e1 := base64.RawURLEncoding.DecodeString(strs[0])

	if e1 != nil {
		err = fmt.Errorf("facebook: fail to decode signed request sig with error %v", e1)
		return
	}

	payload, e2 := base64.RawURLEncoding.DecodeString(strs[1])

	if e2 != nil {
		err = fmt.Errorf("facebook: fail to decode signed request payload with error is %v", e2)
		return
	}

	err = json.Unmarshal(payload, &res)

	if err != nil {
		err = fmt.Errorf("facebook: signed request payload is not a valid json string with error %v", err)
		return
	}

	var hashMethod string
	err = res.DecodeField("algorithm", &hashMethod)

	if err != nil {
		err = fmt.Errorf("facebook: signed request payload doesn't contains a valid 'algorithm' field")
		return
	}

	hashMethod = strings.ToUpper(hashMethod)

	if hashMethod != "HMAC-SHA256" {
		err = fmt.Errorf("facebook: signed request payload uses an unknown HMAC method; expect 'HMAC-SHA256' but actual is '%v'", hashMethod)
		return
	}

	hash := hmac.New(sha256.New, []byte(app.AppSecret))
	hash.Write([]byte(strs[1])) // note: here uses the payload base64 string, not decoded bytes
	expectedSig := hash.Sum(nil)

	if !hmac.Equal(sig, expectedSig) {
		err = fmt.Errorf("facebook: bad signed request signiture")
		return
	}

	return
}

func (app *App) ParseCode(code string) (token string, err error) {
	token, _, _, err = app.ParseCodeInfo(code, "")
	return
}

func (app *App) ParseCodeInfo(code, machineID string) (token string, expires int, newMachineID string, err error) {
	if code == "" {
		err = fmt.Errorf("facebook: code is empty")
		return
	}

	var res Result
	res, err = defaultSession.sendOauthRequest("/oauth/access_token", Params{
		"client_id":     app.AppId,
		"redirect_uri":  app.RedirectUri,
		"client_secret": app.AppSecret,
		"code":          code,
	})

	if err != nil {
		err = fmt.Errorf("facebook: fail to parse facebook response with error %v", err)
		return
	}

	err = res.DecodeField("access_token", &token)

	if err != nil {
		return
	}

	expiresKey := "expires_in"

	if _, ok := res["expires"]; ok {
		expiresKey = "expires"
	}

	if _, ok := res[expiresKey]; ok {
		err = res.DecodeField(expiresKey, &expires)

		if err != nil {
			return
		}
	}

	if _, ok := res["machine_id"]; ok {
		err = res.DecodeField("machine_id", &newMachineID)
	}

	return
}

func (app *App) ExchangeToken(accessToken string) (token string, expires int, err error) {
	if accessToken == "" {
		err = fmt.Errorf("short lived accessToken is empty")
		return
	}

	var res Result
	res, err = defaultSession.sendOauthRequest("/oauth/access_token", Params{
		"grant_type":        "fb_exchange_token",
		"client_id":         app.AppId,
		"client_secret":     app.AppSecret,
		"fb_exchange_token": accessToken,
	})

	if err != nil {
		err = fmt.Errorf("fail to parse facebook response with error %v", err)
		return
	}

	err = res.DecodeField("access_token", &token)

	if err != nil {
		return
	}

	expiresKey := "expires_in"

	if _, ok := res["expires"]; ok {
		expiresKey = "expires"
	}

	if _, ok := res[expiresKey]; ok {
		err = res.DecodeField(expiresKey, &expires)
	}

	return
}

func (app *App) GetCode(accessToken string) (code string, err error) {
	if accessToken == "" {
		err = fmt.Errorf("facebook: long lived accessToken is empty")
		return
	}

	var res Result
	res, err = defaultSession.sendOauthRequest("/oauth/client_code", Params{
		"client_id":     app.AppId,
		"client_secret": app.AppSecret,
		"redirect_uri":  app.RedirectUri,
		"access_token":  accessToken,
	})

	if err != nil {
		err = fmt.Errorf("facebook: fail to get code from facebook with error %v", err)
		return
	}

	err = res.DecodeField("code", &code)
	return
}

func (app *App) Session(accessToken string) *Session {
	return &Session{
		accessToken:          accessToken,
		app:                  app,
		enableAppsecretProof: app.EnableAppsecretProof,
	}
}

func (app *App) SessionFromSignedRequest(signedRequest string) (session *Session, err error) {
	var res Result

	res, err = app.ParseSignedRequest(signedRequest)

	if err != nil {
		return
	}

	var id, token string

	res.DecodeField("user_id", &id) // it's ok without user id.
	err = res.DecodeField("oauth_token", &token)
	if err == nil {
		session = &Session{
			accessToken:          token,
			app:                  app,
			id:                   id,
			enableAppsecretProof: app.EnableAppsecretProof,
		}
		return
	}

	// cannot get "oauth_token"? try to get "code".
	err = res.DecodeField("code", &token)

	if err != nil {
		// no code? no way to continue.
		err = fmt.Errorf("facebook: cannot find 'oauth_token' and 'code'; unable to continue")
		return
	}

	token, err = app.ParseCode(token)

	if err != nil {
		return
	}

	session = &Session{
		accessToken:          token,
		app:                  app,
		id:                   id,
		enableAppsecretProof: app.EnableAppsecretProof,
	}
	return
}
