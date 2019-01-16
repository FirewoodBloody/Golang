package session

import (
	"github.com/FirewoodBloody/go.uuid"
	"sync"
)

type MemorySessionMgr struct {
	sessionMap map[string]Session
	r_wlock    sync.RWMutex
}

func (s *MemorySession) Init(addr string, opting ...string) (err error) {
	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.r_wlock.RLock()
	defer s.r_wlock.RUnlock()

	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = ErrSessionNotExist
		return
	}
	return
}

func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.r_wlock.Lock()
	defer s.r_wlock.RLock()

	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	sessionId := id.String()
	session = NewMemorySession(sessionId)

	s.sessionMap[sessionId] = session
	return
}
