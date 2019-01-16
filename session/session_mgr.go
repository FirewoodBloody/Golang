package session

type SessionMgr interface {
	Init(addr string, opting ...string) (err error)
	CreateSession() (session Session, err error)
	Get(sessionId string) (session Session, err error)
}
