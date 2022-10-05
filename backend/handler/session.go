package handler

import (
	"context"
	"encoding/hex"
	"log"
	"math/rand"
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

func newSid() string {
	sid := make([]byte, sidNBytes)

	_, err := rand.Read(sid)

	if err != nil {
		log.Fatalln("unexpected rand.Read error", err)
	}

	return hex.EncodeToString(sid)
}

func newSidCookie(sid string) *http.Cookie {
	return &http.Cookie{
		Name:     cookieName,
		Value:    sid,
		Path:     "/",
		Domain:   "",
		MaxAge:   oneYearInSeconds,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
}

func getSid(r *http.Request) (string, error) {
	c, err := r.Cookie(cookieName)

	if err != nil || len(c.Value) != sidNChars {
		return "", http.ErrNoCookie
	}

	return c.Value, nil
}

func newSession(w http.ResponseWriter) (string, *entity.Session) {
	sid := newSid()
	http.SetCookie(w, newSidCookie(sid))
	return sid, &entity.Session{}
}

func EnableSession(h http.Handler, store repository.Session) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			h.ServeHTTP(w, r)
			return
		}

		var s *entity.Session
		sid, err := getSid(r)
		if err == nil {
			v, err := store.GetSession(sid)

			if err != nil {
				s = &entity.Session{}
			} else {
				s = v
			}
		} else {
			sid, s = newSession(w)
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
	p := ctx.Value(sessionKey{})
	if p == nil {
		log.Fatalln("session is not enabled")
	}
	return p.(*entity.Session)
}
