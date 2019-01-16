package session

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

type RedisSessionMgr struct {
	addr     string
	password string
	pool     *redis.Pool
}

func NewPool(server string, password string) *redis.Pool {
	return &redis.Pool{
		MaxActive:   1000,
		MaxIdle:     64,
		IdleTimeout: 240 * time.Second,
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			/*
					if _,err:=c.Do("AUTH",password);err!=nil{
						c.Close()
						return nil,err
				}
			*/
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func (r *RedisSessionMgr) Init(addr string, opting ...string) (err error) {
	if len(opting) > 0 {
		r.password = opting[0]
	}
	r.pool = NewPool(addr, r.password)
	r.addr = addr
	return
}

func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	return
}
func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	return
}
