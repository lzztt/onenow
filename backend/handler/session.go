package handler

import (
	"context"
	"encoding/hex"
	"errors"
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

type sessionData struct {
	sid     string
	isNew   bool
	session *entity.Session
	store   repository.Session
}

func isValidSid(sid string) bool {
	return len(sid) == sidNChars
}

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

func getSidFromCookie(r *http.Request) (string, error) {
	c, err := r.Cookie(cookieName)

	if err != nil || !isValidSid(c.Value) {
		return "", http.ErrNoCookie
	}

	return c.Value, nil
}

func EnableSession(h http.Handler, store repository.Session) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			h.ServeHTTP(w, r)
			return
		}

		d := &sessionData{
			session: nil,
			isNew:   false,
			store:   store,
		}

		var err error
		d.sid, err = getSidFromCookie(r)
		if err != nil {
			d.sid = newSid()
			d.isNew = true
			http.SetCookie(w, newSidCookie(d.sid))
		}

		r = r.WithContext(context.WithValue(r.Context(), sessionKey{}, d))

		h.ServeHTTP(w, r)

		if d.session == nil {
			return
		}

		err = store.SaveSession(d.sid, d.session)
		if err != nil {
			log.Println(err)
		}
	})
}

func GetSession(ctx context.Context) (*entity.Session, error) {
	b, ok := ctx.Value(sessionKey{}).(*sessionData)
	if !ok {
		return nil, errors.New("session is not enabled")
	}

	if b.session != nil {
		return b.session, nil
	}

	if b.isNew {
		b.session = &entity.Session{}
		return b.session, nil
	}

	v, err := b.store.GetSession(b.sid)
	if err == nil {
		b.session = v
		return b.session, nil
	} else if err == repository.ErrNoSession {
		b.session = &entity.Session{}
		return b.session, nil
	}

	return nil, err
}
