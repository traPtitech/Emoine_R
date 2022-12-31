package handler

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"net/http"
	"os"

	"github.com/dvsekhvalnov/jose2go/base64url"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type OAuthParams struct {
	CodeChallenge       string `json:"codeChallenge,omitempty"`
	CodeChallengeMethod string `json:"codeChallengeMethod,omitempty"`
	ClientID            string `json:"clientID,omitempty"`
	ResponseType        string `json:"responseType,omitempty"`
}

func OAuthGenerateCodeHandler(c echo.Context) error {
	sess, err := session.Get("sessions", c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "セッションの読み込みに失敗しました")
	}
	clientID := os.Getenv("CLIENT_ID")
	params := OAuthParams{ResponseType: "code", ClientID: clientID, CodeChallengeMethod: "S256"}

	bytesCodeVerifier, err := randBytes(43)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to generate random bytes: %w", err).Error())
	}

	codeVerifier := string(bytesCodeVerifier)
	bytesCodeChallenge := sha256.Sum256(bytesCodeVerifier[:])
	codeChallenge := base64url.Encode(bytesCodeChallenge[:])
	pkceParams.CodeChallenge = codeChallenge
}

func OAuthCallbackHandler(c echo.Context) error {

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