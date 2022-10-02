package repository

import "one.now/backend/entity"

type Session interface {
	GetSession(string) (*entity.Session, error)
	SaveSession(string, *entity.Session) error
}

type InMemorySession struct {
	sessions map[string]*entity.Session
}

func (s InMemorySession) GetSession(sid string) (*entity.Session, error) {
	v, ok := s.sessions[sid]
	if !ok {
		v = &entity.Session{}
		s.sessions[sid] = v
	}

	return v, nil
}

func (s InMemorySession) SaveSession(sid string, v *entity.Session) error {
	s.sessions[sid] = v
	return nil
}

func NewInMemorySession() InMemorySession {
	return InMemorySession{
		sessions: make(map[string]*entity.Session),
	}
}
