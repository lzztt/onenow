package handler

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"time"

	"one.now/backend/entity"
	"one.now/backend/repository"
)

const (
	cookieName       = "__Secure-SID"
	sidNBytes        = 16
	sidNChars        = sidNBytes * 2
	oneYearInSeconds = int(time.Hour/time.Second) * 24 * 365
)

type sessionKey struct{}

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

func getSid(w http.ResponseWriter, r *http.Request) (string, error) {
	c, err := r.Cookie(cookieName)

	if err != nil || len(c.Value) != sidNChars {
		c, err = newSidCookie()

		if err != nil {
			return "", err
		}

		http.SetCookie(w, c)
	}

	return c.Value, nil
}

func EnableSession(h http.Handler, store repository.Session) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			h.ServeHTTP(w, r)
			return
		}

		sid, err := getSid(w, r)
		if err != nil {
			log.Println(err)

			h.ServeHTTP(w, r)
			return
		}

		s, err := store.GetSession(sid)
		if err != nil {
			s = &entity.Session{}
		}

		r = r.WithContext(context.WithValue(r.Context(), sessionKey{}, s))

		h.ServeHTTP(w, r)

		err = store.SaveSession(sid, s)
		if err != nil {
			log.Println(err)
		}
	})
}

func getSession(ctx context.Context) *entity.Session {
	return ctx.Value(sessionKey{}).(*entity.Session)
}
