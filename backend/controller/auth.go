package controller

import (
	"context"

	"one.now/backend/entity"
)

type AuthCtrler struct {
	allowedEmail string
}

func (s AuthCtrler) Login(ctx context.Context, session *entity.Session, email string) bool {
	ok := session.Loggedin
	if !ok {
		ok = email == s.allowedEmail
		if ok {
			session.Loggedin = ok
		}
	}

	return ok
}

func (s AuthCtrler) Logout(ctx context.Context, session *entity.Session) {
	session.Loggedin = false
}

func NewAuthCtrler(email string) AuthCtrler {
	return AuthCtrler{
		allowedEmail: email,
	}
}
