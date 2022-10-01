package handler

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"net/http"
	"time"
)

const (
	cookieName       = "SID"
	sidNBytes        = 16
	sidNChars        = sidNBytes * 2
	oneYearInSeconds = int(time.Hour/time.Second) * 24 * 365
)

type sessionKey struct{}

type Session struct {
	Loggedin bool
}

var sessions = map[string]*Session{}

func newSid() (string, error) {
	sid := make([]byte, sidNBytes)

	_, err := rand.Read(sid)

	for retries := 3; err != nil && retries > 0; retries-- {
		_, err = rand.Read(sid)
	}

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(sid), nil
}

func newSidCookie() (*http.Cookie, error) {
	sid, err := newSid()
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:     cookieName,
		Value:    sid,
		Path:     "/",
		Domain:   "",
		MaxAge:   oneYearInSeconds,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}, nil
}

func EnableSession(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			c, err := r.Cookie(cookieName)
			if err != nil || len(c.Value) != sidNChars {
				c, err = newSidCookie()
				if err != nil {
					log.Println(err)

					h.ServeHTTP(w, r)
					return
				}

				http.SetCookie(w, c)
				sessions[c.Value] = &Session{}
			}

			r = r.Clone(context.WithValue(r.Context(), sessionKey{}, c.Value))
		}

		h.ServeHTTP(w, r)
	})
}

func GetSession(ctx context.Context) (*Session, error) {
	v := ctx.Value(sessionKey{})
	if v == nil {
		return nil, errors.New("no session key")
	}

	sid := v.(string)
	if len(sid) != sidNChars {
		return nil, errors.New("invalid session key")
	}

	s, ok := sessions[sid]
	if !ok {
		s = &Session{}
		sessions[sid] = s
	}

	return s, nil
}
