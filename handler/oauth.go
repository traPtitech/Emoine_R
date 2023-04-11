package handler

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	urlPrefix string = "https://q.trap.jp/api/v3"
)

type OAuthParams struct {
	CodeChallenge       string `json:"codeChallenge"`
	CodeChallengeMethod string `json:"codeChallengeMethod"`
	ClientID            string `json:"clientID"`
	ResponseType        string `json:"responseType"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func OAuthGenerateCodeHandler(c echo.Context) error {
	sess, err := session.Get(SessionKey, c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "セッションの読み込みに失敗しました")
	}
	codeVerifier, err := randBytes(43)
	if err != nil {
		return c.String(http.StatusInternalServerError, "codeVerifierの作成に失敗: "+err.Error())
	}

	h := sha256.Sum256(codeVerifier[:])
	codeChallenge := base64.RawURLEncoding.EncodeToString(h[:])
	
	params := OAuthParams{ResponseType: "code", ClientID: ClientID, CodeChallengeMethod: "S256"}
	params.CodeChallenge = codeChallenge

	sess.Values["codeVerifier"] = string(codeVerifier)
	sess.Options = &SessionOptionsDefault
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, "セッションエラー: "+err.Error())
	}
	return c.JSON(http.StatusOK, params)
}

func OAuthCallbackHandler(c echo.Context) error {
	code := c.QueryParam("code")
	if code == "" {
		return c.String(http.StatusBadRequest, "コードがみつかりません")
	}

	sess, err := session.Get(SessionKey, c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "セッションの読み込みに失敗しました: "+err.Error())
	}
	codeVerifier := sess.Values["codeVerifier"].(string)

	res, err := collectOAuthResponse(code, codeVerifier)
	if err != nil {
		return c.String(http.StatusInternalServerError, "セッションの読み込みに失敗しました: "+err.Error())
	}

	sess.Values["accessToken"] = res.AccessToken
	sess.Values["refreshToken"] = res.RefreshToken

	myUserId, err := GetMyUserId(res.AccessToken)
	if err != nil {
		return c.String(http.StatusInternalServerError, "useridを取得できません: "+err.Error())
	}

	sess.Values["userid"] = myUserId

	sess.Options = &SessionOptionsDefault
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, "セッションエラー: "+err.Error())
	}

	return c.String(http.StatusOK, "認証に成功しました")
}

func collectOAuthResponse(code string, codeVerifier string) (AuthResponse, error) {
	form := url.Values{
		"grant_type": {"authorization_code"},
		"client_id": {ClientID},
		"code": {code},
		"code_verifier": {codeVerifier},
	}
	reqBody := strings.NewReader(form.Encode())

	req, err := http.NewRequest("POST", urlPrefix+"/oauth2/token", reqBody)
	if err != nil {
		return AuthResponse{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return AuthResponse{}, err
	}
	if res.StatusCode != 200 {
		return AuthResponse{}, fmt.Errorf("Access Tokenを受け取れませんでした(Status:%d %s)", res.StatusCode, res.Status)
	}
	var authRes AuthResponse
	err = json.NewDecoder(res.Body).Decode(&authRes)
	if err != nil {
		return AuthResponse{}, err
	}

	defer res.Body.Close()
	return authRes, nil
}

func GetMyUserId(accessToken string) (string, error) {
	req, err := http.NewRequest("GET", urlPrefix+"/users/me", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ユーザーIDの取得に失敗しました(Status:%d %s)", res.StatusCode, res.Status)
	}

	type TraqJSON struct {
		Name string `json:"name"`
	}

	var traqJSON TraqJSON
	err = json.NewDecoder(res.Body).Decode(&traqJSON)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	return traqJSON.Name, nil
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_.~"

func randBytes(n int) ([]byte, error) {
	buf := make([]byte, n)
	max := big.NewInt(int64(len(letters)))
	for i := range buf {
		r, err := rand.Int(rand.Reader, max)
		if err != nil {
			return nil, fmt.Errorf("乱数生成に失敗しました %w", err)
		}
		buf[i] = letters[r.Int64()]
	}

	return buf, nil
}
