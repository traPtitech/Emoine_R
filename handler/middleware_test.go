package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type IsCalled bool

func (b *IsCalled) InnerFunc(c echo.Context) error {
	*b = true;
	return c.NoContent(http.StatusOK)
}

func setCookieHeader(c echo.Context) {
	cookie := c.Response().Header().Get("Set-Cookie")
	c.Response().Header().Del("Set-Cookie")
	c.Request().Header.Set("Cookie", cookie)
}

func TestCheckLogin(t *testing.T) {
	t.Parallel()

	e := echo.New()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	sess, err := sessions.NewCookieStore([]byte("secret")).New(req, "session")
	if err != nil {
		t.Fatal(err)
	}

	// sessionの値をテストケースに合わせて書き換える
	sess.Values["userid"] = "aaa";

	err = sess.Save(req, rec)
	if err != nil {
		t.Fatalf("failed to save session: %v", err)
	}

	setCookieHeader(c)

	isCalled := IsCalled(true)

	CheckLoginInstance := CheckLogin(isCalled.InnerFunc)
	e.HTTPErrorHandler(CheckLoginInstance(c), c)
}
/*
func TestCheckLogin(t *testing.T) {
	t.Parallel()

	type test struct {
		description string
		userid      string
		isCalled    bool
		statusCode  int
	}
	testcases := []test{
		{"未登録:403", "", false, http.StatusForbidden}, {"登録済み:200", "id", true, http.StatusOK},
	}
	
	for _, tc := range testcases {
		e := echo.New()
		// e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

		req := httptest.NewRequest("GET", "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		sess, err := sessions.NewCookieStore([]byte("secret")).New(req, "session")
		if err != nil {
			t.Fatal(err)
		}
		
		// sess, _ := session.Get("session", c)
		err = sess.Save(req, rec)
		if err != nil {
			t.Fatalf("failed to save session: %v", err)
		}

		setCookieHeader(c)
		isCalled := IsCalled(false)
		CheckLoginInstance := CheckLogin(isCalled.InnerFunc)
		e.HTTPErrorHandler(CheckLoginInstance(c), c)

		assert.Equal(t, tc.statusCode, rec.Code)
		assert.Equal(t, tc.isCalled, isCalled, tc.description)
	}
}*/

func TestCheckIsAdmin(t *testing.T) {
	t.Parallel()
}